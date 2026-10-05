package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Param represents type parameter.
type Param struct {
	Name  token.Ident
	Value Expr // Default value of the parameter
}

func buildParamList(tokens *tokenStream) ([]Param, error) {
	if _, ok := tokens.peek(0).(token.LParen); !ok {
		return nil, nil
	}
	if _, ok := tokens.peek(1).(token.RParen); ok {
		return nil, token.Error{
			Msg:  "empty parameter list",
			Toks: []token.Token{tokens.peek(0)},
		}
	}

	tokens.next()
	params := []Param{}
	for {
		t := tokens.next()
		name, ok := t.(token.Ident)
		if !ok {
			return nil, unexpected(t, "identifier")
		}
		p := Param{Name: name}

		switch t := tokens.peek(0).(type) {
		case token.Ass:
			tokens.next()
			expr, err := buildExpr(tokens, nil)
			if err != nil {
				return nil, err
			}
			p.Value = expr
		case token.Comma, token.RParen:
		default:
			tokens.next()
			return nil, unexpected(t, "'=', ')' or ','")
		}
		params = append(params, p)

		switch t := tokens.next().(type) {
		case token.Comma:
		case token.RParen:
			return params, nil
		default:
			return nil, unexpected(t, "',' or ')'")
		}
	}
}
