package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Arg struct represents instantiation or type argument.
type Arg struct {
	Name  tok.Token // tok.Ident or nil
	Value Expr
}

func buildArgList(toks []tok.Token, c *ctx) ([]Arg, error) {
	if _, ok := toks[c.i].(tok.LeftParen); !ok {
		return nil, nil
	}
	if _, ok := toks[c.i+1].(tok.RightParen); ok {
		return nil, fmt.Errorf(
			"%s: empty argument list", tok.Loc(toks[c.i]),
		)
	}

	args := []Arg{}
	a := Arg{}

	const (
		Name = iota
		Ass
		Comma
		Val
	)
	state := Name

tokenLoop:
	for {
		c.i++
		switch state {
		case Name:
			switch t := toks[c.i].(type) {
			case tok.Ident:
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
			case tok.Ass:
				state = Val
			default:
				return nil, unexpected(t, "=")
			}
		case Comma:
			switch t := toks[c.i].(type) {
			case tok.Comma:
				state = Name
			case tok.RightParen:
				c.i++
				break tokenLoop
			default:
				return nil, unexpected(t, ", or )")
			}
		case Val:
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
