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

func MakeExpr(astExpr ast.Expr, file *File, s Scope) (Expr, error) {
	var err error
	var expr Expr

	switch e := astExpr.(type) {
	case ast.BinaryExpr:
		expr, err = MakeBinaryExpr(e, file, s)
	case ast.BitString:
		expr, err = MakeBitString(e, file)
	case ast.Call:
		expr, err = MakeCall(e, file, s)
	case ast.Ident:
		expr = MakeDeclaredIdentifier(e, file, s)
	case ast.QualIdent:
		expr = MakeQualifiedIdentifier(e, file, s)
	case ast.Int:
		expr, err = MakeInt(e, file)
	case ast.List:
		expr, err = MakeList(e, file, s)
	case ast.Bool:
		expr = MakeBool(e, file)
	case ast.Float:
		expr, err = MakeFloat(e, file)
	case ast.String:
		expr = MakeString(e, file)
	case ast.Time:
		expr, err = MakeTime(e, file, s)
	case ast.UnaryExpr:
		expr, err = MakeUnaryExpr(e, file, s)
	case nil:
		return nil, nil
	default:
		panic(fmt.Sprintf("unimplemented type %T", astExpr))
	}

	return expr, err
}

type BinaryExpr struct {
	file *File
	ast  ast.BinaryExpr

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
					return nil, Error{token.Error{
						Msg:  fmt.Sprintf("negative value of left shift %d", y),
						Toks: []token.Token{op},
					}, be.file}
				}
				v = x << y
			default:
				return nil, Error{token.Error{
					Msg:  fmt.Sprintf("right operand of left shift must be of type integer, current type %s", y.Type()),
					Toks: []token.Token{be.ast.Y.Tok()},
				}, be.file}
			}
		case token.RShift:
			switch y := y.(type) {
			case types.Int:
				if y < 0 {
					return nil, Error{token.Error{
						Msg:  fmt.Sprintf("negative value of right shift %d", y),
						Toks: []token.Token{op},
					}, be.file}
				}
				v = x >> y
			default:
				return nil, Error{token.Error{
					Msg:  fmt.Sprintf("right operand of right shift must be of type integer, current type %s", y.Type()),
					Toks: []token.Token{be.ast.Y.Tok()},
				}, be.file}
			}
		case token.Colon:
			switch y := y.(type) {
			case types.Int:
				v = types.SingleRange{Start: int64(x), End: int64(y)}
			default:
				return nil, Error{token.Error{
					Msg:  fmt.Sprintf("right bound of range must be of type integer, current type %s", y.Type()),
					Toks: []token.Token{be.ast.Y.Tok()},
				}, be.file}
			}
		}
	case types.Range:
		switch op.(type) {
		case token.Colon:
			return nil, Error{token.Error{
				Msg:  "left bound of range must be of type integer, current type range",
				Toks: []token.Token{be.ast.X.Tok()},
			}, be.file}
		}
	}

	if v != nil {
		return v, nil
	}

	return nil, Error{token.Error{
		Msg: fmt.Sprintf(
			"unimplemented binary expression evaluation for %s operator, left operand type %s, right operand type %s, please report this error on %s",
			op.Name(), x.Type(), y.Type(), util.RepoIssueUrl,
		),
		Toks: []token.Token{op},
	}, be.file}
}

func MakeBinaryExpr(be ast.BinaryExpr, file *File, s Scope) (BinaryExpr, error) {
	x, err := MakeExpr(be.X, file, s)
	if err != nil {
		return BinaryExpr{}, err
	}

	y, err := MakeExpr(be.Y, file, s)
	if err != nil {
		return BinaryExpr{}, err
	}

	return BinaryExpr{file: file, ast: be, x: x, op: be.Op, y: y}, nil
}

type BitString struct {
	x types.BitStr
}

func (bs BitString) Eval() (types.Value, error) {
	return bs.x, nil
}

func MakeBitString(e ast.BitString, file *File) (BitString, error) {
	src := file.Src
	x, err := types.MakeBitStr(token.Text(e.X, src))
	if err != nil {
		return BitString{}, Error{token.Error{Msg: fmt.Sprintf("make bit string: %v", err), Toks: []token.Token{e.X}}, file}
	}

	return BitString{x: x}, nil
}

type Call struct {
	file     *File
	funcName string
	args     []Expr
}

func (c Call) Eval() (types.Value, error) {
	var value types.Value
	var err error
	switch c.funcName {
	case "bool":
		value, err = evalBool(c)
	case "ceil":
		value, err = evalCeil(c)
	case "floor":
		value, err = evalFloor(c)
	case "log2":
		value, err = evalLog2(c)
	case "log10":
		value, err = evalLog10(c)
	default:
		panic("should never happen")
	}
	if e, ok := err.(token.Error); ok {
		return nil, Error{e, c.file}
	}
	return value, err
}

func MakeCall(e ast.Call, file *File, s Scope) (Call, error) {
	src := file.Src
	c := Call{file: file, funcName: token.Text(e.Name, src), args: []Expr{}}

	for _, a := range e.Args {
		expr, err := MakeExpr(a, file, s)
		if err != nil {
			return c, err
		}
		c.args = append(c.args, expr)
	}

	err := assertCall(c)
	if err != nil {
		return c, Error{token.Error{Msg: err.Error(), Toks: []token.Token{e.Name}}, file}
	}

	return c, nil
}

