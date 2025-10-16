package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/cnst"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Block struct {
	Func

	Masters int64
	Reset   string
	Width   int64

	Sizes     types.Sizes
	AddrSpace types.SingleRange

	Consts cnst.Container

	Blackboxes []*Blackbox
	Configs    []*Config
	Groups     []*Group
	Irqs       []*Irq
	Masks      []*Mask
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
	return b.AddrSpace.Start
}
