package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashStream(s *fn.Stream) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&s.Func))

	// Params
	for _, p := range s.Params {
		write(&buf, Hash(p))
	}

	// Returns
	for _, r := range s.Returns {
		write(&buf, Hash(r))
	}

	// StbAddr
	write(&buf, s.StbAddr)

	return adler32.Checksum(buf.Bytes())
}
