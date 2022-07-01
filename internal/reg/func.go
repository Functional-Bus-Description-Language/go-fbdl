package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// regFunc registerifies a Func element.
func regFunc(fun *elem.Func, addr int64) int64 {
	/*
		fun := elem.Func{
			Name:    insFun.Name,
			Doc:     insFun.Doc,
			IsArray: insFun.IsArray,
			Count:   insFun.Count,
		}
	*/

	//params := insFun.Elems.GetAllByType("param")
	baseBit := int64(0)
	for _, p := range fun.Params() {
		p := p.(*elem.Param)

		if p.IsArray() {
			p.SetAccess(access.MakeArrayContinuous(p.Count(), addr, baseBit, p.Width()))
		} else {
			p.SetAccess(access.MakeSingle(addr, baseBit, p.Width()))
		}

		if p.Access().EndBit() < busWidth-1 {
			addr += p.Access().RegCount() - 1
			baseBit = p.Access().EndBit() + 1
		} else {
			addr += p.Access().RegCount()
			baseBit = 0
		}

		fun.AddParam(p)
	}

	if len(fun.Params()) == 0 {
		fun.SetStbAddr(addr)
		//addr += 1
	} else {
		fun.SetStbAddr(fun.Params()[len(fun.Params())-1].Access().EndAddr())
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

	//returns := insFun.Elems.GetAllByType("return")
	for _, r := range fun.Returns() {
		r := r.(*elem.Return)

		if r.IsArray() {
			r.SetAccess(access.MakeArrayContinuous(r.Count(), addr, baseBit, r.Width()))
		} else {
			r.SetAccess(access.MakeSingle(addr, baseBit, r.Width()))
		}

		if r.Access().EndBit() < busWidth-1 {
			addr += r.Access().RegCount() - 1
			baseBit = r.Access().EndBit() + 1
		} else {
			addr += r.Access().RegCount()
			baseBit = 0
		}

		fun.AddReturn(r)
	}

	if len(fun.Returns()) > 0 {
		fun.SetAckAddr(fun.Returns()[len(fun.Returns())-1].Access().EndAddr())
	}

	if len(fun.Params()) == 0 && len(fun.Returns()) == 0 {
		addr += 1
	} else {
		var lastAccess access.Access
		if len(fun.Returns()) > 0 {
			lastAccess = fun.Returns()[len(fun.Returns())-1].Access()
		} else {
			lastAccess = fun.Params()[len(fun.Params())-1].Access()
		}
		if lastAccess.EndBit() < busWidth-1 {
			addr += 1
		}
	}

	return addr
}
