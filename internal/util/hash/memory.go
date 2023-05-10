package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashMemory(m *elem.Memory) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&m.Elem))

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
