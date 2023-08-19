package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regStream registerifies a Stream element.
func regStream(s *fn.Stream, addr int64) int64 {
	if len(s.Params) == 0 && len(s.Returns) == 0 {
		return regEmptyStream(s, addr)
	} else if len(s.Returns) > 0 {
		return regUpstream(s, addr)
	} else {
		panic("not yet supported")
	}
}

// regEmptyStream registerifies empty stream.
// Empty stream is treated as downstream.
func regEmptyStream(s *fn.Stream, addr int64) int64 {
	s.StbAddr = addr
	return addr + 1
}

func regUpstream(s *fn.Stream, addr int64) int64 {
	var a access.Access

	returns := s.Returns
	baseBit := int64(0)
	for _, r := range returns {
		if r.IsArray {
			a = access.MakeArrayContinuous(r.Count, addr, baseBit, r.Width)
		} else {
			a = access.MakeSingle(addr, baseBit, r.Width)
		}

		if a.EndBit() < busWidth-1 {
			addr += a.RegCount() - 1
			baseBit = a.EndBit() + 1
		} else {
			addr += a.RegCount()
			baseBit = 0
		}

		r.Access = a
	}

	s.StbAddr = returns[len(returns)-1].Access.EndAddr()

	lastAccess := returns[len(returns)-1].Access
	if lastAccess.EndBit() < busWidth-1 {
		addr += 1
	}

	return addr
}
