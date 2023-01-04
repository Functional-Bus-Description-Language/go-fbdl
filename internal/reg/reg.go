package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/block"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
	"log"
	"sort"
)

var busWidth int64

func Registerify(bus *elem.Block, addTimestamp bool) {
	busWidth = bus.Width
	access.Init(busWidth)

	// addr is currently block internal access address, not global address.
	// 0 is reserved for ID, even if ID is not generated.
	addr := int64(1)

	addr = regFunctionalities(bus, addr)

	timestampAddr := addr
	if addTimestamp {
		addr += 1
	}

	sizes := access.Sizes{}

	sizes.Compact = addr
	sizes.Own = addr

	for _, sb := range bus.Subblocks {
		sbSizes := regBlock(sb)
		sizes.Compact += sb.Count * sbSizes.Compact
		sizes.BlockAligned += sb.Count * sbSizes.BlockAligned
	}

	sizes.BlockAligned = util.AlignToPowerOf2(sizes.BlockAligned + sizes.Own)

	bus.Sizes = sizes

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(bus, 0)

	if block.HasElement(bus, "ID") {
		log.Fatalf("'ID' is reserved element name in main bus")
	}
	id := id()
	id.Access = access.MakeSingle(0, 0, id.Width)
	hash := int64(hash.Hash(bus))
	if busWidth < 32 {
		hash = hash & ((1 << busWidth) - 1)
	}
	// Ignore error, the value has been trimmed to the proper width.
	dflt, _ := val.BitStrFromInt(val.Int(hash), busWidth)
	id.Default = fbdlVal.MakeBitStr(dflt)
	bus.Statics = append(bus.Statics, id)

	if addTimestamp {
		if block.HasElement(bus, "TIMESTAMP") {
			log.Fatalf("'TIMESTAMP' is reserved element name in main bus")
		}
		ts := timestamp()
		ts.Access = access.MakeSingle(timestampAddr, 0, busWidth)
		bus.Statics = append(bus.Statics, ts)
	}
}

func regFunctionalities(blk *elem.Block, addr int64) int64 {
	gp := gap.Pool{}

	addr = regProcs(blk, addr)
	addr = regStreams(blk, addr)
	//addr = regGroups(blk, addr)
	addr = regConfigs(blk, addr, &gp)
	addr = regMasks(blk, addr)
	addr = regStatics(blk, addr, &gp)
	addr = regStatuses(blk, addr, &gp)

	return addr
}

/*
func regGroups(blk *elem.Block, insBlk *ins.Element, addr int64) int64 {
	var grp elem.Group
	for _, g := range insBlk.Grps {
		if g.IsStatus() && g.IsArray() {
			grp, addr = regGroupStatusArray(blk, g, addr)
		} else {
			panic("not yet implemented")
		}

		blkAddGroup(blk, grp)
		for _, st := range grp.Statuses() {
			blkAddStatus(blk, st)
		}
	}

	return addr
}
*/

func regProcs(blk *elem.Block, addr int64) int64 {
	for _, fun := range blk.Procs {
		addr = regProc(fun, addr)
	}

	return addr
}

func regStreams(blk *elem.Block, addr int64) int64 {
	for _, stream := range blk.Streams {
		addr = regStream(stream, addr)
	}

	return addr
}

func regMasks(blk *elem.Block, addr int64) int64 {
	for _, mask := range blk.Masks {
		addr = regMask(mask, addr)
	}

	return addr
}

func regStatics(blk *elem.Block, addr int64, gp *gap.Pool) int64 {
	statics := []*elem.Static{}

	for _, st := range blk.Statics {
		// Omit elements that have been already registerified as group members.
		if st.Access != nil {
			continue
		}
		statics = append(statics, st)
	}

	sortFunc := func(sts []*elem.Static) func(int, int) bool {
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
		addr = regStatic(st, addr, gp)
	}

	return addr
}

func regStatuses(blk *elem.Block, addr int64, gp *gap.Pool) int64 {
	atomicSts := []*elem.Status{}
	nonAtomicSts := []*elem.Status{}

	for _, st := range blk.Statuses {
		// Omit elements that have been already registerified as group members.
		if st.Access != nil {
			continue
		}

		if st.Atomic {
			atomicSts = append(atomicSts, st)
		} else {
			nonAtomicSts = append(nonAtomicSts, st)
		}
	}

	sortFunc := func(sts []*elem.Status) func(int, int) bool {
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
		addr = regAtomicStatus(st, addr, gp)
	}
	for _, st := range nonAtomicSts {
		addr = regNonAtomicStatus(st, addr, gp)
	}

	return addr
}

func regConfigs(blk *elem.Block, addr int64, gp *gap.Pool) int64 {
	for _, cfg := range blk.Configs {
		// Omit elements that have been already registerified as group members.
		if cfg.Access != nil {
			continue
		}

		addr = regConfig(cfg, addr, gp)
	}

	return addr
}

func regBlock(blk *elem.Block) access.Sizes {
	addr := int64(0)

	addr = regFunctionalities(blk, addr)
	sizes := access.Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, sb := range blk.Subblocks {
		b := regBlock(sb)
		sizes.Compact += sb.Count * b.Compact
		sizes.BlockAligned += sb.Count * b.BlockAligned
	}

	sizes.BlockAligned = util.AlignToPowerOf2(addr + sizes.BlockAligned)

	blk.Sizes = sizes

	return sizes
}
