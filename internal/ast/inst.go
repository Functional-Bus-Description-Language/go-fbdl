package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Argument struct {
	Name  token.Ident
	Value Expr
}

type Property struct {
	Name  token.Property
	Value Expr
}

type Body struct {
	Consts []Const
	Insts  []Instantiation
	Props  []Property
}

type Instantiation struct {
	Doc   Doc
	Name  token.Ident
	Count Expr // If not nil, then it is a list
	Body  Body
	Type  token.Token // Basic type, identifier or qualified identifier
	Args  []Argument
}
