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
		panic("downstream registerification not yet implemented")
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

		if acs.GetEndBit() < busWidth-1 {
			addr += acs.GetRegCount() - 1
			baseBit = acs.GetEndBit() + 1
		} else {
			addr += acs.GetRegCount()
			baseBit = 0
		}

		r.Access = acs
	}

	s.StbAddr = returns[len(returns)-1].Access.GetEndAddr()

	lastAccess := returns[len(returns)-1].Access
	if lastAccess.GetEndBit() < busWidth-1 {
		addr += 1
	}

	return addr
}
