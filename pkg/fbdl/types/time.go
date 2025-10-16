package types

type Time struct {
	S  int64
	Ns int64
}

func (t Time) IsZero() bool {
	return t.S == 0 && t.Ns == 0
}