type Int struct {
	x int64
}

func (i Int) Eval() (types.Value, error) {
	return types.Int(i.x), nil
}

func MakeInt(e ast.Int, file *File) (Int, error) {
	src := file.Src
	x, err := strconv.ParseInt(token.Text(e.X, src), 0, 64)
	if err != nil {
		return Int{}, Error{token.Error{Msg: fmt.Sprintf("make int: %v", err), Toks: []token.Token{e.X}}, file}
	}

	return Int{x: x}, nil
}

type List struct {
	file  *File
	exprs []Expr
}

func (l List) Eval() (types.Value, error) {
	vals := []types.Value{}

	for _, expr := range l.exprs {
		v, err := expr.Eval()
		if err != nil {
			return types.Int(0), err
		}

		vals = append(vals, v)
	}

	return types.List(vals), nil
}

func MakeList(el ast.List, file *File, s Scope) (List, error) {
	exprs := []Expr{}

	for _, e := range el.Xs {
		e, err := MakeExpr(e, file, s)
		if err != nil {
			return List{}, err
		}
		exprs = append(exprs, e)
	}

	return List{file: file, exprs: exprs}, nil
}

type Bool struct {
	x bool
}

func (b Bool) Eval() (types.Value, error) {
	return types.Bool(b.x), nil
}

func MakeBool(e ast.Bool, file *File) Bool {
	src := file.Src
	text := token.Text(e.X, src)
	return Bool{x: text == "true"}
}

type Float struct {
	x float64
}

func (f Float) Eval() (types.Value, error) {
	return types.Float(f.x), nil
}

func MakeFloat(e ast.Float, file *File) (Float, error) {
	src := file.Src
	text := token.Text(e.X, src)
	x, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return Float{}, Error{token.Error{Msg: fmt.Sprintf("make float: %v", err), Toks: []token.Token{e.X}}, file}
	}

	return Float{x: x}, nil
}

type DeclaredIdentifier struct {
	ast   ast.Ident
	scope Scope
	x     string
}

func (di DeclaredIdentifier) Eval() (types.Value, error) {
	c, err := di.scope.GetConst(di.x)
	if err != nil {
		return types.Int(0), token.Error{
			Msg:  fmt.Sprintf("%v", err),
			Toks: []token.Token{di.ast.Name},
		}
	}

	x, err := c.Value.Eval()
	if err != nil {
		return types.Int(0), err
	}
	return x, nil
}

func MakeDeclaredIdentifier(e ast.Ident, file *File, s Scope) DeclaredIdentifier {
	return DeclaredIdentifier{
		ast:   e,
		x:     token.Text(e.Name, file.Src),
		scope: s,
	}
}

type QualifiedIdentifier struct {
	ast   ast.QualIdent
	scope Scope
	x     string
}

func (qi QualifiedIdentifier) Eval() (types.Value, error) {
	c, err := qi.scope.GetConst(qi.x)
	if err != nil {
		return types.Int(0), token.Error{
			Msg:  fmt.Sprintf("%v", err),
			Toks: []token.Token{qi.ast.Name},
		}
	}

	x, err := c.Value.Eval()
	if err != nil {
		return types.Int(0), err
	}
	return x, nil
}

func MakeQualifiedIdentifier(e ast.QualIdent, file *File, s Scope) QualifiedIdentifier {
	return QualifiedIdentifier{
		ast:   e,
		scope: s,
		x:     token.Text(e.Name, file.Src),
	}
}

type String struct {
	x string
}

func (s String) Eval() (types.Value, error) {
	return types.Str(s.x), nil
}

func MakeString(e ast.String, file *File) String {
	src := file.Src
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

func MakeTime(e ast.Time, file *File, s Scope) (Time, error) {
	src := file.Src
	txt := token.Text(e.X, src)

	intLiteral, unit, ok := strings.Cut(txt, "'")
	if !ok {
		return Time{}, Error{token.Error{Msg: "make time literal: missing apostrophe", Toks: []token.Token{e.X}}, file}
	}

	x, err := strconv.ParseInt(intLiteral, 10, 64)
	if err != nil {
		return Time{}, Error{token.Error{Msg: fmt.Sprintf("make time literal: integer literal: %v", err), Toks: []token.Token{e.X}}, file}
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
	tok  token.Token
	file *File
	op   UnaryOperator
	x    Expr
}

func (ue UnaryExpr) Eval() (types.Value, error) {
	x, err := ue.x.Eval()
	if err != nil {
		return types.Int(0), err
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

	return types.Int(0), Error{token.Error{Msg: fmt.Sprintf("unary expression, unknown operand type '%s'", x.Type()), Toks: []token.Token{ue.tok}}, ue.file}
}

func MakeUnaryExpr(e ast.UnaryExpr, file *File, s Scope) (UnaryExpr, error) {
	src := file.Src
	var op UnaryOperator
	switch text := token.Text(e.Op, src); text {
	case "+":
		op = UnaryPlus
	case "-":
		op = UnaryMinus
	default:
		return UnaryExpr{}, Error{token.Error{Msg: fmt.Sprintf("make unary expression: invalid operator %s", text)}, file}
	}

	x, err := MakeExpr(e.X, file, s)
	if err != nil {
		return UnaryExpr{}, err
	}

	return UnaryExpr{file: file, tok: e.Op, op: op, x: x}, nil
}
