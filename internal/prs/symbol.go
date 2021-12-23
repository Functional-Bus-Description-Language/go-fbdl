package prs

type Symbol interface {
	Searchable
	Name() string
	LineNumber() uint32
	SetParent(s Symbol)
	Parent() Searchable
	SetFile(f *File)
	File() *File
	FilePath() string
}

type base struct {
	file       *File
	lineNumber uint32
	name       string
	parent     Searchable
}

func (b base) Name() string {
	return b.name
}

func (b base) LineNumber() uint32 {
	return b.lineNumber
}

func (b *base) SetParent(s Symbol) {
	b.parent = s
}

func (b base) Parent() Searchable {
	return b.parent
}

func (b *base) SetFile(f *File) {
	if b.file != nil {
		panic("should never happen")
	}

	b.file = f
}

func (b *base) File() *File {
	return b.file
}

func (b base) FilePath() string {
	if b.file != nil {
		return b.file.Path
	}

	return b.parent.(Symbol).FilePath()
}
