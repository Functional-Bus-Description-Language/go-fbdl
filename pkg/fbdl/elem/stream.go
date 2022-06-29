package elem

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

func (s Stream) IsDownstream() bool {
	if len(s.Params) > 0 {
		return true
	} else if len(s.Returns) > 0 {
		return false
	}
	// Empty stream is treated as downstream.
	return true
}

func (s Stream) IsUpstream() bool { return !s.IsDownstream() }

func (s Stream) StartAddr() int64 {
	if len(s.Params) > 0 {
		return s.Params[0].Access.StartAddr()
	} else if len(s.Returns) > 0 {
		return s.Returns[0].Access.StartAddr()
	}

	// For empty stream return strobe address.
	return s.StbAddr
}

// IsEmpty returns true if stream has no params and no returns.
// Empty stream is treated as downstream.
func (s Stream) IsEmpty() bool {
	if len(s.Params) == 0 && len(s.Returns) == 0 {
		return true
	}
	return false
}
