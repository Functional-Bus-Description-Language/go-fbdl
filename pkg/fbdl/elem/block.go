package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Block represents block element as well as bus element.
type Block struct {
	Name      string
	Doc       string
	IsArray   bool
	Count     int64
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

// Status returns pointer to the Status if status with given name exists
// within the block. Otherwise it returns nil.
func (b *Block) Status(name string) *Status {
	for _, s := range b.Statuses {
		if s.Name == name {
			return s
		}
	}
	return nil
}

// TODO: Check if it is used anywhere.
func (b *Block) HasElement(name string) bool {
	for i, _ := range b.Statuses {
		if b.Statuses[i].Name == name {
			return true
		}
	}

	return false
}
