package elem

type Stream struct {
	Elem

	Params  []*Param
	Returns []*Return

	StbAddr int64
}

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
