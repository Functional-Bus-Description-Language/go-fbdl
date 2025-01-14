package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Stream struct {
	Func

	Delay *val.Time

	Params  []*Param
	Returns []*Return

	StbAddr int64
}

func (s Stream) Type() string { return "stream" }

// IsDownstream returns true if Stream has only params or Stream has not params and no returns.
// Empty stream is treated as a downstream.
func (s *Stream) IsDownstream() bool {
	if len(s.Params) > 0 {
		return true
	} else if len(s.Returns) > 0 {
		return false
	}
	return true
}

// IsUpstream returns true if Stream has only returns.
func (s *Stream) IsUpstream() bool {
	return !s.IsDownstream()
}

func (s *Stream) StartAddr() int64 {
	if len(s.Params) > 0 {
		return s.Params[0].Access.StartAddr()
	} else if len(s.Returns) > 0 {
		return s.Returns[0].Access.StartAddr()
	}

	// For empty stream return strobe address.
	return s.StbAddr
}
