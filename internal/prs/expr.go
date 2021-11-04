package prs

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ts"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
	"strconv"
)

type Expression interface {
	Eval() (val.Value, error)
}

func MakeExpression(n ts.Node, s Searchable) (Expression, error) {
	var err error = nil
	var expr Expression

	switch t := n.Type(); t {
	case "binary_operation":
		expr, err = MakeBinaryOperation(n, s)
	case "decimal_literal":
		expr, err = MakeDecimalLiteral(n)
	case "expression_list":
		expr, err = MakeExpressionList(n, s)
	case "false":
		expr = MakeFalse()
	case "identifier":
		expr = MakeIdentifier(n, s)
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
	left, right Expression
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
			default:
				panic("operator not yet supported")
			}
		}
	}

	return val.Int(0), fmt.Errorf("unknown operand type")
}

func MakeBinaryOperation(n ts.Node, s Searchable) (BinaryOperation, error) {
	left, err := MakeExpression(n.Child(0), s)
	if err != nil {
		return BinaryOperation{}, fmt.Errorf("make binary operation: left operand: %v", err)
	}

	right, err := MakeExpression(n.Child(2), s)
	if err != nil {
		return BinaryOperation{}, fmt.Errorf("make binary operation: right operand: %v", err)
	}

	var operator BinaryOperator
	switch op := n.Child(1).Content(); op {
	case "+":
		operator = Add
	default:
		return BinaryOperation{}, fmt.Errorf("make binary operation: invalid operator %s", op)
	}

	return BinaryOperation{left: left, right: right, operator: operator}, nil
}

type DecimalLiteral struct {
	v int64
}

func (dl DecimalLiteral) Eval() (val.Value, error) {
	return val.Int(dl.v), nil
}

func MakeDecimalLiteral(n ts.Node) (DecimalLiteral, error) {
	v, err := strconv.ParseInt(n.Content(), 10, 32)
	if err != nil {
		return DecimalLiteral{}, fmt.Errorf("make decimal literal: %v", err)
	}

	return DecimalLiteral{v: v}, nil
}

type ExpressionList struct {
	exprs []Expression
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

	return val.List{Items: vals}, nil
}

func MakeExpressionList(n ts.Node, s Searchable) (ExpressionList, error) {
	exprs := []Expression{}

	itemIdx := 0
	for i := 1; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "," || t == "]" {
			continue
		}

		e, err := MakeExpression(nc, s)
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

type Identifier struct {
	v string
	s Searchable
}

func (i Identifier) Eval() (val.Value, error) {
	id, err := i.s.GetSymbol(i.v)
	if err != nil {
		return val.Int(0), fmt.Errorf("evaluating identifier '%s': %v", i.v, err)
	}

	if c, ok := id.(*Constant); ok {
		v, err := c.Value.Eval()
		if err != nil {
			return val.Int(0), fmt.Errorf("evaluating constant identifier '%s': %v", i.v, err)
		}
		return v, nil
	} else {
		panic("not yet implemented")
	}
}

func MakeIdentifier(n ts.Node, s Searchable) Identifier {
	return Identifier{v: n.Content(), s: s}
}

type PrimaryExpression struct {
	v Expression
}

func (pe PrimaryExpression) Eval() (val.Value, error) {
	v, err := pe.v.Eval()
	if err != nil {
		return val.Int(0), fmt.Errorf("primary expression: %v", err)
	}

	return v, nil
}

func MakePrimaryExpression(n ts.Node, s Searchable) (PrimaryExpression, error) {
	v, err := MakeExpression(n.Child(0), s)
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
	return StringLiteral{v: n.Content()}
}

type Subscript struct {
	name string
	idx  Expression
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

	sym, err := s.s.GetSymbol(s.name)
	if err != nil {
		return val.Int(0), fmt.Errorf("subscript evaluation, cannot find symbol '%s'", s.name)
	}

	cons, ok := sym.(*Constant)
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

	idx, err := MakeExpression(n.Child(2), s)
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
	operand  Expression
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

	operand, err := MakeExpression(n.Child(1), s)
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
