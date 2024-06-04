package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regProc registerifies a Proc functionality.
func regProc(proc *fn.Proc, addr int64) int64 {
	var acs access.Access

	params := proc.Params
	baseBit := int64(0)
	for _, p := range params {
		if p.IsArray {
			acs = access.MakeArrayNRegs(p.Count, addr, baseBit, p.Width)
		} else {
			acs = access.MakeSingle(addr, baseBit, p.Width)
		}

		if acs.GetEndBit() < busWidth-1 {
			addr += acs.GetRegCount() - 1
			baseBit = acs.GetEndBit() + 1
		} else {
			addr += acs.GetRegCount()
			baseBit = 0
		}

		p.Access = acs
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
			acs = access.MakeArrayNRegs(r.Count, addr, baseBit, r.Width)
		} else {
			acs = access.MakeSingle(addr, baseBit, r.Width)
		}

		if acs.GetEndBit() < busWidth-1 {
			addr += acs.GetRegCount() - 1
			baseBit = acs.GetEndBit() + 1
		} else {
			addr += acs.GetRegCount()
			baseBit = 0
		}

		r.Access = acs
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
		if lastAccess.GetEndBit() < busWidth-1 {
			addr += 1
		}
	}

	return addr
}
