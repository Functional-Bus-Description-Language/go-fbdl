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
func buildExpr(s []token.Token, c *ctx, leftOp token.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := s[c.i].(type) {
	case token.Neg, token.Sub, token.Add:
		expr, err = buildUnaryExpr(s, c)
	case token.Ident:
		switch s[c.i+1].(type) {
		case token.LeftParen:
			expr, err = buildCallExpr(s, c)
		default:
			expr, err = buildIdent(s, c)
		}
	case token.Bool:
		expr, err = buildBool(s, c)
	case token.Int:
		expr, err = buildInt(s, c)
	case token.Real:
		expr, err = buildReal(s, c)
	case token.LeftParen:
		expr, err = buildParenExpr(s, c)
	case token.LeftBracket:
		expr, err = buildExprList(s, c)
	default:
		return Ident{}, unexpected(t, "expression")
	}

	if err != nil {
		return expr, err
	}

	for {
		var rightOp token.Operator
		if op, ok := s[c.i].(token.Operator); ok {
			rightOp = op
		} else {
			return expr, nil
		}

		if (leftOp == nil) ||
			(leftOp != nil && (leftOp.Precedence() < rightOp.Precedence())) {
			be := BinaryExpr{X: expr, Op: rightOp}
			c.i++
			expr, err = buildExpr(s, c, rightOp)
			if err != nil {
				return expr, err
			}
			be.Y = expr
			expr = be
		} else if leftOp.Precedence() >= rightOp.Precedence() {
			return expr, nil
		}
	}
}

func buildIdent(s []token.Token, c *ctx) (Ident, error) {
	id := Ident{Name: s[c.i]}
	c.i++
	return id, nil
}

func buildBool(s []token.Token, c *ctx) (Bool, error) {
	b := Bool{Val: s[c.i].(token.Bool)}
	c.i++
	return b, nil
}

func buildInt(s []token.Token, c *ctx) (Int, error) {
	int_ := Int{Val: s[c.i].(token.Int)}
	c.i++
	return int_, nil
}

func buildReal(s []token.Token, c *ctx) (Real, error) {
	r := Real{Val: s[c.i].(token.Real)}
	c.i++
	return r, nil
}

func buildParenExpr(s []token.Token, c *ctx) (ParenExpr, error) {
	pe := ParenExpr{Lparen: s[c.i].(token.LeftParen)}
	var (
		err  error
		expr Expr
	)
	c.i++
	expr, err = buildExpr(s, c, nil)
	if err != nil {
		return pe, err
	}
	pe.X = expr

	if rp, ok := s[c.i].(token.RightParen); ok {
		pe.Rparen = rp
		c.i++
	} else {
		return pe, unexpected(s[c.i], ")")
	}

	return pe, nil
}

func buildExprList(s []token.Token, c *ctx) (ExprList, error) {
	el := ExprList{Lbracket: s[c.i].(token.LeftBracket)}
	prevExpr := false
	lbi := c.i // Left bracket token index
	c.i++

tokenLoop:
	for {
		switch t := s[c.i].(type) {
		case token.RightBracket:
			el.Rbracket = t
			c.i++
			break tokenLoop
		case token.Comma:
			if c.i == lbi+1 {
				return el, unexpected(t, "expression")
			}
			prevExpr = false
			c.i++
		default:
			if prevExpr {
				return el, unexpected(t, ", or ]")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(s, c, nil)
			if err != nil {
				return el, err
			}
			el.Exprs = append(el.Exprs, expr)
			prevExpr = true
		}
	}

	return el, nil
}

func buildCallExpr(s []token.Token, c *ctx) (CallExpr, error) {
	call := CallExpr{
		Name:   s[c.i].(token.Ident),
		Lparen: s[c.i+1].(token.LeftParen),
	}
	lpi := c.i // Left parenthesis token index
	c.i += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := s[c.i].(type) {
		case token.RightParen:
			call.Rparen = t
			c.i++
			break tokenLoop
		case token.Comma:
			if c.i == lpi+2 {
				return call, unexpected(t, "expression")
			}
			prevExpr = false
			c.i++
		default:
			if prevExpr {
				return call, unexpected(t, ", or )")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(s, c, nil)
			if err != nil {
				return call, err
			}
			call.Args = append(call.Args, expr)
			prevExpr = true
		}
	}

	return call, nil
}

func buildUnaryExpr(s []token.Token, c *ctx) (UnaryExpr, error) {
	un := UnaryExpr{Op: s[c.i]}
	c.i++
	x, err := buildExpr(s, c, nil)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
