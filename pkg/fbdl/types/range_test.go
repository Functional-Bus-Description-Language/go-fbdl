package types

import "testing"

func TestSingleRangeWidth(t *testing.T) {
	var tests = []struct {
		sr   SingleRange
		want int64
	}{
		{SingleRange{Start: 0, End: 1}, 1},
		{SingleRange{Start: 0, End: 14}, 4},
		{SingleRange{Start: 0, End: 15}, 4},
		{SingleRange{Start: 0, End: 16}, 5},
		{SingleRange{Start: 130, End: 255}, 8},
		{SingleRange{Start: 245, End: 256}, 9},
	}

	for i, test := range tests {
		if test.sr.BitWidth() != test.want {
			t.Fatalf("%d: got %d, want %d", i, test.sr.BitWidth(), test.want)
		}
	}
}

func TestArrayRangeWidth(t *testing.T) {
	var tests = []struct {
		mr   ArrayRange
		want int64
	}{
		{
			ArrayRange{
				SingleRange{Start: 0, End: 1}, SingleRange{Start: 0, End: 15},
			},
			4,
		},
		{
			ArrayRange{
				SingleRange{Start: 0, End: 1023}, SingleRange{Start: 400, End: 510},
			},
			10,
		},
		{
			ArrayRange{
				SingleRange{Start: 0, End: 7}, SingleRange{Start: 10, End: 36}, SingleRange{Start: 40, End: 250},
			},
			8,
		},
	}

	for i, test := range tests {
		if test.mr.BitWidth() != test.want {
			t.Fatalf("%d: got %d, want %d", i, test.mr.BitWidth(), test.want)
		}
	}
}
