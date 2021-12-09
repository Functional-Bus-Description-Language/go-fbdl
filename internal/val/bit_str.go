package val

import (
	"fmt"
	"math"
	"strconv"
)

// BitStr represents FBDL bit string type.
type BitStr string

func (bs BitStr) Type() string {
	return "bit string"
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
			break
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
			break
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
			break
		default:
			return BitStr(""), fmt.Errorf("invalid character '%c' in hex bit literal", s[i])
		}
	}

	return BitStr("x" + s[1:]), nil
}

// BitStrFromInt converts val.Int to BitStr.
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

	if i > 0 {
		var s string
		if width%4 == 0 {
			s = "x\"" + strconv.FormatInt(i, 16) + "\""
		} else if width%3 == 0 {
			s = "o\"" + strconv.FormatInt(i, 8) + "\""
		} else {
			s = "b\"" + strconv.FormatInt(i, 2) + "\""
		}

		return BitStr(s), nil
	}

	// Negative value handling
	panic("BitStrFromInt, negative value handling not yet implemented")

	//return BitStr(""), nil
}
