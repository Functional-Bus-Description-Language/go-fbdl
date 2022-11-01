package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashStream(s *elem.Stream) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&s.Elem))

	// Params
	for _, p := range s.Params {
		write(Hash(p))
	}

	// Returns
	for _, r := range s.Returns {
		write(Hash(r))
	}

	// StbAddr
	write(s.StbAddr)

	return adler32.Checksum(buf.Bytes())
}
