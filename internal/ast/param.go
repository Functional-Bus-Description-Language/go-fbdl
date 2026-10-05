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
	p := Param{}

	type State int
	const (
		Name State = iota
		Ass
		Val
		Comma
	)
	state := Name

tokenLoop:
	for {
		switch state {
		case Name:
			switch t := tokens.next().(type) {
			case token.Ident:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := tokens.next().(type) {
			case token.Ass:
				state = Val
			case token.Comma:
				params = append(params, p)
				p = Param{}
				state = Name
			case token.RParen:
				params = append(params, p)
				break tokenLoop
			default:
				return nil, unexpected(t, "'=', ')' or ','")
			}
		case Val:
			expr, err := buildExpr(tokens, nil)
			if err != nil {
				return nil, err
			}
			p.Value = expr
			params = append(params, p)
			p = Param{}
			state = Comma
		case Comma:
			switch t := tokens.next().(type) {
			case token.Comma:
				state = Name
			case token.RParen:
				break tokenLoop
			default:
				return nil, unexpected(t, "',' or ')'")
			}
		}
	}

	return params, nil
}
