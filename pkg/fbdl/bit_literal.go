package fbdl

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
