package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Arg represents instantiation or type argument.
// ValueFirstTok token might be useful to get argument location when Name is nil.
type Arg struct {
	Name          tok.Token // tok.Ident or nil
	Value         Expr
	ValueFirstTok tok.Token
}

// ArgList represents argument list.
type ArgList struct {
	LParen tok.LParen
	Args   []Arg
	RParen tok.RParen
}

func (al ArgList) Len() int {
	return len(al.Args)
}

func buildArgList(ctx *context) (ArgList, error) {
	if _, ok := ctx.tok().(tok.LParen); !ok {
		return ArgList{}, nil
	}

	argList := ArgList{
		LParen: ctx.tok().(tok.LParen),
		Args:   []Arg{},
	}

	if _, ok := ctx.nextTok().(tok.RParen); ok {
		return argList, tok.Error{
			Msg:  "empty argument list",
			Toks: []tok.Token{tok.Join(ctx.tok(), ctx.nextTok())},
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
			switch t := ctx.tok().(type) {
			case tok.Ident:
				switch ctx.nextTok().(type) {
				case tok.Ass:
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
				arg.ValueFirstTok = ctx.tok()
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
			switch t := ctx.tok().(type) {
			case tok.Ass:
				state = Val
			default:
				return argList, unexpected(t, "'='")
			}
		case Comma:
			switch t := ctx.tok().(type) {
			case tok.Comma:
				state = Name
			case tok.RParen:
				argList.RParen = t
				ctx.idx++
				break tokenLoop
			default:
				return argList, unexpected(t, "',' or ')'")
			}
		case Val:
			arg.ValueFirstTok = ctx.tok()
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
