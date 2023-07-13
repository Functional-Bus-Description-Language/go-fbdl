package token

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		idx  int // Test index, useful for navigation
		src  string
		want Stream
	}{
		{
			0,
			"\n\n",
			Stream{
				Token{Kind: NEWLINE, Pos: Position{Start: 0, End: 0, Line: 1, Column: 1}},
				Token{Kind: NEWLINE, Pos: Position{Start: 1, End: 1, Line: 2, Column: 1}},
			},
		},
		{
			1,
			"# Comment line",
			Stream{
				Token{Kind: COMMENT, Pos: Position{Start: 0, End: 13, Line: 1, Column: 1}},
			},
		},
		{
			2,
			"# Comment\n",
			Stream{
				Token{Kind: COMMENT, Pos: Position{Start: 0, End: 8, Line: 1, Column: 1}},
				Token{Kind: NEWLINE, Pos: Position{Start: 9, End: 9, Line: 1, Column: 10}},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			Stream{
				Token{Kind: COMMENT, Pos: Position{Start: 0, End: 10, Line: 1, Column: 1}},
				Token{Kind: NEWLINE, Pos: Position{Start: 11, End: 11, Line: 1, Column: 12}},
				Token{Kind: COMMENT, Pos: Position{Start: 12, End: 22, Line: 2, Column: 1}},
			},
		},
		{
			4,
			"const A = true",
			Stream{
				Token{Kind: CONST, Pos: Position{Start: 0, End: 4, Line: 1, Column: 1}},
				Token{Kind: IDENT, Pos: Position{Start: 6, End: 6, Line: 1, Column: 7}},
				Token{Kind: ASS, Pos: Position{Start: 8, End: 8, Line: 1, Column: 9}},
				Token{Kind: BOOL, Pos: Position{Start: 10, End: 13, Line: 1, Column: 11}},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			Stream{
				Token{Kind: IDENT, Pos: Position{Start: 0, End: 2, Line: 1, Column: 1}},
				Token{Kind: MASK, Pos: Position{Start: 4, End: 7, Line: 1, Column: 5}},
				Token{Kind: SEMICOLON, Pos: Position{Start: 8, End: 8, Line: 1, Column: 9}},
				Token{Kind: ATOMIC, Pos: Position{Start: 10, End: 15, Line: 1, Column: 11}},
				Token{Kind: ASS, Pos: Position{Start: 17, End: 17, Line: 1, Column: 18}},
				Token{Kind: BOOL, Pos: Position{Start: 19, End: 23, Line: 1, Column: 20}},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			Stream{
				Token{Kind: IDENT, Pos: Position{Start: 0, End: 0, Line: 1, Column: 1}},
				Token{Kind: IRQ, Pos: Position{Start: 2, End: 4, Line: 1, Column: 3}},
				Token{Kind: SEMICOLON, Pos: Position{Start: 5, End: 5, Line: 1, Column: 6}},
				Token{Kind: ADD_ENABLE, Pos: Position{Start: 7, End: 16, Line: 1, Column: 8}},
				Token{Kind: ASS, Pos: Position{Start: 18, End: 18, Line: 1, Column: 19}},
				Token{Kind: BOOL, Pos: Position{Start: 20, End: 23, Line: 1, Column: 21}},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			Stream{
				Token{Kind: TYPE, Pos: Position{Start: 0, End: 3, Line: 1, Column: 1}},
				Token{Kind: IDENT, Pos: Position{Start: 5, End: 9, Line: 1, Column: 6}},
				Token{Kind: LPAREN, Pos: Position{Start: 10, End: 10, Line: 1, Column: 11}},
				Token{Kind: IDENT, Pos: Position{Start: 11, End: 11, Line: 1, Column: 12}},
				Token{Kind: ASS, Pos: Position{Start: 13, End: 13, Line: 1, Column: 14}},
				Token{Kind: INT, Pos: Position{Start: 15, End: 16, Line: 1, Column: 16}},
				Token{Kind: RPAREN, Pos: Position{Start: 17, End: 17, Line: 1, Column: 18}},
				Token{Kind: CONFIG, Pos: Position{Start: 19, End: 24, Line: 1, Column: 20}},
				Token{Kind: SEMICOLON, Pos: Position{Start: 25, End: 25, Line: 1, Column: 26}},
				Token{Kind: WIDTH, Pos: Position{Start: 27, End: 31, Line: 1, Column: 28}},
				Token{Kind: ASS, Pos: Position{Start: 33, End: 33, Line: 1, Column: 34}},
				Token{Kind: IDENT, Pos: Position{Start: 35, End: 35, Line: 1, Column: 36}},
			},
		},
		{
			8,
			"s static; groups = [\"a\", \"b\"]",
			Stream{
				Token{Kind: IDENT, Pos: Position{Start: 0, End: 0, Line: 1, Column: 1}},
				Token{Kind: STATIC, Pos: Position{Start: 2, End: 7, Line: 1, Column: 3}},
				Token{Kind: SEMICOLON, Pos: Position{Start: 8, End: 8, Line: 1, Column: 9}},
				Token{Kind: GROUPS, Pos: Position{Start: 10, End: 15, Line: 1, Column: 11}},
				Token{Kind: ASS, Pos: Position{Start: 17, End: 17, Line: 1, Column: 18}},
				Token{Kind: LBRACK, Pos: Position{Start: 19, End: 19, Line: 1, Column: 20}},
				Token{Kind: STRING, Pos: Position{Start: 20, End: 22, Line: 1, Column: 21}},
				Token{Kind: COMMA, Pos: Position{Start: 23, End: 23, Line: 1, Column: 24}},
				Token{Kind: STRING, Pos: Position{Start: 25, End: 27, Line: 1, Column: 26}},
				Token{Kind: RBRACK, Pos: Position{Start: 28, End: 28, Line: 1, Column: 29}},
			},
		},
		{
			9,
			"import foo \"path\"",
			Stream{
				Token{Kind: IMPORT, Pos: Position{Start: 0, End: 5, Line: 1, Column: 1}},
				Token{Kind: IDENT, Pos: Position{Start: 7, End: 9, Line: 1, Column: 8}},
				Token{Kind: STRING, Pos: Position{Start: 11, End: 16, Line: 1, Column: 12}},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			Stream{
				Token{Kind: CONST, Pos: Position{Start: 0, End: 4, Line: 1, Column: 1}},
				Token{Kind: IDENT, Pos: Position{Start: 6, End: 6, Line: 1, Column: 7}},
				Token{Kind: ASS, Pos: Position{Start: 8, End: 8, Line: 1, Column: 9}},
				Token{Kind: INT, Pos: Position{Start: 10, End: 10, Line: 1, Column: 11}},
				Token{Kind: EXP, Pos: Position{Start: 11, End: 12, Line: 1, Column: 12}},
				Token{Kind: INT, Pos: Position{Start: 13, End: 13, Line: 1, Column: 14}},
				Token{Kind: SUB, Pos: Position{Start: 15, End: 15, Line: 1, Column: 16}},
				Token{Kind: INT, Pos: Position{Start: 17, End: 17, Line: 1, Column: 18}},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			Stream{
				Token{Kind: CONST, Pos: Position{Start: 0, End: 4, Line: 1, Column: 1}},
				Token{Kind: IDENT, Pos: Position{Start: 6, End: 7, Line: 1, Column: 7}},
				Token{Kind: ASS, Pos: Position{Start: 9, End: 9, Line: 1, Column: 10}},
				Token{Kind: INT, Pos: Position{Start: 11, End: 13, Line: 1, Column: 12}},
				Token{Kind: SHL, Pos: Position{Start: 15, End: 16, Line: 1, Column: 16}},
				Token{Kind: INT, Pos: Position{Start: 18, End: 20, Line: 1, Column: 19}},
			},
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expecting %d", test.idx, i)
		}

		got, err := Parse([]byte(test.src))
		if err != nil {
			t.Fatalf("Test %d: error != nil: %v", i, err)
		}

		if len(got) != len(test.want) {
			t.Fatalf(
				"\nTest: %d\n\nCode:\n%s\n\nInvalid number of tokens: got %d, want %d",
				i, test.src, len(got), len(test.want),
			)
		}

		for j, tok := range test.want {
			if got[j] != tok {
				t.Fatalf(
					"\nTest: %d\n\nCode:\n%s\n\nToken: %d\n got: %+v\nwant: %+v",
					i, test.src, j, got[j], tok,
				)
			}
		}
	}
}

func TestParseError(t *testing.T) {
	var tests = []struct {
		idx int // Test index, useful for navigation
		src string
		err error
	}{
		{
			0,
			"\n ", fmt.Errorf("2:1: space character ' ' not allowed for indent"),
		},
		{
			1,
			";\n", fmt.Errorf("1:1: extra ';' at the end of line"),
		},
		{
			2,
			" ; \n", fmt.Errorf("1:2: extra ';' at the end of line"),
		},
		{
			3,
			";;", fmt.Errorf("1:2: redundant ';'"),
		},
		{
			4,
			"b\"01-uUwWxXzZ3\"", fmt.Errorf("1:14: invalid character '3' in binary bit string literal"),
		},
		{
			5,
			"B\"0", fmt.Errorf("1:1: missing terminating '\"' in binary bit string literal"),
		},
		{
			6,
			"o\"01234567-uUwWxXzZ8\"", fmt.Errorf("1:20: invalid character '8' in octal bit string literal"),
		},
		{
			7,
			"O\"0", fmt.Errorf("1:1: missing terminating '\"' in octal bit string literal"),
		},
		{
			8,
			"x\"0123456789aAbBcCdDeEfF-uUwWxXzZ8g\"", fmt.Errorf("1:35: invalid character 'g' in hex bit string literal"),
		},
		{
			9,
			"X\"0", fmt.Errorf("1:1: missing terminating '\"' in hex bit string literal"),
		},
		{
			10,
			",,", fmt.Errorf("1:2: redundant ','"),
		},
		{
			11,
			"1.2.3", fmt.Errorf("1:4: second point character '.' in number literal"),
		},
		{
			12,
			"1e2.", fmt.Errorf("1:4: point character '.' after exponent in number literal"),
		},
		{
			13,
			"1e2d", fmt.Errorf("1:4: invalid character 'd' in number literal"),
		},
		{
			14,
			"\n\"str", fmt.Errorf("2:1: unterminated string literal"),
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expecting %d", test.idx, i)
		}

		_, err := Parse([]byte(test.src))
		if err == nil {
			t.Fatalf("%d: err = nil, expecting != nil", i)
		}

		if err.Error() != test.err.Error() {
			t.Fatalf("%d:\n got: %v\nwant: %v", i, err, test.err)
		}
	}
}
