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
	for {
		t := tokens.next()
		name, ok := t.(token.Property)
		if !ok {
			return nil, unexpected(t, "property name")
		}

		t = tokens.next()
		if _, ok := t.(token.Ass); !ok {
			return nil, unexpected(t, "'='")
		}

		expr, err := buildExpr(tokens, nil)
		if err != nil {
			return nil, err
		}
		props = append(props, Prop{Name: name, Value: expr})

		switch t := tokens.peek(0).(type) {
		case token.Newline, token.Eof:
			return props, nil
		case token.Semicolon:
			tokens.next()
		default:
			return nil, unexpected(t, "';' or newline")
		}
	}
}
