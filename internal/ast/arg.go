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
		switch state {
		case Name:
			switch t := tokens.peek(0).(type) {
			case token.Ident:
				switch tokens.peek(1).(type) {
				case token.Ass:
					arg.Name = t
					tokens.next()
					state = Ass
				default:
					arg.Name = nil
					arg.ValueFirstTok = t
					expr, err := buildExpr(tokens, nil)
					if err != nil {
						return argList, err
					}
					arg.Value = expr
					argList.Args = append(argList.Args, arg)
					state = Comma
				}
			default:
				arg.Name = nil
				arg.ValueFirstTok = tokens.peek(0)
				expr, err := buildExpr(tokens, nil)
				if err != nil {
					return argList, err
				}
				arg.Value = expr
				argList.Args = append(argList.Args, arg)
				state = Comma
			}
		case Ass:
			switch t := tokens.peek(0).(type) {
			case token.Ass:
				tokens.next()
				state = Val
			default:
				return argList, unexpected(t, "'='")
			}
		case Comma:
			switch t := tokens.peek(0).(type) {
			case token.Comma:
				tokens.next()
				state = Name
			case token.RParen:
				argList.RParen = t
				tokens.next()
				break tokenLoop
			default:
				return argList, unexpected(t, "',' or ')'")
			}
		case Val:
			arg.ValueFirstTok = tokens.peek(0)
			expr, err := buildExpr(tokens, nil)
			if err != nil {
				return argList, err
			}
			arg.Value = expr
			argList.Args = append(argList.Args, arg)
			state = Comma
		}
	}

	return argList, nil
}
