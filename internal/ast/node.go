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
		Import token.Import
		Name   token.Ident
		Path   token.String
	}
)

func (si SingleImport) importNode() {}

type Const interface {
	constNode()
}

// Const types
type (
	SingleConst struct {
		Const token.Const
		Name  token.Ident
		Ass   token.Ass
		Expr  Expr
	}
)

func (sc SingleConst) constNode() {}

type File struct {
	Comments []Comment
	Imports  []Import
	Consts   []Const
}
