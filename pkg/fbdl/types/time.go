package types

// The Time struct represents time type from the FBDL specification.
type Time struct {
	S  int64
	Ns int64
}

func (t Time) IsZero() bool {
	return t.S == 0 && t.Ns == 0
}
