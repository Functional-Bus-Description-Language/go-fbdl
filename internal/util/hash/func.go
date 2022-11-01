package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashFunc(f *elem.Func) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&f.Elem))

	// Params
	for _, p := range f.Params {
		write(&buf, Hash(p))
	}

	// Returns
	for _, r := range f.Returns {
		write(&buf, Hash(r))
	}

	// StbAddr
	write(&buf, f.StbAddr)

	// AckAddr
	write(&buf, f.AckAddr)

	return adler32.Checksum(buf.Bytes())
}
