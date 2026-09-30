package types

import (
	"fmt"
	"math"
	"strconv"
)

// BitStr represents bit string type from the FBDL specification.
//
// BitStr (bit string) is used for representing init values.
// BitStr type is needed for 2 reasons:
//
//  1. To support default value for registers with arbitrary width.
//  2. To support meta logic values supported in Hardware Description Languages.
type BitStr string

func (bs BitStr) Type() string {
	return "bit string"
}

func (bs BitStr) Bytes() []byte {
	return []byte(bs)
}

// BitWidth returns bit width of the bit string.
func (bs BitStr) BitWidth() int64 {
	var width int64

	width = int64(len(bs)) - 3

	switch string(bs)[0] {
	case 'b':
		width *= 1
	case 'o':
		width *= 3
	case 'x':
		width *= 4
	default:
		panic("should never happen")
	}

	return width
}

// CharWidth returns character width of the bit string excluding format specifier and leading and trailing '"'.
func (bs BitStr) CharWidth() int64 {
	return int64(len(bs)) - 3
}

func (bs BitStr) IsBin() bool {
	return bs[0] == 'b'
}

func (bs BitStr) IsOctal() bool {
	return bs[0] == 'o'
}

func (bs BitStr) IsHex() bool {
	return bs[0] == 'x'
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

	for i := range widthDiff {
		s[2+i] = '0'
	}
	for i := range bs.BitWidth() {
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

		for i := range bitWidthDiff / 4 {
			s[2+i] = '0'
		}
		for i := range bs.CharWidth() {
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

	for i := range bs.CharWidth() {
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
		for j := range chunkWidth {
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

// ValueLiteral returns the internal value pattern of bit string represented as a string.
// For example, ValueLiteral for x"AB" returns AB, for b"1100" returns 1100.
func (bs BitStr) ValueLiteral() string {
	return string(bs[2 : len(bs)-1])
}

func MakeBitStr(s string) (BitStr, error) {
	format := s[0]
	bs := BitStr("")
	var err error

	switch format {
	case 'b', 'B', 'o', 'O', 'x', 'X':
		break
	default:
		return bs, fmt.Errorf("invalid bit literal format '%c'", format)
	}

	if s[1] != '"' {
		return bs, fmt.Errorf("missing '\"' at beginning of bit literal")
	}

	if s[len(s)-1] != '"' {
		return bs, fmt.Errorf("missing '\"' at end of bit literal")
	}

	switch format {
	case 'b', 'B':
		bs, err = makeBinBitStr(s)
		if err != nil {
			return bs, fmt.Errorf("make bit literal: %v", err)
		}
	case 'o', 'O':
		bs, err = makeOctalBitStr(s)
		if err != nil {
			return bs, fmt.Errorf("make bit literal: %v", err)
		}
	case 'x', 'X':
		bs, err = makeHexBitStr(s)
		if err != nil {
			return bs, fmt.Errorf("make bit literal: %v", err)
		}
	}

	return bs, nil
}

func makeBinBitStr(s string) (BitStr, error) {
	for i := 2; i < len(s)-1; i++ {
		switch s[i] {
		case '0', '1':
		case 'h', 'H', 'l', 'L', 'u', 'U', 'x', 'X', 'w', 'W', 'z', 'Z', '-':
			continue
		default:
			return BitStr(""), fmt.Errorf("invalid character '%c' in binary bit literal", s[i])
		}
	}

	return BitStr("b" + s[1:]), nil
}

func makeOctalBitStr(s string) (BitStr, error) {
	for i := 2; i < len(s)-1; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7':
		case 'h', 'H', 'l', 'L', 'u', 'U', 'x', 'X', 'w', 'W', 'z', 'Z', '-':
			continue
		default:
			return BitStr(""), fmt.Errorf("invalid character '%c' in hex bit literal", s[i])
		}
	}

	return BitStr("o" + s[1:]), nil
}

func makeHexBitStr(s string) (BitStr, error) {
	for i := 2; i < len(s)-1; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case 'a', 'b', 'c', 'd', 'e', 'f':
		case 'A', 'B', 'C', 'D', 'E', 'F':
		case 'h', 'H', 'l', 'L', 'u', 'U', 'x', 'X', 'w', 'W', 'z', 'Z', '-':
			continue
		default:
			return BitStr(""), fmt.Errorf("invalid character '%c' in hex bit literal", s[i])
		}
	}

	return BitStr("x" + s[1:]), nil
}

// BitStrFromInt converts Int to BitStr.
// It only checks whether given value can be represented with given width.
// It uses U2 encoding for negative values.
func BitStrFromInt(v Int, width int64) (BitStr, error) {
	i := int64(v)

	max := int64(math.Pow(float64(2), float64(width))) - int64(1)
	min := -int64(math.Pow(float64(2), float64(width-1)))

	if i > max {
		return BitStr(""),
			fmt.Errorf(
				"value %d is too large to be converted to bit string of width %d, max = %d",
				i, width, max,
			)
	} else if i < min {
		return BitStr(""),
			fmt.Errorf(
				"value %d is too small to be converted to bit string of width %d, min = %d",
				i, width, min,
			)
	}

	if i >= 0 {
		var s string
		if width%4 == 0 {
			s = fmt.Sprintf("x\"%0*x\"", width/4, i)
		} else if width%3 == 0 {
			s = fmt.Sprintf("o\"%0*o\"", width/3, i)
		} else {
			s = fmt.Sprintf("b\"%0*b\"", width, i)
		}

		return BitStr(s), nil
	}

	// Negative value handling
	panic("BitStrFromInt, negative value handling not yet implemented")

	//return BitStr(""), nil
}
