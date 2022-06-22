package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Stream represents stream element.
// Stream with params (or empty stream) is a downstream.
// Stream with returns is an upstream.
type Stream struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	StbAddr int64 // Strobe address

	// Properties
	// Currently stream has no properties.

	Params  []*Param
	Returns []*Return
}

func (s Stream) IsUpstream() bool { return !s.IsDownstream() }

func (s Stream) IsDownstream() bool {
	if len(s.Params) > 0 {
		return true
	} else if len(s.Returns) > 0 {
		return false
	}
	// Empty stream is treated as downstream.
	return true
}

// IsEmpty returns true if stream has no params and no returns.
// Empty stream is treated as downstream.
func (s Stream) IsEmpty() bool {
	if len(s.Params) == 0 && len(s.Returns) == 0 {
		return true
	}
	return false
}

// regStream registerifies a Stream element.
func regStream(insStream *ins.Element, addr int64) (*Stream, int64) {
	stream := Stream{
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
func regEmptyStream(stream *Stream, addr int64) (*Stream, int64) {
	stream.StbAddr = addr
	return stream, addr + 1
}

func regUpstream(stream *Stream, addr int64, returns []*ins.Element) (*Stream, int64) {
	baseBit := int64(0)
	for _, ret := range returns {
		r := Return{
			Name:    ret.Name,
			Doc:     ret.Doc,
			IsArray: ret.IsArray,
			Count:   ret.Count,
			Width:   int64(ret.Props["width"].(val.Int)),
		}

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

		stream.Returns = append(stream.Returns, &r)
	}

	stream.StbAddr = stream.Returns[len(stream.Returns)-1].Access.EndAddr()

	lastAccess := stream.Returns[len(stream.Returns)-1].Access
	if lastAccess.EndBit() < busWidth-1 {
		addr += 1
	}

	return stream, addr
}
