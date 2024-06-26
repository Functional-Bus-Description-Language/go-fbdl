package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Param represents type parameter.
type Param struct {
	Name  tok.Ident
	Value Expr // Default value of the parameter
}

func buildParamList(ctx *context) ([]Param, error) {
	if _, ok := ctx.tok().(tok.LParen); !ok {
		return nil, nil
	}
	if _, ok := ctx.nextTok().(tok.RParen); ok {
		return nil, tok.Error{
			Msg:  "empty parameter list",
			Toks: []tok.Token{tok.Join(ctx.tok(), ctx.nextTok())},
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
			switch t := ctx.tok().(type) {
			case tok.Ident:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := ctx.tok().(type) {
			case tok.Ass:
				state = Val
			case tok.Comma:
				params = append(params, p)
				p = Param{}
				state = Name
			case tok.RParen:
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
			switch t := ctx.tok().(type) {
			case tok.Comma:
				state = Name
			case tok.RParen:
				ctx.idx++
				break tokenLoop
			default:
				return nil, unexpected(t, "',' or ')'")
			}
		}
	}

	return params, nil
}
