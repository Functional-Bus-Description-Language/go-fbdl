package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

// Inst represents element instantiation.
type Inst struct {
	base

	typ     string
	isArray bool
	count   Expr

	props   PropContainer
	symbols SymbolContainer

	args         []Arg
	resolvedArgs map[string]Expr
}

func (i Inst) Kind() SymbolKind { return ElemInst }
func (i Inst) Type() string     { return i.typ }
func (i Inst) IsArray() bool    { return i.isArray }
func (i Inst) Count() Expr      { return i.count }

func (i *Inst) GetSymbol(name string, kind SymbolKind) (Symbol, error) {
	sym, ok := i.symbols.Get(name, kind)
	if ok {
		return sym, nil
	}

	if v, ok := i.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	if i.parent != nil {
		return i.parent.GetSymbol(name, kind)
	}

	return i.file.GetSymbol(name, kind)
}

func (i Inst) Args() []Arg                         { return i.args }
func (i *Inst) SetResolvedArgs(ra map[string]Expr) { i.resolvedArgs = ra }
func (i Inst) ResolvedArgs() map[string]Expr       { return i.resolvedArgs }
func (i Inst) Props() PropContainer                { return i.props }
func (i Inst) Symbols() SymbolContainer            { return i.symbols }

func (i Inst) File() *File {
	if i.file != nil {
		return i.file
	}

	if s, ok := i.parent.(Symbol); ok {
		return s.File()
	}

	panic("should never happen")
}

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

	for j, p := range i.props {
		if err := util.IsValidProperty(p.Name, i.typ); err != nil {
			return fmt.Errorf("line %d: %v", p.LineNum, err)
		}

		if err := checkPropConflict(i.typ, p, i.props[0:j]); err != nil {
			return fmt.Errorf("%v", err)
		}
	}

	return nil
}
