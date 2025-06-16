package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/addrSpace"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/cnst"
)

type Block struct {
	Func

	Masters int64
	Reset   string
	Width   int64

	Sizes     access.Sizes
	AddrSpace addrSpace.AddrSpace

	Consts cnst.Container

	Blackboxes []*Blackbox
	Configs    []*Config
	Irqs       []*Irq
	Masks      []*Mask
	Memories   []*Memory
	Procs      []*Proc
	Statics    []*Static
	Statuses   []*Status
	Streams    []*Stream
	Subblocks  []*Block
}

func (b Block) Type() string { return "block" }

// StartAddr returns block start address.
// In case of array of blocks it returns the start address of the first block.
func (b *Block) StartAddr() int64 {
	return addrSpace.Start(b.AddrSpace)
}
