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

type symbol struct {
	file   *File
	line   int
	name   string
	doc    string
	parent Searchable
}

func (s symbol) Name() string       { return s.name }
func (s symbol) Line() int          { return s.line }
func (s symbol) Doc() string        { return s.doc }
func (s symbol) Parent() Searchable { return s.parent }
func (s symbol) File() *File        { return s.file }

func (s *symbol) setParent(p Searchable) {
	if s.parent != nil {
		panic("should never happen")
	}
	s.parent = p
}

func (s *symbol) setFile(f *File) {
	if s.file != nil {
		panic("should never happen")
	}
	s.file = f
}
