package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

// Inst struct represents functionality instantiation.
type Inst struct {
	symbol

	typ   string
	count Expr

	args         []Arg
	resolvedArgs map[string]Expr

	props PropContainer
	symbolContainer
}

func (i Inst) Kind() SymbolKind { return FuncInst }
func (i Inst) Type() string     { return i.typ }
func (i Inst) IsArray() bool    { return i.count != nil }
func (i Inst) Count() Expr      { return i.count }

func (i *Inst) GetConst(name string) (*Const, error) {
	sym, ok := i.symbolContainer.GetConst(name)
	if ok {
		return sym, nil
	}

	if v, ok := i.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	return i.scope.GetConst(name)
}

func (i *Inst) GetInst(name string) (*Inst, error) {
	sym, ok := i.symbolContainer.GetInst(name)
	if ok {
		return sym, nil
	}

	return i.scope.GetInst(name)
}

func (i *Inst) GetType(name string) (*Type, error) {
	sym, ok := i.symbolContainer.GetType(name)
	if ok {
		return sym, nil
	}

	return i.scope.GetType(name)
}

func (i Inst) Args() []Arg                         { return i.args }
func (i *Inst) SetResolvedArgs(ra map[string]Expr) { i.resolvedArgs = ra }
func (i Inst) ResolvedArgs() map[string]Expr       { return i.resolvedArgs }
func (i Inst) Props() PropContainer                { return i.props }
func (i Inst) Symbols() []Symbol                   { return i.symbolContainer.Symbols() }

func (i Inst) File() *File {
	if i.file != nil {
		return i.file
	}

	if s, ok := i.scope.(Symbol); ok {
		return s.File()
	}

	panic("should never happen")
}

func (i Inst) Params() []Param {
	panic("should never happen, element definition cannot have parameters")
}

// buildInsts builds list of Insts based on the list of ast.Instance.
func buildInsts(astInsts []ast.Instance, src []byte) ([]*Inst, error) {
	insts := make([]*Inst, 0, len(astInsts))
	cache := make(map[string]*Inst)

	for _, ai := range astInsts {
		i, err := buildInst(ai, src)
		if err != nil {
			return nil, err
		}

		if first, ok := cache[i.name]; ok {
			return nil, tok.Error{
				Tok: ai.Name,
				Msg: fmt.Sprintf(
					"reinstantiation of '%s', first instantiation line %d column %d",
					i.name, first.Line(), first.Col(),
				),
			}
		}

		cache[i.name] = i
		insts = append(insts, i)
	}

	return insts, nil
}

func buildInst(ai ast.Instance, src []byte) (*Inst, error) {
	i := &Inst{}

	i.line = ai.Name.Line()
	i.col = ai.Name.Column()
	i.name = tok.Text(ai.Name, src)
	i.doc = ai.Doc.Text(src)

	v, err := MakeExpr(ai.Count, src, i)
	if err != nil {
		return nil, err
	}
	i.count = v

	i.typ = tok.Text(ai.Type, src)

	args, err := buildArgList(ai.Args, src, i)
	if err != nil {
		return nil, err
	}
	i.args = args

	if util.IsBaseType(i.typ) && len(i.args) > 0 {
		return nil, tok.Error{
			Tok: ai.Type,
			Msg: fmt.Sprintf("base type '%s' does not accept argument list", i.typ),
		}
	}

	props, syms, err := buildBody(ai.Body, src, i)
	if err != nil {
		return nil, err
	}

	if util.IsBaseType(i.typ) {
		for j, p := range props {
			if err := util.IsValidProperty(p.Name, i.typ); err != nil {
				return nil, tok.Error{
					Tok: ai.Body.Props[j].Name,
					Msg: err.Error(),
				}
			}

			if err := checkPropConflict(i.typ, p, props[0:j]); err != nil {
				return nil, tok.Error{
					Tok: ai.Body.Props[j].Name,
					Msg: err.Error(),
				}
			}
		}
	}
	i.props = props

	for _, s := range syms.Consts {
		s.setScope(i)
	}
	for _, s := range syms.Insts {
		s.setScope(i)
	}
	for _, s := range syms.Types {
		s.setScope(i)
	}
	i.symbolContainer = syms

	return i, nil
}
