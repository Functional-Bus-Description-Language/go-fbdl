package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Param struct represents type parameter node.
type Param struct {
	Name  token.Ident
	Value Expr // Default value of the parameter
}

// The Type struct represents type definition node.
type Type struct {
	Doc    Doc
	Name   token.Ident
	Params Param
	Count  Expr        // If Count is not nil, then the type is a list
	Type   token.Token // Basic type, identifier or qualified identifier
	Args   []Arg
	Body   Body
}
