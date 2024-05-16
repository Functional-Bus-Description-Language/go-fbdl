package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Inst represents functionality instantiation.
type Inst struct {
	Doc     Doc
	Name    tok.Ident
	Count   Expr      // If not nil, then it is a list
	Type    tok.Token // Basic type, identifier or qualified identifier
	ArgList ArgList
	Body    Body
}

func buildInst(ctx *context) (Inst, error) {
	inst := Inst{Name: ctx.tok().(tok.Ident)}
	ctx.idx++

	// Count
	if _, ok := ctx.tok().(tok.LBracket); ok {
		ctx.idx++
		expr, err := buildExpr(ctx, nil)
		if err != nil {
			return inst, err
		}
		inst.Count = expr
		if _, ok := ctx.tok().(tok.RBracket); !ok {
			return inst, unexpected(ctx.tok(), "']'")
		}
		ctx.idx++
	}

	// Type
	switch t := ctx.tok().(type) {
	case tok.Functionality, tok.Ident, tok.QualIdent:
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
	switch t := ctx.tok().(type) {
	case tok.Semicolon:
		ctx.idx++
		props, err := buildPropAssignments(ctx)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	case tok.Newline:
		if _, ok := ctx.nextTok().(tok.Indent); ok {
			ctx.idx += 2
			body, err := buildBody(ctx)
			if err != nil {
				return inst, err
			}
			inst.Body = body
		}
	case tok.Eof:
		break
	default:
		return inst, unexpected(t, "';' or newline")
	}

	return inst, nil
}
