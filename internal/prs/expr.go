package prs

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ts"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type Expr interface {
	Eval() (val.Value, error)
}

func MakeExpr(n ts.Node, s Searchable) (Expr, error) {
	var err error = nil
	var expr Expr

	switch t := n.Type(); t {
	case "binary_operation":
		expr, err = MakeBinaryOperation(n, s)
	case "bit_literal":
		expr, err = MakeBitLiteral(n)
	case "call":
		expr, err = MakeCall(n, s)
	case "declared_identifier":
		expr = MakeDeclaredIdentifier(n, s)
	case "decimal_literal":
		expr, err = MakeDecimalLiteral(n)
	case "expression_list":
		expr, err = MakeExprList(n, s)
	case "false":
		expr = MakeFalse()
	case "float_literal":
		expr, err = MakeFloatLiteral(n)
	case "hex_literal":
		expr, err = MakeHexLiteral(n)
	case "primary_expression":
		expr, err = MakePrimaryExpression(n, s)
	case "string_literal":
		expr = MakeStringLiteral(n)
	case "subscript":
		expr, err = MakeSubscript(n, s)
	case "true":
		expr = MakeTrue()
	case "unary_operation":
		expr, err = MakeUnaryOperation(n, s)
	case "zero_literal":
		expr = MakeZeroLiteral()
	default:
		return DecimalLiteral{}, fmt.Errorf("unsupported expression type '%s'", t)
	}

	return expr, err
}

type BinaryOperator uint8

const (
	Add BinaryOperator = iota
	Subtract
	Multiply
	Divide
	Modulo
	Power
	LeftShift
	RightShift
)

type BinaryOperation struct {
	left, right Expr
	operator    BinaryOperator
}

func (bo BinaryOperation) Eval() (val.Value, error) {
	left, err := bo.left.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("binary operation: left operand: %v", err)
	}
	right, err := bo.right.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("binary operation: right operand: %v", err)
	}

	if left, ok := left.(val.Int); ok {
		if right, ok := right.(val.Int); ok {
			switch bo.operator {
			case Add:
				return val.Int(left + right), nil
			case Subtract:
				return val.Int(left - right), nil
			case Multiply:
				return val.Int(left * right), nil
			case Divide:
				if left%right == 0 {
					return val.Int(left / right), nil
				} else {
					panic("not yet implement, needs float point type")
				}
			case Modulo:
				return val.Int(left % right), nil
			case Power:
				return val.Int(int64(math.Pow(float64(left), float64(right)))), nil
			case LeftShift:
				return val.Int(left << right), nil
			case RightShift:
				return val.Int(left >> right), nil
			default:
				panic("operator not yet supported")
			}
		}
	}

	return val.Int(0), fmt.Errorf("unknown operand type")
}

func MakeBinaryOperation(n ts.Node, s Searchable) (BinaryOperation, error) {
	left, err := MakeExpr(n.Child(0), s)
	if err != nil {
		return BinaryOperation{}, fmt.Errorf("make binary operation: left operand: %v", err)
	}

	right, err := MakeExpr(n.Child(2), s)
	if err != nil {
		return BinaryOperation{}, fmt.Errorf("make binary operation: right operand: %v", err)
	}

	var operator BinaryOperator
	switch op := n.Child(1).Content(); op {
	case "+":
		operator = Add
	case "-":
		operator = Subtract
	case "*":
		operator = Multiply
	case "/":
		operator = Divide
	case "%":
		operator = Modulo
	case "**":
		operator = Power
	case "<<":
		operator = LeftShift
	case ">>":
		operator = RightShift
	default:
		return BinaryOperation{}, fmt.Errorf("make binary operation: invalid operator %s", op)
	}

	return BinaryOperation{left: left, right: right, operator: operator}, nil
}

type BitLiteral struct {
	v val.BitStr
}

func (bl BitLiteral) Eval() (val.Value, error) {
	return bl.v, nil
}

func MakeBitLiteral(n ts.Node) (BitLiteral, error) {
	v, err := val.MakeBitStr(n.Content())
	if err != nil {
		return BitLiteral{}, fmt.Errorf("make bit literal: %v", err)
	}

	return BitLiteral{v: v}, nil
}

type Call struct {
	funcName string
	args     []Expr
}

func (c Call) Eval() (val.Value, error) {
	switch c.funcName {
	case "ceil":
		return evalCeil(c)
	case "log2":
		return evalLog2(c)
	}

	panic("should never happen")
}

func MakeCall(n ts.Node, s Searchable) (Call, error) {
	c := Call{funcName: n.Child(0).Content(), args: []Expr{}}

	argIdx := 0
	for i := 2; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "," || t == ")" {
			continue
		}

		e, err := MakeExpr(nc, s)
		if err != nil {
			return c, fmt.Errorf("make call: argument %d: %v", argIdx, err)
		}
		c.args = append(c.args, e)

		argIdx += 1
	}

	err := assertCall(c)
	if err != nil {
		return c, fmt.Errorf("make call: %v", err)
	}

	return c, nil
}

type DecimalLiteral struct {
	v int64
}

func (dl DecimalLiteral) Eval() (val.Value, error) {
	return val.Int(dl.v), nil
}

func MakeDecimalLiteral(n ts.Node) (DecimalLiteral, error) {
	v, err := strconv.ParseInt(n.Content(), 10, 64)
	if err != nil {
		return DecimalLiteral{}, fmt.Errorf("make decimal literal: %v", err)
	}

	return DecimalLiteral{v: v}, nil
}

type ExpressionList struct {
	exprs []Expr
}

