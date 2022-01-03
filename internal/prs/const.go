package prs

import (
	"strings"
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
