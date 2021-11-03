// Package reg implements code responsible for registerificaiton.
// This includes packing functionalities into registers and assigning addresses.
package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
	"log"
	"sort"
)

var busWidth uint

func Registerify(insBus *ins.Element) *BlockElement {
	if insBus == nil {
		log.Println("registerification: there is no main bus; returning nil")
		return nil
	}

	busWidth = uint(insBus.Properties["width"].(val.Int).V)

	regBus := BlockElement{
		InsElem:            insBus,
		BlockElements:      make(map[string]*BlockElement),
		FunctionalElements: make(map[string]*FunctionalElement),
	}

	// addr is current block internal access address, not global address.
	// 0 and 1 are reserved for x_uuid_x and x_timestamp_x.
	addr := uint(2)

	addr = registerifyFunctionalities(&regBus, addr)

	regBus.Sizes.Compact = addr
	regBus.Sizes.Own = addr

	for _, e := range insBus.Elements {
		if e.BaseType == "block" {
			sizes := registerifyBlock(e)
			count := uint(1)
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
	regBus.FunctionalElements["x_uuid_x"] = x_timestamp_x()

	if regBus.hasElement("x_timestamp_x") {
		panic("x_timestamp_x is reserved element name")
	}
	regBus.FunctionalElements["x_timestamp_x"] = x_timestamp_x()

	regBus.Sizes.BlockAligned = util.AlignToPowerOf2(
		regBus.Sizes.BlockAligned + regBus.Sizes.Own,
	)

	// Base address property is not yet supported, so it starts from 0.
	assignGlobalAccessAddresses(&regBus, 0)

	return &regBus
}

func registerifyFunctionalities(elem *BlockElement, addr uint) uint {
	if len(elem.InsElem.Elements) == 0 {
		return addr
	}

	addr = registerifyStatuses(elem, addr)

	return addr
}

// Current approach is trivial. Even groups are not respected.
func registerifyStatuses(elem *BlockElement, addr uint) uint {
	// Collect names instead of elements because the results must be reproducable.
	// Keys from a dictionary are returned in random order, so collect names and sort them.
	names := []string{}
	for name, ie := range elem.InsElem.Elements {
		if ie.BaseType == "status" {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		st := elem.InsElem.Elements[name]
		e := FunctionalElement{InsElem: st}

		width := uint(st.Properties["width"].(val.Int).V)

		if st.IsArray {
			e.Access = MakeAccessArray(st.Count, addr, width)
		} else {
			e.Access = MakeAccessSingle(addr, width)
		}
		addr += e.Access.Count()

		elem.FunctionalElements[st.Name] = &e
	}

	return addr
}

func registerifyBlock(block *ins.Element) Sizes {
	addr := uint(0)

	b := BlockElement{
		InsElem:            block,
		BlockElements:      make(map[string]*BlockElement),
		FunctionalElements: make(map[string]*FunctionalElement),
	}

	addr = registerifyFunctionalities(&b, addr)
	sizes := Sizes{BlockAligned: 0, Own: addr, Compact: addr}

	for _, e := range block.Elements {
		if e.BaseType == "block" {
			s := registerifyBlock(e)
			count := uint(1)
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
