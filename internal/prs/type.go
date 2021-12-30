package prs

import (
	"strings"
)

// Parameter represents parameter in the type definition parameter list, not 'param' element.
type Parameter struct {
	Name         string
	HasDfltValue bool
	DfltValue    Expr
}

// Type represents type definition.
type Type struct {
	base

	typ        string
	properties map[string]Property
	symbols    SymbolContainer

	params       []Parameter
	args         []Argument
	resolvedArgs map[string]Expr
}

func (t *Type) GetSymbol(name string) (Symbol, error) {
	if strings.Contains(name, ".") {
		panic("To be implemented")
	}

	sym, ok := t.symbols.Get(name)
	if ok {
		return sym, nil
	}

	if v, ok := t.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	if t.parent != nil {
		return t.parent.GetSymbol(name)
	}

	return t.file.GetSymbol(name)
}

func (t Type) Type() string                        { return t.typ }
func (t Type) Args() []Argument                    { return t.args }
func (t Type) Params() []Parameter                 { return t.params }
func (t *Type) SetResolvedArgs(ra map[string]Expr) { t.resolvedArgs = ra }
func (t Type) ResolvedArgs() map[string]Expr       { return t.resolvedArgs }
func (t Type) Properties() map[string]Property     { return t.properties }
func (t Type) Symbols() SymbolContainer            { return t.symbols }
