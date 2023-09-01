package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Type struct represents type definition.
type Type struct {
	Doc    Doc
	Name   tok.Ident
	Params []Param
	Count  Expr      // If Count is not nil, then the type is a list
	Type   tok.Token // Basic type, identifier or qualified identifier
	Args   []Arg
	Body   Body
}

func (t Type) eq(t2 Type) bool {
	if !t.Doc.eq(t2.Doc) ||
		t.Name != t2.Name ||
		len(t.Params) != len(t2.Params) ||
		t.Count != t2.Count ||
		t.Type != t2.Type ||
		len(t.Args) != len(t2.Args) ||
		!t.Body.eq(t2.Body) {
		return false
	}

	for n := range t.Params {
		if t.Params[n] != t2.Params[n] {
			return false
		}
	}

	for n := range t.Args {
		if t.Args[n] != t2.Args[n] {
			return false
		}
	}

	return true
}

func buildType(toks []tok.Token, c *ctx) (Type, error) {
	typ := Type{}
	c.i++

	// Name
	if t, ok := toks[c.i].(tok.Ident); ok {
		typ.Name = t
	} else {
		return typ, unexpected(toks[c.i], "identifier")
	}
	c.i++

	// Parameter List
	params, err := buildParamList(toks, c)
	if err != nil {
		return typ, err
	}
	typ.Params = params

	// Count
	if _, ok := toks[c.i].(tok.LeftBracket); ok {
		c.i++
		expr, err := buildExpr(toks, c, nil)
		if err != nil {
			return typ, err
		}
		typ.Count = expr
		if _, ok := toks[c.i].(tok.RightBracket); !ok {
			return typ, unexpected(toks[c.i], "']'")
		}
		c.i++
	}

	// Type
	switch t := toks[c.i].(type) {
	case tok.Functionality, tok.Ident, tok.QualIdent:
		typ.Type = t
		c.i++
	default:
		return typ, unexpected(t, "functionality type")
	}

	// Argument List
	args, err := buildArgList(toks, c)
	if err != nil {
		return typ, err
	}
	typ.Args = args

	// Body
	switch t := toks[c.i].(type) {
	case tok.Semicolon:
		c.i++
		props, err := buildPropAssignments(toks, c)
		if err != nil {
			return typ, err
		}
		typ.Body.Props = props
	case tok.Newline:
		if _, ok := toks[c.i+1].(tok.Indent); ok {
			c.i += 2
			body, err := buildBody(toks, c)
			if err != nil {
				return typ, err
			}
			typ.Body = body
		}
	case tok.Eof:
		// Do nothing.
	default:
		return typ, unexpected(t, "';' or newline")
	}

	return typ, nil
}
