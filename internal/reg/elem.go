package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
)

type Element struct {
	InsElem  *ins.Element
	Access   Access
	Sizes    Sizes
	Elements map[string]*Element
}

func (e *Element) Constants() map[string]val.Value { return e.InsElem.Constants }
