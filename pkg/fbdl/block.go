package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Block represents block element as well as bus element.
type Block struct {
	Name      string
	IsArray   bool
	Count     int64
	Sizes     Sizes
	AddrSpace AddrSpace

	// Properties
	Doc     string
	Masters int64
	Width   int64

	// Constants
	IntConsts map[string]int64
	StrConsts map[string]string

	// Elements
	Configs   []*Config
	Funcs     []*Func
	Masks     []*Mask
	Statuses  []*Status
	Subblocks []*Block

	Groups []Group `json:"-"`
}

func (b *Block) addSubblock(sb *Block) { b.Subblocks = append(b.Subblocks, sb) }
func (b *Block) addConfig(c *Config)   { b.Configs = append(b.Configs, c) }
func (b *Block) addFunc(f *Func)       { b.Funcs = append(b.Funcs, f) }
func (b *Block) addMask(m *Mask)       { b.Masks = append(b.Masks, m) }
func (b *Block) addStatus(s *Status)   { b.Statuses = append(b.Statuses, s) }
func (b *Block) addGroup(g Group)      { b.Groups = append(b.Groups, g) }

func (b *Block) hasElement(name string) bool {
	for i, _ := range b.Statuses {
		if b.Statuses[i].Name == name {
			return true
		}
	}

	return false
}

func (b *Block) addIntConst(name string, value int64) {
	if b.IntConsts == nil {
		b.IntConsts = map[string]int64{name: value}
	}

	b.IntConsts[name] = value
}

func (b *Block) addStrConst(name, value string) {
	if b.StrConsts == nil {
		b.StrConsts = map[string]string{name: value}
	}
	b.StrConsts[name] = value
}

func (b *Block) addConsts(insBlk *ins.Element) {
	for name, v := range insBlk.Consts {
		switch v.(type) {
		case val.BitStr:
			panic("not yet implemented")
		case val.Bool:
			panic("not yet implemented")
		case val.Int:
			b.addIntConst(name, int64(v.(val.Int)))
		case val.List:
			panic("not yet implemented")
		case val.Str:
			b.addStrConst(name, string(v.(val.Str)))
		default:
			panic("should never happen")
		}
	}
}
