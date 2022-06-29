package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

// regStream registerifies a Stream element.
func regStream(s *elem.Stream, addr int64) int64 {
	/*
		stream := elem.Stream{
			Name:    insStream.Name,
			Doc:     insStream.Doc,
			IsArray: insStream.IsArray,
			Count:   insStream.Count,
		}
	*/

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
func regEmptyStream(s *elem.Stream, addr int64) int64 {
	s.StbAddr = addr
	return addr + 1
}

func regUpstream(s *elem.Stream, addr int64) int64 {
	baseBit := int64(0)
	for _, ret := range s.Returns {
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
	}

	s.StbAddr = s.Returns[len(s.Returns)-1].Access.EndAddr()

	lastAccess := s.Returns[len(s.Returns)-1].Access
	if lastAccess.EndBit() < busWidth-1 {
		addr += 1
	}

	return addr
}
