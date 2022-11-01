package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashMask(m *elem.Mask) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&m.Elem))

	// Atomic
	write(&buf, m.Atomic)

	// Default
	buf.Write([]byte(m.Default))

	// Groups
	for _, g := range m.Groups {
		buf.Write([]byte(g))
	}

	// Once
	write(&buf, m.Once)

	// Width
	write(&buf, m.Width)

	// Access
	write(&buf, Hash(m.Access))

	return adler32.Checksum(buf.Bytes())
}
