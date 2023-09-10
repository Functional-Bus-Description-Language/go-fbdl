package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regProc registerifies a Proc functionality.
func regProc(proc *fn.Proc, addr int64) int64 {
	var a access.Access

	params := proc.Params
	baseBit := int64(0)
	for _, p := range params {
		if p.IsArray {
			a = access.MakeArrayContinuous(p.Count, addr, baseBit, p.Width)
		} else {
			a = access.MakeSingle(addr, baseBit, p.Width)
		}

		if a.EndBit() < busWidth-1 {
			addr += a.GetRegCount() - 1
			baseBit = a.EndBit() + 1
		} else {
			addr += a.GetRegCount()
			baseBit = 0
		}

		p.Access = a
	}

	if len(params) > 0 {
		callAddr := params[len(params)-1].Access.GetEndAddr()
		proc.CallAddr = &callAddr
	} else if len(proc.Returns) > 0 {
		if proc.Delay != nil {
			callAddr := addr
			proc.CallAddr = &callAddr
		}
	} else {
		callAddr := addr
		proc.CallAddr = &callAddr
	}

	returns := proc.Returns
	for _, r := range returns {
		if r.IsArray {
			a = access.MakeArrayContinuous(r.Count, addr, baseBit, r.Width)
		} else {
			a = access.MakeSingle(addr, baseBit, r.Width)
		}

		if a.EndBit() < busWidth-1 {
			addr += a.GetRegCount() - 1
			baseBit = a.EndBit() + 1
		} else {
			addr += a.GetRegCount()
			baseBit = 0
		}

		r.Access = a
	}

	if len(returns) > 0 {
		exitAddr := returns[len(returns)-1].Access.GetEndAddr()
		proc.ExitAddr = &exitAddr
	} else {
		if proc.Delay != nil {
			exitAddr := *proc.CallAddr
			proc.ExitAddr = &exitAddr
		}
	}

	if len(params) == 0 && len(returns) == 0 {
		addr += 1
	} else {
		var lastAccess access.Access
		if len(returns) > 0 {
			lastAccess = returns[len(returns)-1].Access
		} else {
			lastAccess = params[len(params)-1].Access
		}
		if lastAccess.EndBit() < busWidth-1 {
			addr += 1
		}
	}

	return addr
}
