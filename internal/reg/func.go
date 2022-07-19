package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// regFunc registerifies a Func element.
func regFunc(fun *elem.Func, addr int64) int64 {
	var a access.Access

	params := fun.Params()
	baseBit := int64(0)
	for _, p := range params {
		p := p.(*elem.Param)

		if p.IsArray() {
			a = access.MakeArrayContinuous(p.Count(), addr, baseBit, p.Width())
		} else {
			a = access.MakeSingle(addr, baseBit, p.Width())
		}

		if a.EndBit() < busWidth-1 {
			addr += a.RegCount() - 1
			baseBit = a.EndBit() + 1
		} else {
			addr += a.RegCount()
			baseBit = 0
		}

		p.SetAccess(a)
	}

	if len(params) == 0 {
		fun.SetStbAddr(addr)
		//addr += 1
	} else {
		fun.SetStbAddr(params[len(params)-1].Access().EndAddr())
		// If the last register is not fully occupied go to next address.
		// TODO: This is a potential place for adding a gap struct instance
		// for further address space optimization.
		/*
			lastAccess := params[len(params)-1].Access
			if lastAccess.EndBit() < busWidth-1 {
				addr += 1
			}
		*/
	}

	returns := fun.Returns()
	for _, r := range returns {
		r := r.(*elem.Return)

		if r.IsArray() {
			a = access.MakeArrayContinuous(r.Count(), addr, baseBit, r.Width())
		} else {
			a = access.MakeSingle(addr, baseBit, r.Width())
		}

		if a.EndBit() < busWidth-1 {
			addr += a.RegCount() - 1
			baseBit = a.EndBit() + 1
		} else {
			addr += a.RegCount()
			baseBit = 0
		}

		r.SetAccess(a)
		fun.AddReturn(r)
	}

	if len(returns) > 0 {
		fun.SetAckAddr(returns[len(returns)-1].Access().EndAddr())
	}

	if len(params) == 0 && len(returns) == 0 {
		addr += 1
	} else {
		var lastAccess access.Access
		if len(returns) > 0 {
			lastAccess = returns[len(returns)-1].Access()
		} else {
			lastAccess = params[len(params)-1].Access()
		}
		if lastAccess.EndBit() < busWidth-1 {
			addr += 1
		}
	}

	return addr
}
