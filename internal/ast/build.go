package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// leftOp is the operator on the left side of the expression.
func buildExpr(s []token.Token, i int, leftOp token.Operator) (int, Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := s[i].(type) {
	case token.Neg, token.Sub, token.Add:
		i, expr, err = buildUnaryExpr(s, i)
	case token.Ident:
		switch s[i+1].(type) {
		case token.LeftParen:
			i, expr, err = buildCallExpr(s, i)
		default:
			i, expr, err = buildIdent(s, i)
		}
	case token.Bool:
		i, expr, err = buildBool(s, i)
	case token.Int:
		i, expr, err = buildInt(s, i)
	case token.Real:
		i, expr, err = buildReal(s, i)
	case token.LeftParen:
		i, expr, err = buildParenExpr(s, i)
	default:
		return 0, Ident{}, fmt.Errorf(
			"%s: unexpected %s, expected expression",
			token.Loc(t), t.Kind(),
		)
	}

	if err != nil {
		return 0, expr, err
	}

	for {
		var rightOp token.Operator
		if op, ok := s[i].(token.Operator); ok {
			rightOp = op
		} else {
			return i, expr, nil
		}

		if (leftOp == nil) ||
			(leftOp != nil && (leftOp.Precedence() < rightOp.Precedence())) {
			be := BinaryExpr{X: expr, Op: rightOp}
			i, expr, err = buildExpr(s, i+1, rightOp)
			if err != nil {
				return 0, expr, err
			}
			be.Y = expr
			expr = be
		} else if leftOp.Precedence() >= rightOp.Precedence() {
			return i, expr, nil
		}
	}
}

func buildIdent(s []token.Token, i int) (int, Ident, error) {
	id := Ident{Name: s[i]}
	return i + 1, id, nil
}

func buildBool(s []token.Token, i int) (int, Bool, error) {
	b := Bool{Val: s[i].(token.Bool)}
	return i + 1, b, nil
}

func buildInt(s []token.Token, i int) (int, Int, error) {
	int_ := Int{Val: s[i].(token.Int)}
	return i + 1, int_, nil
}

func buildReal(s []token.Token, i int) (int, Real, error) {
	r := Real{Val: s[i].(token.Real)}
	return i + 1, r, nil
}

func buildParenExpr(s []token.Token, i int) (int, ParenExpr, error) {
	pe := ParenExpr{Lparen: s[i].(token.LeftParen)}
	var (
		err  error
		expr Expr
	)
	i, expr, err = buildExpr(s, i+1, nil)
	if err != nil {
		return 0, pe, err
	}
	pe.X = expr

	if rp, ok := s[i].(token.RightParen); ok {
		pe.Rparen = rp
	} else {
		return 0, pe, fmt.Errorf(
			"%s: unexpected %s, expected )",
			token.Loc(s[i]), s[i].Kind(),
		)
	}

	return i + 1, pe, nil
}

func buildCallExpr(s []token.Token, i int) (int, CallExpr, error) {
	call := CallExpr{
		Name:   s[i].(token.Ident),
		Lparen: s[i+1].(token.LeftParen),
	}
	i += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := s[i].(type) {
		case token.RightParen:
			call.Rparen = t
			i++
			break tokenLoop
		case token.Comma:
			prevExpr = false
			i++
		default:
			if prevExpr {
				return 0, call, fmt.Errorf(
					"%s: unexpected %s, expected , or )",
					token.Loc(t), t.Kind(),
				)
			}

			var (
				expr Expr
				err  error
			)
			i, expr, err = buildExpr(s, i, nil)
			if err != nil {
				return 0, call, err
			}
			call.Args = append(call.Args, expr)
			prevExpr = true
		}
	}

	return i, call, nil
}

func buildUnaryExpr(s []token.Token, i int) (int, UnaryExpr, error) {
	un := UnaryExpr{Op: s[i]}
	i, x, err := buildExpr(s, i+1, nil)
	if err != nil {
		return 0, un, err
	}
	un.X = x
	return i, un, nil
}
