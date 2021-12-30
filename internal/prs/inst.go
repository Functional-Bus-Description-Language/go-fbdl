package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

// Inst represents element instantiation.
type Inst struct {
	base

	typ     string
	IsArray bool
	Count   Expr

	properties map[string]Prop
	symbols    SymbolContainer

	args         []Arg
	resolvedArgs map[string]Expr
}

func (i Inst) Type() string {
	return i.typ
}

func (i *Inst) GetSymbol(name string) (Symbol, error) {
	sym, ok := i.symbols.Get(name)
	if ok {
		return sym, nil
	}

	if v, ok := i.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	if i.parent != nil {
		return i.parent.GetSymbol(name)
	}

	return i.file.GetSymbol(name)
}

func (i Inst) Args() []Arg                         { return i.args }
func (i *Inst) SetResolvedArgs(ra map[string]Expr) { i.resolvedArgs = ra }
func (i Inst) ResolvedArgs() map[string]Expr       { return i.resolvedArgs }
func (i Inst) Props() map[string]Prop              { return i.properties }
func (i Inst) Symbols() SymbolContainer            { return i.symbols }

func (i Inst) Params() []Param {
	panic("should never happen, element definition cannot have parameters")
}

// validate checks whether given element definition is valid.
// For example, whether given properties are valid for given element type.
func (i Inst) validate() error {
	if !util.IsBaseType(i.typ) {
		return nil
	}

	// Checks specific for base type only.
	if len(i.args) != 0 {
		return fmt.Errorf("base type '%s' does not accept arguments", i.typ)
	}

	for prop, v := range i.properties {
		if err := util.IsValidProperty(prop, i.typ); err != nil {
			return fmt.Errorf("line %d: %v", v.LineNumber, err)
		}
	}

	return nil
}
