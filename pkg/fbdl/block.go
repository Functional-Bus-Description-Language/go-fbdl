package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
)

// Block represents block element as well as bus element.
type Block struct {
	Name      string
	Doc       string
	IsArray   bool
	Count     int64
	Sizes     Sizes
	AddrSpace AddrSpace

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

func (b *Block) addConfig(c *Config)   { b.Configs = append(b.Configs, c) }
func (b *Block) addFunc(f *Func)       { b.Funcs = append(b.Funcs, f) }
func (b *Block) addMask(m *Mask)       { b.Masks = append(b.Masks, m) }
func (b *Block) addStatus(s *Status)   { b.Statuses = append(b.Statuses, s) }
func (b *Block) addStream(s *Stream)   { b.Streams = append(b.Streams, s) }
func (b *Block) addSubblock(sb *Block) { b.Subblocks = append(b.Subblocks, sb) }

func (b *Block) addGroup(g Group) { b.Groups = append(b.Groups, g) }

func (b *Block) hasElement(name string) bool {
	for i, _ := range b.Statuses {
		if b.Statuses[i].Name == name {
			return true
		}
	}

	return false
}

func (b *Block) addConsts(insBlk *ins.Element) {
	for name, v := range insBlk.Consts {
		b.addConst(name, v)
	}
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
