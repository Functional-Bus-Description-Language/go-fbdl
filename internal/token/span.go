package token

type span struct {
	start int
	end   int
}

func (s span) Start() int { return s.start }
func (s span) End() int   { return s.end }

// Text returns the text of the inclusive source span.
func (s span) Text(src []byte) string {
	return string(src[s.start : s.end+1])
}
