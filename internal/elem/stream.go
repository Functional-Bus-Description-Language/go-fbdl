package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	fbdl "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

type stream struct {
	Elem

	// Properties
	// Currently stream has no properties.

	Params  []fbdl.Param
	Returns []fbdl.Return

	StbAddr int64 // Strobe address
}

// Stream represents stream element.
// Stream with params (or empty stream) is a downstream.
// Stream with returns is an upstream.
type Stream struct {
	stream
}

func (s *Stream) Type() string { return "stream" }

func (s *Stream) SetStbAddr(a int64) { s.stream.StbAddr = a }
func (s *Stream) StbAddr() int64     { return s.stream.StbAddr }

func (s *Stream) AddParam(p *Param)    { s.stream.Params = append(s.stream.Params, p) }
func (s *Stream) Params() []fbdl.Param { return s.stream.Params }

func (s *Stream) AddReturn(r *Return)    { s.stream.Returns = append(s.stream.Returns, r) }
func (s *Stream) Returns() []fbdl.Return { return s.stream.Returns }

func (s *Stream) HasElement(name string) bool {
	for i := range s.stream.Params {
		if s.stream.Params[i].Name() == name {
			return true
		}
	}
	for i := range s.stream.Returns {
		if s.stream.Returns[i].Name() == name {
			return true
		}
	}
	return false
}

func (s *Stream) IsDownstream() bool {
	if len(s.stream.Params) > 0 {
		return true
	} else if len(s.stream.Returns) > 0 {
		return false
	}
	// Empty stream is treated as downstream.
	return true
}

func (s *Stream) IsUpstream() bool { return !s.IsDownstream() }

func (s *Stream) StartAddr() int64 {
	if len(s.stream.Params) > 0 {
		return s.stream.Params[0].Access().StartAddr()
	} else if len(s.stream.Returns) > 0 {
		return s.stream.Returns[0].Access().StartAddr()
	}

	// For empty stream return strobe address.
	return s.stream.StbAddr
}

// IsEmpty returns true if stream has no params and no returns.
// Empty stream is treated as downstream.
func (s *Stream) IsEmpty() bool {
	if len(s.stream.Params) == 0 && len(s.stream.Returns) == 0 {
		return true
	}
	return false
}

func (s *Stream) Hash() uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(s.Elem.Hash())

	// Params
	for _, p := range s.Params() {
		write(p.Hash())
	}

	// Returns
	for _, r := range s.Returns() {
		write(r.Hash())
	}

	// StbAddr
	write(s.StbAddr())

	return adler32.Checksum(buf.Bytes())
}
