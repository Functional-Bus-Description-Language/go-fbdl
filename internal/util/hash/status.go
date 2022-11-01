package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashStatus(s *elem.Status) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&s.Elem))

	// Atomic
	write(&buf, s.Atomic)

	// Groups
	for _, g := range s.Groups {
		buf.Write([]byte(g))
	}

	// Once
	write(&buf, s.Once)

	// Width
	write(&buf, s.Width)

	// Access
	write(&buf, Hash(s.Access))

	return adler32.Checksum(buf.Bytes())
}
