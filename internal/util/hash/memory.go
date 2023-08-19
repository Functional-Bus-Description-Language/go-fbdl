package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashMemory(m *fn.Memory) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&m.Func))

	// Access
	buf.Write([]byte(m.Access))

	// ByteWriteEnable
	write(&buf, m.ByteWriteEnable)

	// ReadLatency
	write(&buf, m.ReadLatency)

	// Size
	write(&buf, m.Size)

	// Width
	write(&buf, m.Width)

	return adler32.Checksum(buf.Bytes())
}
