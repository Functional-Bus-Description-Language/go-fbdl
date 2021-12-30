package prs

type Symbol interface {
	Name() string
	LineNumber() uint32
	SetParent(s Searchable)
	Parent() Searchable
	SetFile(f *File)
	File() *File
}

type base struct {
	file       *File
	lineNumber uint32
	name       string
	parent     Searchable
}

func (b base) Name() string       { return b.name }
func (b base) LineNumber() uint32 { return b.lineNumber }
func (b base) Parent() Searchable { return b.parent }
func (b base) File() *File        { return b.file }

func (b *base) SetParent(s Searchable) {
	if b.parent != nil {
		panic("should never happen")
	}

	b.parent = s
}

func (b *base) SetFile(f *File) {
	if b.file != nil {
		panic("should never happen")
	}

	b.file = f
}
