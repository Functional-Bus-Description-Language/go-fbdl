package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

// regStream registerifies a Stream element.
func regStream(insStream *ins.Element, addr int64) (*elem.Stream, int64) {
	stream := elem.Stream{
		Name:    insStream.Name,
		Doc:     insStream.Doc,
		IsArray: insStream.IsArray,
		Count:   insStream.Count,
	}

	params := insStream.Elems.GetAllByType("param")
	returns := insStream.Elems.GetAllByType("return")

	if len(params) == 0 && len(returns) == 0 {
		return regEmptyStream(&stream, addr)
	} else if len(returns) > 0 {
		return regUpstream(&stream, addr, returns)
	} else {
		panic("not yet supported")
	}
}

// regEmptyStream registerifies empty stream.
// Empty stream is treated as downstream.
func regEmptyStream(stream *elem.Stream, addr int64) (*elem.Stream, int64) {
	stream.StbAddr = addr
	return stream, addr + 1
}

func regUpstream(stream *elem.Stream, addr int64, returns []*ins.Element) (*elem.Stream, int64) {
	baseBit := int64(0)
	for _, ret := range returns {
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

		stream.Returns = append(stream.Returns, r)
	}

	stream.StbAddr = stream.Returns[len(stream.Returns)-1].Access.EndAddr()

	lastAccess := stream.Returns[len(stream.Returns)-1].Access
	if lastAccess.EndBit() < busWidth-1 {
		addr += 1
	}

	return stream, addr
}
