package token

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		src  string
		want Stream
	}{
		{ // 0
			"\n\n",
			Stream{
				Token{Kind: NEWLINE, Pos: Position{Start: 0, End: 0, Line: 1, Column: 1}},
				Token{Kind: NEWLINE, Pos: Position{Start: 1, End: 1, Line: 2, Column: 1}},
			},
		},
	}

	for i, test := range tests {
		got, err := Parse([]byte(test.src))
		if err != nil {
			t.Errorf("%d: error is not nil: %v", i, err)
		}

		for j, tok := range test.want {
			if got[j] != tok {
				t.Errorf("\nTest %d, token %d:\n got: %+v\nwant: %+v", i, j, got[j], tok)
			}
		}
	}
}

func TestParseError(t *testing.T) {
	var tests = []struct {
		src string
		err error
	}{
		{ // 0
			"\n ",
			fmt.Errorf("2:1: space character ' ' not allowed for indent"),
		},
	}

	for i, test := range tests {
		_, err := Parse([]byte(test.src))
		if err == nil {
			t.Errorf("%d: err = nil, expecting != nil", i)
		}

		if err.Error() != test.err.Error() {
			t.Errorf("%d:\n got: %v\nwant: %v", i, err, test.err)
		}
	}
}
