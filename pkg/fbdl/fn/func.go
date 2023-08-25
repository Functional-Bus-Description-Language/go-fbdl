package fn

type Functionality interface {
	isFunctionality()
	GetName() string
	Type() string
}

type Func struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
}

func (f Func) isFunctionality() {}
func (f Func) GetName() string  { return f.Name }
