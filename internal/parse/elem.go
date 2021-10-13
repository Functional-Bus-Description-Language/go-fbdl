package parse

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/expr"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/value"
	"strings"
)

// Argument represents argument in the argument list.
type Argument struct {
	HasName bool
	Name    string
	Value   expr.Expression
}

// Parameter represents parameter in the parameter list, not 'param' element.
type Parameter struct {
	Name            string
	HasDefaultValue bool
	DefaultValue    expr.Expression
}

type Property struct {
	LineNumber uint32
	Value      expr.Expression
}

type ElementInstantiationType uint8

const (
	Anonymous ElementInstantiationType = iota
	Definitive
)

type Element interface {
	Symbol
	Type() string
	Args() []Argument
	Params() []Parameter
	SetResolvedArgs(args map[string]value.Value)
	ResolvedArgs() map[string]value.Value
	Properties() map[string]Property
	Symbols() map[string]Symbol
}

type ElementDefinition struct {
	base

	type_             string
	InstantiationType ElementInstantiationType
	IsArray           bool
	Count             expr.Expression

	properties map[string]Property
	symbols    map[string]Symbol

	params       []Parameter
	args         []Argument
	resolvedArgs map[string]value.Value
}

func (e ElementDefinition) Type() string {
	return e.type_
}

func (e *ElementDefinition) GetSymbol(s string) (Symbol, error) {
	if strings.Contains(s, ".") {
		panic("To be implemented")
	}

	if sym, ok := e.symbols[s]; ok {
		return sym, nil
	}

	if e.parent != nil {
		return e.parent.GetSymbol(s)
	}

	return e.file.GetSymbol(s)
}

func (e ElementDefinition) Args() []Argument {
	return e.args
}

func (e ElementDefinition) Params() []Parameter {
	return e.params
}

func (e *ElementDefinition) SetResolvedArgs(ra map[string]value.Value) {
	e.resolvedArgs = ra
}

func (e ElementDefinition) ResolvedArgs() map[string]value.Value {
	return e.resolvedArgs
}

func (e ElementDefinition) Properties() map[string]Property {
	return e.properties
}

func (e ElementDefinition) Symbols() map[string]Symbol {
	return e.symbols
}

type TypeDefinition struct {
	base

	type_      string
	properties map[string]Property
	symbols    map[string]Symbol

	params       []Parameter
	args         []Argument
	resolvedArgs map[string]value.Value
}

func (t *TypeDefinition) GetSymbol(s string) (Symbol, error) {
	if strings.Contains(s, ".") {
		panic("To be implemented")
	}

	if sym, ok := t.symbols[s]; ok {
		return sym, nil
	}

	if t.parent != nil {
		return t.parent.GetSymbol(s)
	}

	return t.file.GetSymbol(s)
}

func (t TypeDefinition) Type() string {
	return t.type_
}

func (t TypeDefinition) Args() []Argument {
	return t.args
}

func (t TypeDefinition) Params() []Parameter {
	return t.params
}

func (t *TypeDefinition) SetResolvedArgs(ra map[string]value.Value) {
	t.resolvedArgs = ra
}

func (t TypeDefinition) ResolvedArgs() map[string]value.Value {
	return t.resolvedArgs
}

func (t TypeDefinition) Properties() map[string]Property {
	return t.properties
}

func (t TypeDefinition) Symbols() map[string]Symbol {
	return t.symbols
}
