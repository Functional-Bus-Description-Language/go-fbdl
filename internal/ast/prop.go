package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Prop represents functionality property.
type Prop struct {
	Name  token.Property
	Value Expr
}

func buildPropAssignments(tokens *tokenStream) ([]Prop, error) {
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

tokenLoop:
	for {
		switch state {
		case Prop:
			switch t := tokens.next().(type) {
			case token.Property:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "property name")
			}
		case Ass:
			switch t := tokens.next().(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "'='")
			}
		case Exp:
			expr, err := buildExpr(tokens, nil)
			if err != nil {
				return nil, err
			}
			p.Value = expr
			props = append(props, p)
			state = Semicolon
		case Semicolon:
			switch t := tokens.peek(0).(type) {
			case token.Newline, token.Eof:
				break tokenLoop
			case token.Semicolon:
				tokens.next()
				state = Prop
			default:
				return nil, unexpected(t, "';' or newline")
			}
		}
	}

	return props, nil
}
