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

func buildInst(toks []tok.Token, ctx *context) (Inst, error) {
	inst := Inst{Name: toks[ctx.idx].(tok.Ident)}
	ctx.idx++

	// Count
	if _, ok := toks[ctx.idx].(tok.LeftBracket); ok {
		ctx.idx++
		expr, err := buildExpr(toks, ctx, nil)
		if err != nil {
			return inst, err
		}
		inst.Count = expr
		if _, ok := toks[ctx.idx].(tok.RightBracket); !ok {
			return inst, unexpected(toks[ctx.idx], "']'")
		}
		ctx.idx++
	}

	// Type
	switch t := toks[ctx.idx].(type) {
	case tok.Functionality, tok.Ident, tok.QualIdent:
		inst.Type = t
		ctx.idx++
	default:
		return inst, unexpected(t, "functionality type")
	}

	// Argument List
	argList, err := buildArgList(toks, ctx)
	if err != nil {
		return inst, err
	}
	inst.ArgList = argList

	// Body
	switch t := toks[ctx.idx].(type) {
	case tok.Semicolon:
		ctx.idx++
		props, err := buildPropAssignments(toks, ctx)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	case tok.Newline:
		if _, ok := toks[ctx.idx+1].(tok.Indent); ok {
			ctx.idx += 2
			body, err := buildBody(toks, ctx)
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
