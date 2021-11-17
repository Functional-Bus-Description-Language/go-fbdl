package fbdl

// Struct represents func element.
type Func struct {
	Name    string
	IsArray bool
	Count   int64

	// Properties
	Doc string

	Params []*Param
}

func (f *Func) StartAddr() int64 {
	return f.Params[0].Access.StartAddr()
}

func (f *Func) EndAddr() int64 {
	return f.Params[len(f.Params)-1].Access.EndAddr()
}
