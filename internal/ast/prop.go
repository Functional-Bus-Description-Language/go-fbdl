package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Prop represents functionality property.
type Prop struct {
	Name  tok.Property
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

	// Decrement context index as it is incremented at the beginnig of the for loop.
	ctx.idx--
tokenLoop:
	for {
		ctx.idx++
		switch state {
		case Prop:
			switch t := ctx.tok().(type) {
			case tok.Property:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "property name")
			}
		case Ass:
			switch t := ctx.tok().(type) {
			case tok.Ass:
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
			switch t := ctx.tok().(type) {
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