func (el ExpressionList) Eval() (val.Value, error) {
	vals := []val.Value{}

	for i, expr := range el.exprs {
		v, err := expr.Eval()
		if err != nil {
			return val.Int(0), fmt.Errorf("expression list evaluation, index %d: %v", i, err)
		}

		vals = append(vals, v)
	}

	return val.List(vals), nil
}

func MakeExprList(n ts.Node, s Searchable) (ExpressionList, error) {
	exprs := []Expr{}

	itemIdx := 0
	for i := 1; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "," || t == "]" {
			continue
		}

		e, err := MakeExpr(nc, s)
		if err != nil {
			return ExpressionList{}, fmt.Errorf("make expression list: item %d: %v", itemIdx, err)
		}
		exprs = append(exprs, e)

		itemIdx += 1
	}

	return ExpressionList{exprs: exprs}, nil
}

type False struct{}

func (f False) Eval() (val.Value, error) {
	return val.Bool(false), nil
}

func MakeFalse() False {
	return False{}
}

type FloatLiteral struct {
	v float64
}

func (fl FloatLiteral) Eval() (val.Value, error) {
	return val.Float(fl.v), nil
}

func MakeFloatLiteral(n ts.Node) (FloatLiteral, error) {
	v, err := strconv.ParseFloat(n.Content(), 64)
	if err != nil {
		return FloatLiteral{}, fmt.Errorf("make float literal: %v", err)
	}

	return FloatLiteral{v: v}, nil
}

type DeclaredIdentifier struct {
	v string
	s Searchable
}

func (di DeclaredIdentifier) Eval() (val.Value, error) {
	id, err := di.s.GetSymbol(di.v, ConstDef)
	if err != nil {
		return val.Int(0), fmt.Errorf("evaluating identifier '%s': %v", di.v, err)
	}

	if c, ok := id.(*Const); ok {
		v, err := c.Value.Eval()
		if err != nil {
			return val.Int(0), fmt.Errorf("evaluating constant identifier '%s': %v", di.v, err)
		}
		return v, nil
	} else {
		panic("not yet implemented")
	}
}

func MakeDeclaredIdentifier(n ts.Node, s Searchable) DeclaredIdentifier {
	return DeclaredIdentifier{v: n.Content(), s: s}
}

type HexLiteral struct {
	v int64
}

func (hl HexLiteral) Eval() (val.Value, error) {
	return val.Int(hl.v), nil
}

func MakeHexLiteral(n ts.Node) (HexLiteral, error) {
	v, err := strconv.ParseInt(n.Content(), 0, 64)
	if err != nil {
		return HexLiteral{}, fmt.Errorf("make hex literal: %v", err)
	}

	return HexLiteral{v: v}, nil
}

type PrimaryExpression struct {
	v Expr
}

func (pe PrimaryExpression) Eval() (val.Value, error) {
	v, err := pe.v.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("primary expression: %v", err)
	}

	return v, nil
}

func MakePrimaryExpression(n ts.Node, s Searchable) (PrimaryExpression, error) {
	v, err := MakeExpr(n.Child(0), s)
	if err != nil {
		return PrimaryExpression{}, fmt.Errorf("make primary expression: %v", err)
	}

	return PrimaryExpression{v: v}, nil
}

type StringLiteral struct {
	v string
}

func (sl StringLiteral) Eval() (val.Value, error) {
	return val.Str(sl.v), nil
}

func MakeStringLiteral(n ts.Node) StringLiteral {
	return StringLiteral{v: n.Content()[1 : len(n.Content())-1]}
}

type Subscript struct {
	name string
	idx  Expr
	s    Searchable
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

func MakeSubscript(n ts.Node, s Searchable) (Subscript, error) {
	name := n.Child(0).Content()

	idx, err := MakeExpr(n.Child(2), s)
	if err != nil {
		return Subscript{}, fmt.Errorf("make subscript: %v", err)
	}

	return Subscript{name: name, idx: idx, s: s}, nil
}

type True struct{}

func (t True) Eval() (val.Value, error) {
	return val.Bool(true), nil
}

func MakeTrue() True {
	return True{}
}

type UnaryOperator uint8

const (
	UnaryPlus = iota
	UnaryMinus
)

type UnaryOperation struct {
	operator UnaryOperator
	operand  Expr
}

func (uo UnaryOperation) Eval() (val.Value, error) {
	operand, err := uo.operand.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("unary operation: operand: %v", err)
	}

	if operand, ok := operand.(val.Int); ok {
		switch uo.operator {
		case UnaryPlus:
			return val.Int(operand), nil
		case UnaryMinus:
			return val.Int(-operand), nil
		default:
			panic("operator not yet supported")
		}
	}

	return val.Int(0), fmt.Errorf("unknown operand type")
}

func MakeUnaryOperation(n ts.Node, s Searchable) (UnaryOperation, error) {
	var operator UnaryOperator
	switch op := n.Child(0).Content(); op {
	case "+":
		operator = UnaryPlus
	case "-":
		operator = UnaryMinus
	default:
		return UnaryOperation{}, fmt.Errorf("make unary operation: invalid operator %s", op)
	}

	operand, err := MakeExpr(n.Child(1), s)
	if err != nil {
		return UnaryOperation{}, fmt.Errorf("make unary operation: operand: %v", err)
	}

	return UnaryOperation{operator: operator, operand: operand}, nil
}

type ZeroLiteral struct {
	v int64
}

func (zl ZeroLiteral) Eval() (val.Value, error) {
	return val.Int(zl.v), nil
}

func MakeZeroLiteral() ZeroLiteral {
	return ZeroLiteral{v: 0}
}
