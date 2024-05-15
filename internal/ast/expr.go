package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Expr interface represents generic expression.
//
// The Tok method returns token which position spans the whole expression.
// It is useful for error messages.
// No assumptions shall be made on the returned token type.
type Expr interface {
	expr()
	Tok() tok.Token
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
		LeftBracket  tok.LeftBracket
		Xs           []Expr
		RightBracket tok.RightBracket
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
		LeftParen  tok.LeftParen
		X          Expr
		RightParen tok.RightParen
	}
)

func (be BinaryExpr) expr()          {}
func (be BinaryExpr) Tok() tok.Token { return tok.Join(be.X.Tok(), be.Y.Tok()) }

func (bs BitString) expr()          {}
func (bs BitString) Tok() tok.Token { return bs.X }

func (b Bool) expr()          {}
func (b Bool) Tok() tok.Token { return b.X }

func (c Call) expr()          {}
func (c Call) Tok() tok.Token { return c.Name }

func (l List) expr()          {}
func (l List) Tok() tok.Token { return tok.Join(l.LeftBracket, l.RightBracket) }

func (i Ident) expr()          {}
func (i Ident) Tok() tok.Token { return i.Name }

func (i Int) expr()          {}
func (i Int) Tok() tok.Token { return i.X }

func (r Real) expr()          {}
func (r Real) Tok() tok.Token { return r.X }

func (s String) expr()          {}
func (s String) Tok() tok.Token { return s.X }

func (t Time) expr()          {}
func (t Time) Tok() tok.Token { return t.X }

func (ue UnaryExpr) expr()          {}
func (ue UnaryExpr) Tok() tok.Token { return tok.Join(ue.Op, ue.X.Tok()) }

func (pe ParenExpr) expr()          {}
func (pe ParenExpr) Tok() tok.Token { return tok.Join(pe.LeftParen, pe.RightParen) }

// leftOp is the operator on the left side of the expression.
func buildExpr(toks []tok.Token, ctx *context, leftOp tok.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := toks[ctx.idx].(type) {
	case tok.Neg, tok.Sub, tok.Add:
		expr, err = buildUnaryExpr(toks, ctx)
	case tok.Ident:
		switch toks[ctx.idx+1].(type) {
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
		if op, ok := toks[ctx.idx].(tok.Operator); ok {
			rightOp = op
		} else {
			return expr, nil
		}

		if (leftOp == nil) ||
			(leftOp != nil && (leftOp.Precedence() < rightOp.Precedence())) {
			be := BinaryExpr{X: expr, Op: rightOp}
			ctx.idx++
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
	id := Ident{Name: toks[ctx.idx]}
	ctx.idx++
	return id, nil
}

func buildBool(toks []tok.Token, ctx *context) (Bool, error) {
	b := Bool{toks[ctx.idx].(tok.Bool)}
	ctx.idx++
	return b, nil
}

func buildInt(toks []tok.Token, ctx *context) (Int, error) {
	int_ := Int{toks[ctx.idx].(tok.Int)}
	ctx.idx++
	return int_, nil
}

func buildReal(toks []tok.Token, ctx *context) (Real, error) {
	r := Real{toks[ctx.idx].(tok.Real)}
	ctx.idx++
	return r, nil
}

func buildString(toks []tok.Token, ctx *context) (String, error) {
	s := String{toks[ctx.idx].(tok.String)}
	ctx.idx++
	return s, nil
}

func buildTime(toks []tok.Token, ctx *context) (Time, error) {
	t := Time{toks[ctx.idx].(tok.Time)}
	ctx.idx++
	return t, nil
}

func buildBitString(toks []tok.Token, ctx *context) (BitString, error) {
	s := BitString{toks[ctx.idx].(tok.BitString)}
	ctx.idx++
	return s, nil
}

func buildParenExpr(toks []tok.Token, ctx *context) (ParenExpr, error) {
	pe := ParenExpr{}
	var (
		err  error
		expr Expr
	)

	pe.LeftParen = toks[ctx.idx].(tok.LeftParen)

	ctx.idx++
	expr, err = buildExpr(toks, ctx, nil)
	if err != nil {
		return pe, err
	}
	pe.X = expr

	if _, ok := toks[ctx.idx].(tok.RightParen); ok {
		pe.RightParen = toks[ctx.idx].(tok.RightParen)
		ctx.idx++
	} else {
		return pe, unexpected(toks[ctx.idx], "')'")
	}

	return pe, nil
}

func buildList(toks []tok.Token, ctx *context) (List, error) {
	l := List{}
	prevExpr := false
	l.LeftBracket = toks[ctx.idx].(tok.LeftBracket)
	lbi := ctx.idx // Left bracket token index
	ctx.idx++

tokenLoop:
	for {
		switch t := toks[ctx.idx].(type) {
		case tok.RightBracket:
			l.RightBracket = t
			ctx.idx++
			break tokenLoop
		case tok.Comma:
			if ctx.idx == lbi+1 {
				return l, unexpected(t, "expression")
			}
			prevExpr = false
			ctx.idx++
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
	call := Call{Name: toks[ctx.idx].(tok.Ident)}
	lpi := ctx.idx // Left parenthesis token index
	ctx.idx += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := toks[ctx.idx].(type) {
		case tok.RightParen:
			ctx.idx++
			break tokenLoop
		case tok.Comma:
			if ctx.idx == lpi+2 {
				return call, unexpected(t, "expression")
			}
			prevExpr = false
			ctx.idx++
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
	op := toks[ctx.idx].(tok.Operator)
	un := UnaryExpr{Op: op}
	ctx.idx++
	x, err := buildExpr(toks, ctx, op)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
