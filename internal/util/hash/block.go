package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashBlock(b *elem.Block) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&b.Elem))

	// Masters
	write(&buf, b.Masters)

	// Width
	write(&buf, b.Width)

	// ConstContainer
	write(&buf, Hash(&b.ConstContainer))

	// Configs
	for _, c := range b.Configs {
		write(&buf, Hash(c))
	}
	// Masks
	for _, m := range b.Masks {
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
