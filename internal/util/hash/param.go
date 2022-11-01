package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashParam(p *elem.Param) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&p.Elem))

	// Groups
	for _, g := range p.Groups {
		buf.Write([]byte(g))
	}

	// Width
	write(p.Width)

	// Access
	write(Hash(p.Access))

	return adler32.Checksum(buf.Bytes())
}
