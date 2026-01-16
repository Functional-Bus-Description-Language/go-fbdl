package reg

import (
	"log"
	"sort"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/block"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

var busAlign int64
var busWidth int64

func Registerify(bus *fn.Block, addTimestamp bool) {
	busAlign = bus.Align
	busWidth = bus.Width
	types.Init(busWidth)

	// addr is currently block internal access address, not global address.
	// 0 is reserved for ID, even if ID is not generated.
	addr := int64(1)

	addr = regFunctionalities(bus, addr)

	timestampAddr := addr
	if addTimestamp {
		addr += 1
	}

	sizes := types.Sizes{}

	sizes.Compact = addr
	sizes.Own = addr

	for _, sb := range bus.Subblocks {
		sbSizes := regBlock(sb)
		sizes.Compact += sb.Count * sbSizes.Compact
		sizes.BlockAligned += sb.Count * sbSizes.BlockAligned
	}

	bus.Sizes = alignBlockSize(sizes, bus.Align)

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(bus, 0)

	if block.HasFunctionality(bus, "ID") {
		log.Fatalf("'ID' is reserved functionality name in main bus")
	}
	id := id()
	id.Access = types.MakeSingleAccess(0, 0, id.Width)
	hash := int64(hash.Hash(bus))
	if busWidth < 32 {
		hash = hash & ((1 << busWidth) - 1)
	}
	// Ignore error, the value has been trimmed to the proper width.
	val, _ := val.BitStrFromInt(val.Int(hash), busWidth)
	id.InitValue = types.MakeBitStr(val)
	bus.Statics = append(bus.Statics, id)

	if addTimestamp {
		if block.HasFunctionality(bus, "TIMESTAMP") {
			log.Fatalf("'TIMESTAMP' is reserved functionality name in main bus")
		}
		ts := timestamp()
		ts.Access = types.MakeSingleAccess(timestampAddr, 0, busWidth)
		bus.Statics = append(bus.Statics, ts)
	}
}

func regFunctionalities(blk *fn.Block, addr int64) int64 {
	gp := gap.Pool{}

	addr = regProcs(blk, addr)
	addr = regStreams(blk, addr)
	//addr = regGroups(blk, addr)

	addr = regConfigs(blk, addr, &gp)
	addr = regMasks(blk, addr)
	addr = regStatics(blk, addr, &gp)
	addr = regStatuses(blk, addr, &gp)

	// Registerify irqs as the last ones.
	// Single irqs have a width of 1, so they can easily fit gaps.
	addr = regIrqs(blk, addr, &gp)

	return addr
}

/*
func regGroups(blk *fn.Block, insBlk *ins.Element, addr int64) int64 {
	var grp fn.Group
	for _, g := range insBlk.Grps {
		if g.IsStatus() && g.IsArray() {
			grp, addr = regGroupStatusArray(blk, g, addr)
		} else {
			panic("unimplemented")
		}

		blkAddGroup(blk, grp)
		for _, st := range grp.Statuses() {
			blkAddStatus(blk, st)
		}
	}

	return addr
}
*/

func regProcs(blk *fn.Block, addr int64) int64 {
	for _, fun := range blk.Procs {
		addr = regProc(fun, addr)
	}

	return addr
}

func regStreams(blk *fn.Block, addr int64) int64 {
	for _, stream := range blk.Streams {
		addr = regStream(stream, addr)
	}

	return addr
}

func regMasks(blk *fn.Block, addr int64) int64 {
	for _, mask := range blk.Masks {
		addr = regMask(mask, addr)
	}

	return addr
}

func regStatics(blk *fn.Block, addr int64, gp *gap.Pool) int64 {
	statics := []*fn.Static{}
	statics = append(statics, blk.Statics...)

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
		addr = regStatic(st, addr, gp)
	}

	return addr
}

func regStatuses(blk *fn.Block, addr int64, gp *gap.Pool) int64 {
	atomicSts := []*fn.Status{}
	nonAtomicSts := []*fn.Status{}

	for _, st := range blk.Statuses {
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
		addr = regAtomicStatus(st, addr, gp)
	}
	for _, st := range nonAtomicSts {
		addr = regNonAtomicStatus(st, addr, gp)
	}

	return addr
}

func regIrqs(blk *fn.Block, addr int64, gp *gap.Pool) int64 {
	for _, irq := range blk.Irqs {
		addr = regIrq(irq, addr, gp)
	}
	return addr
}

func regConfigs(blk *fn.Block, addr int64, gp *gap.Pool) int64 {
	atomicCfgs := []*fn.Config{}
	nonAtomicCfgs := []*fn.Config{}

	for _, cfg := range blk.Configs {
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
		addr = regAtomicConfig(cfg, addr, gp)
	}
	for _, cfg := range nonAtomicCfgs {
		addr = regNonAtomicConfig(cfg, addr, gp)
	}

	return addr
}

func regBlock(blk *fn.Block) types.Sizes {
	addr := int64(0)

	addr = regFunctionalities(blk, addr)
	sizes := types.Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, sb := range blk.Subblocks {
		b := regBlock(sb)
		sizes.Compact += sb.Count * b.Compact
		sizes.BlockAligned += sb.Count * b.BlockAligned
	}

	align := blk.Align
	if align == 0 {
		align = busAlign
	}

	blk.Sizes = alignBlockSize(sizes, align)

	return blk.Sizes
}

func alignBlockSize(sizes types.Sizes, align int64) types.Sizes {
	if align == 0 {
		sizes.BlockAligned = util.AlignToPowerOf2(util.AlignToPowerOf2(sizes.Own) + sizes.BlockAligned)
	} else {
		sizes.BlockAligned = util.AlignToMultipleOf(
			util.AlignToMultipleOf(sizes.Own, align)+sizes.BlockAligned,
			align,
		)
	}

	return sizes
}
