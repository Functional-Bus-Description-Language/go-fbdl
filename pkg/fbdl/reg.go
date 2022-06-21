package fbdl

import (
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

var busWidth int64

func Registerify(insBus *ins.Element) *Block {
	if insBus == nil {
		log.Println("registerification: there is no Main bus; returning nil")
		return nil
	}

	busWidth = int64(insBus.Props["width"].(val.Int))

	regBus := Block{
		Name:    "Main",
		Doc:     insBus.Doc,
		IsArray: insBus.IsArray,
		Count:   int64(insBus.Count),
		Masters: int64(insBus.Props["masters"].(val.Int)),
		Width:   int64(insBus.Props["width"].(val.Int)),
	}

	regBus.addConsts(insBus)

	// addr is current block internal access address, not global address.
	// 0 and 1 are reserved for ID and TIMESTAMP.
	addr := int64(2)

	addr = regFunctionalities(&regBus, insBus, addr)

	regBus.Sizes.Compact = addr
	regBus.Sizes.Own = addr

	for _, e := range insBus.Elems {
		if e.Type == "block" {
			sb, sizes := regBlock(e)
			regBus.Sizes.Compact += e.Count * sizes.Compact
			regBus.Sizes.BlockAligned += e.Count * sizes.BlockAligned
			regBus.addSubblock(sb)
		}
	}

	id, _ := insBus.Elems.Get("ID")
	regBus.addStatus(
		&Status{
			Name:    id.Name,
			Doc:     id.Doc,
			Count:   id.Count,
			Access:  makeAccessSingle(0, 0, busWidth),
			Atomic:  bool(id.Props["atomic"].(val.Bool)),
			Width:   int64(id.Props["width"].(val.Int)),
			Default: MakeBitStr(id.Props["default"].(val.BitStr)),
		},
	)

	ts, _ := insBus.Elems.Get("TIMESTAMP")
	regBus.addStatus(
		&Status{
			Name:    ts.Name,
			Doc:     ts.Doc,
			Count:   ts.Count,
			Access:  makeAccessSingle(1, 0, busWidth),
			Atomic:  bool(ts.Props["atomic"].(val.Bool)),
			Width:   int64(ts.Props["width"].(val.Int)),
			Default: MakeBitStr(ts.Props["default"].(val.BitStr)),
		},
	)

	regBus.Sizes.BlockAligned = util.AlignToPowerOf2(
		regBus.Sizes.BlockAligned + regBus.Sizes.Own,
	)

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(&regBus, 0)

	return &regBus
}

func regFunctionalities(blk *Block, insBlk *ins.Element, addr int64) int64 {
	if len(insBlk.Elems) == 0 {
		return addr
	}

	gp := gapPool{}

	addr = regFuncs(blk, insBlk, addr)
	addr = regGroups(blk, insBlk, addr)
	addr = regConfigs(blk, insBlk, addr, &gp)
	addr = regMasks(blk, insBlk, addr)
	addr = regStatuses(blk, insBlk, addr, &gp)

	return addr
}

func regGroups(blk *Block, insBlk *ins.Element, addr int64) int64 {
	var grp Group
	for _, g := range insBlk.Grps {
		if g.IsStatus() && g.IsArray() {
			grp, addr = regGroupStatusArray(blk, g, addr)
		} else {
			panic("not yet implemented")
		}

		blk.addGroup(grp)
		for _, st := range grp.Statuses() {
			blk.addStatus(st)
		}
	}

	return addr
}

func regFuncs(blk *Block, insBlk *ins.Element, addr int64) int64 {
	insFuncs := insBlk.Elems.GetAllByType("func")

	var fun *Func

	for _, insFun := range insFuncs {
		fun, addr = regFunc(insFun, addr)
		blk.addFunc(fun)
	}

	return addr
}

func regMasks(blk *Block, insBlk *ins.Element, addr int64) int64 {
	insMasks := insBlk.Elems.GetAllByType("mask")

	var mask *Mask

	for _, insMask := range insMasks {
		mask, addr = regMask(insMask, addr)
		blk.addMask(mask)
	}

	return addr
}

func regStatuses(blk *Block, insBlk *ins.Element, addr int64, gp *gapPool) int64 {
	insStatuses := insBlk.Elems.GetAllByType("status")

	var st *Status

	for _, insSt := range insStatuses {
		if insSt.Name == "ID" || insSt.Name == "TIMESTAMP" {
			continue
		}
		// Omit elements that have been already registerified as group members.
		if blk.hasElement(insSt.Name) {
			continue
		}
		st, addr = regStatus(insSt, addr, gp)
		blk.addStatus(st)
	}

	return addr
}

func regConfigs(blk *Block, insBlk *ins.Element, addr int64, gp *gapPool) int64 {
	insConfigs := insBlk.Elems.GetAllByType("config")

	var cfg *Config

	for _, insCfg := range insConfigs {
		// Omit elements that have been already registerified as group members.
		if blk.hasElement(insCfg.Name) {
			continue
		}
		cfg, addr = regConfig(insCfg, addr, gp)
		blk.addConfig(cfg)
	}

	return addr
}

func regBlock(insBlk *ins.Element) (*Block, Sizes) {
	addr := int64(0)

	b := Block{
		Name:    insBlk.Name,
		Doc:     insBlk.Doc,
		IsArray: insBlk.IsArray,
		Count:   int64(insBlk.Count),
		Masters: int64(insBlk.Props["masters"].(val.Int)),
	}

	b.addConsts(insBlk)

	addr = regFunctionalities(&b, insBlk, addr)
	sizes := Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, e := range insBlk.Elems {
		if e.Type == "block" {
			sb, s := regBlock(e)
			sizes.Compact += e.Count * s.Compact
			sizes.BlockAligned += e.Count * s.BlockAligned
			b.addSubblock(sb)
		}
	}

	sizes.BlockAligned = util.AlignToPowerOf2(addr + sizes.BlockAligned)

	b.Sizes = sizes

	return &b, sizes
}
