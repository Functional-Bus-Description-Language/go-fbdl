package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Comment struct {
	Comments []token.Comment
}

type Import interface {
	importNode()
}

// Import types
type (
	SingleImport struct {
		Name token.Ident
		Path token.String
	}
)

func (si SingleImport) importNode() {}

type Const interface {
	constNode()
}

// Const types
type (
	SingleConst struct {
		Name token.Ident
		Expr Expr
	}

	MultiConst struct {
		Names []token.Ident
		Exprs []Expr
	}
)

func (sc SingleConst) constNode() {}
func (mc MultiConst) constNode()  {}

type File struct {
	Comments []Comment
	Imports  []Import
	Consts   []Const
}
