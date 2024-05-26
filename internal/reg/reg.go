package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/block"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
	"log"
	"sort"
)

var busWidth int64

func Registerify(bus *fn.Block, addTimestamp bool) {
	busWidth = bus.Width
	access.Init(busWidth)

	// addr is currently block internal access address, not global address.
	// 0 is reserved for ID, even if ID is not generated.
	addr := makeAddr(1, busWidth)

	regFunctionalities(bus, &addr)

	timestampAddr := addr.value
	if addTimestamp {
		addr.inc(1)
	}

	sizes := access.Sizes{}

	sizes.Compact = addr.value
	sizes.Own = addr.value

	for _, sb := range bus.Subblocks {
		sbSizes := regBlock(sb)
		sizes.Compact += sb.Count * sbSizes.Compact
		sizes.BlockAligned += sb.Count * sbSizes.BlockAligned
	}

	sizes.BlockAligned = util.AlignToPowerOf2(sizes.BlockAligned + util.AlignToPowerOf2(sizes.Own))

	bus.Sizes = sizes

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(bus, 0)

	if block.HasFunctionality(bus, "ID") {
		log.Fatalf("'ID' is reserved functionality name in main bus")
	}
	id := id()
	id.Access = access.MakeSingle(0, 0, id.Width)
	hash := int64(hash.Hash(bus))
	if busWidth < 32 {
		hash = hash & ((1 << busWidth) - 1)
	}
	// Ignore error, the value has been trimmed to the proper width.
	val, _ := val.BitStrFromInt(val.Int(hash), busWidth)
	id.InitValue = fbdlVal.MakeBitStr(val)
	bus.Statics = append(bus.Statics, id)

	if addTimestamp {
		if block.HasFunctionality(bus, "TIMESTAMP") {
			log.Fatalf("'TIMESTAMP' is reserved functionality name in main bus")
		}
		ts := timestamp()
		ts.Access = access.MakeSingle(timestampAddr, 0, busWidth)
		bus.Statics = append(bus.Statics, ts)
	}
}

func regFunctionalities(blk *fn.Block, addr *address) {
	gp := gap.Pool{}

	regProcs(blk, addr)
	regStreams(blk, addr)
	//regGroups(blk, addr)
	regConfigs(blk, addr, &gp)
	regMasks(blk, addr)
	regStatics(blk, addr, &gp)
	regStatuses(blk, addr, &gp)
}

/*
func regGroups(blk *fn.Block, insBlk *ins.Element, addr *address) {
	var grp fn.Group
	for _, g := range insBlk.Grps {
		if g.IsStatus() && g.IsArray() {
			grp regGroupStatusArray(blk, g, addr)
		} else {
			panic("unimplemented")
		}

		blkAddGroup(blk, grp)
		for _, st := range grp.Statuses() {
			blkAddStatus(blk, st)
		}
	}

}
*/

func regProcs(blk *fn.Block, addr *address) {
	for _, fun := range blk.Procs {
		regProc(fun, addr)
	}
}

func regStreams(blk *fn.Block, addr *address) {
	for _, stream := range blk.Streams {
		regStream(stream, addr)
	}
}

func regMasks(blk *fn.Block, addr *address) {
	for _, mask := range blk.Masks {
		regMask(mask, addr)
	}
}

func regStatics(blk *fn.Block, addr *address, gp *gap.Pool) {
	statics := []*fn.Static{}

	for _, st := range blk.Statics {
		// Omit functionalities that have been already registerified as group members.
		if st.Access != nil {
			continue
		}
		statics = append(statics, st)
	}

	sortFunc := func(sts []*fn.Static) func(int, int) bool {
		return func(i, j int) bool {
			if sts[i].IsArray && !sts[j].IsArray {
				return true
			} else if !sts[i].IsArray && sts[j].IsArray {
				return false
			}

			if sts[i].Width > sts[j].Width {
				return true
			}
			return false
		}
	}

	sort.SliceStable(statics, sortFunc(statics))

	for _, st := range statics {
		regStatic(st, addr, gp)
	}
}

func regStatuses(blk *fn.Block, addr *address, gp *gap.Pool) {
	atomicSts := []*fn.Status{}
	nonAtomicSts := []*fn.Status{}

	for _, st := range blk.Statuses {
		// Omit functionalities that have been already registerified as group members.
		if st.Access != nil {
			continue
		}

		if st.Atomic {
			atomicSts = append(atomicSts, st)
		} else {
			nonAtomicSts = append(nonAtomicSts, st)
		}
	}

	sortFunc := func(sts []*fn.Status) func(int, int) bool {
		return func(i, j int) bool {
			if sts[i].IsArray && !sts[j].IsArray {
				return true
			} else if !sts[i].IsArray && sts[j].IsArray {
				return false
			}

			if sts[i].Width > sts[j].Width {
				return true
			}
			return false
		}
	}

	sort.SliceStable(atomicSts, sortFunc(atomicSts))
	sort.SliceStable(nonAtomicSts, sortFunc(nonAtomicSts))

	for _, st := range atomicSts {
		regAtomicStatus(st, addr, gp)
	}
	for _, st := range nonAtomicSts {
		regNonAtomicStatus(st, addr, gp)
	}
}

func regConfigs(blk *fn.Block, addr *address, gp *gap.Pool) {
	atomicCfgs := []*fn.Config{}
	nonAtomicCfgs := []*fn.Config{}

	for _, cfg := range blk.Configs {
		// Omit functionalities that have been already registerified as group members.
		if cfg.Access != nil {
			continue
		}

		if cfg.Atomic {
			atomicCfgs = append(atomicCfgs, cfg)
		} else {
			nonAtomicCfgs = append(nonAtomicCfgs, cfg)
		}
	}

	sortFunc := func(cfgs []*fn.Config) func(int, int) bool {
		return func(i, j int) bool {
			if cfgs[i].IsArray && !cfgs[j].IsArray {
				return true
			} else if !cfgs[i].IsArray && cfgs[j].IsArray {
				return false
			}

			if cfgs[i].Width > cfgs[j].Width {
				return true
			}
			return false
		}
	}

	sort.SliceStable(atomicCfgs, sortFunc(atomicCfgs))
	sort.SliceStable(nonAtomicCfgs, sortFunc(nonAtomicCfgs))

	for _, cfg := range atomicCfgs {
		regAtomicConfig(cfg, addr, gp)
	}
	for _, cfg := range nonAtomicCfgs {
		regNonAtomicConfig(cfg, addr, gp)
	}
}

func regBlock(blk *fn.Block) access.Sizes {
	addr := makeAddr(0, busWidth)

	regFunctionalities(blk, &addr)
	sizes := access.Sizes{BlockAligned: 0, Own: addr.value, Compact: addr.value}

	for _, sb := range blk.Subblocks {
		b := regBlock(sb)
		sizes.Compact += sb.Count * b.Compact
		sizes.BlockAligned += sb.Count * b.BlockAligned
	}

	sizes.BlockAligned = util.AlignToPowerOf2(util.AlignToPowerOf2(addr.value) + sizes.BlockAligned)

	blk.Sizes = sizes

	return sizes
}
