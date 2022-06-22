package access

import (
	"fmt"
)

type Mask struct {
	Upper, Lower int64
}

func (am Mask) Width() int64 { return am.Upper - am.Lower + 1 }

// Uint64 converts Mask to uint64.
// If mask can't be represented on 64 bits it panics.
// The returned mask is always shifted to the right. For example, the result for
// Mask{Upper: 2, Lower: 1} is 3 (0b11), not 6 (0b110).
func (am Mask) Uint64() uint64 {
	if am.Width() > 64 {
		panic(fmt.Sprintf("cannot convert access mask of width %d to uint64", am.Width()))
	}
	return (1 << (am.Upper - am.Lower + 1)) - 1
}
