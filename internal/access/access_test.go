package access

import (
	"reflect"
	"testing"
)

func init() {
	busWidth = 32
}

func TestMakeArrayContinuous(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		startBit  int64
		width     int64
		want      Access
	}{
		{0, 1, 0, 32, ArrayContinuous{
			ac: ac{
				StartAddr: 0,
				EndAddr: 0,
				StartMask: makeMask(0, 31),
				EndMask: makeMask(0, 31),
				ItemCount: 1,
				ItemWidth: 32,
			},
		}},
		{1, 4, 0, 5, ArrayContinuous{
			ac: ac{
				StartAddr: 1,
				EndAddr: 1,
				StartMask: makeMask(0, 4),
				EndMask: makeMask(15, 19),
				ItemCount: 4,
				ItemWidth: 5,
			},
		}},
		{2, 2, 20, 23, ArrayContinuous{
			ac: ac{
				StartAddr: 2,
				EndAddr: 4,
				StartMask: makeMask(0, 4),
				EndMask: makeMask(15, 19),
				ItemCount: 2,
				ItemWidth: 23,
			},

			ArrayContinuous{
				ItemCount: 2,
				ItemWidth: 23,
				startBit:  20,
			},
		}},
		{3, 2, 20, 22,
			ArrayContinuous{
				regCount:  2,
				ItemCount: 2,
				ItemWidth: 22,
				startAddr: 3,
				startBit:  20,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayContinuous(test.count, test.startAddr, test.startBit, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeArrayMultiplePacked(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		width     int64
		want      Access
	}{
		{0, 1, 32,
			ArrayMultiple{
				regCount:       1,
				ItemCount:      1,
				ItemWidth:      32,
				ItemsPerAccess: 1,
				startAddr:      0,
				startBit:       0,
			},
		},
		{1, 4, 8,
			ArrayMultiple{
				regCount:       1,
				ItemCount:      4,
				ItemWidth:      8,
				ItemsPerAccess: 4,
				startAddr:      1,
				startBit:       0,
			},
		},
		{2, 3, 16,
			ArrayMultiple{
				regCount:       2,
				ItemCount:      3,
				ItemWidth:      16,
				ItemsPerAccess: 2,
				startAddr:      2,
				startBit:       0,
			},
		},
		{3, 4, 4,
			ArrayMultiple{
				regCount:       1,
				ItemCount:      4,
				ItemWidth:      4,
				ItemsPerAccess: 8,
				startAddr:      3,
				startBit:       0,
			},
		},
		{4, 5, 8,
			ArrayMultiple{
				regCount:       2,
				ItemCount:      5,
				ItemWidth:      8,
				ItemsPerAccess: 4,
				startAddr:      4,
				startBit:       0,
			},
		},
		{5, 10, 7,
			ArrayMultiple{
				regCount:       3,
				ItemCount:      10,
				ItemWidth:      7,
				ItemsPerAccess: 4,
				startAddr:      5,
				startBit:       0,
			},
		},
		{6, 50, 3,
			ArrayMultiple{
				regCount:       5,
				ItemCount:      50,
				ItemWidth:      3,
				ItemsPerAccess: 10,
				startAddr:      6,
				startBit:       0,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayMultiplePacked(test.count, test.startAddr, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
