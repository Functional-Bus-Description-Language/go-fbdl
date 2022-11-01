package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

func hashAccessSizes(sizes access.Sizes) uint32 {
	buf := bytes.Buffer{}

	write(&buf, sizes.BlockAligned)
	write(&buf, sizes.Compact)
	write(&buf, sizes.Own)

	return adler32.Checksum(buf.Bytes())
}

func hashAccessAddrSpace(as access.AddrSpace) uint32 {
	buf := bytes.Buffer{}

	write(&buf, as.Start())
	write(&buf, as.End())
	write(&buf, as.IsArray())
	write(&buf, as.Count())

	return adler32.Checksum(buf.Bytes())
}

func hashAccessAccess(a access.Access) uint32 {
	buf := bytes.Buffer{}

	write(&buf, a.StartAddr())
	write(&buf, a.EndAddr())
	write(&buf, a.StartBit())
	write(&buf, a.EndBit())

	return adler32.Checksum(buf.Bytes())
}
