package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashMask(m *fn.Mask) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&m.Func))

	// Atomic
	write(&buf, m.Atomic)

	// InitValue
	buf.Write([]byte(m.InitValue))

	// Width
	write(&buf, m.Width)

	// Access
	write(&buf, Hash(m.Access))

	return adler32.Checksum(buf.Bytes())
}
