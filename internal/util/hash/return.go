package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashReturn(r *elem.Return) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&r.Elem))

	// Groups
	for _, g := range r.Groups {
		buf.Write([]byte(g))
	}

	// Width
	write(r.Width)

	// Access
	write(Hash(r.Access))

	return adler32.Checksum(buf.Bytes())
}
