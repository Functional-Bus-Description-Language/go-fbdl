package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Block represents block element as well as bus element.
type Block struct {
	Elem

	Sizes     access.Sizes
	AddrSpace access.AddrSpace

	// Properties
	Masters int64
	Width   int64

	// Constants
	ConstContainer

	// Elements
	Configs   []*Config
	Funcs     []*Func
	Masks     []*Mask
	Statuses  []*Status
	Streams   []*Stream
	Subblocks []*Block

	Groups []Group `json:"-"`
}

func (b *Block) Type() string { return "block" }

// Status returns pointer to the Status if status with given name exists
// within the block. Otherwise it returns nil.
func (b *Block) Status(name string) (*Status, bool) {
	for _, s := range b.Statuses {
		if s.Name() == name {
			return s, true
		}
	}
	return nil, false
}

func (b *Block) HasElement(name string) bool {
	for i, _ := range b.Configs {
		if b.Configs[i].Name() == name {
			return true
		}
	}
	for i, _ := range b.Funcs {
		if b.Funcs[i].Name == name {
			return true
		}
	}
	for i, _ := range b.Masks {
		if b.Masks[i].Name() == name {
			return true
		}
	}
	for i, _ := range b.Statuses {
		if b.Statuses[i].Name() == name {
			return true
		}
	}
	for i, _ := range b.Streams {
		if b.Streams[i].Name == name {
			return true
		}
	}
	for i, _ := range b.Subblocks {
		if b.Subblocks[i].Name() == name {
			return true
		}
	}

	return false
}

func (b *Block) Hash() int64 {
	return 0
}

// ElemsWithGroups return list of inner elements belonging to any group.
func (b *Block) ElemsWithGroups() []Groupable {
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
