package fbdl

import (
	_ "fmt"
)

type Symbol interface {
	Name() string
	LineNumber() uint32
	SetParent(s Symbol)
	Parent() Symbol
}

type common struct {
	File       *File
	lineNumber uint32
	name       string
	parent     Symbol
}

func (c common) Name() string {
	return c.name
}

func (c common) LineNumber() uint32 {
	return c.lineNumber
}

func (c *common) SetParent(s Symbol) {
	if c.parent == nil {
		c.parent = s
	} else {
		panic("should never happen")
	}
}

func (c common) Parent() Symbol {
	return c.parent
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

func IsBaseType(t string) bool {
	base_types := [...]string{"block", "bus", "config", "func", "mask", "param", "status"}

	for i, _ := range base_types {
		if t == base_types[i] {
			return true
		}
	}

	return false
}

func IsValidProperty(t string, p string) bool {
	validProps := map[string][]string{
		"block":  []string{"doc"},
		"bus":    []string{"doc", "masters", "width"},
		"config": []string{"atomic", "default", "doc", "groups", "range", "once", "width"},
		"func":   []string{"doc"},
		"mask":   []string{"atomic", "default", "doc", "groups", "width"},
		"param":  []string{"default", "doc", "range", "width"},
		"status": []string{"atomic", "doc", "groups", "once", "width"},
	}

	if list, ok := validProps[t]; ok {
		for i, _ := range list {
			if p == list[i] {
				return true
			}
		}
	} else {
		panic("should never happen")
	}

	return false
}

// Argument represents argument in the argument list.
type Argument struct {
	HasName bool
	Name    string
	Value   Expression
}

type Element struct {
	common
	IsArray           bool
	Count             Expression
	Type              string
	InstantiationType ElementInstantiationType
	Properties        map[string]Property
	Symbols           map[string]Symbol
	Arguments         []Argument
}

type Import struct {
	Path       string
	ImportName string
	Package    *Package
}

// Parameter represents parameter in the parameter list, not 'param' element.
type Parameter struct {
	Name            string
	HasDefaultValue bool
	DefaultValue    Expression
}

type Property struct {
	LineNumber uint32
	Value      Expression
}

type Type struct {
	common
	Parameters []Parameter
	Arguments  []Argument
	Type       string
	Properties map[string]Property
	Symbols    map[string]Symbol
}
