package access

import (
	"fmt"
)

type Mask struct {
	Upper, Lower int64
}

func (m Mask) Width() int64 { return m.Upper - m.Lower + 1 }

// Uint64 converts Mask to uint64.
// If mask can't be represented on 64 bits it panics.
// The returned mask is always shifted to the right. For example, the result for
// Mask{Upper: 2, Lower: 1} is 3 (0b11), not 6 (0b110).
func (m Mask) Uint64() uint64 {
	if m.Width() > 64 {
		panic(fmt.Sprintf("cannot convert access mask of width %d to uint64", m.Width()))
	}
	return (1 << (m.Upper - m.Lower + 1)) - 1
}
