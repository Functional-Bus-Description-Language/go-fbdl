package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashConfig(c *fn.Config) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&c.Func))

	// Atomic
	write(&buf, c.Atomic)

	// InitValue
	buf.Write([]byte(c.InitValue))

	// Groups
	for _, g := range c.Groups {
		buf.Write([]byte(g))
	}

	// Width
	write(&buf, c.Width)

	// Access
	write(&buf, Hash(c.Access))

	return adler32.Checksum(buf.Bytes())
}
