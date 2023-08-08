package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Type struct represents type definition.
type Type struct {
	Doc    Doc
	Name   token.Ident
	Params []Param
	Count  Expr        // If Count is not nil, then the type is a list
	Type   token.Token // Basic type, identifier or qualified identifier
	Args   []Arg
	Body   Body
}

func buildType(toks []token.Token, c *ctx) (Type, error) {
	typ := Type{}
	c.i++

	if t, ok := toks[c.i].(token.Ident); ok {
		typ.Name = t
	} else {
		return typ, unexpected(toks[c.i], "identifier")
	}
	c.i++

	params, err := buildParamList(toks, c)
	if err != nil {
		return typ, err
	}
	typ.Params = params

	if _, ok := toks[c.i].(token.LeftBracket); ok {
		c.i++
		expr, err := buildExpr(toks, c, nil)
		if err != nil {
			return typ, err
		}
		typ.Count = expr
		if _, ok := toks[c.i].(token.RightBracket); !ok {
			return typ, unexpected(toks[c.i], "]")
		}
		c.i++
	}

	return typ, nil
}
