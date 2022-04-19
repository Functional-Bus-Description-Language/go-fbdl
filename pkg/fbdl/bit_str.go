package fbdl

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"strconv"
)

// BitStr (bit string) is used for representing default values.
// BitStr type is needed for 2 reasons:
//
//   1. To support default value for registers with arbitrary width.
//   2. To support meta logic values supported in Hardware Description Languages.
type BitStr string

func MakeBitStr(bs val.BitStr) BitStr { return BitStr(string(bs)) }

// BitWidth returns bit width of the bit string.
func (bs BitStr) BitWidth() int64 {
	valBs := val.BitStr(string(bs))
	return valBs.BitWidth()
}

// CharWidth returns character width of the bit string excluding format specifier and leading and trailing '"'.
func (bs BitStr) CharWidth() int64 {
	valBs := val.BitStr(string(bs))
	return valBs.CharWidth()
}

func (bs BitStr) IsBin() bool {
	if bs[0] == 'b' {
		return true
	}
	return false
}

func (bs BitStr) IsOctal() bool {
	if bs[0] == 'o' {
		return true
	}
	return false
}

func (bs BitStr) IsHex() bool {
	if bs[0] == 'x' {
		return true
	}
	return false
}

// Extend extends BitStr to given width and returns new BitStr.
// If the provided width is lesser than the current width it panics.
// Additional bits are added at the beginning and have value '0'.
// For example, extending b"1" to width 2 returns b"01".
func (bs BitStr) Extend(width int64) BitStr {
	if width < bs.BitWidth() {
		panic("cannot extend bit string width to lesser value")
	}

	if width == bs.BitWidth() {
		return bs
	}

	switch string(bs)[0] {
	case 'b':
		return extendBin(bs, width)
	case 'o':
		panic("not yet implemented")
	case 'x':
		return extendHex(bs, width)
	default:
		panic("should never happen")
	}
}

func extendBin(bs BitStr, width int64) BitStr {
	s := make([]byte, width+3)

	s[0] = 'b'
	s[1] = '"'

	widthDiff := width - bs.BitWidth()

	for i := int64(0); i < widthDiff; i++ {
		s[2+i] = '0'
	}
	for i := int64(0); i < bs.BitWidth(); i++ {
		s[2+widthDiff+i] = string(bs)[2+i]
	}

	s[len(s)-1] = '"'

	return BitStr(string(s))
}

func extendHex(bs BitStr, width int64) BitStr {
	bitWidthDiff := width - bs.BitWidth()

	if bitWidthDiff%4 == 0 {
		s := make([]byte, width/4+3)

		s[0] = 'x'
		s[1] = '"'

		for i := int64(0); i < bitWidthDiff/4; i++ {
			s[2+i] = '0'
		}
		for i := int64(0); i < bs.CharWidth(); i++ {
			s[2+bitWidthDiff/4+i] = string(bs)[2+i]
		}

		s[len(s)-1] = '"'
		return BitStr(string(s))
	}

	return extendBin(bs.ToBin(), width)
}

func (bs BitStr) ToBin() BitStr {
	if bs.IsBin() {
		return bs
	}

	s := make([]byte, bs.BitWidth()+3)
	s[0] = 'b'
	s[1] = '"'

	chunkStart := int64(0)
	chunkWidth := int64(4)
	if bs.IsOctal() {
		chunkStart = 1
		chunkWidth = 3
	}

	for i := int64(0); i < bs.CharWidth(); i++ {
		var chunk [4]byte
		char := bs[2+i]
		switch char {
		case '1':
			chunk = [4]byte{'0', '0', '0', '1'}
		case '2':
			chunk = [4]byte{'0', '0', '1', '0'}
		case '3':
			chunk = [4]byte{'0', '0', '1', '1'}
		case '4':
			chunk = [4]byte{'0', '1', '0', '0'}
		case '5':
			chunk = [4]byte{'0', '1', '0', '1'}
		case '6':
			chunk = [4]byte{'0', '1', '1', '0'}
		case '7':
			chunk = [4]byte{'0', '1', '1', '1'}
		case '8':
			chunk = [4]byte{'1', '0', '0', '0'}
		case '9':
			chunk = [4]byte{'1', '0', '0', '1'}
		case 'a', 'A':
			chunk = [4]byte{'1', '0', '1', '0'}
		case 'b', 'B':
			chunk = [4]byte{'1', '0', '1', '1'}
		case 'c', 'C':
			chunk = [4]byte{'1', '1', '0', '0'}
		case 'd', 'D':
			chunk = [4]byte{'1', '1', '0', '1'}
		case 'e', 'E':
			chunk = [4]byte{'1', '1', '1', '0'}
		case 'f', 'F':
			chunk = [4]byte{'1', '1', '1', '1'}
		case '0', 'h', 'H', 'l', 'L', 'u', 'U', 'x', 'X', 'w', 'W', 'z', 'Z', '-':
			chunk = [4]byte{char, char, char, char}
		}
		for j := int64(0); j < chunkWidth; j++ {
			s[2+chunkWidth*i+j] = chunk[chunkStart+j]
		}
	}

	s[len(s)-1] = '"'

	return BitStr(string(s))
}

// Uint64 converts bit string to uint64.
// If conversion is not possible, for example because of meta values within
// the bit string, it panics.
func (bs BitStr) Uint64() uint64 {
	base := 2
	if bs.IsOctal() {
		base = 8
	} else if bs.IsHex() {
		base = 16
	}

	u, err := strconv.ParseUint(string(bs[2:len(bs)-1]), base, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot parse bit string '%s' to uint64: %v", bs, err))
	}

	return u
}
