package prs

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type Expr interface {
	Eval() (val.Value, error)
}

func MakeExpr(astExpr ast.Expr, src []byte, s Scope) (Expr, error) {
	var err error = nil
	var expr Expr

	switch e := astExpr.(type) {
	case ast.BinaryExpr:
		expr, err = MakeBinaryExpr(e, src, s)
	case ast.BitString:
		expr, err = MakeBitString(e, src)
	case ast.Call:
		expr, err = MakeCall(e, src, s)
	case ast.Ident:
		expr = MakeDeclaredIdentifier(e, src, s)
	case ast.Int:
		expr, err = MakeInt(e, src)
	case ast.List:
		expr, err = MakeList(e, src, s)
	case ast.Bool:
		expr = MakeBool(e, src)
	case ast.Float:
		expr, err = MakeFloat(e, src)
	case ast.String:
		expr = MakeString(e, src)
	case ast.Time:
		expr, err = MakeTime(e, src, s)
	case ast.UnaryExpr:
		expr, err = MakeUnaryExpr(e, src, s)
	case nil:
		return nil, nil
	default:
		panic(fmt.Sprintf("unimplemented type %T", astExpr))
	}

	return expr, err
}

type BinaryExpr struct {
	ast ast.BinaryExpr

	x  Expr
	op tok.Operator
	y  Expr
}

func (be BinaryExpr) Eval() (val.Value, error) {
	x, err := be.x.Eval()
	if err != nil {
		return nil, err
	}
	y, err := be.y.Eval()
	if err != nil {
		return nil, err
	}
	op := be.op // Operator

	var v val.Value

	switch x := x.(type) {
	case val.Int:
		switch y := y.(type) {
		case val.Int:
			switch op.(type) {
			case tok.Add:
				v = val.Int(x + y)
			case tok.Sub:
				v = val.Int(x - y)
			case tok.Mul:
				v = val.Int(x * y)
			case tok.Div:
				if x%y == 0 {
					v = val.Int(x / y)
				} else {
					// Float must be returned here
					panic("unimplemented")
				}
			case tok.Rem:
				v = val.Int(x % y)
			case tok.Exp:
				v = val.Int(int64(math.Pow(float64(x), float64(y))))
			case tok.LShift:
				v = val.Int(x << y)
			case tok.RShift:
				v = val.Int(x >> y)
			case tok.Colon:
				v = val.Range{L: int64(x), R: int64(y)}
			}
		case val.Float:
			switch op.(type) {
			case tok.Colon:
				return nil, tok.Error{
					Msg:  "right bound of range must be of type integer, current type float",
					Toks: []tok.Token{be.ast.Y.Tok()},
				}
			}
		}
	case val.Range:
		switch op.(type) {
		case tok.Colon:
			return nil, tok.Error{
				Msg:  "left bound of range must be of type integer, current type range",
				Toks: []tok.Token{be.ast.X.Tok()},
			}
		}
	}

	if v != nil {
		return v, nil
	}

	return nil, tok.Error{
		Msg: fmt.Sprintf(
			"unimplemented binary expression evaluation for %s operator, left operand type %s, right operand type %s, please report this error on %s",
			op.Name(), x.Type(), y.Type(), util.RepoIssueUrl,
		),
		Toks: []tok.Token{op},
	}
}

func MakeBinaryExpr(be ast.BinaryExpr, src []byte, s Scope) (BinaryExpr, error) {
	x, err := MakeExpr(be.X, src, s)
	if err != nil {
		return BinaryExpr{}, fmt.Errorf("make binary expression: left operand: %v", err)
	}

	y, err := MakeExpr(be.Y, src, s)
	if err != nil {
		return BinaryExpr{}, fmt.Errorf("make binary expression: right operand: %v", err)
	}

	return BinaryExpr{ast: be, x: x, op: be.Op, y: y}, nil
}

type BitString struct {
	x val.BitStr
}

func (bs BitString) Eval() (val.Value, error) {
	return bs.x, nil
}

