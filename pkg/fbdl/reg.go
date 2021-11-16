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

func registerifyFunctionalities(block *Block, insElem *ins.Element, addr int64) int64 {
	if len(insElem.Elements) == 0 {
		return addr
	}

	addr = registerifyFuncs(block, insElem, addr)
	addr = registerifyStatuses(block, insElem, addr)

	return addr
}

func registerifyFuncs(block *Block, insElem *ins.Element, addr int64) int64 {
	funcs := insElem.Elements.GetAllByBaseType("func")

	for _, f := range funcs {
		addr = registerifyFunc(block, f, addr)
	}

	return addr
}

func registerifyFunc(block *Block, insElem *ins.Element, addr int64) int64 {
	f := Func{
		Name:    insElem.Name,
		IsArray: insElem.IsArray,
		Count:   insElem.Count,
	}

	if doc, ok := insElem.Properties["doc"]; ok {
		f.Doc = string(doc.(val.Str))
	}

	block.addFunc(&f)

	params := insElem.Elements.GetAllByBaseType("param")

	baseBit := int64(0)
	for _, param := range params {
		p := Param{
			Name:    param.Name,
			IsArray: param.IsArray,
			Count:   param.Count,
			//Doc: string(param.Properties["doc"].(val.Str)),
			Width: int64(param.Properties["width"].(val.Int)),
		}

		if p.IsArray {
			panic("parram array not yet supported")
		} else {
			p.Access = makeAccessSingle(addr, baseBit, p.Width)
			as := p.Access.(*AccessSingle)
			if as.LastMask.Upper < busWidth-1 {
				addr += as.Count() - 1
				baseBit = as.LastMask.Upper + 1
			} else {
				addr += as.Count()
				baseBit = 0
			}
		}

		f.Params = append(f.Params, &p)
	}

	// If the last register is not fully occupied go to next address.
	// TODO: This is a potential place for adding a gap struct instance
	// for further address space optimization.
	lastAs := f.Params[len(f.Params)-1].Access.(*AccessSingle)
	if lastAs.LastMask.Upper < busWidth-1 {
		addr += 1
	}

	return addr
}

// Current approach is trivial. Even groups are not respected.
func registerifyStatuses(block *Block, insElem *ins.Element, addr int64) int64 {
	statuses := insElem.Elements.GetAllByBaseType("status")

	for _, st := range statuses {
		if st.Name == "x_uuid_x" || st.Name == "x_timestamp_x" {
			continue
		}

		s := Status{
			Name:   st.Name,
			Count:  insElem.Count,
			Atomic: bool(st.Properties["atomic"].(val.Bool)),
			Width:  int64(st.Properties["width"].(val.Int)),
		}

		width := int64(st.Properties["width"].(val.Int))

		if st.IsArray {
			s.Access = makeAccessArray(st.Count, addr, width)
		} else {
			s.Access = makeAccessSingle(addr, 0, width)
		}
		addr += s.Access.Count()

		block.addStatus(&s)
	}

	return addr
}

func registerifyBlock(insBlock *ins.Element) (*Block, Sizes) {
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
