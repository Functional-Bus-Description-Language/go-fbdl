package tok

type position struct {
	start  int
	end    int
	line   int
	column int
	src    []byte
	path   string
}

func (pos position) Start() int   { return pos.start }
func (pos position) End() int     { return pos.end }
func (pos position) Line() int    { return pos.line }
func (pos position) Column() int  { return pos.column }
func (pos position) Src() []byte  { return pos.src }
func (pos position) Path() string { return pos.path }
