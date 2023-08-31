package prs

import "fmt"

type SymbolKind uint8

const (
	ConstDef SymbolKind = iota // Constant Definition
	TypeDef                    // Type Definition
	FuncInst                   // Functionality Instantiation
)

type Symbol interface {
	Name() string
	Kind() SymbolKind
	Line() int
	Col() int
	Doc() string

	setScope(s Scope)
	Scope() Scope

	setFile(f *File)
	File() *File
}

type symbol struct {
	file  *File
	line  int
	col   int // Column of first character
	name  string
	doc   string
	scope Scope
}

func (s symbol) Name() string { return s.name }
func (s symbol) Line() int    { return s.line }
func (s symbol) Col() int     { return s.col }
func (s symbol) Doc() string  { return s.doc }
func (s symbol) Scope() Scope { return s.scope }
func (s symbol) File() *File  { return s.file }

func (sym *symbol) setScope(s Scope) {
	if sym.scope != nil {
		panic(fmt.Sprintf("resetting scope for symbol '%s'", sym.name))
	}
	sym.scope = s
}

func (s *symbol) setFile(f *File) {
	if s.file != nil {
		panic(fmt.Sprintf("resetting file for symbol '%s'", s.name))
	}
	s.file = f
}
