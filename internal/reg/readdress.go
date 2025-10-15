package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func readrressAccesses(blk *fn.Block) {
	offset := blk.AddrSpace.Start

	for _, b := range blk.Blackboxes {
		b.AddrSpace = b.AddrSpace.Shift(offset)

	}
	for _, c := range blk.Configs {
		c.Access.StartAddr += offset
		c.Access.EndAddr += offset
	}
	for range blk.Groups {
		panic("unimplemented")
	}
	for _, i := range blk.Irqs {
		i.Access.StartAddr += offset
		i.Access.EndAddr += offset
	}
	for _, m := range blk.Masks {
		m.Access.StartAddr += offset
		m.Access.EndAddr += offset
	}
	for _, proc := range blk.Procs {
		for _, p := range proc.Params {
			p.Access.StartAddr += offset
			p.Access.EndAddr += offset
		}
		for _, r := range proc.Returns {
			r.Access.StartAddr += offset
			r.Access.EndAddr += offset
		}
	}
	for _, s := range blk.Statics {
		s.Access.StartAddr += offset
		s.Access.EndAddr += offset
	}
	for _, s := range blk.Statuses {
		s.Access.StartAddr += offset
		s.Access.EndAddr += offset
	}
	for _, stream := range blk.Streams {
		for _, p := range stream.Params {
			p.Access.StartAddr += offset
			p.Access.EndAddr += offset
		}
		for _, r := range stream.Returns {
			r.Access.StartAddr += offset
			r.Access.EndAddr += offset
		}
	}
	for _, subblk := range blk.Subblocks {
		readrressAccesses(subblk)
	}
}
