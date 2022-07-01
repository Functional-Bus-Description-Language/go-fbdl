package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/iface"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type blk struct {
	Elem

	// Properties
	Masters int64
	Width   int64

	// Constants
	//ConstContainer

	// Elements
	Configs   []iface.Config
	Funcs     []iface.Func
	Masks     []iface.Mask
	Statuses  []iface.Status
	Streams   []iface.Stream
	Subblocks []iface.Block

	Sizes     access.Sizes
	AddrSpace access.AddrSpace

	//Groups []Group `json:"-"`
}

// Block represents block element as well as bus element.
type Block struct {
	blk
}

func (b *Block) Type() string { return "block" }

func (b *Block) SetMasters(m int64) { b.blk.Masters = m }
func (b *Block) Masters() int64     { return b.blk.Masters }

func (b *Block) SetWidth(m int64) { b.blk.Width = m }
func (b *Block) Width() int64     { return b.blk.Width }

func (b *Block) AddConfig(c *Config)     { b.blk.Configs = append(b.blk.Configs, c) }
func (b *Block) Configs() []iface.Config { return b.blk.Configs }

func (b *Block) AddFunc(f *Func)     { b.blk.Funcs = append(b.blk.Funcs, f) }
func (b *Block) Funcs() []iface.Func { return b.blk.Funcs }

func (b *Block) AddMask(m *Mask)     { b.blk.Masks = append(b.blk.Masks, m) }
func (b *Block) Masks() []iface.Mask { return b.blk.Masks }

func (b *Block) AddStatus(s *Status)      { b.blk.Statuses = append(b.blk.Statuses, s) }
func (b *Block) Statuses() []iface.Status { return b.blk.Statuses }

func (b *Block) AddStream(s *Stream)     { b.blk.Streams = append(b.blk.Streams, s) }
func (b *Block) Streams() []iface.Stream { return b.blk.Streams }

func (b *Block) addSubblock(sb *Block)    { b.blk.Subblocks = append(b.blk.Subblocks, sb) }
func (b *Block) Subblocks() []iface.Block { return b.blk.Subblocks }

func (b *Block) SetSizes(s access.Sizes) { b.blk.Sizes = s }
func (b *Block) Sizes() access.Sizes     { return b.blk.Sizes }

func (b *Block) SetAddrSpace(as access.AddrSpace) { b.blk.AddrSpace = as }
func (b *Block) AddrSpace() access.AddrSpace      { return b.blk.AddrSpace }

// Status returns pointer to the Status if status with given name exists
// within the block. Otherwise it returns nil.
/*
func (b *Block) Status(name string) (*Status, bool) {
	for _, s := range b.Statuses {
		if s.Name() == name {
			return s, true
		}
	}
	return nil, false
}
*/

func (b *Block) HasElement(name string) bool {
	for i, _ := range b.blk.Configs {
		if b.blk.Configs[i].Name() == name {
			return true
		}
	}
	for i, _ := range b.blk.Funcs {
		if b.blk.Funcs[i].Name() == name {
			return true
		}
	}
	for i, _ := range b.blk.Masks {
		if b.blk.Masks[i].Name() == name {
			return true
		}
	}
	for i, _ := range b.blk.Statuses {
		if b.blk.Statuses[i].Name() == name {
			return true
		}
	}
	for i, _ := range b.blk.Streams {
		if b.blk.Streams[i].Name() == name {
			return true
		}
	}
	/*
		for i, _ := range b.Subblocks {
			if b.Subblocks[i].Name() == name {
				return true
			}
		}
	*/

	return false
}

func (b *Block) Hash() int64 {
	return 0
}

// ElemsWithGroups return list of inner elements belonging to any group.
func (b *Block) ElemsWithGroups() []iface.Groupable {
	elemsWithGrps := []iface.Groupable{}
	for _, c := range b.blk.Configs {
		if len(c.Groups()) > 0 {
			elemsWithGrps = append(elemsWithGrps, c)
		}
	}
	for _, m := range b.blk.Masks {
		if len(m.Groups()) > 0 {
			elemsWithGrps = append(elemsWithGrps, m)
		}
	}
	for _, s := range b.blk.Statuses {
		if len(s.Groups()) > 0 {
			elemsWithGrps = append(elemsWithGrps, s)
		}
	}

	return elemsWithGrps
}
