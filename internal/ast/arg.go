package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Arg struct represents instantiation or type argument.
type Arg struct {
	Name  token.Token // token.Ident or nil
	Value Expr
}

func buildArgList(toks []token.Token, c *ctx) ([]Arg, error) {
	if _, ok := toks[c.i].(token.LeftParen); !ok {
		return nil, nil
	}
	if _, ok := toks[c.i+1].(token.RightParen); ok {
		return nil, fmt.Errorf(
			"%s: empty argument list", token.Loc(toks[c.i]),
		)
	}

	args := []Arg{}
	a := Arg{}

	const (
		Name = iota
		Ass
		Comma
		Exp
	)
	state := Name

tokenLoop:
	for {
		c.i++
		switch state {
		case Name:
			switch t := toks[c.i].(type) {
			case token.Ident:
				a.Name = t
				state = Ass
			default:
				a.Name = nil
				expr, err := buildExpr(toks, c, nil)
				if err != nil {
					return nil, err
				}
				c.i--
				a.Value = expr
				args = append(args, a)
				state = Comma
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "=")
			}
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
		case Exp:
			expr, err := buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			c.i--
			a.Value = expr
			args = append(args, a)
			state = Comma
		}
	}

	return args, nil
}
