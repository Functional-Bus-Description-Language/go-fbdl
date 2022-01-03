package prs

import (
	"strings"
)

// Param represents parameter in the type definition parameter list, not 'param' element.
type Param struct {
	Name         string
	HasDfltValue bool
	DfltValue    Expr
}

// Type represents type definition.
type Type struct {
	base

	typ     string
	props   map[string]Prop
	symbols SymbolContainer

	params       []Param
	args         []Arg
	resolvedArgs map[string]Expr
}

func (t *Type) GetSymbol(name string, kind SymbolKind) (Symbol, error) {
	if strings.Contains(name, ".") {
		panic("To be implemented")
	}

	sym, ok := t.symbols.Get(name, kind)
	if ok {
		return sym, nil
	}

	if v, ok := t.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	if t.parent != nil {
		return t.parent.GetSymbol(name, kind)
	}

	return t.file.GetSymbol(name, kind)
}

func (t Type) Kind() SymbolKind                    { return TypeDef }
func (t Type) Type() string                        { return t.typ }
func (t Type) Args() []Arg                         { return t.args }
func (t Type) Params() []Param                     { return t.params }
func (t *Type) SetResolvedArgs(ra map[string]Expr) { t.resolvedArgs = ra }
func (t Type) ResolvedArgs() map[string]Expr       { return t.resolvedArgs }
func (t Type) Props() map[string]Prop              { return t.props }
func (t Type) Symbols() SymbolContainer            { return t.symbols }
