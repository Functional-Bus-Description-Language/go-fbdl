package parser

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func buildBody(astBody ast.Body, file *File, scope Scope) (PropContainer, symbolContainer, error) {
	pc := PropContainer{}
	sc := symbolContainer{}

	for _, ap := range astBody.Props {
		p := Prop{src: file.Src}

		p.NameTok = ap.Name
		p.Name = ap.Name.Text(file.Src)
		v, err := MakeExpr(ap.Value, file, scope)
		if err != nil {
			return nil, sc, err
		}
		p.Value = v
		p.ValueTok = ap.Value.Token()

		if ok := pc.Add(p); !ok {
			return nil, sc, token.Error{
				Msg:  fmt.Sprintf("reassignment to '%s' property", p.Name),
				Toks: []token.Token{ap.Name},
			}
		}
	}

	// Handle constants
	consts, err := buildConsts(astBody.Consts, file, scope)
	if err != nil {
		return nil, sc, err
	}
	for i, c := range consts {
		if ok := sc.addConst(c); !ok {
			return nil, sc, token.Error{
				Msg:  fmt.Sprintf("redefinition of symbol '%s'", c.Name()),
				Toks: []token.Token{astBody.Consts[i].Name},
			}
		}
	}

	// Handle types
	types, err := buildTypes(astBody.Types, file)
	if err != nil {
		return nil, sc, err
	}
	for i, t := range types {
		if ok := sc.addType(t); !ok {
			return nil, sc, token.Error{
				Msg:  fmt.Sprintf("redefinition of symbol '%s'", t.Name()),
				Toks: []token.Token{astBody.Types[i].Name},
			}
		}
	}

	// Handle instantiations
	insts, err := buildInsts(astBody.Insts, file)
	if err != nil {
		return nil, sc, err
	}
	for i, ins := range insts {
		if ok := sc.addInst(ins); !ok {
			return nil, sc, token.Error{
				Msg:  fmt.Sprintf("redefinition of symbol '%s'", ins.Name()),
				Toks: []token.Token{astBody.Insts[i].Name},
			}
		}
	}

	return pc, sc, nil
}
