package value

import "testing"

func TestSingleRangeWidth(t *testing.T) {
	var tests = []struct {
		sr   SingleRange
		want int64
	}{
		{SingleRange{Left: 0, Right: 1}, 1},
		{SingleRange{Left: 0, Right: 14}, 4},
		{SingleRange{Left: 0, Right: 15}, 4},
		{SingleRange{Left: 0, Right: 16}, 5},
		{SingleRange{Left: 130, Right: 255}, 8},
		{SingleRange{Left: 245, Right: 256}, 9},
	}

	for i, test := range tests {
		if test.sr.Width() != test.want {
			t.Fatalf("%d: got %d, want %d", i, test.sr.Width(), test.want)
		}
	}
}

func TestMultiRangeWidth(t *testing.T) {
	var tests = []struct {
		mr   MultiRange
		want int64
	}{
		{
			MultiRange{
				SingleRange{Left: 0, Right: 1}, SingleRange{Left: 0, Right: 15},
			},
			4,
		},
		{
			MultiRange{
				SingleRange{Left: 0, Right: 1023}, SingleRange{Left: 400, Right: 510},
			},
			10,
		},
		{
			MultiRange{
				SingleRange{Left: 0, Right: 7}, SingleRange{Left: 10, Right: 36}, SingleRange{Left: 40, Right: 250},
			},
			8,
		},
	}

	for i, test := range tests {
		if test.mr.Width() != test.want {
			t.Fatalf("%d: got %d, want %d", i, test.mr.Width(), test.want)
		}
	}
}
