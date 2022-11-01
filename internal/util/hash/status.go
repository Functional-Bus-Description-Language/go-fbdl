package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashStatus(s *elem.Status) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&s.Elem))

	// Atomic
	write(s.Atomic)

	// Groups
	for _, g := range s.Groups {
		buf.Write([]byte(g))
	}

	// Once
	write(s.Once)

	// Width
	write(s.Width)

	// Access
	write(Hash(s.Access))

	return adler32.Checksum(buf.Bytes())
}
