// Package hash implements hash calculation for public types.
// The types are public, however their hash functions should not be public.
package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"
	"io"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

func write(buf io.Writer, data any) {
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		panic(err)
	}
}

func AccessSizes(sizes access.Sizes) uint32 {
	buf := bytes.Buffer{}

	write(&buf, sizes.BlockAligned)
	write(&buf, sizes.Compact)
	write(&buf, sizes.Own)

	return adler32.Checksum(buf.Bytes())
}

func AccessAddrSpace(as access.AddrSpace) uint32 {
	buf := bytes.Buffer{}

	write(&buf, as.Start())
	write(&buf, as.End())
	write(&buf, as.IsArray())
	write(&buf, as.Count())

	return adler32.Checksum(buf.Bytes())
}

func AccessAccess(a access.Access) uint32 {
	buf := bytes.Buffer{}

	write(&buf, a.StartAddr())
	write(&buf, a.EndAddr())
	write(&buf, a.StartBit())
	write(&buf, a.EndBit())

	return adler32.Checksum(buf.Bytes())
}
