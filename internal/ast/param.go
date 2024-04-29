package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Param represents type parameter.
type Param struct {
	Name  tok.Ident
	Value Expr // Default value of the parameter
}

func buildParamList(toks []tok.Token, ctx *context) ([]Param, error) {
	if _, ok := toks[ctx.i].(tok.LeftParen); !ok {
		return nil, nil
	}
	if _, ok := toks[ctx.i+1].(tok.RightParen); ok {
		return nil, tok.Error{
			Msg:  "empty parameter list",
			Toks: []tok.Token{tok.Join(toks[ctx.i], toks[ctx.i+1])},
		}
	}

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
		ctx.i++
		switch state {
		case Name:
			switch t := toks[ctx.i].(type) {
			case tok.Ident:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[ctx.i].(type) {
			case tok.Ass:
				state = Val
			case tok.Comma:
				params = append(params, p)
				p = Param{}
				state = Name
			case tok.RightParen:
				params = append(params, p)
				ctx.i++
				break tokenLoop
			default:
				return nil, unexpected(t, "'=', ')' or ','")
			}
		case Val:
			expr, err := buildExpr(toks, ctx, nil)
			if err != nil {
				return nil, err
			}
			ctx.i--
			p.Value = expr
			params = append(params, p)
			p = Param{}
			state = Comma
		case Comma:
			switch t := toks[ctx.i].(type) {
			case tok.Comma:
				state = Name
			case tok.RightParen:
				ctx.i++
				break tokenLoop
			default:
				return nil, unexpected(t, "',' or ')'")
			}
		}
	}

	return params, nil
}
