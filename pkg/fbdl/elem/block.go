package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Block struct {
	Elem

	Masters int64
	Width   int64

	ConstContainer

	Configs   []*Config
	Funcs     []*Func
	Masks     []*Mask
	Statics   []*Static
	Statuses  []*Status
	Streams   []*Stream
	Subblocks []*Block

	Sizes     access.Sizes
	AddrSpace access.AddrSpace
}

func (b *Block) GroupedElems() []Groupable {
	elemsWithGrps := []Groupable{}

	for _, c := range b.Configs {
		if len(c.Groups) > 0 {
			elemsWithGrps = append(elemsWithGrps, c)
		}
	}
	for _, m := range b.Masks {
		if len(m.Groups) > 0 {
			elemsWithGrps = append(elemsWithGrps, m)
		}
	}
	for _, s := range b.Statuses {
		if len(s.Groups) > 0 {
			elemsWithGrps = append(elemsWithGrps, s)
		}
	}

	return elemsWithGrps
}
