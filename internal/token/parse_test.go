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
				Token{Kind: NEWLINE, Start: 0, End: 0, Line: 1, Column: 1},
				Token{Kind: NEWLINE, Start: 1, End: 1, Line: 2, Column: 1},
				Token{Kind: EOF, Start: 2, End: 2, Line: 3, Column: 1},
			},
		},
		{
			1,
			"# Comment line",
			Stream{
				Token{Kind: COMMENT, Start: 0, End: 13, Line: 1, Column: 1},
				Token{Kind: EOF, Start: 14, End: 14, Line: 1, Column: 15},
			},
		},
		{
			2,
			"# Comment\n",
			Stream{
				Token{Kind: COMMENT, Start: 0, End: 8, Line: 1, Column: 1},
				Token{Kind: NEWLINE, Start: 9, End: 9, Line: 1, Column: 10},
				Token{Kind: EOF, Start: 10, End: 10, Line: 2, Column: 1},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			Stream{
				Token{Kind: COMMENT, Start: 0, End: 10, Line: 1, Column: 1},
				Token{Kind: NEWLINE, Start: 11, End: 11, Line: 1, Column: 12},
				Token{Kind: COMMENT, Start: 12, End: 22, Line: 2, Column: 1},
				Token{Kind: EOF, Start: 23, End: 23, Line: 2, Column: 12},
			},
		},
		{
			4,
			"const A = true",
			Stream{
				Token{Kind: CONST, Start: 0, End: 4, Line: 1, Column: 1},
				Token{Kind: IDENT, Start: 6, End: 6, Line: 1, Column: 7},
				Token{Kind: ASS, Start: 8, End: 8, Line: 1, Column: 9},
				Token{Kind: BOOL, Start: 10, End: 13, Line: 1, Column: 11},
				Token{Kind: EOF, Start: 14, End: 14, Line: 1, Column: 15},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			Stream{
				Token{Kind: IDENT, Start: 0, End: 2, Line: 1, Column: 1},
				Token{Kind: MASK, Start: 4, End: 7, Line: 1, Column: 5},
				Token{Kind: SEMICOLON, Start: 8, End: 8, Line: 1, Column: 9},
				Token{Kind: ATOMIC, Start: 10, End: 15, Line: 1, Column: 11},
				Token{Kind: ASS, Start: 17, End: 17, Line: 1, Column: 18},
				Token{Kind: BOOL, Start: 19, End: 23, Line: 1, Column: 20},
				Token{Kind: EOF, Start: 24, End: 24, Line: 1, Column: 25},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			Stream{
				Token{Kind: IDENT, Start: 0, End: 0, Line: 1, Column: 1},
				Token{Kind: IRQ, Start: 2, End: 4, Line: 1, Column: 3},
				Token{Kind: SEMICOLON, Start: 5, End: 5, Line: 1, Column: 6},
				Token{Kind: ADD_ENABLE, Start: 7, End: 16, Line: 1, Column: 8},
				Token{Kind: ASS, Start: 18, End: 18, Line: 1, Column: 19},
				Token{Kind: BOOL, Start: 20, End: 23, Line: 1, Column: 21},
				Token{Kind: EOF, Start: 24, End: 24, Line: 1, Column: 25},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			Stream{
				Token{Kind: TYPE, Start: 0, End: 3, Line: 1, Column: 1},
				Token{Kind: IDENT, Start: 5, End: 9, Line: 1, Column: 6},
				Token{Kind: LPAREN, Start: 10, End: 10, Line: 1, Column: 11},
				Token{Kind: IDENT, Start: 11, End: 11, Line: 1, Column: 12},
				Token{Kind: ASS, Start: 13, End: 13, Line: 1, Column: 14},
				Token{Kind: INT, Start: 15, End: 16, Line: 1, Column: 16},
				Token{Kind: RPAREN, Start: 17, End: 17, Line: 1, Column: 18},
				Token{Kind: CONFIG, Start: 19, End: 24, Line: 1, Column: 20},
				Token{Kind: SEMICOLON, Start: 25, End: 25, Line: 1, Column: 26},
				Token{Kind: WIDTH, Start: 27, End: 31, Line: 1, Column: 28},
				Token{Kind: ASS, Start: 33, End: 33, Line: 1, Column: 34},
				Token{Kind: IDENT, Start: 35, End: 35, Line: 1, Column: 36},
				Token{Kind: EOF, Start: 36, End: 36, Line: 1, Column: 37},
			},
		},
		{
			8,
			"s static; groups = [\"a\", \"b\"]",
			Stream{
				Token{Kind: IDENT, Start: 0, End: 0, Line: 1, Column: 1},
				Token{Kind: STATIC, Start: 2, End: 7, Line: 1, Column: 3},
				Token{Kind: SEMICOLON, Start: 8, End: 8, Line: 1, Column: 9},
				Token{Kind: GROUPS, Start: 10, End: 15, Line: 1, Column: 11},
				Token{Kind: ASS, Start: 17, End: 17, Line: 1, Column: 18},
				Token{Kind: LBRACK, Start: 19, End: 19, Line: 1, Column: 20},
				Token{Kind: STRING, Start: 20, End: 22, Line: 1, Column: 21},
				Token{Kind: COMMA, Start: 23, End: 23, Line: 1, Column: 24},
				Token{Kind: STRING, Start: 25, End: 27, Line: 1, Column: 26},
				Token{Kind: RBRACK, Start: 28, End: 28, Line: 1, Column: 29},
				Token{Kind: EOF, Start: 29, End: 29, Line: 1, Column: 30},
			},
		},
		{
			9,
			"import foo \"path\"",
			Stream{
				Token{Kind: IMPORT, Start: 0, End: 5, Line: 1, Column: 1},
				Token{Kind: IDENT, Start: 7, End: 9, Line: 1, Column: 8},
				Token{Kind: STRING, Start: 11, End: 16, Line: 1, Column: 12},
				Token{Kind: EOF, Start: 17, End: 17, Line: 1, Column: 18},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			Stream{
				Token{Kind: CONST, Start: 0, End: 4, Line: 1, Column: 1},
				Token{Kind: IDENT, Start: 6, End: 6, Line: 1, Column: 7},
				Token{Kind: ASS, Start: 8, End: 8, Line: 1, Column: 9},
				Token{Kind: INT, Start: 10, End: 10, Line: 1, Column: 11},
				Token{Kind: EXP, Start: 11, End: 12, Line: 1, Column: 12},
				Token{Kind: INT, Start: 13, End: 13, Line: 1, Column: 14},
				Token{Kind: SUB, Start: 15, End: 15, Line: 1, Column: 16},
				Token{Kind: INT, Start: 17, End: 17, Line: 1, Column: 18},
				Token{Kind: EOF, Start: 18, End: 18, Line: 1, Column: 19},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			Stream{
				Token{Kind: CONST, Start: 0, End: 4, Line: 1, Column: 1},
				Token{Kind: IDENT, Start: 6, End: 7, Line: 1, Column: 7},
				Token{Kind: ASS, Start: 9, End: 9, Line: 1, Column: 10},
				Token{Kind: INT, Start: 11, End: 13, Line: 1, Column: 12},
				Token{Kind: SHL, Start: 15, End: 16, Line: 1, Column: 16},
				Token{Kind: INT, Start: 18, End: 20, Line: 1, Column: 19},
				Token{Kind: EOF, Start: 21, End: 21, Line: 1, Column: 22},
			},
		},
		{
			12,
			"p proc; delay=10 ns",
			Stream{
				Token{Kind: IDENT, Start: 0, End: 0, Line: 1, Column: 1},
				Token{Kind: PROC, Start: 2, End: 5, Line: 1, Column: 3},
				Token{Kind: SEMICOLON, Start: 6, End: 6, Line: 1, Column: 7},
				Token{Kind: DELAY, Start: 8, End: 12, Line: 1, Column: 9},
				Token{Kind: ASS, Start: 13, End: 13, Line: 1, Column: 14},
				Token{Kind: TIME, Start: 14, End: 18, Line: 1, Column: 15},
				Token{Kind: EOF, Start: 19, End: 19, Line: 1, Column: 20},
			},
		},
		{
			13,
			"b [a&&true]block",
			Stream{
				Token{Kind: IDENT, Start: 0, End: 0, Line: 1, Column: 1},
				Token{Kind: LBRACK, Start: 2, End: 2, Line: 1, Column: 3},
				Token{Kind: IDENT, Start: 3, End: 3, Line: 1, Column: 4},
				Token{Kind: LAND, Start: 4, End: 5, Line: 1, Column: 5},
				Token{Kind: BOOL, Start: 6, End: 9, Line: 1, Column: 7},
				Token{Kind: RBRACK, Start: 10, End: 10, Line: 1, Column: 11},
				Token{Kind: BLOCK, Start: 11, End: 15, Line: 1, Column: 12},
				Token{Kind: EOF, Start: 16, End: 16, Line: 1, Column: 17},
			},
		},
		{
			14,
			"const C_1 = 0xaf| 0x11",
			Stream{
				Token{Kind: CONST, Start: 0, End: 4, Line: 1, Column: 1},
				Token{Kind: IDENT, Start: 6, End: 8, Line: 1, Column: 7},
				Token{Kind: ASS, Start: 10, End: 10, Line: 1, Column: 11},
				Token{Kind: INT, Start: 12, End: 15, Line: 1, Column: 13},
				Token{Kind: OR, Start: 16, End: 16, Line: 1, Column: 17},
				Token{Kind: INT, Start: 18, End: 21, Line: 1, Column: 19},
				Token{Kind: EOF, Start: 22, End: 22, Line: 1, Column: 23},
			},
		},
		{
			15,
			"Main bus\n\ti irq\n\t\tadd-enable = true",
			Stream{
				Token{Kind: IDENT, Start: 0, End: 3, Line: 1, Column: 1},
				Token{Kind: BUS, Start: 5, End: 7, Line: 1, Column: 6},
				Token{Kind: NEWLINE, Start: 8, End: 8, Line: 1, Column: 9},
				Token{Kind: INDENT_INC, Start: 9, End: 9, Line: 2, Column: 1},
				Token{Kind: IDENT, Start: 10, End: 10, Line: 2, Column: 2},
				Token{Kind: IRQ, Start: 12, End: 14, Line: 2, Column: 4},
				Token{Kind: NEWLINE, Start: 15, End: 15, Line: 2, Column: 7},
				Token{Kind: INDENT_INC, Start: 16, End: 17, Line: 3, Column: 1},
				Token{Kind: ADD_ENABLE, Start: 18, End: 27, Line: 3, Column: 3},
				Token{Kind: ASS, Start: 29, End: 29, Line: 3, Column: 14},
				Token{Kind: BOOL, Start: 31, End: 34, Line: 3, Column: 16},
				Token{Kind: EOF, Start: 35, End: 35, Line: 3, Column: 20},
			},
		},
		{
			16,
			"type t static\n\twidth=7\n\nMain bus",
			Stream{
				Token{Kind: TYPE, Start: 0, End: 3, Line: 1, Column: 1},
				Token{Kind: IDENT, Start: 5, End: 5, Line: 1, Column: 6},
				Token{Kind: STATIC, Start: 7, End: 12, Line: 1, Column: 8},
				Token{Kind: NEWLINE, Start: 13, End: 13, Line: 1, Column: 14},
				Token{Kind: INDENT_INC, Start: 14, End: 14, Line: 2, Column: 1},
				Token{Kind: WIDTH, Start: 15, End: 19, Line: 2, Column: 2},
				Token{Kind: ASS, Start: 20, End: 20, Line: 2, Column: 7},
				Token{Kind: INT, Start: 21, End: 21, Line: 2, Column: 8},
				Token{Kind: NEWLINE, Start: 22, End: 22, Line: 2, Column: 9},
				Token{Kind: NEWLINE, Start: 23, End: 23, Line: 3, Column: 1},
				Token{Kind: IDENT, Start: 24, End: 27, Line: 4, Column: 1},
				Token{Kind: BUS, Start: 29, End: 31, Line: 4, Column: 6},
				Token{Kind: EOF, Start: 32, End: 32, Line: 4, Column: 9},
			},
		},
		{
			17,
			"Main bus\n\t# Comment\n\tc config\n\t\twidth = 6\n\t# Comment 2\n\ts stream",
			Stream{
				Token{Kind: IDENT, Start: 0, End: 3, Line: 1, Column: 1},
				Token{Kind: BUS, Start: 5, End: 7, Line: 1, Column: 6},
				Token{Kind: NEWLINE, Start: 8, End: 8, Line: 1, Column: 9},
				Token{Kind: INDENT_INC, Start: 9, End: 9, Line: 2, Column: 1},
				Token{Kind: COMMENT, Start: 10, End: 18, Line: 2, Column: 2},
				Token{Kind: NEWLINE, Start: 19, End: 19, Line: 2, Column: 11},
				Token{Kind: IDENT, Start: 21, End: 21, Line: 3, Column: 2},
				Token{Kind: CONFIG, Start: 23, End: 28, Line: 3, Column: 4},
				Token{Kind: NEWLINE, Start: 29, End: 29, Line: 3, Column: 10},
				Token{Kind: INDENT_INC, Start: 30, End: 31, Line: 4, Column: 1},
				Token{Kind: WIDTH, Start: 32, End: 36, Line: 4, Column: 3},
				Token{Kind: ASS, Start: 38, End: 38, Line: 4, Column: 9},
				Token{Kind: INT, Start: 40, End: 40, Line: 4, Column: 11},
				Token{Kind: NEWLINE, Start: 41, End: 41, Line: 4, Column: 12},
				Token{Kind: COMMENT, Start: 43, End: 53, Line: 5, Column: 2},
				Token{Kind: NEWLINE, Start: 54, End: 54, Line: 5, Column: 13},
				Token{Kind: IDENT, Start: 56, End: 56, Line: 6, Column: 2},
				Token{Kind: STREAM, Start: 58, End: 63, Line: 6, Column: 4},
				Token{Kind: EOF, Start: 64, End: 64, Line: 6, Column: 10},
			},
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		got, err := Parse([]byte(test.src))
		if err != nil {
			t.Fatalf("Test %d: err != nil: %v", i, err)
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
		{
			15,
			"\t", fmt.Errorf("1:1: tab character '\t' not allowed for alignment"),
		},
		{
			16,
			"Main bus\n\tc config;\twidth = 7", fmt.Errorf("2:11: tab character '\t' not allowed for alignment"),
		},
		{
			17,
			"Main bus\n\t c config", fmt.Errorf("2:2: space character ' ' right after tab character '\t'"),
		},
		{
			18,
			"Main bus\n\t\tc config", fmt.Errorf("2:1: multi indent increase"),
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		_, err := Parse([]byte(test.src))
		if err == nil {
			t.Fatalf("%d: err == nil, expected != nil", i)
		}

		if err.Error() != test.err.Error() {
			t.Fatalf("%d:\n got: %v\nwant: %v", i, err, test.err)
		}
	}
}
