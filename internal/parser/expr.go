package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

type Expr interface {
	Eval() (types.Value, error)
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
	case ast.QualIdent:
		expr = MakeQualifiedIdentifier(e, src, s)
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
	op token.Operator
	y  Expr
}

func (be BinaryExpr) Eval() (types.Value, error) {
	x, err := be.x.Eval()
	if err != nil {
		return nil, err
	}
	y, err := be.y.Eval()
	if err != nil {
		return nil, err
	}
	op := be.op // Operator

	var v types.Value

	switch x := x.(type) {
	case types.Int:
		switch op.(type) {
		case token.Add:
			switch y := y.(type) {
			case types.Int:
				v = x + y
			}
		case token.Sub:
			switch y := y.(type) {
			case types.Int:
				v = x - y
			}
		case token.Mul:
			switch y := y.(type) {
			case types.Int:
				v = x * y
			}
		case token.Div:
			switch y := y.(type) {
			case types.Int:
				if x%y == 0 {
					v = x / y
				} else {
					v = types.Float(float64(x) / float64(y))
				}
			}
		case token.Rem:
			switch y := y.(type) {
			case types.Int:
				v = x % y
			}
		case token.Exp:
			switch y := y.(type) {
			case types.Int:
				v = types.Int(int64(math.Pow(float64(x), float64(y))))
			}
		case token.LShift:
			switch y := y.(type) {
			case types.Int:
				if y < 0 {
					return nil, token.Error{
						Msg:  fmt.Sprintf("negative value of left shift %d", y),
						Toks: []token.Token{op},
					}
				}
				v = x << y
			default:
				return nil, token.Error{
					Msg:  fmt.Sprintf("right operand of left shift must be of type integer, current type %s", y.Type()),
					Toks: []token.Token{be.ast.Y.Tok()},
				}
			}
		case token.RShift:
			switch y := y.(type) {
			case types.Int:
				if y < 0 {
					return nil, token.Error{
						Msg:  fmt.Sprintf("negative value of right shift %d", y),
						Toks: []token.Token{op},
					}
				}
				v = x >> y
			default:
				return nil, token.Error{
					Msg:  fmt.Sprintf("right operand of right shift must be of type integer, current type %s", y.Type()),
					Toks: []token.Token{be.ast.Y.Tok()},
				}
			}
		case token.Colon:
			switch y := y.(type) {
			case types.Int:
				v = types.SingleRange{Start: int64(x), End: int64(y)}
			default:
				return nil, token.Error{
					Msg:  fmt.Sprintf("right bound of range must be of type integer, current type %s", y.Type()),
					Toks: []token.Token{be.ast.Y.Tok()},
				}
			}
		}
	case types.Range:
		switch op.(type) {
		case token.Colon:
			return nil, token.Error{
				Msg:  "left bound of range must be of type integer, current type range",
				Toks: []token.Token{be.ast.X.Tok()},
			}
		}
	}

	if v != nil {
		return v, nil
	}

	return nil, token.Error{
		Msg: fmt.Sprintf(
			"unimplemented binary expression evaluation for %s operator, left operand type %s, right operand type %s, please report this error on %s",
			op.Name(), x.Type(), y.Type(), util.RepoIssueUrl,
		),
		Toks: []token.Token{op},
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
	x types.BitStr
}

func (bs BitString) Eval() (types.Value, error) {
	return bs.x, nil
}

func MakeBitString(e ast.BitString, src []byte) (BitString, error) {
	x, err := types.MakeBitStr(token.Text(e.X, src))
	if err != nil {
		return BitString{}, fmt.Errorf("make bit string: %v", err)
	}

	return BitString{x: x}, nil
}

type Call struct {
	funcName string
	args     []Expr
}

func (c Call) Eval() (types.Value, error) {
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
	c := Call{funcName: token.Text(e.Name, src), args: []Expr{}}

	for i, a := range e.Args {
		expr, err := MakeExpr(a, src, s)
		if err != nil {
			return c, fmt.Errorf("make call: argument %d: %v", i, err)
		}
		c.args = append(c.args, expr)
	}

	err := assertCall(c)
	if err != nil {
		return c, token.Error{Msg: err.Error(), Toks: []token.Token{e.Name}}
	}

	return c, nil
}

type Int struct {
	x int64
}

func (i Int) Eval() (types.Value, error) {
	return types.Int(i.x), nil
}

func MakeInt(e ast.Int, src []byte) (Int, error) {
	x, err := strconv.ParseInt(token.Text(e.X, src), 0, 64)
	if err != nil {
		return Int{}, fmt.Errorf("make int: %v", err)
	}

	return Int{x: x}, nil
}

type List struct {
	exprs []Expr
}

func (l List) Eval() (types.Value, error) {
	vals := []types.Value{}

	for i, expr := range l.exprs {
		v, err := expr.Eval()
		if err != nil {
			return types.Int(0), fmt.Errorf("list evaluation, index %d: %v", i, err)
		}

		vals = append(vals, v)
	}

	return types.List(vals), nil
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

func (b Bool) Eval() (types.Value, error) {
	return types.Bool(b.x), nil
}

func MakeBool(e ast.Bool, src []byte) Bool {
	text := token.Text(e.X, src)
	return Bool{x: text == "true"}
}

type Float struct {
	x float64
}

func (f Float) Eval() (types.Value, error) {
	return types.Float(f.x), nil
}

func MakeFloat(e ast.Float, src []byte) (Float, error) {
	text := token.Text(e.X, src)
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

func (di DeclaredIdentifier) Eval() (types.Value, error) {
	c, err := di.s.GetConst(di.x)
	if err != nil {
		return types.Int(0), fmt.Errorf("evaluating identifier '%s': %v", di.x, err)
	}

	x, err := c.Value.Eval()
	if err != nil {
		return types.Int(0), fmt.Errorf("evaluating constant identifier '%s': %v", di.x, err)
	}
	return x, nil
}

func MakeDeclaredIdentifier(e ast.Ident, src []byte, s Scope) DeclaredIdentifier {
	return DeclaredIdentifier{x: token.Text(e.Name, src), s: s}
}

type QualifiedIdentifier struct {
	x string
	s Scope
}

func (qi QualifiedIdentifier) Eval() (types.Value, error) {
	c, err := qi.s.GetConst(qi.x)
	if err != nil {
		return types.Int(0), fmt.Errorf("evaluating qualified identifier '%s': %v", qi.x, err)
	}

	x, err := c.Value.Eval()
	if err != nil {
		return types.Int(0), fmt.Errorf("evaluating constant qualified identifier '%s': %v", qi.x, err)
	}
	return x, nil
}

func MakeQualifiedIdentifier(e ast.QualIdent, src []byte, s Scope) QualifiedIdentifier {
	return QualifiedIdentifier{x: token.Text(e.Name, src), s: s}
}

type String struct {
	x string
}

func (s String) Eval() (types.Value, error) {
	return types.Str(s.x), nil
}

func MakeString(e ast.String, src []byte) String {
	txt := token.Text(e.X, src)
	return String{x: txt[1 : len(txt)-1]}
}

/*
type Subscript struct {
	name string
	idx  Expr
	s    Scope
}

func (s Subscript) Eval() (types.Value, error) {
	idx, err := s.idx.Eval()
	if err != nil {
		return types.Int(0), fmt.Errorf("subscript index evaluation:%v", err)
	}

	i, ok := idx.(types.Int)
	if !ok {
		return types.Int(0), fmt.Errorf("index must be of type 'integer', current type '%s'", idx.Type())
	}

	sym, err := s.s.GetSymbol(s.name, ConstDef)
	if err != nil {
		return types.Int(0), fmt.Errorf("subscript evaluation, cannot find symbol '%s'", s.name)
	}

	cons, ok := sym.(*Const)
	if !ok {
		return types.Int(0), fmt.Errorf("subscript evaluation, symbol '%s' is not a constant, type '%T'", s.name, sym)
	}

	exprList, ok := cons.Value.(ExpressionList)
	if !ok {
		return types.Int(0),
			fmt.Errorf("subscript evaluation, constant '%s' is not expression list, type '%T'", s.name, cons.Value)
	}

	if int(i) >= len(exprList.exprs) {
		return types.Int(0), fmt.Errorf("list '%s', index %d out of range", s.name, i)
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
	txt := token.Text(e.X, src)

	aux := strings.Fields(txt)
	intLiteral := aux[0]
	unit := aux[1]

	x, err := strconv.ParseInt(intLiteral, 10, 64)
	if err != nil {
		return Time{}, fmt.Errorf("make time literal: integer literal: %v", err)
	}

	return Time{Int{x}, unit}, nil
}

func (tim Time) Eval() (types.Value, error) {
	v, _ := tim.v.Eval()

	var t types.Time

	switch tim.unit {
	case "s":
		t = types.Time{S: int64(v.(types.Int)), Ns: 0}
	case "ms":
		t = types.Time{S: 0, Ns: 1000000 * int64(v.(types.Int))}
	case "us":
		t = types.Time{S: 0, Ns: 1000 * int64(v.(types.Int))}
	case "ns":
		t = types.Time{S: 0, Ns: int64(v.(types.Int))}
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

func (ue UnaryExpr) Eval() (types.Value, error) {
	x, err := ue.x.Eval()
	if err != nil {
		return types.Int(0), fmt.Errorf("unary expression, operand: %v", err)
	}

	if x, ok := x.(types.Int); ok {
		switch ue.op {
		case UnaryPlus:
			return x, nil
		case UnaryMinus:
			return -x, nil
		default:
			panic("operator not yet supported")
		}
	}

	return types.Int(0), fmt.Errorf("unary expression, unknown operand type '%s'", x.Type())
}

func MakeUnaryExpr(e ast.UnaryExpr, src []byte, s Scope) (UnaryExpr, error) {
	var op UnaryOperator
	switch text := token.Text(e.Op, src); text {
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
