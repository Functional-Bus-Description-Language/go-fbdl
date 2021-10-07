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

func (e ElementType) String() string {
	switch e {
	case Block:
		return "block"
	case Bus:
		return "bus"
	case Config:
		return "config"
	case Func:
		return "func"
	case Mask:
		return "mask"
	case Param:
		return "param"
	case Status:
		return "status"
	default:
		panic("invalid element type")
	}
}

func IsValidProperty(e ElementType, p string) bool {
	validProps := map[ElementType][]string{
		Block:  []string{"doc"},
		Bus:    []string{"doc", "masters", "width"},
		Config: []string{"atomic", "default", "doc", "groups", "range", "once", "width"},
		Func:   []string{"doc"},
		Mask:   []string{"atomic", "default", "doc", "groups", "width"},
		Param:  []string{"default", "doc", "range", "width"},
		Status: []string{"atomic", "doc", "groups", "once", "width"},
	}

	if list, ok := validProps[e]; ok {
		for i, _ := range list {
			if p == list[i] {
				return true
			}
		}
	}

	return false
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

// Parameter represents parameter in the parameter list, not 'param' element.
type Parameter struct {
	Name            string
	HasDefaultValue bool
	DefaultValue    Expression
}

// Argument represents argument in the argument list.
type Argument struct {
	HasName bool
	Name    string
	Value   Expression
}

type Property struct {
	LineNumber uint32
	Value      Expression
}

type Element struct {
	common
	IsArray           bool
	Count             Expression
	Parent            *Symbol
	Type              ElementType
	InstantiationType ElementInstantiationType
	Properties        map[string]Property
	Symbols           map[string]Symbol
}
