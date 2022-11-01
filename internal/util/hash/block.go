package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashBlock(b *elem.Block) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&b.Elem))

	// Masters
	write(b.Masters)

	// Width
	write(b.Width)

	// ConstContainer
	write(Hash(&b.ConstContainer))

	// Configs
	for _, c := range b.Configs {
		write(Hash(c))
	}
	// Funcs
	for _, f := range b.Funcs {
		write(Hash(f))
	}
	// Masks
	for _, m := range b.Masks {
		write(Hash(m))
	}
	// Statics
	for _, s := range b.Statics {
		write(Hash(s))
	}
	// Statuses
	for _, s := range b.Statuses {
		write(Hash(s))
	}
	// Streams
	for _, s := range b.Streams {
		write(Hash(s))
	}
	// Subblocks
	for _, s := range b.Subblocks {
		write(Hash(s))
	}

	// Sizes
	write(Hash(b.Sizes))

	// AddrSpace
	write(Hash(b.AddrSpace))

	return adler32.Checksum(buf.Bytes())
}
