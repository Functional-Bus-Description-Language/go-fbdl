package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"log"
)

var busWidth int64

func Registerify(insBus *ins.Element) *Block {
	if insBus == nil {
		log.Println("registerification: there is no main bus; returning nil")
		return nil
	}

	busWidth = int64(insBus.Properties["width"].(val.Int))

	regBus := Block{
		Name:    "main",
		IsArray: insBus.IsArray,
		Count:   int64(insBus.Count),
		Masters: int64(insBus.Properties["masters"].(val.Int)),
		Width:   int64(insBus.Properties["width"].(val.Int)),
	}

	// addr is current block internal access address, not global address.
	// 0 and 1 are reserved for x_uuid_x and x_timestamp_x.
	addr := int64(2)

	addr = registerifyFunctionalities(&regBus, insBus, addr)

	regBus.Sizes.Compact = addr
	regBus.Sizes.Own = addr

	for _, e := range insBus.Elements {
		if e.BaseType == "block" {
			sb, sizes := registerifyBlock(e)
			regBus.Sizes.Compact += e.Count * sizes.Compact
			regBus.Sizes.BlockAligned += e.Count * sizes.BlockAligned
			regBus.addSubblock(sb)
		}
	}

	uuid, _ := insBus.Elements.Get("x_uuid_x")
	regBus.addStatus(
		&Status{
			Name:    uuid.Name,
			Count:   uuid.Count,
			Access:  makeAccessSingle(0, 0, busWidth),
			Atomic:  bool(uuid.Properties["atomic"].(val.Bool)),
			Width:   int64(uuid.Properties["width"].(val.Int)),
			Default: MakeBitStr(uuid.Properties["default"].(val.BitStr)),
		},
	)

	ts, _ := insBus.Elements.Get("x_timestamp_x")
	regBus.addStatus(
		&Status{
			Name:    ts.Name,
			Count:   ts.Count,
			Access:  makeAccessSingle(1, 0, busWidth),
			Atomic:  bool(ts.Properties["atomic"].(val.Bool)),
			Width:   int64(ts.Properties["width"].(val.Int)),
			Default: MakeBitStr(ts.Properties["default"].(val.BitStr)),
		},
	)

	regBus.Sizes.BlockAligned = util.AlignToPowerOf2(
		regBus.Sizes.BlockAligned + regBus.Sizes.Own,
	)

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(&regBus, 0)

	return &regBus
}

func registerifyFunctionalities(blk *Block, insBlk *ins.Element, addr int64) int64 {
	if len(insBlk.Elements) == 0 {
		return addr
	}

	addr = registerifyFuncs(blk, insBlk, addr)
	addr = registerifyStatuses(blk, insBlk, addr)

	return addr
}

func registerifyFuncs(blk *Block, insBlk *ins.Element, addr int64) int64 {
	insFuncs := insBlk.Elements.GetAllByBaseType("func")

	var fun *Func

	for _, insFun := range insFuncs {
		fun, addr = registerifyFunc(insFun, addr)
		blk.addFunc(fun)
	}

	return addr
}

// Current approach is trivial. Even groups are not respected.
func registerifyStatuses(blk *Block, insBlk *ins.Element, addr int64) int64 {
	insStatuses := insBlk.Elements.GetAllByBaseType("status")

	var st *Status

	for _, insSt := range insStatuses {
		if insSt.Name == "x_uuid_x" || insSt.Name == "x_timestamp_x" {
			continue
		}

		st, addr = registerifyStatus(insSt, addr)
		blk.addStatus(st)
	}

	return addr
}

func registerifyBlock(insBlk *ins.Element) (*Block, Sizes) {
	addr := int64(0)

	b := Block{
		Name:    insBlk.Name,
		IsArray: insBlk.IsArray,
		Count:   int64(insBlk.Count),
	}

	addr = registerifyFunctionalities(&b, insBlk, addr)
	sizes := Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, e := range insBlk.Elements {
		if e.BaseType == "block" {
			sb, s := registerifyBlock(e)
			sizes.Compact += e.Count * s.Compact
			sizes.BlockAligned += e.Count * s.BlockAligned
			b.addSubblock(sb)
		}
	}

	sizes.BlockAligned = util.AlignToPowerOf2(addr + sizes.BlockAligned)

	b.Sizes = sizes

	return &b, sizes
}
