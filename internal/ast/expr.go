package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Expr interface represents generic expression.
type Expr interface {
	expr()
}

// Expression nodes
type (
	BinaryExpr struct {
		X  Expr
		Op tok.Operator
		Y  Expr
	}

	BitString struct {
		X tok.BitString
	}

	Bool struct {
		X tok.Bool
	}

	// Function Call
	Call struct {
		Name tok.Ident
		Args []Expr
	}

	List struct {
		Xs []Expr
	}

	Ident struct {
		Name tok.Token
	}

	Int struct {
		X tok.Int
	}

	Real struct {
		X tok.Real
	}

	String struct {
		X tok.String
	}

	Time struct {
		X tok.Time
	}

	UnaryExpr struct {
		Op tok.Token
		X  Expr
	}

	ParenExpr struct {
		X Expr
	}
)

func (be BinaryExpr) expr() {}
func (bs BitString) expr()  {}
func (b Bool) expr()        {}
func (c Call) expr()        {}
func (l List) expr()        {}
func (i Ident) expr()       {}
func (i Int) expr()         {}
func (r Real) expr()        {}
func (s String) expr()      {}
func (t Time) expr()        {}
func (ue UnaryExpr) expr()  {}
func (pe ParenExpr) expr()  {}

// leftOp is the operator on the left side of the expression.
func buildExpr(toks []tok.Token, ctx *context, leftOp tok.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := toks[ctx.i].(type) {
	case tok.Neg, tok.Sub, tok.Add:
		expr, err = buildUnaryExpr(toks, ctx)
	case tok.Ident:
		switch toks[ctx.i+1].(type) {
		case tok.LeftParen:
			expr, err = buildCallExpr(toks, ctx)
		default:
			expr, err = buildIdent(toks, ctx)
		}
	case tok.Bool:
		expr, err = buildBool(toks, ctx)
	case tok.Int:
		expr, err = buildInt(toks, ctx)
	case tok.Real:
		expr, err = buildReal(toks, ctx)
	case tok.String:
		expr, err = buildString(toks, ctx)
	case tok.Time:
		expr, err = buildTime(toks, ctx)
	case tok.BitString:
		expr, err = buildBitString(toks, ctx)
	case tok.LeftParen:
		expr, err = buildParenExpr(toks, ctx)
	case tok.LeftBracket:
		expr, err = buildList(toks, ctx)
	default:
		return Ident{}, unexpected(t, "expression")
	}

	if err != nil {
		return expr, err
	}

	for {
		var rightOp tok.Operator
		if op, ok := toks[ctx.i].(tok.Operator); ok {
			rightOp = op
		} else {
			return expr, nil
		}

		if (leftOp == nil) ||
			(leftOp != nil && (leftOp.Precedence() < rightOp.Precedence())) {
			be := BinaryExpr{X: expr, Op: rightOp}
			ctx.i++
			expr, err = buildExpr(toks, ctx, rightOp)
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

func buildIdent(toks []tok.Token, ctx *context) (Ident, error) {
	id := Ident{Name: toks[ctx.i]}
	ctx.i++
	return id, nil
}

func buildBool(toks []tok.Token, ctx *context) (Bool, error) {
	b := Bool{toks[ctx.i].(tok.Bool)}
	ctx.i++
	return b, nil
}

func buildInt(toks []tok.Token, ctx *context) (Int, error) {
	int_ := Int{toks[ctx.i].(tok.Int)}
	ctx.i++
	return int_, nil
}

func buildReal(toks []tok.Token, ctx *context) (Real, error) {
	r := Real{toks[ctx.i].(tok.Real)}
	ctx.i++
	return r, nil
}

func buildString(toks []tok.Token, ctx *context) (String, error) {
	s := String{toks[ctx.i].(tok.String)}
	ctx.i++
	return s, nil
}

func buildTime(toks []tok.Token, ctx *context) (Time, error) {
	t := Time{toks[ctx.i].(tok.Time)}
	ctx.i++
	return t, nil
}

func buildBitString(toks []tok.Token, ctx *context) (BitString, error) {
	s := BitString{toks[ctx.i].(tok.BitString)}
	ctx.i++
	return s, nil
}

func buildParenExpr(toks []tok.Token, ctx *context) (ParenExpr, error) {
	pe := ParenExpr{}
	var (
		err  error
		expr Expr
	)
	ctx.i++
	expr, err = buildExpr(toks, ctx, nil)
	if err != nil {
		return pe, err
	}
	pe.X = expr

	if _, ok := toks[ctx.i].(tok.RightParen); ok {
		ctx.i++
	} else {
		return pe, unexpected(toks[ctx.i], "')'")
	}

	return pe, nil
}

func buildList(toks []tok.Token, ctx *context) (List, error) {
	l := List{}
	prevExpr := false
	lbi := ctx.i // Left bracket token index
	ctx.i++

tokenLoop:
	for {
		switch t := toks[ctx.i].(type) {
		case tok.RightBracket:
			ctx.i++
			break tokenLoop
		case tok.Comma:
			if ctx.i == lbi+1 {
				return l, unexpected(t, "expression")
			}
			prevExpr = false
			ctx.i++
		default:
			if prevExpr {
				return l, unexpected(t, "',' or ']'")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(toks, ctx, nil)
			if err != nil {
				return l, err
			}
			l.Xs = append(l.Xs, expr)
			prevExpr = true
		}
	}

	return l, nil
}

func buildCallExpr(toks []tok.Token, ctx *context) (Call, error) {
	call := Call{Name: toks[ctx.i].(tok.Ident)}
	lpi := ctx.i // Left parenthesis token index
	ctx.i += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := toks[ctx.i].(type) {
		case tok.RightParen:
			ctx.i++
			break tokenLoop
		case tok.Comma:
			if ctx.i == lpi+2 {
				return call, unexpected(t, "expression")
			}
			prevExpr = false
			ctx.i++
		default:
			if prevExpr {
				return call, unexpected(t, "',' or ')'")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(toks, ctx, nil)
			if err != nil {
				return call, err
			}
			call.Args = append(call.Args, expr)
			prevExpr = true
		}
	}

	return call, nil
}

func buildUnaryExpr(toks []tok.Token, ctx *context) (UnaryExpr, error) {
	un := UnaryExpr{Op: toks[ctx.i]}
	ctx.i++
	x, err := buildExpr(toks, ctx, nil)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
