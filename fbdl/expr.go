package fbdl

import (
	"fmt"
	"strconv"
)

/*
type ValueType uint8

const (
	Bool ValueType = iota
	Integer
	Real
	String
)
*/

type Value interface {
	IsValue() bool
}

type Bool struct {
	V bool
}

func (b Bool) IsValue() bool {
	return true
}

type Integer struct {
	V int32
}

func (i Integer) IsValue() bool {
	return true
}

type String struct {
	v string
}

func (s String) IsValue() bool {
	return true
}

type Expression interface {
	Value() (Value, error)
}

func MakeExpression(n Node) (Expression, error) {
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

func (bo BinaryOperation) Value() (Value, error) {
	left, err := bo.left.Value()
	if err != nil {
		return Bool{}, fmt.Errorf("binary operation: left operand: %v", err)
	}
	right, err := bo.right.Value()
	if err != nil {
		return Bool{}, fmt.Errorf("binary operation: right operand: %v", err)
	}

	if left, ok := left.(Integer); ok {
		if right, ok := right.(Integer); ok {
			switch bo.operator {
			case Add:
				return Integer{V: left.V + right.V}, nil
			default:
				panic("operator not yet supported")
			}
		}
	}

	return Bool{}, fmt.Errorf("unknown operand type")
}

func MakeBinaryOperation(n Node) (BinaryOperation, error) {
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

func (dl DecimalLiteral) Value() (Value, error) {
	return Integer{V: dl.v}, nil
}

func MakeDecimalLiteral(n Node) (DecimalLiteral, error) {
	v, err := strconv.ParseInt(n.Content(), 10, 32)
	if err != nil {
		return DecimalLiteral{}, fmt.Errorf("make decimal literal: %v", err)
	}

	return DecimalLiteral{v: int32(v)}, nil
}

type False struct {
	v bool
}

func (f False) Value() (Value, error) {
	return Bool{V: false}, nil
}

func MakeFalse() False {
	return False{v: false}
}

type Identifier struct {
	v string
}

func (i Identifier) Value() (Value, error) {
	// TODO: implement
	return Bool{V: false}, nil
}

func MakeIdentifier(n Node) Identifier {
	return Identifier{v: n.Content()}
}

type PrimaryExpression struct {
	v Expression
}

func (pe PrimaryExpression) Value() (Value, error) {
	v, err := pe.v.Value()
	if err != nil {
		return Bool{}, fmt.Errorf("primary expression: %v", err)
	}

	return v, nil
}

func MakePrimaryExpression(n Node) (PrimaryExpression, error) {
	v, err := MakeExpression(n.Child(0))
	if err != nil {
		return PrimaryExpression{}, fmt.Errorf("make primary expression: %v", err)
	}

	return PrimaryExpression{v: v}, nil
}

type ZeroLiteral struct {
	v int32
}

func (zl ZeroLiteral) Value() (Value, error) {
	return Integer{V: zl.v}, nil
}

func MakeZeroLiteral() ZeroLiteral {
	return ZeroLiteral{v: 0}
}
