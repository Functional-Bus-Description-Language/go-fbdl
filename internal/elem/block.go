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
	Configs []iface.Config
	//funcs     []iface.Func
	//masks     []iface.Mask
	//statuses  []iface.Status
	//streams   []iface.Stream
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

/*
func (b *Block) addFunc(f *Func)     { b.Funcs = append(b.Funcs, f) }
func (b *Block) Funcs() []iface.Func { return b.Funcs }

func (b *Block) addMask(m *Mask)     { b.Masks = append(b.Masks, m) }
func (b *Block) Masks() []iface.Mask { return b.Masks }

func (b *Block) addStatus(s *Status)      { b.Statuses = append(b.Statuses, s) }
func (b *Block) Statuses() []iface.Status { return b.Statuses }

func (b *Block) addStream(s *Stream)     { b.Streams = append(b.Streams, s) }
func (b *Block) Streams() []iface.Stream { return b.Streams }
*/

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
	/*
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
	*/

	return false
}

func (b *Block) Hash() int64 {
	return 0
}

// ElemsWithGroups return list of inner elements belonging to any group.
/*
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
*/
