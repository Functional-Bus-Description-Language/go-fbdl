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
				Msg:  fmt.Sprintf("reassignment to '%s' property", p.Name),
				Toks: []tok.Token{ap.Name},
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
				Msg:  fmt.Sprintf("redefinition of symbol '%s'", c.Name()),
				Toks: []tok.Token{astBody.Consts[i].Name},
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
				Msg:  fmt.Sprintf("redefinition of symbol '%s'", t.Name()),
				Toks: []tok.Token{astBody.Types[i].Name},
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
				Msg:  fmt.Sprintf("redefinition of symbol '%s'", ins.Name()),
				Toks: []tok.Token{astBody.Insts[i].Name},
			}
		}
	}

	return pc, sc, nil
}
