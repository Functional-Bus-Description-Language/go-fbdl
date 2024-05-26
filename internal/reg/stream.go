package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regStream registerifies a Stream functionality.
func regStream(s *fn.Stream, addr *address) {
	if len(s.Params) == 0 && len(s.Returns) == 0 {
		regEmptyStream(s, addr)
	} else if len(s.Returns) > 0 {
		regUpstream(s, addr)
	} else {
		regDownstream(s, addr)
	}
}

// regEmptyStream registerifies empty stream.
// Empty stream is treated as downstream.
func regEmptyStream(s *fn.Stream, addr *address) {
	s.StbAddr = addr.value
	addr.inc(1)
}

func regUpstream(s *fn.Stream, addr *address) {
	var acs access.Access

	returns := s.Returns
	baseBit := int64(0)
	for _, r := range returns {
		if r.IsArray {
			acs = access.MakeArrayNRegs(r.Count, addr.value, baseBit, r.Width)
		} else {
			acs = access.MakeSingle(addr.value, baseBit, r.Width)
		}

		if acs.GetEndBit() < busWidth-1 {
			addr.inc(acs.GetRegCount() - 1)
			baseBit = acs.GetEndBit() + 1
		} else {
			addr.inc(acs.GetRegCount())
			baseBit = 0
		}

		r.Access = acs
	}

	s.StbAddr = returns[len(returns)-1].Access.GetEndAddr()

	lastAccess := returns[len(returns)-1].Access
	if lastAccess.GetEndBit() < busWidth-1 {
		addr.inc(1)
	}
}

func regDownstream(s *fn.Stream, addr *address) {
	var acs access.Access

	params := s.Params
	baseBit := int64(0)
	for _, p := range params {
		if p.IsArray {
			acs = access.MakeArrayNRegs(p.Count, addr.value, baseBit, p.Width)
		} else {
			acs = access.MakeSingle(addr.value, baseBit, p.Width)
		}

		if acs.GetEndBit() < busWidth-1 {
			addr.inc(acs.GetRegCount() - 1)
			baseBit = acs.GetEndBit() + 1
		} else {
			addr.inc(acs.GetRegCount())
			baseBit = 0
		}

		p.Access = acs
	}

	s.StbAddr = params[len(params)-1].Access.GetEndAddr()

	lastAccess := params[len(params)-1].Access
	if lastAccess.GetEndBit() < busWidth-1 {
		addr.inc(1)
	}
}
