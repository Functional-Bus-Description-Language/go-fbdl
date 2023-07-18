package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Import interface {
	importNode()
}

type SingleImport struct {
	Import token.Token
	Name   token.Token
	Path   token.Token
}

func (si SingleImport) importNode() {}

type Comment struct {
	Comments []token.Token
}

func (c Comment) add(t token.Token) {
	c.Comments = append(c.Comments, t)
}

type File struct {
	Comments []Comment
	Imports  []Import
	//Consts ConstDecl
}
