package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Struct represents func element.
type Func struct {
	Name     string
	Doc      string
	IsArray  bool
	Count    int64
	CallAddr int64

	// Properties

	Params []*Param
}

func (f *Func) StartAddr() int64 {
	if len(f.Params) == 0 {
		return f.CallAddr
	} else {
		return f.Params[0].Access.StartAddr()
	}
}

func (f *Func) EndAddr() int64 {
	return f.CallAddr
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

func registerifyFunc(insFun *ins.Element, addr int64) (*Func, int64) {
	fun := Func{
		Name:    insFun.Name,
		Doc:     insFun.Doc,
		IsArray: insFun.IsArray,
		Count:   insFun.Count,
	}

	if doc, ok := insFun.Props["doc"]; ok {
		fun.Doc = string(doc.(val.Str))
	}

	params := insFun.Elems.GetAllByType("param")

	baseBit := int64(0)
	for _, param := range params {
		p := Param{
			Name:    param.Name,
			Doc:     param.Doc,
			IsArray: param.IsArray,
			Count:   param.Count,
			//Doc: string(param.Props["doc"].(val.Str)),
			Width: int64(param.Props["width"].(val.Int)),
		}

		if p.IsArray {
			p.Access = makeAccessArrayContinuous(p.Count, addr, baseBit, p.Width)
		} else {
			p.Access = makeAccessSingle(addr, baseBit, p.Width)
		}

		if p.Access.EndBit() < busWidth-1 {
			addr += p.Access.RegCount() - 1
			baseBit = p.Access.EndBit() + 1
		} else {
			addr += p.Access.RegCount()
			baseBit = 0
		}

		fun.Params = append(fun.Params, &p)
	}

	if len(fun.Params) == 0 {
		fun.CallAddr = addr
		addr += 1
	} else {
		fun.CallAddr = fun.Params[len(fun.Params)-1].Access.EndAddr()
		// If the last register is not fully occupied go to next address.
		// TODO: This is a potential place for adding a gap struct instance
		// for further address space optimization.
		lastAccess := fun.Params[len(fun.Params)-1].Access
		if lastAccess.EndBit() < busWidth-1 {
			addr += 1
		}
	}

	return &fun, addr
}
