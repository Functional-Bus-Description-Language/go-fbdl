package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regStream registerifies a Stream functionality.
func regStream(s *fn.Stream, addr int64) int64 {
	if len(s.Params) == 0 && len(s.Returns) == 0 {
		return regEmptyStream(s, addr)
	} else if len(s.Returns) > 0 {
		return regUpstream(s, addr)
	} else {
		return regDownstream(s, addr)
	}
}

// regEmptyStream registerifies empty stream.
// Empty stream is treated as downstream.
func regEmptyStream(s *fn.Stream, addr int64) int64 {
	s.StbAddr = addr
	return addr + 1
}

func regUpstream(s *fn.Stream, addr int64) int64 {
	var acs access.Access

	returns := s.Returns
	baseBit := int64(0)
	for _, r := range returns {
		if r.IsArray {
			acs = access.MakeArrayNRegs(r.Count, addr, baseBit, r.Width)
		} else {
			acs = access.MakeSingle(addr, baseBit, r.Width)
		}

		if acs.EndBit < busWidth-1 {
			addr += acs.RegCount - 1
			baseBit = acs.EndBit + 1
		} else {
			addr += acs.RegCount
			baseBit = 0
		}

		r.Access = acs
	}

	s.StbAddr = returns[len(returns)-1].Access.EndAddr

	lastAccess := returns[len(returns)-1].Access
	if lastAccess.EndBit < busWidth-1 {
		addr += 1
	}

	return addr
}

func regDownstream(s *fn.Stream, addr int64) int64 {
	var acs access.Access

	params := s.Params
	baseBit := int64(0)
	for _, p := range params {
		if p.IsArray {
			acs = access.MakeArrayNRegs(p.Count, addr, baseBit, p.Width)
		} else {
			acs = access.MakeSingle(addr, baseBit, p.Width)
		}

		if acs.EndBit < busWidth-1 {
			addr += acs.RegCount - 1
			baseBit = acs.EndBit + 1
		} else {
			addr += acs.RegCount
			baseBit = 0
		}

		p.Access = acs
	}

	s.StbAddr = params[len(params)-1].Access.EndAddr

	lastAccess := params[len(params)-1].Access
	if lastAccess.EndBit < busWidth-1 {
		addr += 1
	}

	return addr
}
