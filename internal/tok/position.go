package tok

type position struct {
	start  int
	end    int
	line   int
	column int
}

func (pos position) Start() int  { return pos.start }
func (pos position) End() int    { return pos.end }
func (pos position) Line() int   { return pos.line }
func (pos position) Column() int { return pos.column }
