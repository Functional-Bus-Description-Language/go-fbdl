package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Param struct {
	Name  token.Ident
	Value Expr // Default value of the parameter
}

type TypeDefinition struct {
	Name   token.Ident
	Count  Expr // If Count is not nil, then the type is a list
	Params Param
}
