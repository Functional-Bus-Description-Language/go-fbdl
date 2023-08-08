package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Param struct represents type parameter.
type Param struct {
	Name  token.Ident
	Value Expr // Default value of the parameter
}

func buildParamList(toks []token.Token, c *ctx) ([]Param, error) {
	if _, ok := toks[c.i].(token.LeftParen); !ok {
		return nil, nil
	}
	if _, ok := toks[c.i+1].(token.RightParen); ok {
		return nil, fmt.Errorf(
			"%s: empty parameter list", token.Loc(toks[c.i]),
		)
	}

	params := []Param{}
	p := Param{}

	const (
		Name = iota
		Ass
		Val
		Comma
	)
	state := Name

tokenLoop:
	for {
		c.i++
		switch state {
		case Name:
			switch t := toks[c.i].(type) {
			case token.Ident:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case token.Ass:
				state = Val
			case token.Comma:
				params = append(params, p)
				p = Param{}
				state = Name
			case token.RightParen:
				params = append(params, p)
				c.i++
				break tokenLoop
			default:
				return nil, unexpected(t, "=, ) or ,")
			}
		case Val:
			expr, err := buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			c.i--
			p.Value = expr
			params = append(params, p)
			p = Param{}
			state = Comma
		case Comma:
			switch t := toks[c.i].(type) {
			case token.Comma:
				state = Name
			case token.RightParen:
				c.i++
				break tokenLoop
			default:
				return nil, unexpected(t, ", or )")
			}
		}
	}

	return params, nil
}
