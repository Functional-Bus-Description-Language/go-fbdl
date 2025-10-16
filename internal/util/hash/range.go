package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/value"
)

func hashRange(ran value.Range) uint32 {
	switch r := ran.(type) {
	case value.SingleRange:
		return hashSingleRange(r)
	case value.MultiRange:
		panic("unimplemented")
	default:
		panic("unimplemented")
	}
}

func hashSingleRange(sr value.SingleRange) uint32 {
	buf := bytes.Buffer{}

	write(&buf, sr.Start)
	write(&buf, sr.End)

	return adler32.Checksum(buf.Bytes())
}
