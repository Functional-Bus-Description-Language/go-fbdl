package val

import (
	"testing"
)

func TestToBin(t *testing.T) {
	var tests = []struct {
		in   BitStr
		want BitStr
	}{
		{in: BitStr(`b"0101"`), want: BitStr(`b"0101"`)},
		{in: BitStr(`o"234"`), want: BitStr(`b"010011100"`)},
		{in: BitStr(`o"77"`), want: BitStr(`b"111111"`)},
		{in: BitStr(`o"22"`), want: BitStr(`b"010010"`)},
		{in: BitStr(`o"hh"`), want: BitStr(`b"hhhhhh"`)},
		{in: BitStr(`o"uu"`), want: BitStr(`b"uuuuuu"`)},
		{in: BitStr(`o"WW"`), want: BitStr(`b"WWWWWW"`)},
		{in: BitStr(`x"1"`), want: BitStr(`b"0001"`)},
		{in: BitStr(`x"ab"`), want: BitStr(`b"10101011"`)},
		{in: BitStr(`x"cd"`), want: BitStr(`b"11001101"`)},
		{in: BitStr(`x"ef"`), want: BitStr(`b"11101111"`)},
		{in: BitStr(`x"-"`), want: BitStr(`b"----"`)},
		{in: BitStr(`x"LL"`), want: BitStr(`b"LLLLLLLL"`)},
	}

	for i, test := range tests {
		got := test.in.ToBin()

		if got != test.want {
			t.Errorf("[%d]: got %v, want %v", i, got, test.want)
		}
	}
}

func TestExtendBin(t *testing.T) {
	var tests = []struct {
		in            BitStr
		extendedWidth int64
		want          BitStr
	}{
		{in: BitStr(`b"0"`), extendedWidth: 2, want: BitStr(`b"00"`)},
		{in: BitStr(`b"1"`), extendedWidth: 2, want: BitStr(`b"01"`)},
		{in: BitStr(`b"111"`), extendedWidth: 3, want: BitStr(`b"111"`)},
		{in: BitStr(`b"101"`), extendedWidth: 5, want: BitStr(`b"00101"`)},
	}

	for i, test := range tests {
		got := test.in.Extend(test.extendedWidth)

		if got != test.want {
			t.Errorf("[%d]: got %v, want %v", i, got, test.want)
		}
	}
}

func TestExtendHex(t *testing.T) {
	var tests = []struct {
		in            BitStr
		extendedWidth int64
		want          BitStr
	}{
		{in: BitStr(`x"0"`), extendedWidth: 8, want: BitStr(`x"00"`)},
		{in: BitStr(`x"F"`), extendedWidth: 8, want: BitStr(`x"0F"`)},
		{in: BitStr(`x"a0"`), extendedWidth: 12, want: BitStr(`x"0a0"`)},
		{in: BitStr(`x"abcd"`), extendedWidth: 20, want: BitStr(`x"0abcd"`)},
		{in: BitStr(`x"f"`), extendedWidth: 5, want: BitStr(`b"01111"`)},
		{in: BitStr(`x"u"`), extendedWidth: 7, want: BitStr(`b"000uuuu"`)},
		{in: BitStr(`x"a0"`), extendedWidth: 10, want: BitStr(`b"0010100000"`)},
	}

	for i, test := range tests {
		got := test.in.Extend(test.extendedWidth)

		if got != test.want {
			t.Errorf("[%d]: got %v, want %v", i, got, test.want)
		}
	}
}
