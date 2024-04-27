package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Type struct represents type definition.
type Type struct {
	Doc    Doc
	Name   tok.Ident
	Params []Parameter
	Count  Expr      // If Count is not nil, then the type is a list
	Type   tok.Token // Basic type, identifier or qualified identifier
	Args   []Argument
	Body   Body
}

func buildType(toks []tok.Token, ctx *context) (Type, error) {
	typ := Type{}
	ctx.i++

	// Name
	if t, ok := toks[ctx.i].(tok.Ident); ok {
		typ.Name = t
	} else {
		return typ, unexpected(toks[ctx.i], "identifier")
	}
	ctx.i++

	// Parameter List
	params, err := buildParamList(toks, ctx)
	if err != nil {
		return typ, err
	}
	typ.Params = params

	// Count
	if _, ok := toks[ctx.i].(tok.LeftBracket); ok {
		ctx.i++
		expr, err := buildExpr(toks, ctx, nil)
		if err != nil {
			return typ, err
		}
		typ.Count = expr
		if _, ok := toks[ctx.i].(tok.RightBracket); !ok {
			return typ, unexpected(toks[ctx.i], "']'")
		}
		ctx.i++
	}

	// Type
	switch t := toks[ctx.i].(type) {
	case tok.Functionality, tok.Ident, tok.QualIdent:
		typ.Type = t
		ctx.i++
	default:
		return typ, unexpected(t, "functionality type")
	}

	// Argument List
	args, err := buildArgList(toks, ctx)
	if err != nil {
		return typ, err
	}
	typ.Args = args

	// Body
	switch t := toks[ctx.i].(type) {
	case tok.Semicolon:
		ctx.i++
		props, err := buildPropAssignments(toks, ctx)
		if err != nil {
			return typ, err
		}
		typ.Body.Props = props
	case tok.Newline:
		if _, ok := toks[ctx.i+1].(tok.Indent); ok {
			ctx.i += 2
			body, err := buildBody(toks, ctx)
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
