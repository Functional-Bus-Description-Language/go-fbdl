package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// BitStr (bit string) is used for representing default values.
// BitStr type is needed for 2 reasons:
//
//   1. To support default value for registers with arbitrary width.
//   2. To support meta logic values supported in Hardware Description Languages.
type BitStr string

func MakeBitStr(bs val.BitStr) BitStr { return BitStr(string(bs)) }

func (bs BitStr) IsBin() bool {
	if bs[0] == 'b' || bs[0] == 'B' {
		return true
	}

	return false
}

func (bs BitStr) IsOctal() bool {
	if bs[0] == 'o' || bs[0] == 'O' {
		return true
	}

	return false
}

func (bs BitStr) IsHex() bool {
	if bs[0] == 'x' || bs[0] == 'X' {
		return true
	}

	return false
}
