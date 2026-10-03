package token

type position struct {
	start int
	end   int
	src   []byte
	path  string
}

func (pos position) Start() int   { return pos.start }
func (pos position) End() int     { return pos.end }
func (pos position) Src() []byte  { return pos.src }
func (pos position) Path() string { return pos.path }

func (pos position) Line() int {
	line := 1
	for _, b := range pos.src[0:pos.start] {
		if b == '\n' {
			line++
		}
	}
	return line
}

func (pos position) Column() int {
	col := 1
	for i := pos.start - 1; i >= 0; i-- {
		if pos.src[i] == '\n' {
			break
		} else {
			col++
		}
	}
	return col
}
