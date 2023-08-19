package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashIrq(i *fn.Irq) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&i.Func))

	// AddEnable
	write(&buf, i.AddEnable)

	// Clear
	buf.Write([]byte(i.Clear))

	// EnableInitValue
	buf.Write([]byte(i.EnableInitValue))

	// EnableResetValue
	buf.Write([]byte(i.EnableResetValue))

	// Groups
	for _, g := range i.Groups {
		buf.Write([]byte(g))
	}

	// InTrigger
	buf.Write([]byte(i.InTrigger))

	// OutTrigger
	buf.Write([]byte(i.OutTrigger))

	// Access
	write(&buf, Hash(i.Access))

	return adler32.Checksum(buf.Bytes())
}
