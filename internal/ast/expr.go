package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Expr interface represents generic expression.
//
// The Tok method returns token which position spans the whole expression.
// It is useful for error messages.
// No assumptions shall be made on the returned token type.
type Expr interface {
	expr()
	Tok() token.Token
}

// Expression nodes
type (
	BinaryExpr struct {
		X  Expr
		Op token.Operator
		Y  Expr
	}

	BitString struct {
		X token.BitString
	}

	Bool struct {
		X token.Bool
	}

	// Function Call
	Call struct {
		Name token.Ident
		Args []Expr
	}

	List struct {
		LBracket token.LBracket
		Xs       []Expr
		RBracket token.RBracket
	}

	Ident struct {
		Name token.Token
	}

	QualIdent struct {
		Name token.Token
	}

	Int struct {
		X token.Int
	}

	Float struct {
		X token.Float
	}

	String struct {
		X token.String
	}

	Time struct {
		X token.Time
	}

	UnaryExpr struct {
		Op token.Token
		X  Expr
	}

	ParenExpr struct {
		LParen token.LParen
		X      Expr
		RParen token.RParen
	}
)

func (be BinaryExpr) expr()            {}
func (be BinaryExpr) Tok() token.Token { return be.Op }

func (bs BitString) expr()            {}
func (bs BitString) Tok() token.Token { return bs.X }

func (b Bool) expr()            {}
func (b Bool) Tok() token.Token { return b.X }

func (c Call) expr()            {}
func (c Call) Tok() token.Token { return c.Name }

func (l List) expr()            {}
func (l List) Tok() token.Token { return l.LBracket }

func (i Ident) expr()            {}
func (i Ident) Tok() token.Token { return i.Name }

func (qi QualIdent) expr()            {}
func (qi QualIdent) Tok() token.Token { return qi.Name }

func (i Int) expr()            {}
func (i Int) Tok() token.Token { return i.X }

func (f Float) expr()            {}
func (f Float) Tok() token.Token { return f.X }

func (s String) expr()            {}
func (s String) Tok() token.Token { return s.X }

func (t Time) expr()            {}
func (t Time) Tok() token.Token { return t.X }

func (ue UnaryExpr) expr()            {}
func (ue UnaryExpr) Tok() token.Token { return ue.Op }

func (pe ParenExpr) expr()            {}
func (pe ParenExpr) Tok() token.Token { return pe.LParen }

// leftOp is the operator on the left side of the expression.
func buildExpr(tokens *tokenStream, leftOp token.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := tokens.peek(0).(type) {
	case token.Neg, token.Sub, token.Add:
		expr, err = buildUnaryExpr(tokens)
	case token.Ident:
		switch tokens.peek(1).(type) {
		case token.LParen:
			expr, err = buildCallExpr(tokens)
		default:
			expr, err = buildIdent(tokens)
		}
	case token.QualIdent:
		switch tokens.peek(1).(type) {
		case token.LParen:
			expr, err = buildCallExpr(tokens)
		default:
			expr, err = buildQualIdent(tokens)
		}
	case token.Bool:
		expr, err = buildBool(tokens)
	case token.Int:
		expr, err = buildInt(tokens)
	case token.Float:
		expr, err = buildFloat(tokens)
	case token.String:
		expr, err = buildString(tokens)
	case token.Time:
		expr, err = buildTime(tokens)
	case token.BitString:
		expr, err = buildBitString(tokens)
	case token.LParen:
		expr, err = buildParenExpr(tokens)
	case token.LBracket:
		expr, err = buildList(tokens)
	default:
		return Ident{}, unexpected(t, "expression")
	}

	if err != nil {
		return expr, err
	}

	for {
		var rightOp token.Operator
		if op, ok := tokens.peek(0).(token.Operator); ok {
			rightOp = op
		} else {
			return expr, nil
		}

		if (leftOp == nil) ||
			(leftOp != nil && (leftOp.Precedence() < rightOp.Precedence())) {
			be := BinaryExpr{X: expr, Op: rightOp}
			tokens.next()
			expr, err = buildExpr(tokens, rightOp)
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

func buildIdent(tokens *tokenStream) (Ident, error) {
	id := Ident{Name: tokens.next()}
	return id, nil
}

func buildQualIdent(tokens *tokenStream) (QualIdent, error) {
	id := QualIdent{Name: tokens.next()}
	return id, nil
}

func buildBool(tokens *tokenStream) (Bool, error) {
	b := Bool{tokens.next().(token.Bool)}
	return b, nil
}

func buildInt(tokens *tokenStream) (Int, error) {
	int_ := Int{tokens.next().(token.Int)}
	return int_, nil
}

func buildFloat(tokens *tokenStream) (Float, error) {
	r := Float{tokens.next().(token.Float)}
	return r, nil
}

func buildString(tokens *tokenStream) (String, error) {
	s := String{tokens.next().(token.String)}
	return s, nil
}

func buildTime(tokens *tokenStream) (Time, error) {
	t := Time{tokens.next().(token.Time)}
	return t, nil
}

func buildBitString(tokens *tokenStream) (BitString, error) {
	s := BitString{tokens.next().(token.BitString)}
	return s, nil
}

func buildParenExpr(tokens *tokenStream) (ParenExpr, error) {
	pe := ParenExpr{}
	var (
		err  error
		expr Expr
	)

	pe.LParen = tokens.next().(token.LParen)
	expr, err = buildExpr(tokens, nil)
	if err != nil {
		return pe, err
	}
	pe.X = expr

	if rp, ok := tokens.peek(0).(token.RParen); ok {
		pe.RParen = rp
		tokens.next()
	} else {
		return pe, unexpected(tokens.peek(0), "')'")
	}

	return pe, nil
}

func buildList(tokens *tokenStream) (List, error) {
	l := List{}
	prevExpr := false
	l.LBracket = tokens.peek(0).(token.LBracket)
	tokens.next()

tokenLoop:
	for {
		switch t := tokens.peek(0).(type) {
		case token.RBracket:
			l.RBracket = t
			tokens.next()
			break tokenLoop
		case token.Comma:
			if len(l.Xs) == 0 {
				return l, unexpected(t, "expression")
			}
			prevExpr = false
			tokens.next()
		default:
			if prevExpr {
				return l, unexpected(t, "',' or ']'")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(tokens, nil)
			if err != nil {
				return l, err
			}
			l.Xs = append(l.Xs, expr)
			prevExpr = true
		}
	}

	return l, nil
}

func buildCallExpr(tokens *tokenStream) (Call, error) {
	call := Call{Name: tokens.next().(token.Ident)}
	tokens.next()

	prevExpr := false

tokenLoop:
	for {
		switch t := tokens.peek(0).(type) {
		case token.RParen:
			tokens.next()
			break tokenLoop
		case token.Comma:
			if len(call.Args) == 0 {
				return call, unexpected(t, "expression")
			}
			prevExpr = false
			tokens.next()
		default:
			if prevExpr {
				return call, unexpected(t, "',' or ')'")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(tokens, nil)
			if err != nil {
				return call, err
			}
			call.Args = append(call.Args, expr)
			prevExpr = true
		}
	}

	return call, nil
}

func buildUnaryExpr(tokens *tokenStream) (UnaryExpr, error) {
	op := tokens.peek(0).(token.Operator)
	un := UnaryExpr{Op: op}
	tokens.next()
	x, err := buildExpr(tokens, op)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
