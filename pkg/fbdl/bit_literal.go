package fbdl

// BitLiteral is used for representing default values.
// BitLiteral type is needed for 2 reasons:
//
//   1. To support default value for registers with arbitrary width.
//   2. To support meta logic values supported in Hardware Description Languages.
type BitLiteral string

func (bt BitLiteral) IsBin() bool {
	if bt[0] == 'b' || bt[0] == 'B' {
		return true
	}

	return false
}

func (bt BitLiteral) IsOctal() bool {
	if bt[0] == 'o' || bt[0] == 'O' {
		return true
	}

	return false
}

func (bt BitLiteral) IsHex() bool {
	if bt[0] == 'x' || bt[0] == 'X' {
		return true
	}

	return false
}
