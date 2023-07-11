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
		{ // 1
			"# Comment line",
			Stream{
				Token{Kind: COMMENT, Pos: Position{Start: 0, End: 13, Line: 1, Column: 1}},
			},
		},
		{ // 2
			"# Comment\n",
			Stream{
				Token{Kind: COMMENT, Pos: Position{Start: 0, End: 8, Line: 1, Column: 1}},
				Token{Kind: NEWLINE, Pos: Position{Start: 9, End: 9, Line: 1, Column: 10}},
			},
		},
		{ // 3
			"# Comment 1\n# Comment 2",
			Stream{
				Token{Kind: COMMENT, Pos: Position{Start: 0, End: 10, Line: 1, Column: 1}},
				Token{Kind: NEWLINE, Pos: Position{Start: 11, End: 11, Line: 1, Column: 12}},
				Token{Kind: COMMENT, Pos: Position{Start: 12, End: 22, Line: 2, Column: 1}},
			},
		},
	}

	for i, test := range tests {
		got, err := Parse([]byte(test.src))
		if err != nil {
			t.Fatalf("%d: error is not nil: %v", i, err)
		}

		for j, tok := range test.want {
			if got[j] != tok {
				t.Fatalf("\nTest %d, token %d:\n got: %+v\nwant: %+v", i, j, got[j], tok)
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
			"\n ", fmt.Errorf("2:1: space character ' ' not allowed for indent"),
		},
		{ // 1
			";\n", fmt.Errorf("1:1: extra ';' at the end of line"),
		},
		{ // 2
			" ; \n", fmt.Errorf("1:2: extra ';' at the end of line"),
		},
		{ // 3
			";;", fmt.Errorf("1:2: redundant ';'"),
		},
		{ // 4
			"b\"01-uUwWxXzZ3\"", fmt.Errorf("1:14: invalid character '3' in binary bit string literal"),
		},
		{ // 5
			"B\"0", fmt.Errorf("1:1: missing terminating '\"' in binary bit string literal"),
		},
		{ // 6
			"o\"01234567-uUwWxXzZ8\"", fmt.Errorf("1:20: invalid character '8' in octal bit string literal"),
		},
		{ // 7
			"O\"0", fmt.Errorf("1:1: missing terminating '\"' in octal bit string literal"),
		},
		{ // 8
			"x\"0123456789aAbBcCdDeEfF-uUwWxXzZ8g\"", fmt.Errorf("1:35: invalid character 'g' in hex bit string literal"),
		},
		{ // 9
			"X\"0", fmt.Errorf("1:1: missing terminating '\"' in hex bit string literal"),
		},
	}

	for i, test := range tests {
		_, err := Parse([]byte(test.src))
		if err == nil {
			t.Fatalf("%d: err = nil, expecting != nil", i)
		}

		if err.Error() != test.err.Error() {
			t.Fatalf("%d:\n got: %v\nwant: %v", i, err, test.err)
		}
	}
}
