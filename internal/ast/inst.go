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
	Args  []Arg
	Body  Body
}

func (i Inst) eq(i2 Inst) bool {
	if !i.Doc.eq(i2.Doc) ||
		i.Name != i2.Name ||
		i.Count != i2.Count ||
		i.Type != i2.Type ||
		len(i.Args) != len(i2.Args) ||
		!i.Body.eq(i2.Body) {
		return false
	}

	for n := range i.Args {
		if i.Args[n] != i2.Args[n] {
			return false
		}
	}

	return true
}

func buildInst(toks []tok.Token, c *ctx) (Inst, error) {
	inst := Inst{Name: toks[c.i].(tok.Ident)}
	c.i++

	// Count
	if _, ok := toks[c.i].(tok.LeftBracket); ok {
		c.i++
		expr, err := buildExpr(toks, c, nil)
		if err != nil {
			return inst, err
		}
		inst.Count = expr
		if _, ok := toks[c.i].(tok.RightBracket); !ok {
			return inst, unexpected(toks[c.i], "]")
		}
		c.i++
	}

	// Type
	switch t := toks[c.i].(type) {
	case tok.Functionality, tok.Ident, tok.QualIdent:
		inst.Type = t
		c.i++
	default:
		return inst, unexpected(t, "functionality type")
	}

	// Argument List
	args, err := buildArgList(toks, c)
	if err != nil {
		return inst, err
	}
	inst.Args = args

	// Body
	switch t := toks[c.i].(type) {
	case tok.Semicolon:
		c.i++
		props, err := buildPropAssignments(toks, c)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	case tok.Newline:
		if _, ok := toks[c.i+1].(tok.Indent); ok {
			c.i += 2
			body, err := buildBody(toks, c)
			if err != nil {
				return inst, err
			}
			inst.Body = body
		}
	default:
		return inst, unexpected(t, "; or newline")
	}

	return inst, nil
}
