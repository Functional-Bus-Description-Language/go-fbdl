package token

type span struct {
	start int
	end   int
}

func (s span) Start() int { return s.start }
func (s span) End() int   { return s.end }
