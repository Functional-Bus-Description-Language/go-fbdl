package parser

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

// Inst struct represents functionality instantiation.
type Inst struct {
	symbol

	typ   string
	count Expr

	argList      ArgList
	resolvedArgs map[string]Expr

	props PropContainer
	symbolContainer
}

func (inst Inst) Kind() SymbolKind { return FuncInst }
func (inst Inst) Type() string     { return inst.typ }
func (inst Inst) IsArray() bool    { return inst.count != nil }
func (inst Inst) Count() Expr      { return inst.count }

func (inst *Inst) GetConst(name string) (*Const, error) {
	sym, ok := inst.symbolContainer.GetConst(name)
	if ok {
		return sym, nil
	}

	if v, ok := inst.resolvedArgs[name]; ok {
		return &Const{Value: v}, nil
	}

	return inst.scope.GetConst(name)
}

func (inst *Inst) GetInst(name string) (*Inst, error) {
	sym, ok := inst.symbolContainer.GetInst(name)
	if ok {
		return sym, nil
	}

	return inst.scope.GetInst(name)
}

func (inst *Inst) GetType(name string) (*Type, error) {
	sym, ok := inst.symbolContainer.GetType(name)
	if ok {
		return sym, nil
	}

	return inst.scope.GetType(name)
}

func (inst Inst) Args() []Arg                         { return inst.argList.Args }
func (inst *Inst) SetResolvedArgs(ra map[string]Expr) { inst.resolvedArgs = ra }
func (inst Inst) ResolvedArgs() map[string]Expr       { return inst.resolvedArgs }
func (inst Inst) Props() PropContainer                { return inst.props }
func (inst Inst) Symbols() []Symbol                   { return inst.symbolContainer.Symbols() }

func (inst Inst) File() *File {
	if inst.file != nil {
		return inst.file
	}

	if s, ok := inst.scope.(Symbol); ok {
		return s.File()
	}

	panic("should never happen")
}

func (inst Inst) Params() []Param {
	panic("should never happen, element definition cannot have parameters")
}

// buildInsts builds list of Insts based on the list of ast.Inst.
func buildInsts(astInsts []ast.Inst, src []byte) ([]*Inst, error) {
	insts := make([]*Inst, 0, len(astInsts))
	cache := make(map[string]*Inst)

	for _, astInst := range astInsts {
		inst, err := buildInst(astInst, src)
		if err != nil {
			return nil, err
		}

		if first, ok := cache[inst.name]; ok {
			return nil, token.Error{
				Msg: fmt.Sprintf(
					"reinstantiation of '%s', first instantiation line %d column %d",
					inst.name, first.Line(), first.Col(),
				),
				Toks: []token.Token{astInst.Name},
			}
		}

		cache[inst.name] = inst
		insts = append(insts, inst)
	}

	return insts, nil
}

func buildInst(astInst ast.Inst, src []byte) (*Inst, error) {
	inst := &Inst{}

	inst.token = astInst.Name
	inst.name = token.Text(astInst.Name, src)
	inst.doc = astInst.Doc.Text(src)

	v, err := MakeExpr(astInst.Count, src, inst)
	if err != nil {
		return nil, err
	}
	inst.count = v

	inst.typ = token.Text(astInst.Type, src)

	argList, err := buildArgList(astInst.ArgList, src, inst)
	if err != nil {
		return nil, err
	}
	inst.argList = argList

	if util.IsBaseType(inst.typ) && inst.argList.Len() > 0 {
		return nil, token.Error{
			Msg:  fmt.Sprintf("base type '%s' does not accept argument list", inst.typ),
			Toks: []token.Token{astInst.Type},
		}
	}

	props, syms, err := buildBody(astInst.Body, src, inst)
	if err != nil {
		return nil, err
	}

	if util.IsBaseType(inst.typ) {
		for j, p := range props {
			if err := util.IsValidProperty(p.Name, inst.typ); err != nil {
				return nil, token.Error{
					Msg:  err.Error(),
					Toks: []token.Token{astInst.Body.Props[j].Name},
				}
			}

			if err := checkPropConflict(inst.typ, p, props[0:j]); err != nil {
				return nil, err
			}
		}
	}
	inst.props = props

	for _, s := range syms.Consts {
		s.setScope(inst)
	}
	for _, s := range syms.Insts {
		s.setScope(inst)
	}
	for _, s := range syms.Types {
		s.setScope(inst)
	}
	inst.symbolContainer = syms

	return inst, nil
}
