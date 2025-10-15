package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/addrSpace"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func readrressAccesses(blk *fn.Block) {
	startAddr := addrSpace.Start(blk.AddrSpace)

	for _, b := range blk.Blackboxes {
		b.AddrSpace = addrSpace.Readdress(b.AddrSpace, startAddr)

	}
	for _, c := range blk.Configs {
		c.Access.StartAddr += startAddr
		c.Access.EndAddr += startAddr
	}
	for range blk.Groups {
		panic("unimplemented")
	}
	for _, i := range blk.Irqs {
		i.Access.StartAddr += startAddr
		i.Access.EndAddr += startAddr
	}
	for _, m := range blk.Masks {
		m.Access.StartAddr += startAddr
		m.Access.EndAddr += startAddr
	}
	for _, proc := range blk.Procs {
		for _, p := range proc.Params {
			p.Access.StartAddr += startAddr
			p.Access.EndAddr += startAddr
		}
		for _, r := range proc.Returns {
			r.Access.StartAddr += startAddr
			r.Access.EndAddr += startAddr
		}
	}
	for _, s := range blk.Statics {
		s.Access.StartAddr += startAddr
		s.Access.EndAddr += startAddr
	}
	for _, s := range blk.Statuses {
		s.Access.StartAddr += startAddr
		s.Access.EndAddr += startAddr
	}
	for _, stream := range blk.Streams {
		for _, p := range stream.Params {
			p.Access.StartAddr += startAddr
			p.Access.EndAddr += startAddr
		}
		for _, r := range stream.Returns {
			r.Access.StartAddr += startAddr
			r.Access.EndAddr += startAddr
		}
	}
	for _, subblk := range blk.Subblocks {
		readrressAccesses(subblk)
	}
}
