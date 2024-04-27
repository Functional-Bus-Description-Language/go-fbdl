package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Property struct represents functionality property.
type Property struct {
	Name  tok.Property
	Value Expr
}

func buildPropAssignments(toks []tok.Token, ctx *context) ([]Property, error) {
	props := []Property{}
	prop := Property{}

	type State int
	const (
		Prop State = iota
		Ass
		Exp
		Semicolon
	)
	state := Prop

	// Decrement context index as it is incremented at the beginnig of the for loop.
	ctx.i--
tokenLoop:
	for {
		ctx.i++
		switch state {
		case Prop:
			switch t := toks[ctx.i].(type) {
			case tok.Property:
				prop.Name = t
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
			prop.Value = expr
			props = append(props, prop)
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
