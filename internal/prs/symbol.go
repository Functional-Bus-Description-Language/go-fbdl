package prs

type Symbol interface {
	Name() string
	LineNum() uint32
	Doc() string
	SetDoc(c comment)
	SetParent(s Searchable)
	Parent() Searchable
	SetFile(f *File)
	File() *File
}

type base struct {
	file    *File
	lineNum uint32
	name    string
	doc     string
	parent  Searchable
}

func (b base) Name() string       { return b.name }
func (b base) LineNum() uint32    { return b.lineNum }
func (b base) Doc() string        { return b.doc }
func (b base) Parent() Searchable { return b.parent }
func (b base) File() *File        { return b.file }

func (b *base) SetDoc(c comment) {
	if b.doc != "" {
		panic("should never happen")
	}

	b.doc = c.msg
}

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
