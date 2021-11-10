package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"log"
	"sort"
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
			sizes := registerifyBlock(e)
			count := int64(1)
			if e.IsArray {
				count = e.Count
			}
			regBus.Sizes.Compact += count * sizes.Compact
			regBus.Sizes.BlockAligned += count * sizes.BlockAligned
		}
	}

	if regBus.hasElement("x_uuid_x") {
		panic("x_uuid_x is reserved element name")
	}
	regBus.addStatus(x_uuid_x())

	if regBus.hasElement("x_timestamp_x") {
		panic("x_timestamp_x is reserved element name")
	}
	regBus.addStatus(x_timestamp_x())

	regBus.Sizes.BlockAligned = util.AlignToPowerOf2(
		regBus.Sizes.BlockAligned + regBus.Sizes.Own,
	)

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(&regBus, 0)

	return &regBus
}

func registerifyFunctionalities(block *Block, insElem *ins.Element, addr int64) int64 {
	if len(insElem.Elements) == 0 {
		return addr
	}

	addr = registerifyStatuses(block, insElem, addr)

	return addr
}

// Current approach is trivial. Even groups are not respected.
func registerifyStatuses(block *Block, insElem *ins.Element, addr int64) int64 {
	// Collect names instead of elements because the results must be reproducable.
	// Keys from a dictionary are returned in random order, so collect names and sort them.
	names := []string{}
	for name, ie := range insElem.Elements {
		if ie.BaseType == "status" {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		st := insElem.Elements[name]
		s := Status{
			Name:   name,
			Atomic: bool(st.Properties["atomic"].(val.Bool)),
			Width:  int64(st.Properties["width"].(val.Int)),
		}

		width := int64(st.Properties["width"].(val.Int))

		if st.IsArray {
			s.Access = makeAccessArray(st.Count, addr, width)
		} else {
			s.Access = makeAccessSingle(addr, width)
		}
		addr += s.Access.Count()

		block.addStatus(&s)
	}

	return addr
}

func registerifyBlock(insBlock *ins.Element) Sizes {
	addr := int64(0)

	b := Block{
		Name:    insBlock.Name,
		IsArray: insBlock.IsArray,
		Count:   int64(insBlock.Count),
	}

	addr = registerifyFunctionalities(&b, insBlock, addr)
	sizes := Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, e := range insBlock.Elements {
		if e.BaseType == "block" {
			s := registerifyBlock(e)
			count := int64(1)
			if e.IsArray {
				count = e.Count
			}
			sizes.Compact += count * s.Compact
			sizes.BlockAligned += count * s.BlockAligned
		}
	}

	sizes.BlockAligned = util.AlignToPowerOf2(addr + sizes.BlockAligned)

	b.Sizes = sizes

	return sizes
}