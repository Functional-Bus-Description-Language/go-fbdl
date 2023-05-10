package block

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func AddConfig(b *elem.Block, c *elem.Config)  { b.Configs = append(b.Configs, c) }
func AddIrq(b *elem.Block, i *elem.Irq)        { b.Irqs = append(b.Irqs, i) }
func AddMask(b *elem.Block, m *elem.Mask)      { b.Masks = append(b.Masks, m) }
func AddMemory(b *elem.Block, m *elem.Memory)  { b.Memories = append(b.Memories, m) }
func AddProc(b *elem.Block, f *elem.Proc)      { b.Procs = append(b.Procs, f) }
func AddStatic(b *elem.Block, s *elem.Static)  { b.Statics = append(b.Statics, s) }
func AddStatus(b *elem.Block, s *elem.Status)  { b.Statuses = append(b.Statuses, s) }
func AddStream(b *elem.Block, s *elem.Stream)  { b.Streams = append(b.Streams, s) }
func AddSubblock(b *elem.Block, s *elem.Block) { b.Subblocks = append(b.Subblocks, s) }

func HasElement(blk *elem.Block, name string) bool {
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
