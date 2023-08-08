package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Prop struct represents functionality property.
type Prop struct {
	Name  token.Property
	Value Expr
}

func buildPropAssignments(toks []token.Token, c *ctx) ([]Prop, error) {
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
	c.i--
tokenLoop:
	for {
		c.i++
		switch state {
		case Prop:
			switch t := toks[c.i].(type) {
			case token.Property:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "property name")
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "=")
			}
		case Exp:
			expr, err := buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			c.i--
			p.Value = expr
			props = append(props, p)
			state = Semicolon
		case Semicolon:
			switch t := toks[c.i].(type) {
			case token.Newline, token.Eof:
				break tokenLoop
			case token.Semicolon:
				state = Prop
			default:
				return nil, unexpected(t, "; or newline")
			}
		}
	}

	return props, nil
}
