package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Func represents func element.
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

func (f *Func) ParamsStartAddr() int64 {
	if len(f.Params) == 0 {
		return f.StbAddr
	}

	return f.Params[0].Access.StartAddr()
}

// AreAllParamsSingleSingle returns true if accesses to all parameters are of type AccessSingleSingle.
func (f *Func) AreAllParamsSingleSingle() bool {
	for _, p := range f.Params {
		switch p.Access.(type) {
		case access.SingleSingle:
			continue
		default:
			return false
		}
	}
	return true
}

// regFunc registerifies a Func element.
func regFunc(insFun *ins.Element, addr int64) (*Func, int64) {
	fun := Func{
		Name:    insFun.Name,
		Doc:     insFun.Doc,
		IsArray: insFun.IsArray,
		Count:   insFun.Count,
	}

	params := insFun.Elems.GetAllByType("param")
	baseBit := int64(0)
	for _, param := range params {
		p := makeParam(param)

		if p.IsArray {
			p.Access = access.MakeArrayContinuous(p.Count, addr, baseBit, p.Width)
		} else {
			p.Access = access.MakeSingle(addr, baseBit, p.Width)
		}

		if p.Access.EndBit() < busWidth-1 {
			addr += p.Access.RegCount() - 1
			baseBit = p.Access.EndBit() + 1
		} else {
			addr += p.Access.RegCount()
			baseBit = 0
		}

		fun.Params = append(fun.Params, p)
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
		r := makeReturn(ret)

		if r.IsArray {
			r.Access = access.MakeArrayContinuous(r.Count, addr, baseBit, r.Width)
		} else {
			r.Access = access.MakeSingle(addr, baseBit, r.Width)
		}

		if r.Access.EndBit() < busWidth-1 {
			addr += r.Access.RegCount() - 1
			baseBit = r.Access.EndBit() + 1
		} else {
			addr += r.Access.RegCount()
			baseBit = 0
		}

		fun.Returns = append(fun.Returns, r)
	}

	if len(fun.Returns) > 0 {
		fun.AckAddr = fun.Returns[len(fun.Returns)-1].Access.EndAddr()
	}

	if len(fun.Params) == 0 && len(fun.Returns) == 0 {
		addr += 1
	} else {
		var lastAccess access.Access
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
