package block

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func AddBlackbox(b *fn.Block, bb *fn.Blackbox) { b.Blackboxes = append(b.Blackboxes, bb) }
func AddConfig(b *fn.Block, c *fn.Config)      { b.Configs = append(b.Configs, c) }
func AddIrq(b *fn.Block, i *fn.Irq)            { b.Irqs = append(b.Irqs, i) }
func AddMask(b *fn.Block, m *fn.Mask)          { b.Masks = append(b.Masks, m) }
func AddProc(b *fn.Block, f *fn.Proc)          { b.Procs = append(b.Procs, f) }
func AddStatic(b *fn.Block, s *fn.Static)      { b.Statics = append(b.Statics, s) }
func AddStatus(b *fn.Block, s *fn.Status)      { b.Statuses = append(b.Statuses, s) }
func AddStream(b *fn.Block, s *fn.Stream)      { b.Streams = append(b.Streams, s) }
func AddSubblock(b *fn.Block, s *fn.Block)     { b.Subblocks = append(b.Subblocks, s) }

func HasFunctionality(blk *fn.Block, name string) bool {
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
