package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashBlackbox(b *fn.Blackbox) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&b.Func))

	// Size
	write(&buf, b.Size)

	// Sizes
	write(&buf, Hash(b.Sizes))

	// AddrSpace
	write(&buf, Hash(b.AddrSpace))

	return adler32.Checksum(buf.Bytes())
}
