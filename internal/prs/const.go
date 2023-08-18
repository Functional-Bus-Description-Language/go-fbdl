package prs

import (
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Const represents constant definition.
type Const struct {
	base
	Value Expr
	doc   string
}

func (c Const) Kind() SymbolKind { return ConstDef }

func (c Const) GetSymbol(name string, kind SymbolKind) (Symbol, error) {
	if strings.Contains(name, ".") {
		panic("To be implemented")
	}

	if c.parent != nil {
		return c.parent.GetSymbol(name, kind)
	}

	return c.file.GetSymbol(name, kind)
}

// buildConst builds list of Consts based on the list of ast.Const.
func buildConsts(astConsts []ast.Const, src []byte) ([]*Const, error) {
	consts := make([]*Const, 0, len(astConsts))

	for _, ac := range astConsts {
		c := &Const{}

		c.line = ac.Name.Line()
		c.name = tok.Text(ac.Name, src)
		v, err := MakeExpr(ac.Value, src, c)
		if err != nil {
			return nil, err
		}
		c.Value = v
		c.doc = ac.Doc.Text(src)
		consts = append(consts, c)
	}

	return consts, nil
}
