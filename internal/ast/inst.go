package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Inst struct represents functionality instantiation.
type Inst struct {
	Doc   Doc
	Name  tok.Ident
	Count Expr      // If not nil, then it is a list
	Type  tok.Token // Basic type, identifier or qualified identifier
	Args  []Argument
	Body  Body
}

func buildInst(toks []tok.Token, ctx *context) (Inst, error) {
	inst := Inst{Name: toks[ctx.i].(tok.Ident)}
	ctx.i++

	// Count
	if _, ok := toks[ctx.i].(tok.LeftBracket); ok {
		ctx.i++
		expr, err := buildExpr(toks, ctx, nil)
		if err != nil {
			return inst, err
		}
		inst.Count = expr
		if _, ok := toks[ctx.i].(tok.RightBracket); !ok {
			return inst, unexpected(toks[ctx.i], "']'")
		}
		ctx.i++
	}

	// Type
	switch t := toks[ctx.i].(type) {
	case tok.Functionality, tok.Ident, tok.QualIdent:
		inst.Type = t
		ctx.i++
	default:
		return inst, unexpected(t, "functionality type")
	}

	// Argument List
	args, err := buildArgList(toks, ctx)
	if err != nil {
		return inst, err
	}
	inst.Args = args

	// Body
	switch t := toks[ctx.i].(type) {
	case tok.Semicolon:
		ctx.i++
		props, err := buildPropAssignments(toks, ctx)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	case tok.Newline:
		if _, ok := toks[ctx.i+1].(tok.Indent); ok {
			ctx.i += 2
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
