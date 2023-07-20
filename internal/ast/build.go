package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func buildExpr(s []token.Token, i int) (int, Expr, error) {
	switch t := s[i].(type) {
	case token.Neg, token.Sub, token.Add:
		return buildUnaryExpr(s, i)
	case token.Ident:
		switch s[i+1].(type) {
		case token.LeftParen:
			return buildCallExpr(s, i)
		default:
			return buildIdent(s, i)
		}
	case token.Int:
		return buildInt(s, i)
	default:
		return 0, Ident{}, fmt.Errorf(
			"%s: unexpected %s, expected expression",
			token.Loc(t), t.Kind(),
		)
	}
}

func buildIdent(s []token.Token, i int) (int, Ident, error) {
	id := Ident{Name: s[i]}
	return i + 1, id, nil
}

func buildInt(s []token.Token, i int) (int, Int, error) {
	int_ := Int{Val: s[i].(token.Int)}
	return i + 1, int_, nil
}

func buildCallExpr(s []token.Token, i int) (int, CallExpr, error) {
	call := CallExpr{
		Name:   s[i].(token.Ident),
		Lparen: s[i+1].(token.LeftParen),
	}
	i += 2

tokenLoop:
	for {
		switch t := s[i].(type) {
		case token.RightParen:
			call.Rparen = t
			i++
			break tokenLoop
		case token.Comma:
			i++
		default:
			var (
				expr Expr
				err  error
			)
			i, expr, err = buildExpr(s, i)
			if err != nil {
				return 0, call, err
			}
			call.Args = append(call.Args, expr)
		}
	}

	return i, call, nil
}

func buildUnaryExpr(s []token.Token, i int) (int, UnaryExpr, error) {
	un := UnaryExpr{Op: s[i]}
	i, x, err := buildExpr(s, i+1)
	if err != nil {
		return 0, un, err
	}
	un.X = x
	return i, un, nil
}
