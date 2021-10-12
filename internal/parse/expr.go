package parse

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ts"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/value"
	"strconv"
)

type Expression interface {
	Value() (value.Value, error)
}

func MakeExpression(n ts.Node) (Expression, error) {
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

func (bo BinaryOperation) Value() (value.Value, error) {
	left, err := bo.left.Value()
	if err != nil {
		return value.Bool{}, fmt.Errorf("binary operation: left operand: %v", err)
	}
	right, err := bo.right.Value()
	if err != nil {
		return value.Bool{}, fmt.Errorf("binary operation: right operand: %v", err)
	}

	if left, ok := left.(value.Integer); ok {
		if right, ok := right.(value.Integer); ok {
			switch bo.operator {
			case Add:
				return value.Integer{V: left.V + right.V}, nil
			default:
				panic("operator not yet supported")
			}
		}
	}

	return value.Bool{}, fmt.Errorf("unknown operand type")
}

func MakeBinaryOperation(n ts.Node) (BinaryOperation, error) {
	left, err := MakeExpression(n.Child(0))
	if err != nil {
		return BinaryOperation{}, fmt.Errorf("make binary operation: left operand: %v", err)
	}

	right, err := MakeExpression(n.Child(2))
	if err != nil {
		return BinaryOperation{}, fmt.Errorf("make binary operation: right operand: %v", err)
	}

	var operator BinaryOperator
	switch op := n.Child(1).Content(); op {
	case "+":
		operator = Add
	default:
		return BinaryOperation{}, fmt.Errorf("make binary_operation: invalid operand %s", op)
	}

	return BinaryOperation{left: left, right: right, operator: operator}, nil
}

type DecimalLiteral struct {
	v int32
}

func (dl DecimalLiteral) Value() (value.Value, error) {
	return value.Integer{V: dl.v}, nil
}

func MakeDecimalLiteral(n ts.Node) (DecimalLiteral, error) {
	v, err := strconv.ParseInt(n.Content(), 10, 32)
	if err != nil {
		return DecimalLiteral{}, fmt.Errorf("make decimal literal: %v", err)
	}

	return DecimalLiteral{v: int32(v)}, nil
}

type False struct{}

func (f False) Value() (value.Value, error) {
	return value.Bool{V: false}, nil
}

func MakeFalse() False {
	return False{}
}

type Identifier struct {
	v string
}

func (i Identifier) Value() (value.Value, error) {
	// TODO: implement
	return value.Bool{V: false}, nil
}

func MakeIdentifier(n ts.Node) Identifier {
	return Identifier{v: n.Content()}
}

type PrimaryExpression struct {
	v Expression
}

func (pe PrimaryExpression) Value() (value.Value, error) {
	v, err := pe.v.Value()
	if err != nil {
		return value.Bool{}, fmt.Errorf("primary expression: %v", err)
	}

	return v, nil
}

func MakePrimaryExpression(n ts.Node) (PrimaryExpression, error) {
	v, err := MakeExpression(n.Child(0))
	if err != nil {
		return PrimaryExpression{}, fmt.Errorf("make primary expression: %v", err)
	}

	return PrimaryExpression{v: v}, nil
}

type True struct{}

func (t True) Value() (value.Value, error) {
	return value.Bool{V: true}, nil
}

func MakeTrue() True {
	return True{}
}

type ZeroLiteral struct {
	v int32
}

func (zl ZeroLiteral) Value() (value.Value, error) {
	return value.Integer{V: zl.v}, nil
}

func MakeZeroLiteral() ZeroLiteral {
	return ZeroLiteral{v: 0}
}
