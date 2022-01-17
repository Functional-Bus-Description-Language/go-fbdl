package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Struct represents func element.
type Func struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	StbAddr int64 // Strobe address
	AckAddr int64 // Acknowledgment address

	// Properties

	Params  []*Param
	Returns []*Return
}

/*
func (f *Func) HasReturns() bool {
	if len(f.Returns) == 0 {
		return false
	}
	return true
}
*/

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
			Width:   int64(param.Props["width"].(val.Int)),
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
		fun.StbAddr = addr
		//addr += 1
	} else {
		fun.StbAddr = fun.Params[len(fun.Params)-1].Access.EndAddr()
		// If the last register is not fully occupied go to next address.
		// TODO: This is a potential place for adding a gap struct instance
		// for further address space optimization.
		/*
			lastAccess := fun.Params[len(fun.Params)-1].Access
			if lastAccess.EndBit() < busWidth-1 {
				addr += 1
			}
		*/
	}

	returns := insFun.Elems.GetAllByType("return")
	for _, ret := range returns {
		r := Return{
			Name:    ret.Name,
			Doc:     ret.Doc,
			IsArray: ret.IsArray,
			Count:   ret.Count,
			Width:   int64(ret.Props["width"].(val.Int)),
		}

		if r.IsArray {
			r.Access = makeAccessArrayContinuous(r.Count, addr, baseBit, r.Width)
		} else {
			r.Access = makeAccessSingle(addr, baseBit, r.Width)
		}

		if r.Access.EndBit() < busWidth-1 {
			addr += r.Access.RegCount() - 1
			baseBit = r.Access.EndBit() + 1
		} else {
			addr += r.Access.RegCount()
			baseBit = 0
		}

		fun.Returns = append(fun.Returns, &r)
	}

	if len(fun.Returns) > 0 {
		fun.AckAddr = fun.Returns[len(fun.Returns)-1].Access.EndAddr()
	}

	if len(fun.Params) == 0 && len(fun.Returns) == 0 {
		addr += 1
	} else {
		var lastAccess Access
		if len(fun.Returns) > 0 {
			lastAccess = fun.Returns[len(fun.Returns)-1].Access
		} else {
			lastAccess = fun.Params[len(fun.Params)-1].Access
		}
		if lastAccess.EndBit() < busWidth-1 {
			addr += 1
		}
	}

	return &fun, addr
}
