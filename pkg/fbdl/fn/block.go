package fn

import (
	"encoding/json"
	"fmt"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/cnst"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Block struct {
	Func

	Align   int64
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

func (blk Block) Type() string { return "block" }

// StartAddr returns block start address.
// In case of array of blocks it returns the start address of the first block.
func (blk *Block) StartAddr() int64 {
	return blk.AddrSpace.Start
}

func (blk *Block) HasFunctionality(name string) bool {
	for i := range blk.Configs {
		if blk.Configs[i].Name == name {
			return true
		}
	}
	for i := range blk.Masks {
		if blk.Masks[i].Name == name {
			return true
		}
	}
	for i := range blk.Procs {
		if blk.Procs[i].Name == name {
			return true
		}
	}
	for i := range blk.Statics {
		if blk.Statics[i].Name == name {
			return true
		}
	}
	for i := range blk.Statuses {
		if blk.Statuses[i].Name == name {
			return true
		}
	}
	for i := range blk.Streams {
		if blk.Streams[i].Name == name {
			return true
		}
	}
	for i := range blk.Subblocks {
		if blk.Subblocks[i].Name == name {
			return true
		}
	}

	return false
}

func (blk *Block) ToJSON() []byte {
	bytes, err := json.MarshalIndent(blk, "", "  ")
	if err != nil {
		panic(fmt.Sprintf("'%s' block json marshalling failed: %v", blk.Name, err))
	}

	return bytes
}

func (blk *Block) Hash() uint32 {
	bytes := blk.ToJSON()

	return adler32.Checksum(bytes)
}
