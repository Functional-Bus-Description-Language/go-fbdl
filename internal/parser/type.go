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

func (typ *Type) GetConst(name string) (*Const, error) {
	sym, ok := typ.symbolContainer.GetConst(name)
	if ok {
		return sym, nil
	}

	if v, ok := typ.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	return typ.scope.GetConst(name)
}

func (typ *Type) GetInst(name string) (*Inst, error) {
	sym, ok := typ.symbolContainer.GetInst(name)
	if ok {
		return sym, nil
	}

	return typ.scope.GetInst(name)
}

func (typ *Type) GetType(name string) (*Type, error) {
	sym, ok := typ.symbolContainer.GetType(name)
	if ok {
		return sym, nil
	}

	return typ.scope.GetType(name)
}

func (typ Type) Kind() SymbolKind                    { return TypeDef }
func (typ Type) Type() string                        { return typ.typ }
func (typ Type) Args() []Arg                         { return typ.args.Args }
func (typ Type) Params() []Param                     { return typ.params }
func (typ *Type) SetResolvedArgs(ra map[string]Expr) { typ.resolvedArgs = ra }
func (typ Type) ResolvedArgs() map[string]Expr       { return typ.resolvedArgs }
func (typ Type) Props() PropContainer                { return typ.props }
func (typ Type) Symbols() []Symbol                   { return typ.symbolContainer.Symbols() }
func (typ Type) IsArray() bool                       { return false }
func (typ Type) Count() Expr                         { return typ.count }

// buildTypes builds list of Types based on the list of ast.Type.
func buildTypes(astTypes []ast.Type, src []byte) ([]*Type, error) {
	types := make([]*Type, 0, len(astTypes))
	cache := make(map[string]*Type)

	for _, astType := range astTypes {
		typ, err := buildType(astType, src)
		if err != nil {
			return nil, err
		}

		if first, ok := cache[typ.name]; ok {
			return nil, token.Error{
				Msg: fmt.Sprintf(
					"redefinition of type '%s', first definition line %d column %d",
					typ.name, first.Line(), first.Col(),
				),
				Toks: []token.Token{astType.Name, first.token},
			}
		}

		cache[typ.name] = typ
		types = append(types, typ)
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
