package parser

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

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
	Tok() token.Token
	Loc() string

	setScope(s Scope)
	Scope() Scope

	setFile(f *File)
	File() *File
}

type symbol struct {
	file  *File
	token token.Token // Symbol name token
	name  string
	doc   string
	scope Scope
}

func (s symbol) Name() string     { return s.name }
func (s symbol) Line() int        { return s.token.Line() }
func (s symbol) Col() int         { return s.token.Column() }
func (s symbol) Doc() string      { return s.doc }
func (s symbol) Scope() Scope     { return s.scope }
func (s symbol) File() *File      { return s.file }
func (s symbol) Tok() token.Token { return s.token }

func (sym symbol) Loc() string {
	return fmt.Sprintf("%d:%d", sym.Line(), sym.Col())
}

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
