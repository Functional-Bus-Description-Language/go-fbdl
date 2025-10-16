package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

func hashAccessSizes(sizes types.Sizes) uint32 {
	buf := bytes.Buffer{}

	write(&buf, sizes.BlockAligned)
	write(&buf, sizes.Compact)
	write(&buf, sizes.Own)

	return adler32.Checksum(buf.Bytes())
}

func hashAccessAccess(acs types.Access) uint32 {
	buf := bytes.Buffer{}

	write(&buf, acs.Type)

	write(&buf, acs.RegCount)
	write(&buf, acs.RegWidth)

	write(&buf, acs.ItemCount)
	write(&buf, acs.ItemWidth)

	write(&buf, acs.StartAddr)
	write(&buf, acs.EndAddr)

	write(&buf, acs.StartBit)
	write(&buf, acs.EndBit)

	write(&buf, acs.StartRegWidth)
	write(&buf, acs.EndRegWidth)

	return adler32.Checksum(buf.Bytes())
}
