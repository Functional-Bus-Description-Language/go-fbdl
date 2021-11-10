package val

import (
	"fmt"
)

// BitStr represents FBDL bit string type.
type BitStr string

func (bs BitStr) Type() string {
	return "bit literal"
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
		case '0', '1', 'h', 'H', 'l', 'L', 'u', 'U', 'x', 'X', 'w', 'W', 'z', 'Z', '-':
			break
		default:
			return BitStr(""), fmt.Errorf("invalid character '%c' in binary bit literal", s[i])
		}
	}

	return BitStr(s), nil
}

func makeOctalBitStr(s string) (BitStr, error) {
	return BitStr(""), fmt.Errorf("makeOctalBitStr not yet implemented")
}

func makeHexBitStr(s string) (BitStr, error) {
	return BitStr(""), fmt.Errorf("makeHexBitStr not yet implemented")
}
