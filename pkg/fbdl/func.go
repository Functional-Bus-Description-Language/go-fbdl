package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Struct represents func element.
type Func struct {
	Name    string
	IsArray bool
	Count   int64

	// Properties
	Doc string

	Params []*Param
}

func (f *Func) StartAddr() int64 {
	return f.Params[0].Access.StartAddr()
}

func (f *Func) EndAddr() int64 {
	return f.Params[len(f.Params)-1].Access.EndAddr()
}

// AreAllParamsSingleSingle returns true if accesses to all parameters are of type AccessSingleSingle.
func (f *Func) AreAllParamsSingleSingle() bool {
	for _, p := range f.Params {
		switch p.Access.(type) {
		case AccessSingleSingle:
			continue
		default:
			return false
		}
	}
	return true
}

func registerifyFunc(blk *Block, insElem *ins.Element, addr int64) int64 {
	f := Func{
		Name:    insElem.Name,
		IsArray: insElem.IsArray,
		Count:   insElem.Count,
	}

	if doc, ok := insElem.Properties["doc"]; ok {
		f.Doc = string(doc.(val.Str))
	}

	blk.addFunc(&f)

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
			p.Access = makeAccessArrayContinuous(p.Count, addr, baseBit, p.Width)
		} else {
			p.Access = makeAccessSingle(addr, baseBit, p.Width)
		}

		if p.Access.EndBit() < busWidth-1 {
			addr += p.Access.Count() - 1
			baseBit = p.Access.EndBit() + 1
		} else {
			addr += p.Access.Count()
			baseBit = 0
		}

		f.Params = append(f.Params, &p)
	}

	// If the last register is not fully occupied go to next address.
	// TODO: This is a potential place for adding a gap struct instance
	// for further address space optimization.
	lastAccess := f.Params[len(f.Params)-1].Access
	if lastAccess.EndBit() < busWidth-1 {
		addr += 1
	}

	return addr
}
