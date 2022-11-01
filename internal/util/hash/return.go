package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashReturn(r *elem.Return) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&r.Elem))

	// Groups
	for _, g := range r.Groups {
		buf.Write([]byte(g))
	}

	// Width
	write(&buf, r.Width)

	// Access
	write(&buf, Hash(r.Access))

	return adler32.Checksum(buf.Bytes())
}
