package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Const interface {
	constNode()
}

type SingleConst struct {
	Doc  Doc
	Name token.Ident
	Expr Expr
}

func (sc SingleConst) constNode() {}

func (sc SingleConst) eq(sc2 SingleConst) bool {
	return sc.Doc.eq(sc2.Doc) && sc.Name == sc2.Name && sc.Expr == sc2.Expr
}

type MultiConst struct {
	Consts []SingleConst
}

func (mc MultiConst) constNode() {}

func (mc MultiConst) eq(mc2 MultiConst) bool {
	if len(mc.Consts) != len(mc2.Consts) {
		return false
	}

	for i, c := range mc.Consts {
		if !c.eq(mc2.Consts[i]) {
			return false
		}
	}

	return true
}
