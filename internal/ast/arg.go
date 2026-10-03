package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Arg represents instantiation or type argument.
// ValueFirstTok token might be useful to get argument location when Name is nil.
type Arg struct {
	Name          token.Token // token.Ident or nil
	Value         Expr
	ValueFirstTok token.Token
}

// ArgList represents argument list.
type ArgList struct {
	LParen token.LParen
	Args   []Arg
	RParen token.RParen
}

func (al ArgList) Len() int {
	return len(al.Args)
}

func buildArgList(ctx *context) (ArgList, error) {
	if _, ok := ctx.token().(token.LParen); !ok {
		return ArgList{}, nil
	}

	argList := ArgList{
		LParen: ctx.token().(token.LParen),
		Args:   []Arg{},
	}

	if _, ok := ctx.nextTok().(token.RParen); ok {
		return argList, token.Error{
			Msg:  "empty argument list",
			Toks: []token.Token{ctx.token()},
		}
	}

	arg := Arg{}

	type State int
	const (
		Name State = iota
		Ass
		Comma
		Val
	)
	state := Name

tokenLoop:
	for {
		ctx.idx++
		switch state {
		case Name:
			switch t := ctx.token().(type) {
			case token.Ident:
				switch ctx.nextTok().(type) {
				case token.Ass:
					arg.Name = t
					state = Ass
				default:
					arg.Name = nil
					arg.ValueFirstTok = t
					expr, err := buildExpr(ctx, nil)
					if err != nil {
						return argList, err
					}
					ctx.idx--
					arg.Value = expr
					argList.Args = append(argList.Args, arg)
					state = Comma
				}
			default:
				arg.Name = nil
				arg.ValueFirstTok = ctx.token()
				expr, err := buildExpr(ctx, nil)
				if err != nil {
					return argList, err
				}
				ctx.idx--
				arg.Value = expr
				argList.Args = append(argList.Args, arg)
				state = Comma
			}
		case Ass:
			switch t := ctx.token().(type) {
			case token.Ass:
				state = Val
			default:
				return argList, unexpected(t, "'='")
			}
		case Comma:
			switch t := ctx.token().(type) {
			case token.Comma:
				state = Name
			case token.RParen:
				argList.RParen = t
				ctx.idx++
				break tokenLoop
			default:
				return argList, unexpected(t, "',' or ')'")
			}
		case Val:
			arg.ValueFirstTok = ctx.token()
			expr, err := buildExpr(ctx, nil)
			if err != nil {
				return argList, err
			}
			ctx.idx--
			arg.Value = expr
			argList.Args = append(argList.Args, arg)
			state = Comma
		}
	}

	return argList, nil
}
