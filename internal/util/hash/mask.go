package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashMask(m *elem.Mask) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&m.Elem))

	// Atomic
	write(m.Atomic)

	// Default
	buf.Write([]byte(m.Default))

	// Groups
	for _, g := range m.Groups {
		buf.Write([]byte(g))
	}

	// Once
	write(m.Once)

	// Width
	write(m.Width)

	// Access
	write(Hash(m.Access))

	return adler32.Checksum(buf.Bytes())
}
