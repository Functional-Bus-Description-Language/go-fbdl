package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Inst represents functionality instantiation.
type Inst struct {
	Doc     Doc
	Name    token.Ident
	Count   Expr        // If not nil, then it is a list
	Type    token.Token // Basic type, identifier or qualified identifier
	ArgList ArgList
	Body    Body
}

func buildInst(ctx *context) (Inst, error) {
	inst := Inst{Name: ctx.token().(token.Ident)}
	ctx.idx++

	// Count
	if _, ok := ctx.token().(token.LBracket); ok {
		ctx.idx++
		expr, err := buildExpr(ctx, nil)
		if err != nil {
			return inst, err
		}
		inst.Count = expr
		if _, ok := ctx.token().(token.RBracket); !ok {
			return inst, unexpected(ctx.token(), "']'")
		}
		ctx.idx++
	}

	// Type
	switch t := ctx.token().(type) {
	case token.Functionality, token.Ident, token.QualIdent:
		inst.Type = t
		ctx.idx++
	default:
		return inst, unexpected(t, "functionality type")
	}

	// Argument List
	argList, err := buildArgList(ctx)
	if err != nil {
		return inst, err
	}
	inst.ArgList = argList

	// Body
	switch t := ctx.token().(type) {
	case token.Semicolon:
		ctx.idx++
		props, err := buildPropAssignments(ctx)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	case token.Newline:
		if _, ok := ctx.nextTok().(token.Indent); ok {
			ctx.idx += 2
			body, err := buildBody(ctx)
			if err != nil {
				return inst, err
			}
			inst.Body = body
		}
	case token.Eof:
		break
	default:
		return inst, unexpected(t, "';' or newline")
	}

	return inst, nil
}
