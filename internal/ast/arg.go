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
	LeftParen  tok.LeftParen
	Args       []Arg
	RightParen tok.RightParen
}

func (al ArgList) Len() int {
	return len(al.Args)
}

func buildArgList(toks []tok.Token, ctx *context) (ArgList, error) {
	if _, ok := toks[ctx.idx].(tok.LeftParen); !ok {
		return ArgList{}, nil
	}

	argList := ArgList{
		LeftParen: toks[ctx.idx].(tok.LeftParen),
		Args:      []Arg{},
	}

	if _, ok := toks[ctx.idx+1].(tok.RightParen); ok {
		return argList, tok.Error{
			Msg:  "empty argument list",
			Toks: []tok.Token{tok.Join(toks[ctx.idx], toks[ctx.idx+1])},
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
			switch t := toks[ctx.idx].(type) {
			case tok.Ident:
				switch toks[ctx.idx+1].(type) {
				case tok.Ass:
					arg.Name = t
					state = Ass
				default:
					arg.Name = nil
					arg.ValueFirstTok = t
					expr, err := buildExpr(toks, ctx, nil)
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
				arg.ValueFirstTok = toks[ctx.idx]
				expr, err := buildExpr(toks, ctx, nil)
				if err != nil {
					return argList, err
				}
				ctx.idx--
				arg.Value = expr
				argList.Args = append(argList.Args, arg)
				state = Comma
			}
		case Ass:
			switch t := toks[ctx.idx].(type) {
			case tok.Ass:
				state = Val
			default:
				return argList, unexpected(t, "'='")
			}
		case Comma:
			switch t := toks[ctx.idx].(type) {
			case tok.Comma:
				state = Name
			case tok.RightParen:
				argList.RightParen = t
				ctx.idx++
				break tokenLoop
			default:
				return argList, unexpected(t, "',' or ')'")
			}
		case Val:
			arg.ValueFirstTok = toks[ctx.idx]
			expr, err := buildExpr(toks, ctx, nil)
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
