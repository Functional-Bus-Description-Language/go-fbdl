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

func buildArgList(tokens *tokenStream) (ArgList, error) {
	if _, ok := tokens.peek(0).(token.LParen); !ok {
		return ArgList{}, nil
	}

	argList := ArgList{
		LParen: tokens.next().(token.LParen),
		Args:   []Arg{},
	}

	if _, ok := tokens.peek(0).(token.RParen); ok {
		return argList, token.Error{
			Msg:  "empty argument list",
			Toks: []token.Token{argList.LParen},
		}
	}

	for {
		arg := Arg{}
		if name, ok := tokens.peek(0).(token.Ident); ok {
			if _, ok := tokens.peek(1).(token.Ass); ok {
				arg.Name = name
				tokens.next()
				tokens.next()
			}
		}

		arg.ValueFirstTok = tokens.peek(0)
		expr, err := buildExpr(tokens, nil)
		if err != nil {
			return argList, err
		}
		arg.Value = expr
		argList.Args = append(argList.Args, arg)

		switch t := tokens.peek(0).(type) {
		case token.Comma:
			tokens.next()
		case token.RParen:
			argList.RParen = t
			tokens.next()
			return argList, nil
		default:
			return argList, unexpected(t, "',' or ')'")
		}
	}
}
