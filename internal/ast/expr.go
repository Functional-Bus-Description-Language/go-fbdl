package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Expr interface {
	exprNode()
}

// Expression nodes
type (
	BinaryExpr struct {
		X  Expr
		Op token.Operator
		Y  Expr
	}

	Bool struct {
		Val token.Bool
	}

	CallExpr struct {
		Name   token.Ident
		Lparen token.LeftParen
		Args   []Expr
		Rparen token.RightParen
	}

	ExprList struct {
		Lbracket token.LeftBracket
		Exprs    []Expr
		Rbracket token.RightBracket
	}

	Ident struct {
		Name token.Token
	}

	Int struct {
		Val token.Int
	}

	Real struct {
		Val token.Real
	}

	UnaryExpr struct {
		Op token.Token
		X  Expr
	}

	ParenExpr struct {
		Lparen token.Token
		X      Expr
		Rparen token.Token
	}
)

func (be BinaryExpr) exprNode() {}
func (b Bool) exprNode()        {}
func (c CallExpr) exprNode()    {}
func (el ExprList) exprNode()   {}
func (i Ident) exprNode()       {}
func (i Int) exprNode()         {}
func (r Real) exprNode()        {}
func (ue UnaryExpr) exprNode()  {}
func (pe ParenExpr) exprNode()  {}

func (c CallExpr) eq(c2 CallExpr) bool {
	if c.Name != c2.Name ||
		c.Lparen != c2.Lparen ||
		c.Rparen != c2.Rparen {
		return false
	}

	if len(c.Args) != len(c2.Args) {
		return false
	}

	for i := range c.Args {
		if c.Args[i] != c2.Args[i] {
			return false
		}
	}

	return true
}


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
	case token.LeftBracket:
		i, expr, err = buildExprList(s, i)
	default:
		return 0, Ident{}, unexpected(t, "expression")
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
		return 0, pe, unexpected(s[i], ")")
	}

	return i + 1, pe, nil
}

func buildExprList(s []token.Token, i int) (int, ExprList, error) {
	el := ExprList{Lbracket: s[i].(token.LeftBracket)}
	prevExpr := false
	lbi := i // Left bracket token index
	i++

tokenLoop:
	for {
		switch t := s[i].(type) {
		case token.RightBracket:
			el.Rbracket = t
			i++
			break tokenLoop
		case token.Comma:
			if i == lbi+1 {
				return 0, el, unexpected(t, "expression")
			}
			prevExpr = false
			i++
		default:
			if prevExpr {
				return 0, el, unexpected(t, ", or ]")
			}

			var (
				expr Expr
				err  error
			)
			i, expr, err = buildExpr(s, i, nil)
			if err != nil {
				return 0, el, err
			}
			el.Exprs = append(el.Exprs, expr)
			prevExpr = true
		}
	}

	return i, el, nil
}

func buildCallExpr(s []token.Token, i int) (int, CallExpr, error) {
	call := CallExpr{
		Name:   s[i].(token.Ident),
		Lparen: s[i+1].(token.LeftParen),
	}
	lpi := i // Left parenthesis token index
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
			if i == lpi+2 {
				return 0, call, unexpected(t, "expression")
			}
			prevExpr = false
			i++
		default:
			if prevExpr {
				return 0, call, unexpected(t, ", or )")
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
