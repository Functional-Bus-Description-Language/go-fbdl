package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/addrSpace"
)

type Block struct {
	Elem

	Masters int64
	Reset   string
	Width   int64

	Sizes     access.Sizes
	AddrSpace addrSpace.AddrSpace

	ConstContainer

	Configs   []*Config
	Masks     []*Mask
	Memories  []*Memory
	Procs     []*Proc
	Statics   []*Static
	Statuses  []*Status
	Streams   []*Stream
	Subblocks []*Block
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

// StartAddr returns block start address.
// In case of array of blocks it returns the start address of the first block.
func (b *Block) StartAddr() int64 {
	return addrSpace.Start(b.AddrSpace)
}
