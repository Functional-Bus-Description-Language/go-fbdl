package reg

import (
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

var busWidth int64

func Registerify(regBus *elem.Block) {
	if regBus == nil {
		log.Println("registerification: there is no Main bus; returning nil")
		return
	}

	busWidth = regBus.Width()
	access.Init(busWidth)

	// addr is currently block internal access address, not global address.
	// 0 and 1 are reserved for ID and TIMESTAMP.
	addr := int64(2)

	addr = regFunctionalities(regBus, addr)

	sizes := access.Sizes{}

	sizes.Compact = addr
	sizes.Own = addr

	for _, sb := range regBus.Subblocks() {
		sizes := regBlock(sb.(*elem.Block))
		sizes.Compact += sb.Count() * sizes.Compact
		sizes.BlockAligned += sb.Count() * sizes.BlockAligned
	}

	sizes.BlockAligned = util.AlignToPowerOf2(sizes.BlockAligned + sizes.Own)

	regBus.SetSizes(sizes)

	id := regBus.Status("ID")
	id.SetAccess(access.MakeSingle(0, 0, id.Width()))

	ts := regBus.Status("TIMESTAMP")
	ts.SetAccess(access.MakeSingle(1, 0, busWidth))

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(regBus, 0)
}

func regFunctionalities(blk *elem.Block, addr int64) int64 {
	gp := gap.Pool{}

	addr = regFuncs(blk, addr)
	addr = regStreams(blk, addr)
	//addr = regGroups(blk, addr)
	addr = regConfigs(blk, addr, &gp)
	addr = regMasks(blk, addr)
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

func regFuncs(blk *elem.Block, addr int64) int64 {
	for _, fun := range blk.Funcs() {
		addr = regFunc(fun.(*elem.Func), addr)
	}

	return addr
}

func regStreams(blk *elem.Block, addr int64) int64 {
	for _, stream := range blk.Streams() {
		addr = regStream(stream.(*elem.Stream), addr)
	}

	return addr
}

func regMasks(blk *elem.Block, addr int64) int64 {
	for _, mask := range blk.Masks() {
		addr = regMask(mask.(*elem.Mask), addr)
	}

	return addr
}

func regStatuses(blk *elem.Block, addr int64, gp *gap.Pool) int64 {
	for _, st := range blk.Statuses() {
		if st.Name() == "ID" || st.Name() == "TIMESTAMP" {
			continue
		}
		// Omit elements that have been already registerified as group members.
		if blk.HasElement(st.Name()) {
			continue
		}
		addr = regStatus(st.(*elem.Status), addr, gp)
	}

	return addr
}

func regConfigs(blk *elem.Block, addr int64, gp *gap.Pool) int64 {
	for _, cfg := range blk.Configs() {
		/*
			// Omit elements that have been already registerified as group members.
			if blk.HasElement(insCfg.Name) {
				continue
			}
		*/
		addr = regConfig(cfg.(*elem.Config), addr, gp)
	}

	return addr
}

func regBlock(blk *elem.Block) access.Sizes {
	addr := int64(0)

	addr = regFunctionalities(blk, addr)
	sizes := access.Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, sb := range blk.Subblocks() {
		s := regBlock(sb.(*elem.Block))
		sizes.Compact += sb.Count() * s.Compact
		sizes.BlockAligned += sb.Count() * s.BlockAligned
	}

	sizes.BlockAligned = util.AlignToPowerOf2(addr + sizes.BlockAligned)

	blk.SetSizes(sizes)

	return sizes
}
