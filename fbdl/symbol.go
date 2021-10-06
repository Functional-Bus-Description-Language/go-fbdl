package fbdl

import (
	"fmt"
)

type Symbol interface {
	Name() string
	LineNumber() uint32
}

type common struct {
	File       *File
	Id         uint32
	lineNumber uint32
	name       string
}

func (c common) Name() string {
	return c.name
}

func (c common) LineNumber() uint32 {
	return c.lineNumber
}

type Constant struct {
	common
	value Expression
}

type ElementInstantiationType uint8

const (
	Anonymous ElementInstantiationType = iota
	Definitive
)

type ElementType uint8

const (
	Block ElementType = iota
	Bus
	Config
	Func
	Mask
	Param
	Status
)

func ToElementType(s string) (ElementType, error) {
	var t ElementType

	switch s {
	case "block":
		t = Block
	case "bus":
		t = Bus
	case "config":
		t = Config
	case "func":
		t = Func
	case "mask":
		t = Mask
	case "param":
		t = Param
	case "status":
		t = Status
	default:
		return Block, fmt.Errorf("invalid element type %s", s)
	}

	return t, nil
}

/*
func IsValidElementName(s string) error {
	switch s {
	case
		"block",
		"bus",
		"config",
		"func",
		"mask",
		"param",
		"status":
		return fmt.Errorf("element name can not be the same as element type keyword")
	}

	return nil
}
*/

type Element struct {
	common
	IsArray           bool
	Count             Expression
	Parent            *Symbol
	Type              ElementType
	InstantiationType ElementInstantiationType
}
