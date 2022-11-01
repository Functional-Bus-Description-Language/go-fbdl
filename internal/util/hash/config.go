package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashConfig(c *elem.Config) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&c.Elem))

	// Atomic
	write(c.Atomic)

	// Default
	buf.Write([]byte(c.Default))

	// Groups
	for _, g := range c.Groups {
		buf.Write([]byte(g))
	}

	// Once
	write(c.Once)

	// Width
	write(c.Width)

	// Access
	write(Hash(c.Access))

	return adler32.Checksum(buf.Bytes())
}
