package prs

import (
	"fmt"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

// Type represents type definition.
type Type struct {
	base

	typ   string
	count Expr

	props   PropContainer
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
func (t Type) Props() PropContainer                { return t.props }
func (t Type) Symbols() SymbolContainer            { return t.symbols }
func (t Type) IsArray() bool                       { return false }
func (t Type) Count() Expr                         { return t.count }

// buildTypes builds list of Types based on the list of ast.Type.
func buildTypes(astTypes []ast.Type, src []byte) ([]*Type, error) {
	types := make([]*Type, 0, len(astTypes))

	for _, at := range astTypes {
		t, err := buildType(at, src)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}

	return types, nil
}

func buildType(at ast.Type, src []byte) (*Type, error) {
	t := &Type{}

	t.line = at.Name.Line()
	t.name = tok.Text(at.Name, src)
	t.doc = at.Doc.Text(src)

	params, err := buildParamList(at.Params, src, t)
	if err != nil {
		return nil, err
	}
	t.params = params

	v, err := MakeExpr(at.Count, src, t)
	if err != nil {
		return nil, err
	}
	t.count = v

	t.typ = tok.Text(at.Type, src)

	args, err := buildArgList(at.Args, src, t)
	if err != nil {
		return nil, err
	}
	t.args = args

	if util.IsBaseType(t.typ) && len(t.args) > 0 {
		return nil, fmt.Errorf(
			"%s: base type '%s' does not accept argument list",
			tok.Loc(at.Type), t.typ,
		)
	}

	props, syms, err := buildBody(at.Body, src, t)
	if err != nil {
		return nil, err
	}

	if util.IsBaseType(t.typ) {
		for j, p := range props {
			if err := util.IsValidProperty(p.Name, t.typ); err != nil {
				return nil, fmt.Errorf("line %d: %v", p.Line, err)
			}

			if err := checkPropConflict(t.typ, p, props[0:j]); err != nil {
				return nil, fmt.Errorf("%v", err)
			}
		}
	}
	t.props = props

	for _, s := range syms {
		s.SetParent(t)
	}
	t.symbols = syms

	return t, nil
}
