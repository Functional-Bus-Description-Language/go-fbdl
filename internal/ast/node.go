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

type File struct {
	Imports []Import
	//Consts ConstDecl
}
