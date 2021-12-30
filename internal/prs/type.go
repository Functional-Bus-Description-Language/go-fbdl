package prs

import (
	"strings"
)

// Parameter represents parameter in the parameter list, not 'param' element.
type Parameter struct {
	Name            string
	HasDefaultValue bool
	DefaultValue    Expression
}

type TypeDefinition struct {
	base

	typ        string
	properties map[string]Property
	symbols    SymbolContainer

	params       []Parameter
	args         []Argument
	resolvedArgs map[string]Expression
}

func (t *TypeDefinition) GetSymbol(name string) (Symbol, error) {
	if strings.Contains(name, ".") {
		panic("To be implemented")
	}

	sym, ok := t.symbols.Get(name)
	if ok {
		return sym, nil
	}

	if v, ok := t.resolvedArgs[name]; ok {
		return &Constant{Value: v}, nil
	}

	if t.parent != nil {
		return t.parent.GetSymbol(name)
	}

	return t.file.GetSymbol(name)
}

func (t TypeDefinition) Type() string {
	return t.typ
}

func (t TypeDefinition) Args() []Argument {
	return t.args
}

func (t TypeDefinition) Params() []Parameter {
	return t.params
}

func (t *TypeDefinition) SetResolvedArgs(ra map[string]Expression) {
	t.resolvedArgs = ra
}

func (t TypeDefinition) ResolvedArgs() map[string]Expression {
	return t.resolvedArgs
}

func (t TypeDefinition) Properties() map[string]Property {
	return t.properties
}

func (t TypeDefinition) Symbols() SymbolContainer {
	return t.symbols
}
