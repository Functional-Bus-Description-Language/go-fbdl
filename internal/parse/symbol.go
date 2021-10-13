package parse

type Symbol interface {
	Name() string
	LineNumber() uint32
	SetParent(s Symbol)
	Parent() Symbol
	GetSymbol(s string) (Symbol, error)
	//Parameters() []Parameter
	SetFile(f *File)
}

type base struct {
	file       *File
	lineNumber uint32
	name       string
	parent     Symbol
}

func (b base) Name() string {
	return b.name
}

func (b base) LineNumber() uint32 {
	return b.lineNumber
}

func (b base) SetParent(s Symbol) {
	b.parent = s
}

func (b base) Parent() Symbol {
	return b.parent
}

func (b *base) SetFile(f *File) {
	if b.file != nil {
		panic("should never happen")
	}

	b.file = f
}
