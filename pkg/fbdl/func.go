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
