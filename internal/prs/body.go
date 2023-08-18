package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

func buildBody(astBody ast.Body, src []byte, parent Searchable) (PropContainer, SymbolContainer, error) {
	pc := PropContainer{}

	for _, ap := range astBody.Props {
		p := Prop{}

		p.Line = ap.Name.Line()
		p.Name = tok.Text(ap.Name, src)
		v, err := MakeExpr(ap.Value, src, parent)
		if err != nil {
			return nil, nil, err
		}
		p.Value = v
		if ok := pc.Add(p); !ok {
			return nil, nil, fmt.Errorf(
				"%s: reassignment to '%s' property",
				tok.Loc(ap.Name), p.Name,
			)
		}
	}

	sc := SymbolContainer{}

	// Handle constants
	consts, err := buildConsts(astBody.Consts, src)
	if err != nil {
		return nil, nil, err
	}
	for i, c := range consts {
		if ok := sc.Add(c); !ok {
			return nil, nil, fmt.Errorf(
				"%s: redefinition of symbol '%s'",
				tok.Loc(astBody.Consts[i].Name), c.Name(),
			)
		}
	}

	// Handle types
	types, err := buildTypes(astBody.Types, src)
	if err != nil {
		return nil, nil, err
	}
	for i, t := range types {
		if ok := sc.Add(t); !ok {
			return nil, nil, fmt.Errorf(
				"%s: redefinition of symbol '%s'",
				tok.Loc(astBody.Types[i].Name), t.Name(),
			)
		}
	}

	// Handle instantiations
	insts, err := buildInsts(astBody.Insts, src)
	if err != nil {
		return nil, nil, err
	}
	for i, ins := range insts {
		if ok := sc.Add(ins); !ok {
			return nil, nil, fmt.Errorf(
				"%s: redefinition of symbol '%s'",
				tok.Loc(astBody.Insts[i].Name), ins.Name(),
			)
		}
	}

	return pc, sc, nil
}
