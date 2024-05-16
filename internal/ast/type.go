package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Type represents type definition.
type Type struct {
	Doc    Doc
	Name   tok.Ident
	Params []Param
	Count  Expr      // If Count is not nil, then the type is a list
	Type   tok.Token // Basic type, identifier or qualified identifier
	Args   ArgList
	Body   Body
}

func buildType(ctx *context) (Type, error) {
	typ := Type{}
	ctx.idx++

	// Name
	if t, ok := ctx.tok().(tok.Ident); ok {
		typ.Name = t
	} else {
		return typ, unexpected(ctx.tok(), "identifier")
	}
	ctx.idx++

	// Parameter List
	params, err := buildParamList(ctx)
	if err != nil {
		return typ, err
	}
	typ.Params = params

	// Count
	if _, ok := ctx.tok().(tok.LBracket); ok {
		ctx.idx++
		expr, err := buildExpr(ctx, nil)
		if err != nil {
			return typ, err
		}
		typ.Count = expr
		if _, ok := ctx.tok().(tok.RBracket); !ok {
			return typ, unexpected(ctx.tok(), "']'")
		}
		ctx.idx++
	}

	// Type
	switch t := ctx.tok().(type) {
	case tok.Functionality, tok.Ident, tok.QualIdent:
		typ.Type = t
		ctx.idx++
	default:
		return typ, unexpected(t, "functionality type")
	}

	// Argument List
	args, err := buildArgList(ctx)
	if err != nil {
		return typ, err
	}
	typ.Args = args

	// Body
	switch t := ctx.tok().(type) {
	case tok.Semicolon:
		ctx.idx++
		props, err := buildPropAssignments(ctx)
		if err != nil {
			return typ, err
		}
		typ.Body.Props = props
	case tok.Newline:
		if _, ok := ctx.nextTok().(tok.Indent); ok {
			ctx.idx += 2
			body, err := buildBody(ctx)
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
