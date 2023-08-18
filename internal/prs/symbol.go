package prs

type SymbolKind uint8

const (
	ConstDef SymbolKind = iota
	TypeDef
	ElemInst
)

type Symbol interface {
	Name() string
	Kind() SymbolKind
	Line() int
	Doc() string

	setParent(s Searchable)
	Parent() Searchable

	setFile(f *File)
	File() *File
}

type base struct {
	file   *File
	line   int
	name   string
	doc    string
	parent Searchable
}

func (b base) Name() string       { return b.name }
func (b base) Line() int          { return b.line }
func (b base) Doc() string        { return b.doc }
func (b base) Parent() Searchable { return b.parent }
func (b base) File() *File        { return b.file }

func (b *base) setParent(s Searchable) {
	if b.parent != nil {
		panic("should never happen")
	}
	b.parent = s
}

func (b *base) setFile(f *File) {
	if b.file != nil {
		panic("should never happen")
	}
	b.file = f
}
