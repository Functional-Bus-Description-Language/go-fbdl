package reg

import (
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	//"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	//fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

var busWidth int64

func Registerify(regBus *elem.Block) {
	if regBus == nil {
		log.Println("registerification: there is no Main bus; returning nil")
		return
	}

	busWidth = regBus.Width
	access.Init(busWidth)

	/*
		regBus := elem.Block{
			Name:    "Main",
			Doc:     insBus.Doc,
			IsArray: insBus.IsArray,
			Count:   int64(insBus.Count),
			Masters: int64(insBus.Props["masters"].(val.Int)),
			Width:   int64(insBus.Props["width"].(val.Int)),
		}
	*/

	/*
		for name, v := range insBus.Consts {
			regBus.AddConst(name, v)
		}
	*/

	// addr is current block internal access address, not global address.
	// 0 and 1 are reserved for ID and TIMESTAMP.
	addr := int64(2)

	addr = regFunctionalities(regBus, addr)

	regBus.Sizes.Compact = addr
	regBus.Sizes.Own = addr

	for _, sb := range regBus.Subblocks {
		sizes := regBlock(sb)
		regBus.Sizes.Compact += sb.Count() * sizes.Compact
		regBus.Sizes.BlockAligned += sb.Count() * sizes.BlockAligned
	}

	/*
		id, _ := insBus.Elems.Get("ID")
		blkAddStatus(&regBus,
			&elem.Status{
				Name:    id.Name,
				Doc:     id.Doc,
				Count:   id.Count,
				Access:  access.MakeSingle(0, 0, busWidth),
				Atomic:  bool(id.Props["atomic"].(val.Bool)),
				Width:   int64(id.Props["width"].(val.Int)),
				Default: fbdlVal.MakeBitStr(id.Props["default"].(val.BitStr)),
			},
		)

		ts, _ := insBus.Elems.Get("TIMESTAMP")
		blkAddStatus(&regBus,
			&elem.Status{
				Name:    ts.Name,
				Doc:     ts.Doc,
				Count:   ts.Count,
				Access:  access.MakeSingle(1, 0, busWidth),
				Atomic:  bool(ts.Props["atomic"].(val.Bool)),
				Width:   int64(ts.Props["width"].(val.Int)),
				Default: fbdlVal.MakeBitStr(ts.Props["default"].(val.BitStr)),
			},
		)
	*/

	regBus.Sizes.BlockAligned = util.AlignToPowerOf2(
		regBus.Sizes.BlockAligned + regBus.Sizes.Own,
	)

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
	for _, fun := range blk.Funcs {
		addr = regFunc(fun, addr)
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

func regStatuses(blk *elem.Block, addr int64, gp *gap.Pool) int64 {
	for _, st := range blk.Statuses {
		if st.Name == "ID" || st.Name == "TIMESTAMP" {
			continue
		}
		/*
			// Omit elements that have been already registerified as group members.
			if blk.HasElement(st.Name) {
				continue
			}
		*/
		addr = regStatus(st, addr, gp)
	}

	return addr
}

func regConfigs(blk *elem.Block, addr int64, gp *gap.Pool) int64 {
	for _, cfg := range blk.Configs {
		/*
			// Omit elements that have been already registerified as group members.
			if blk.HasElement(insCfg.Name) {
				continue
			}
		*/
		addr = regConfig(cfg, addr, gp)
	}

	return addr
}

func regBlock(blk *elem.Block) access.Sizes {
	addr := int64(0)

	/*
		b := elem.Block{
			Name:    insBlk.Name,
			Doc:     insBlk.Doc,
			IsArray: insBlk.IsArray,
			Count:   int64(insBlk.Count),
			Masters: int64(insBlk.Props["masters"].(val.Int)),
		}
	*/

	/*
		for name, v := range insBlk.Consts {
			b.AddConst(name, v)
		}
	*/

	addr = regFunctionalities(blk, addr)
	sizes := access.Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, sb := range blk.Subblocks {
		s := regBlock(sb)
		sizes.Compact += sb.Count() * s.Compact
		sizes.BlockAligned += sb.Count() * s.BlockAligned
	}

	sizes.BlockAligned = util.AlignToPowerOf2(addr + sizes.BlockAligned)

	blk.Sizes = sizes

	return sizes
}
