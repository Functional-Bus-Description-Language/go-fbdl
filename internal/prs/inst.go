package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

// Inst struct represents functionality instantiation.
type Inst struct {
	base

	typ   string
	count Expr

	props   PropContainer
	symbols SymbolContainer

	args         []Arg
	resolvedArgs map[string]Expr
}

func (i Inst) Kind() SymbolKind { return ElemInst }
func (i Inst) Type() string     { return i.typ }
func (i Inst) IsArray() bool    { return i.count != nil }
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

// buildInsts builds list of Insts based on the list of ast.Inst.
func buildInsts(astInsts []ast.Inst, src []byte) ([]*Inst, error) {
	insts := make([]*Inst, 0, len(astInsts))

	for _, ai := range astInsts {
		i, err := buildInst(ai, src)
		if err != nil {
			return nil, err
		}
		insts = append(insts, i)
	}

	return insts, nil
}

func buildInst(ai ast.Inst, src []byte) (*Inst, error) {
	i := &Inst{}

	i.lineNum = uint32(ai.Name.Line())
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
		return nil, fmt.Errorf(
			"%s: base type '%s' does not accept argument list",
			tok.Loc(ai.Type), i.typ,
		)
	}

	props, syms, err := buildBody(ai.Body, src, i)
	if err != nil {
		return nil, err
	}

	if util.IsBaseType(i.typ) {
		for j, p := range props {
			if err := util.IsValidProperty(p.Name, i.typ); err != nil {
				return nil, fmt.Errorf("line %d: %v", p.LineNum, err)
			}

			if err := checkPropConflict(i.typ, p, props[0:j]); err != nil {
				return nil, fmt.Errorf("%v", err)
			}
		}
	}
	i.props = props

	for _, s := range syms {
		s.SetParent(i)
	}
	i.symbols = syms

	return i, nil
}
