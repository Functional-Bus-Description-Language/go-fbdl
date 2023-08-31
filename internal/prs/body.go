package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

func buildBody(astBody ast.Body, src []byte, scope Scope) (PropContainer, symbolContainer, error) {
	pc := PropContainer{}
	sc := symbolContainer{}

	for _, ap := range astBody.Props {
		p := Prop{}

		p.Line = ap.Name.Line()
		p.Col = ap.Name.Column()
		p.Name = tok.Text(ap.Name, src)
		v, err := MakeExpr(ap.Value, src, scope)
		if err != nil {
			return nil, sc, err
		}
		p.Value = v
		if ok := pc.Add(p); !ok {
			return nil, sc, tok.Error{
				Tok: ap.Name,
				Msg: fmt.Sprintf("reassignment to '%s' property", p.Name),
			}
		}
	}

	// Handle constants
	consts, err := buildConsts(astBody.Consts, src, scope)
	if err != nil {
		return nil, sc, err
	}
	for i, c := range consts {
		if ok := sc.addConst(c); !ok {
			return nil, sc, tok.Error{
				Tok: astBody.Consts[i].Name,
				Msg: fmt.Sprintf("redefinition of symbol '%s'", c.Name()),
			}
		}
	}

	// Handle types
	types, err := buildTypes(astBody.Types, src)
	if err != nil {
		return nil, sc, err
	}
	for i, t := range types {
		if ok := sc.addType(t); !ok {
			return nil, sc, tok.Error{
				Tok: astBody.Types[i].Name,
				Msg: fmt.Sprintf("redefinition of symbol '%s'", t.Name()),
			}
		}
	}

	// Handle instantiations
	insts, err := buildInsts(astBody.Insts, src)
	if err != nil {
		return nil, sc, err
	}
	for i, ins := range insts {
		if ok := sc.addInst(ins); !ok {
			return nil, sc, tok.Error{
				Tok: astBody.Insts[i].Name,
				Msg: fmt.Sprintf("redefinition of symbol '%s'", ins.Name()),
			}
		}
	}

	return pc, sc, nil
}
