package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Arg struct represents instantiation or type argument.
// ValueFirstTok token might be useful to get argument location when Name is nil.
type Arg struct {
	Name          tok.Token // tok.Ident or nil
	Value         Expr
	ValueFirstTok tok.Token
}

func buildArgList(toks []tok.Token, ctx *context) ([]Arg, error) {
	if _, ok := toks[ctx.i].(tok.LeftParen); !ok {
		return nil, nil
	}
	if _, ok := toks[ctx.i+1].(tok.RightParen); ok {
		return nil, fmt.Errorf(
			"%s: empty argument list", tok.Loc(toks[ctx.i]),
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
		ctx.i++
		switch state {
		case Name:
			switch t := toks[ctx.i].(type) {
			case tok.Ident:
				switch toks[ctx.i+1].(type) {
				case tok.Ass:
					a.Name = t
					state = Ass
				default:
					a.Name = nil
					a.ValueFirstTok = t
					expr, err := buildExpr(toks, ctx, nil)
					if err != nil {
						return nil, err
					}
					ctx.i--
					a.Value = expr
					args = append(args, a)
					state = Comma
				}
			default:
				a.Name = nil
				a.ValueFirstTok = toks[ctx.i]
				expr, err := buildExpr(toks, ctx, nil)
				if err != nil {
					return nil, err
				}
				ctx.i--
				a.Value = expr
				args = append(args, a)
				state = Comma
			}
		case Ass:
			switch t := toks[ctx.i].(type) {
			case tok.Ass:
				state = Val
			default:
				return nil, unexpected(t, "'='")
			}
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
		case Val:
			a.ValueFirstTok = toks[ctx.i]
			expr, err := buildExpr(toks, ctx, nil)
			if err != nil {
				return nil, err
			}
			ctx.i--
			a.Value = expr
			args = append(args, a)
			state = Comma
		}
	}

	return args, nil
}
