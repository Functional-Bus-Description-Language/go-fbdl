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

type Integer struct {
	v int32
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
	Value() Value
}

func MakeExpression(n Node) (Expression, error) {
	var err error = nil
	var expr Expression

	switch t := n.Type(); t {
	case "decimal_literal":
		expr = MakeDecimalLiteral(n)
	case "primary_expression":
		expr, err = MakePrimaryExpression(n)
	default:
		//var dummy DecimalLiteral
		return DecimalLiteral{}, fmt.Errorf("unsupported expression type '%s'", t)
	}

	return expr, err
}

type DecimalLiteral struct {
	v int32
}

func (d DecimalLiteral) Value() Value {
	return Integer{v: d.v}
}

func MakeDecimalLiteral(n Node) DecimalLiteral {
	v, err := strconv.ParseInt(n.Content(), 10, 32)
	if err != nil {
		panic(err)
	}

	return DecimalLiteral{v: int32(v)}
}

type PrimaryExpression struct {
	v Expression
}

func (pe PrimaryExpression) Value() Value {
	return pe.v.Value()
}

func MakePrimaryExpression(n Node) (PrimaryExpression, error) {
	v, err := MakeExpression(n.Child(0))
	if err != nil {
		panic(err)
	}

	return PrimaryExpression{v: v}, nil
}
