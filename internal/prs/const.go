package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Const represents constant definition.
type Const struct {
	symbol
	Value Expr
}

func (c Const) Kind() SymbolKind { return ConstDef }

// buildConsts builds list of Consts defined in the same scope based on the list of ast.Const.
func buildConsts(astConsts []ast.Const, src []byte, scope Scope) ([]*Const, error) {
	consts := make([]*Const, 0, len(astConsts))
	cache := make(map[string]*Const)

	for _, ac := range astConsts {
		c := &Const{}

		c.tok = ac.Name
		c.name = tok.Text(ac.Name, src)
		v, err := MakeExpr(ac.Value, src, scope)
		if err != nil {
			return nil, err
		}
		c.Value = v
		c.doc = ac.Doc.Text(src)

		if first, ok := cache[c.name]; ok {
			return nil, tok.Error{
				Msg: fmt.Sprintf(
					"redefinition of constant '%s', first definition line %d column %d",
					c.name, first.Line(), first.Col(),
				),
				Toks: []tok.Token{ac.Name},
			}
		}

		cache[c.name] = c
		consts = append(consts, c)
	}

	return consts, nil
}