func MakeBitString(e ast.BitString, src []byte) (BitString, error) {
	x, err := val.MakeBitStr(tok.Text(e.X, src))
	if err != nil {
		return BitString{}, fmt.Errorf("make bit string: %v", err)
	}

	return BitString{x: x}, nil
}

type Call struct {
	funcName string
	args     []Expr
}

func (c Call) Eval() (val.Value, error) {
	switch c.funcName {
	case "bool":
		return evalBool(c)
	case "ceil":
		return evalCeil(c)
	case "floor":
		return evalFloor(c)
	case "log2":
		return evalLog2(c)
	case "log10":
		return evalLog10(c)
	}

	panic("should never happen")
}

func MakeCall(e ast.Call, src []byte, s Scope) (Call, error) {
	c := Call{funcName: tok.Text(e.Name, src), args: []Expr{}}

	for i, a := range e.Args {
		expr, err := MakeExpr(a, src, s)
		if err != nil {
			return c, fmt.Errorf("make call: argument %d: %v", i, err)
		}
		c.args = append(c.args, expr)
	}

	err := assertCall(c)
	if err != nil {
		return c, tok.Error{Msg: err.Error(), Toks: []tok.Token{e.Name}}
	}

	return c, nil
}

type Int struct {
	x int64
}

func (i Int) Eval() (val.Value, error) {
	return val.Int(i.x), nil
}

func MakeInt(e ast.Int, src []byte) (Int, error) {
	x, err := strconv.ParseInt(tok.Text(e.X, src), 0, 64)
	if err != nil {
		return Int{}, fmt.Errorf("make int: %v", err)
	}

	return Int{x: x}, nil
}

type List struct {
	exprs []Expr
}

func (l List) Eval() (val.Value, error) {
	vals := []val.Value{}

	for i, expr := range l.exprs {
		v, err := expr.Eval()
		if err != nil {
			return val.Int(0), fmt.Errorf("list evaluation, index %d: %v", i, err)
		}

		vals = append(vals, v)
	}

	return val.List(vals), nil
}

func MakeList(el ast.List, src []byte, s Scope) (List, error) {
	exprs := []Expr{}

	for i, e := range el.Xs {
		e, err := MakeExpr(e, src, s)
		if err != nil {
			return List{}, fmt.Errorf("make expression list: item %d: %v", i, err)
		}
		exprs = append(exprs, e)
	}

	return List{exprs: exprs}, nil
}

type Bool struct {
	x bool
}

func (b Bool) Eval() (val.Value, error) {
	return val.Bool(b.x), nil
}

func MakeBool(e ast.Bool, src []byte) Bool {
	text := tok.Text(e.X, src)
	return Bool{x: text == "true"}
}

type Float struct {
	x float64
}

func (f Float) Eval() (val.Value, error) {
	return val.Float(f.x), nil
}

func MakeFloat(e ast.Float, src []byte) (Float, error) {
	text := tok.Text(e.X, src)
	x, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return Float{}, fmt.Errorf("make float: %v", err)
	}

	return Float{x: x}, nil
}

type DeclaredIdentifier struct {
	x string
	s Scope
}

func (di DeclaredIdentifier) Eval() (val.Value, error) {
	c, err := di.s.GetConst(di.x)
	if err != nil {
		return val.Int(0), fmt.Errorf("evaluating identifier '%s': %v", di.x, err)
	}

	x, err := c.Value.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("evaluating constant identifier '%s': %v", di.x, err)
	}
	return x, nil
}

func MakeDeclaredIdentifier(e ast.Ident, src []byte, s Scope) DeclaredIdentifier {
	return DeclaredIdentifier{x: tok.Text(e.Name, src), s: s}
}

type String struct {
	x string
}

func (s String) Eval() (val.Value, error) {
	return val.Str(s.x), nil
}

func MakeString(e ast.String, src []byte) String {
	txt := tok.Text(e.X, src)
	return String{x: txt[1 : len(txt)-1]}
}

