package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
)

type BlockElement struct {
	InsElem            *ins.Element
	Sizes              Sizes
	BlockElements      map[string]*BlockElement
	FunctionalElements map[string]*FunctionalElement
}

func (be *BlockElement) Constants() map[string]val.Value { return be.InsElem.Constants }

type FunctionalElement struct {
	InsElem *ins.Element
	Access  Access
}
