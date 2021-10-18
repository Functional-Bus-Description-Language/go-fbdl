package expr

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ts"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
	"strconv"
)

type Expression interface {
	Eval() (val.Value, error)
}

func Make(n ts.Node) (Expression, error) {
	var err error = nil
	var expr Expression

	switch t := n.Type(); t {
	case "binary_operation":
		expr, err = MakeBinaryOperation(n)
	case "decimal_literal":
		expr, err = MakeDecimalLiteral(n)
	case "false":
		expr = MakeFalse()
	case "identifier":
		expr = MakeIdentifier(n)
	case "primary_expression":
		expr, err = MakePrimaryExpression(n)
	case "true":
		expr = MakeTrue()
	case "unary_operation":
		expr, err = MakeUnaryOperation(n)
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
		return val.Bool{}, fmt.Errorf("binary operation: left operand: %v", err)
	}
	right, err := bo.right.Eval()
	if err != nil {
		return val.Bool{}, fmt.Errorf("binary operation: right operand: %v", err)
	}

	if left, ok := left.(val.Int); ok {
		if right, ok := right.(val.Int); ok {
			switch bo.operator {
			case Add:
				return val.Int{V: left.V + right.V}, nil
			default:
				panic("operator not yet supported")
			}
		}
	}

	return val.Bool{}, fmt.Errorf("unknown operand type")
}

func MakeBinaryOperation(n ts.Node) (BinaryOperation, error) {
	left, err := Make(n.Child(0))
	if err != nil {
		return BinaryOperation{}, fmt.Errorf("make binary operation: left operand: %v", err)
	}

	right, err := Make(n.Child(2))
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
	v int32
}

func (dl DecimalLiteral) Eval() (val.Value, error) {
	return val.Int{V: dl.v}, nil
}

func MakeDecimalLiteral(n ts.Node) (DecimalLiteral, error) {
	v, err := strconv.ParseInt(n.Content(), 10, 32)
	if err != nil {
		return DecimalLiteral{}, fmt.Errorf("make decimal literal: %v", err)
	}

	return DecimalLiteral{v: int32(v)}, nil
}

type False struct{}

func (f False) Eval() (val.Value, error) {
	return val.Bool{V: false}, nil
}

func MakeFalse() False {
	return False{}
}

type Identifier struct {
	v string
}

func (i Identifier) Eval() (val.Value, error) {
	// TODO: implement
	return val.Bool{V: false}, nil
}

func MakeIdentifier(n ts.Node) Identifier {
	return Identifier{v: n.Content()}
}

type PrimaryExpression struct {
	v Expression
}

func (pe PrimaryExpression) Eval() (val.Value, error) {
	v, err := pe.v.Eval()
	if err != nil {
		return val.Bool{}, fmt.Errorf("primary expression: %v", err)
	}

	return v, nil
}

func MakePrimaryExpression(n ts.Node) (PrimaryExpression, error) {
	v, err := Make(n.Child(0))
	if err != nil {
		return PrimaryExpression{}, fmt.Errorf("make primary expression: %v", err)
	}

	return PrimaryExpression{v: v}, nil
}

type True struct{}

func (t True) Eval() (val.Value, error) {
	return val.Bool{V: true}, nil
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
		return val.Bool{}, fmt.Errorf("unary operation: operand: %v", err)
	}

	if operand, ok := operand.(val.Int); ok {
		switch uo.operator {
		case UnaryPlus:
			return val.Int{V: operand.V}, nil
		case UnaryMinus:
			return val.Int{V: -operand.V}, nil
		default:
			panic("operator not yet supported")
		}
	}

	return val.Bool{}, fmt.Errorf("unknown operand type")
}

func MakeUnaryOperation(n ts.Node) (UnaryOperation, error) {
	var operator UnaryOperator
	switch op := n.Child(0).Content(); op {
	case "+":
		operator = UnaryPlus
	case "-":
		operator = UnaryMinus
	default:
		return UnaryOperation{}, fmt.Errorf("make unary operation: invalid operator %s", op)
	}

	operand, err := Make(n.Child(1))
	if err != nil {
		return UnaryOperation{}, fmt.Errorf("make unary operation: operand: %v", err)
	}

	return UnaryOperation{operator: operator, operand: operand}, nil
}

type ZeroLiteral struct {
	v int32
}

func (zl ZeroLiteral) Eval() (val.Value, error) {
	return val.Int{V: zl.v}, nil
}

func MakeZeroLiteral() ZeroLiteral {
	return ZeroLiteral{v: 0}
}