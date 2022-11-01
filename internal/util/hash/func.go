package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashFunc(f *elem.Func) uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(Hash(&f.Elem))

	// Params
	for _, p := range f.Params {
		write(Hash(p))
	}

	// Returns
	for _, r := range f.Returns {
		write(Hash(r))
	}

	// StbAddr
	write(f.StbAddr)

	// AckAddr
	write(f.AckAddr)

	return adler32.Checksum(buf.Bytes())
}
