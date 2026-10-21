package token

type span struct {
	start int
	end   int // Exclusive end index.
}

func (s span) Start() int { return s.start }
func (s span) End() int   { return s.end }

// Text returns the text of the source span [start, end).
func (s span) Text(src []byte) string {
	return string(src[s.start:s.end])
}
