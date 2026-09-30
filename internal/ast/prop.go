package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Prop represents functionality property.
type Prop struct {
	Name  token.Property
	Value Expr
}

func buildPropAssignments(ctx *context) ([]Prop, error) {
	props := []Prop{}
	p := Prop{}

	type State int
	const (
		Prop State = iota
		Ass
		Exp
		Semicolon
	)
	state := Prop

	// Decrement context index as it is incremented at the beginning of the for loop.
	ctx.idx--
tokenLoop:
	for {
		ctx.idx++
		switch state {
		case Prop:
			switch t := ctx.token().(type) {
			case token.Property:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "property name")
			}
		case Ass:
			switch t := ctx.token().(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "'='")
			}
		case Exp:
			expr, err := buildExpr(ctx, nil)
			if err != nil {
				return nil, err
			}
			ctx.idx--
			p.Value = expr
			props = append(props, p)
			state = Semicolon
		case Semicolon:
			switch t := ctx.token().(type) {
			case token.Newline, token.Eof:
				break tokenLoop
			case token.Semicolon:
				state = Prop
			default:
				return nil, unexpected(t, "';' or newline")
			}
		}
	}

	return props, nil
}
