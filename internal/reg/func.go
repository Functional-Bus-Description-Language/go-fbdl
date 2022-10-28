package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
)

// regFunc registerifies a Func element.
func regFunc(fun *elem.Func, addr int64) int64 {
	addr, baseBit := regFuncParams(fun, addr)
	addr = regFuncReturns(fun, addr, baseBit)

	params := fun.Params()
	returns := fun.Returns()

	if len(params) == 0 && len(returns) == 0 {
		addr += 1
	} else {
		var lastAccess access.Access
		if len(returns) > 0 {
			lastAccess = returns[len(returns)-1].Access().(access.Access)
		} else {
			lastAccess = params[len(params)-1].Access().(access.Access)
		}

		if lastAccess.EndBit() < busWidth-1 {
			addr += 1
		}
	}

	return addr
}

// regFuncParams returns address and base bit for regFuncReturns.
func regFuncParams(fun *elem.Func, addr int64) (int64, int64) {
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
		baseBit = a.EndBit() + 1
		addr += a.RegCount()

		if baseBit >= busWidth {
			baseBit = 0
		} else {
			addr -= 1
		}

		p.SetAccess(a)
	}

	if len(params) == 0 {
		fun.SetStbAddr(addr)
	} else {
		fun.SetStbAddr(params[len(params)-1].Access().(access.Access).EndAddr())
	}

	return addr, baseBit
}

func regFuncReturns(fun *elem.Func, addr, baseBit int64) int64 {
	var a access.Access

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
	}

	if len(returns) > 0 {
		fun.SetAckAddr(returns[len(returns)-1].Access().(access.Access).EndAddr())
	}

	return addr
}
