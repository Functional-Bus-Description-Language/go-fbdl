package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Comment struct {
	Comments []token.Token
}

type Import interface {
	importNode()
}

// Import types
type (
	SingleImport struct {
		Import token.Token
		Name   token.Token
		Path   token.Token
	}
)

func (si SingleImport) importNode() {}

type Const interface {
	constNode()
}

// Const types
type (
	SingleConst struct {
		Const token.Token
		Name  token.Token
		Ass   token.Token
		Expr  Expr
	}
)

func (sc SingleConst) constNode() {}

type File struct {
	Comments []Comment
	Imports  []Import
	Consts   []Const
}
