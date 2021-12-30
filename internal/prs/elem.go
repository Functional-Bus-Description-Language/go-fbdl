package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

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

type Element interface {
	Searchable
	Symbol
	Type() string
	Args() []Argument
	Params() []Parameter
	SetResolvedArgs(args map[string]Expression)
	ResolvedArgs() map[string]Expression
	Properties() map[string]Property
	Symbols() SymbolContainer
}

type ElementDefinition struct {
	base

	typ     string
	IsArray bool
	Count   Expression

	properties map[string]Property
	symbols    SymbolContainer

	args         []Argument
	resolvedArgs map[string]Expression
}

func (e ElementDefinition) Type() string {
	return e.typ
}

func (e *ElementDefinition) GetSymbol(name string) (Symbol, error) {
	sym, ok := e.symbols.Get(name)
	if ok {
		return sym, nil
	}

	if v, ok := e.resolvedArgs[name]; ok {
		return &Constant{Value: v}, nil
	}

	if e.parent != nil {
		return e.parent.GetSymbol(name)
	}

	return e.file.GetSymbol(name)
}

func (e ElementDefinition) Args() []Argument                          { return e.args }
func (e *ElementDefinition) SetResolvedArgs(ra map[string]Expression) { e.resolvedArgs = ra }
func (e ElementDefinition) ResolvedArgs() map[string]Expression       { return e.resolvedArgs }
func (e ElementDefinition) Properties() map[string]Property           { return e.properties }
func (e ElementDefinition) Symbols() SymbolContainer                  { return e.symbols }

func (e ElementDefinition) Params() []Parameter {
	panic("should never happen, element definition cannot have parameters")
}

// validate checks whether given element definition is valid.
// For example, whether given properties are valid for given element type.
func (e ElementDefinition) validate() error {
	if !util.IsBaseType(e.typ) {
		return nil
	}

	// Checks specific for base type only.
	if len(e.args) != 0 {
		return fmt.Errorf("base type '%s' does not accept arguments", e.typ)
	}

	for prop, v := range e.properties {
		if err := util.IsValidProperty(prop, e.typ); err != nil {
			return fmt.Errorf("line %d: %v", v.LineNumber, err)
		}
	}

	return nil
}