/*
type Subscript struct {
	name string
	idx  Expr
	s    Scope
}

func (s Subscript) Eval() (val.Value, error) {
	idx, err := s.idx.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("subscript index evaluation:%v", err)
	}

	i, ok := idx.(val.Int)
	if !ok {
		return val.Int(0), fmt.Errorf("index must be of type 'integer', current type '%s'", idx.Type())
	}

	sym, err := s.s.GetSymbol(s.name, ConstDef)
	if err != nil {
		return val.Int(0), fmt.Errorf("subscript evaluation, cannot find symbol '%s'", s.name)
	}

	cons, ok := sym.(*Const)
	if !ok {
		return val.Int(0), fmt.Errorf("subscript evaluation, symbol '%s' is not a constant, type '%T'", s.name, sym)
	}

	exprList, ok := cons.Value.(ExpressionList)
	if !ok {
		return val.Int(0),
			fmt.Errorf("subscript evaluation, constant '%s' is not expression list, type '%T'", s.name, cons.Value)
	}

	if int(i) >= len(exprList.exprs) {
		return val.Int(0), fmt.Errorf("list '%s', index %d out of range", s.name, i)
	}

	return exprList.exprs[i].Eval()
}

func MakeSubscript(n ts.Node, s Scope) (Subscript, error) {
	name := n.Child(0).Content()

	idx, err := MakeExpr(n.Child(2), s)
	if err != nil {
		return Subscript{}, fmt.Errorf("make subscript: %v", err)
	}

	return Subscript{name: name, idx: idx, s: s}, nil
}
*/

type Time struct {
	v    Int
	unit string
}

func MakeTime(e ast.Time, src []byte, s Scope) (Time, error) {
	txt := tok.Text(e.X, src)

	aux := strings.Fields(txt)
	intLiteral := aux[0]
	unit := aux[1]

	x, err := strconv.ParseInt(intLiteral, 10, 64)
	if err != nil {
		return Time{}, fmt.Errorf("make time literal: integer literal: %v", err)
	}

	return Time{Int{x}, unit}, nil
}

func (tim Time) Eval() (val.Value, error) {
	v, _ := tim.v.Eval()

	var t val.Time

	switch tim.unit {
	case "s":
		t = val.Time{S: int64(v.(val.Int)), Ns: 0}
	case "ms":
		t = val.Time{S: 0, Ns: 1000000 * int64(v.(val.Int))}
	case "us":
		t = val.Time{S: 0, Ns: 1000 * int64(v.(val.Int))}
	case "ns":
		t = val.Time{S: 0, Ns: int64(v.(val.Int))}
	}

	t.Normalize()
	return t, nil
}

type UnaryOperator uint8

const (
	UnaryPlus = iota
	UnaryMinus
)

type UnaryExpr struct {
	op UnaryOperator
	x  Expr
}

func (ue UnaryExpr) Eval() (val.Value, error) {
	x, err := ue.x.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("unary expression, operand: %v", err)
	}

	if x, ok := x.(val.Int); ok {
		switch ue.op {
		case UnaryPlus:
			return val.Int(x), nil
		case UnaryMinus:
			return val.Int(-x), nil
		default:
			panic("operator not yet supported")
		}
	}

	return val.Int(0), fmt.Errorf("unary expression, unknown operand type '%s'", x.Type())
}

func MakeUnaryExpr(e ast.UnaryExpr, src []byte, s Scope) (UnaryExpr, error) {
	var op UnaryOperator
	switch text := tok.Text(e.Op, src); text {
	case "+":
		op = UnaryPlus
	case "-":
		op = UnaryMinus
	default:
		return UnaryExpr{}, fmt.Errorf("make unary expression: invalid operator %s", text)
	}

	x, err := MakeExpr(e.X, src, s)
	if err != nil {
		return UnaryExpr{}, fmt.Errorf("make unary expression: operand: %v", err)
	}

	return UnaryExpr{op: op, x: x}, nil
}
