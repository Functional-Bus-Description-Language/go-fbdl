package parser

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

// Type represents type definition.
type Type struct {
	symbol

	typ   string
	count Expr

	params       []Param
	args         ArgList
	resolvedArgs map[string]Expr

	props PropContainer
	symbolContainer
}

func (t *Type) GetConst(name string) (*Const, error) {
	sym, ok := t.symbolContainer.GetConst(name)
	if ok {
		return sym, nil
	}

	if v, ok := t.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	return t.scope.GetConst(name)
}

func (t *Type) GetInst(name string) (*Inst, error) {
	sym, ok := t.symbolContainer.GetInst(name)
	if ok {
		return sym, nil
	}

	return t.scope.GetInst(name)
}

func (t *Type) GetType(name string) (*Type, error) {
	sym, ok := t.symbolContainer.GetType(name)
	if ok {
		return sym, nil
	}

	return t.scope.GetType(name)
}

func (t Type) Kind() SymbolKind                    { return TypeDef }
func (t Type) Type() string                        { return t.typ }
func (t Type) Args() []Arg                         { return t.args.Args }
func (t Type) Params() []Param                     { return t.params }
func (t *Type) SetResolvedArgs(ra map[string]Expr) { t.resolvedArgs = ra }
func (t Type) ResolvedArgs() map[string]Expr       { return t.resolvedArgs }
func (t Type) Props() PropContainer                { return t.props }
func (t Type) Symbols() []Symbol                   { return t.symbolContainer.Symbols() }
func (t Type) IsArray() bool                       { return false }
func (t Type) Count() Expr                         { return t.count }

// buildTypes builds list of Types based on the list of ast.Type.
func buildTypes(astTypes []ast.Type, src []byte) ([]*Type, error) {
	types := make([]*Type, 0, len(astTypes))
	cache := make(map[string]*Type)

	for _, at := range astTypes {
		t, err := buildType(at, src)
		if err != nil {
			return nil, err
		}

		if first, ok := cache[t.name]; ok {
			return nil, token.Error{
				Msg: fmt.Sprintf(
					"redefinition of type '%s', first definition line %d column %d",
					t.name, first.Line(), first.Col(),
				),
				Toks: []token.Token{at.Name, first.token},
			}
		}

		cache[t.name] = t
		types = append(types, t)
	}

	return types, nil
}

func buildType(astTyp ast.Type, src []byte) (*Type, error) {
	typ := &Type{}

	typ.token = astTyp.Name
	typ.name = token.Text(astTyp.Name, src)
	typ.doc = astTyp.Doc.Text(src)

	params, err := buildParamList(astTyp.Params, src, typ)
	if err != nil {
		return nil, err
	}
	typ.params = params

	v, err := MakeExpr(astTyp.Count, src, typ)
	if err != nil {
		return nil, err
	}
	typ.count = v

	typ.typ = token.Text(astTyp.Type, src)

	args, err := buildArgList(astTyp.Args, src, typ)
	if err != nil {
		return nil, err
	}
	typ.args = args

	if util.IsBaseType(typ.typ) && len(typ.args.Args) > 0 {
		return nil, token.Error{
			Msg:  fmt.Sprintf("base type '%s' does not accept argument list", typ.typ),
			Toks: []token.Token{astTyp.Type},
		}
	}

	props, syms, err := buildBody(astTyp.Body, src, typ)
	if err != nil {
		return nil, err
	}

	if util.IsBaseType(typ.typ) {
		for j, p := range props {
			if err := util.IsValidProperty(p.Name, typ.typ); err != nil {
				return nil, token.Error{
					Msg:  err.Error(),
					Toks: []token.Token{astTyp.Body.Props[j].Name},
				}
			}

			if err := checkPropConflict(typ.typ, p, props[0:j]); err != nil {
				return nil, token.Error{
					Msg:  err.Error(),
					Toks: []token.Token{astTyp.Body.Props[j].Name},
				}
			}
		}
	}
	typ.props = props

	for _, s := range syms.Consts {
		s.setScope(typ)
	}
	for _, s := range syms.Insts {
		s.setScope(typ)
	}
	for _, s := range syms.Types {
		s.setScope(typ)
	}
	typ.symbolContainer = syms

	return typ, nil
}
