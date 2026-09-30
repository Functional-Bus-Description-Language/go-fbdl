package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Param represents type parameter.
type Param struct {
	Name  token.Ident
	Value Expr // Default value of the parameter
}

func buildParamList(ctx *context) ([]Param, error) {
	if _, ok := ctx.token().(token.LParen); !ok {
		return nil, nil
	}
	if _, ok := ctx.nextTok().(token.RParen); ok {
		return nil, token.Error{
			Msg:  "empty parameter list",
			Toks: []token.Token{token.Join(ctx.token(), ctx.nextTok())},
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
		ctx.idx++
		switch state {
		case Name:
			switch t := ctx.token().(type) {
			case token.Ident:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := ctx.token().(type) {
			case token.Ass:
				state = Val
			case token.Comma:
				params = append(params, p)
				p = Param{}
				state = Name
			case token.RParen:
				params = append(params, p)
				ctx.idx++
				break tokenLoop
			default:
				return nil, unexpected(t, "'=', ')' or ','")
			}
		case Val:
			expr, err := buildExpr(ctx, nil)
			if err != nil {
				return nil, err
			}
			ctx.idx--
			p.Value = expr
			params = append(params, p)
			p = Param{}
			state = Comma
		case Comma:
			switch t := ctx.token().(type) {
			case token.Comma:
				state = Name
			case token.RParen:
				ctx.idx++
				break tokenLoop
			default:
				return nil, unexpected(t, "',' or ')'")
			}
		}
	}

	return params, nil
}
