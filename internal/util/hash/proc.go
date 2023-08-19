package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashProc(p *fn.Proc) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&p.Func))

	// Params
	for _, p := range p.Params {
		write(&buf, Hash(p))
	}

	// Returns
	for _, r := range p.Returns {
		write(&buf, Hash(r))
	}

	// CallAddr
	if p.CallAddr != nil {
		write(&buf, p.CallAddr)
	}

	// ExitAddr
	if p.ExitAddr != nil {
		write(&buf, p.ExitAddr)
	}

	return adler32.Checksum(buf.Bytes())
}
