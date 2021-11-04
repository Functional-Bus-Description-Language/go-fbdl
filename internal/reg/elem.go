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

func (be *BlockElement) hasElement(name string) bool {
	if _, ok := be.BlockElements["x_timestamp_x"]; ok {
		return true
	}

	if _, ok := be.FunctionalElements["x_timestamp_x"]; ok {
		return true
	}

	return false
}

type FunctionalElement struct {
	InsElem *ins.Element
	Access  Access
}
