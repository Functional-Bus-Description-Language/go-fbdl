package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Type represents type definition.
type Type struct {
	Doc    Doc
	Name   token.Ident
	Params []Param
	Count  Expr        // If Count is not nil, then the type is a list
	Type   token.Token // Basic type, identifier or qualified identifier
	Args   ArgList
	Body   Body
}

func buildType(ctx *context) (Type, error) {
	typ := Type{}
	ctx.idx++

	// Name
	if t, ok := ctx.token().(token.Ident); ok {
		typ.Name = t
	} else {
		return typ, unexpected(ctx.token(), "identifier")
	}
	ctx.idx++

	// Parameter List
	params, err := buildParamList(ctx)
	if err != nil {
		return typ, err
	}
	typ.Params = params

	// Count
	if _, ok := ctx.token().(token.LBracket); ok {
		ctx.idx++
		expr, err := buildExpr(ctx, nil)
		if err != nil {
			return typ, err
		}
		typ.Count = expr
		if _, ok := ctx.token().(token.RBracket); !ok {
			return typ, unexpected(ctx.token(), "']'")
		}
		ctx.idx++
	}

	// Type
	switch t := ctx.token().(type) {
	case token.Functionality, token.Ident, token.QualIdent:
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
	switch t := ctx.token().(type) {
	case token.Semicolon:
		ctx.idx++
		props, err := buildPropAssignments(ctx)
		if err != nil {
			return typ, err
		}
		typ.Body.Props = props
	case token.Newline:
		if _, ok := ctx.nextTok().(token.Indent); ok {
			ctx.idx += 2
			body, err := buildBody(ctx)
			if err != nil {
				return typ, err
			}
			typ.Body = body
		}
	case token.Eof:
		// Do nothing.
	default:
		return typ, unexpected(t, "';' or newline")
	}

	return typ, nil
}
