package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Prop struct represents functionality property.
type Prop struct {
	Name  tok.Property
	Value Expr
}

func buildPropAssignments(toks []tok.Token, ctx *context) ([]Prop, error) {
	props := []Prop{}
	p := Prop{}

	const (
		Prop = iota
		Ass
		Exp
		Semicolon
	)
	state := Prop

	// Decrement contex index as it is incremented at the beginnig of the for loop.
	ctx.i--
tokenLoop:
	for {
		ctx.i++
		switch state {
		case Prop:
			switch t := toks[ctx.i].(type) {
			case tok.Property:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "property name")
			}
		case Ass:
			switch t := toks[ctx.i].(type) {
			case tok.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "'='")
			}
		case Exp:
			expr, err := buildExpr(toks, ctx, nil)
			if err != nil {
				return nil, err
			}
			ctx.i--
			p.Value = expr
			props = append(props, p)
			state = Semicolon
		case Semicolon:
			switch t := toks[ctx.i].(type) {
			case tok.Newline, tok.Eof:
				break tokenLoop
			case tok.Semicolon:
				state = Prop
			default:
				return nil, unexpected(t, "';' or newline")
			}
		}
	}

	return props, nil
}
