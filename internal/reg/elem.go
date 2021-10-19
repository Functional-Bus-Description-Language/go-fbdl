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
	AddrSpace          AddrSpace
}

func (be *BlockElement) Constants() map[string]val.Value { return be.InsElem.Constants }

func (be *BlockElement) IsArray() bool { return be.InsElem.IsArray }

func (be *BlockElement) Count() uint { return be.InsElem.Count }

type FunctionalElement struct {
	InsElem *ins.Element
	Access  Access
}
