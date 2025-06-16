package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashStatus(s *fn.Status) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&s.Func))

	// Atomic
	write(&buf, s.Atomic)

	// Width
	write(&buf, s.Width)

	// Access
	write(&buf, Hash(s.Access))

	return adler32.Checksum(buf.Bytes())
}
