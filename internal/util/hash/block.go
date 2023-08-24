package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashBlock(b *fn.Block) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&b.Func))

	// Masters
	write(&buf, b.Masters)

	// Width
	write(&buf, b.Width)

	// cnst.Container
	write(&buf, Hash(&b.Consts))

	// Configs
	for _, c := range b.Configs {
		write(&buf, Hash(c))
	}
	// Irqs
	for _, i := range b.Irqs {
		write(&buf, Hash(i))
	}
	// Masks
	for _, m := range b.Masks {
		write(&buf, Hash(m))
	}
	// Memories
	for _, m := range b.Memories {
		write(&buf, Hash(m))
	}
	// Procs
	for _, p := range b.Procs {
		write(&buf, Hash(p))
	}
	// Statics
	for _, s := range b.Statics {
		write(&buf, Hash(s))
	}
	// Statuses
	for _, s := range b.Statuses {
		write(&buf, Hash(s))
	}
	// Streams
	for _, s := range b.Streams {
		write(&buf, Hash(s))
	}
	// Subblocks
	for _, s := range b.Subblocks {
		write(&buf, Hash(s))
	}

	// Sizes
	write(&buf, Hash(b.Sizes))

	// AddrSpace
	write(&buf, Hash(b.AddrSpace))

	return adler32.Checksum(buf.Bytes())
}
