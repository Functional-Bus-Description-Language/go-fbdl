package token

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		idx  int // Test index, useful for navigation
		src  string
		want []Token
	}{
		{
			0,
			"\n\n",
			[]Token{
				Newline{start: 0, end: 0, line: 1, column: 1},
				Newline{start: 1, end: 1, line: 2, column: 1},
				Eof{start: 2, end: 2, line: 3, column: 1},
			},
		},
		{
			1,
			"# Comment line",
			[]Token{
				Comment{start: 0, end: 13, line: 1, column: 1},
				Eof{start: 14, end: 14, line: 1, column: 15},
			},
		},
		{
			2,
			"# Comment\n",
			[]Token{
				Comment{start: 0, end: 8, line: 1, column: 1},
				Newline{start: 9, end: 9, line: 1, column: 10},
				Eof{start: 10, end: 10, line: 2, column: 1},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			[]Token{
				Comment{start: 0, end: 10, line: 1, column: 1},
				Newline{start: 11, end: 11, line: 1, column: 12},
				Comment{start: 12, end: 22, line: 2, column: 1},
				Eof{start: 23, end: 23, line: 2, column: 12},
			},
		},
		{
			4,
			"const A = true",
			[]Token{
				Const{start: 0, end: 4, line: 1, column: 1},
				Ident{start: 6, end: 6, line: 1, column: 7},
				Ass{start: 8, end: 8, line: 1, column: 9},
				Bool{start: 10, end: 13, line: 1, column: 11},
				Eof{start: 14, end: 14, line: 1, column: 15},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			[]Token{
				Ident{start: 0, end: 2, line: 1, column: 1},
				Mask{start: 4, end: 7, line: 1, column: 5},
				Semicolon{start: 8, end: 8, line: 1, column: 9},
				Atomic{start: 10, end: 15, line: 1, column: 11},
				Ass{start: 17, end: 17, line: 1, column: 18},
				Bool{start: 19, end: 23, line: 1, column: 20},
				Eof{start: 24, end: 24, line: 1, column: 25},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			[]Token{
				Ident{start: 0, end: 0, line: 1, column: 1},
				Irq{start: 2, end: 4, line: 1, column: 3},
				Semicolon{start: 5, end: 5, line: 1, column: 6},
				AddEnable{start: 7, end: 16, line: 1, column: 8},
				Ass{start: 18, end: 18, line: 1, column: 19},
				Bool{start: 20, end: 23, line: 1, column: 21},
				Eof{start: 24, end: 24, line: 1, column: 25},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			[]Token{
				Type{start: 0, end: 3, line: 1, column: 1},
				Ident{start: 5, end: 9, line: 1, column: 6},
				LeftParen{start: 10, end: 10, line: 1, column: 11},
				Ident{start: 11, end: 11, line: 1, column: 12},
				Ass{start: 13, end: 13, line: 1, column: 14},
				Int{start: 15, end: 16, line: 1, column: 16},
				RightParen{start: 17, end: 17, line: 1, column: 18},
				Config{start: 19, end: 24, line: 1, column: 20},
				Semicolon{start: 25, end: 25, line: 1, column: 26},
				Width{start: 27, end: 31, line: 1, column: 28},
				Ass{start: 33, end: 33, line: 1, column: 34},
				Ident{start: 35, end: 35, line: 1, column: 36},
				Eof{start: 36, end: 36, line: 1, column: 37},
			},
		},
		{
			8,
			"s static; groups = [\"a\", \"b\"]",
			[]Token{
				Ident{start: 0, end: 0, line: 1, column: 1},
				Static{start: 2, end: 7, line: 1, column: 3},
				Semicolon{start: 8, end: 8, line: 1, column: 9},
				Groups{start: 10, end: 15, line: 1, column: 11},
				Ass{start: 17, end: 17, line: 1, column: 18},
				LeftBracket{start: 19, end: 19, line: 1, column: 20},
				String{start: 20, end: 22, line: 1, column: 21},
				Comma{start: 23, end: 23, line: 1, column: 24},
				String{start: 25, end: 27, line: 1, column: 26},
				RightBracket{start: 28, end: 28, line: 1, column: 29},
				Eof{start: 29, end: 29, line: 1, column: 30},
			},
		},
		{
			9,
			"import foo \"path\"",
			[]Token{
				Import{start: 0, end: 5, line: 1, column: 1},
				Ident{start: 7, end: 9, line: 1, column: 8},
				String{start: 11, end: 16, line: 1, column: 12},
				Eof{start: 17, end: 17, line: 1, column: 18},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			[]Token{
				Const{start: 0, end: 4, line: 1, column: 1},
				Ident{start: 6, end: 6, line: 1, column: 7},
				Ass{start: 8, end: 8, line: 1, column: 9},
				Int{start: 10, end: 10, line: 1, column: 11},
				Exp{start: 11, end: 12, line: 1, column: 12},
				Int{start: 13, end: 13, line: 1, column: 14},
				Sub{start: 15, end: 15, line: 1, column: 16},
				Int{start: 17, end: 17, line: 1, column: 18},
				Eof{start: 18, end: 18, line: 1, column: 19},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			[]Token{
				Const{start: 0, end: 4, line: 1, column: 1},
				Ident{start: 6, end: 7, line: 1, column: 7},
				Ass{start: 9, end: 9, line: 1, column: 10},
				Int{start: 11, end: 13, line: 1, column: 12},
				LeftShift{start: 15, end: 16, line: 1, column: 16},
				Int{start: 18, end: 20, line: 1, column: 19},
				Eof{start: 21, end: 21, line: 1, column: 22},
			},
		},
		{
			12,
			"p proc; delay=10 ns",
			[]Token{
				Ident{start: 0, end: 0, line: 1, column: 1},
				Proc{start: 2, end: 5, line: 1, column: 3},
				Semicolon{start: 6, end: 6, line: 1, column: 7},
				Delay{start: 8, end: 12, line: 1, column: 9},
				Ass{start: 13, end: 13, line: 1, column: 14},
				Time{start: 14, end: 18, line: 1, column: 15},
				Eof{start: 19, end: 19, line: 1, column: 20},
			},
		},
		{
			13,
			"b [a&&true]block",
			[]Token{
				Ident{start: 0, end: 0, line: 1, column: 1},
				LeftBracket{start: 2, end: 2, line: 1, column: 3},
				Ident{start: 3, end: 3, line: 1, column: 4},
				And{start: 4, end: 5, line: 1, column: 5},
				Bool{start: 6, end: 9, line: 1, column: 7},
				RightBracket{start: 10, end: 10, line: 1, column: 11},
				Block{start: 11, end: 15, line: 1, column: 12},
				Eof{start: 16, end: 16, line: 1, column: 17},
			},
		},
		{
			14,
			"const C_1 = 0xaf| 0x11",
			[]Token{
				Const{start: 0, end: 4, line: 1, column: 1},
				Ident{start: 6, end: 8, line: 1, column: 7},
				Ass{start: 10, end: 10, line: 1, column: 11},
				Int{start: 12, end: 15, line: 1, column: 13},
				BitOr{start: 16, end: 16, line: 1, column: 17},
				Int{start: 18, end: 21, line: 1, column: 19},
				Eof{start: 22, end: 22, line: 1, column: 23},
			},
		},
		{
			15,
			"Main bus\n\ti irq\n\t\tadd-enable = true",
			[]Token{
				Ident{start: 0, end: 3, line: 1, column: 1},
				Bus{start: 5, end: 7, line: 1, column: 6},
				Newline{start: 8, end: 8, line: 1, column: 9},
				IndentInc{start: 9, end: 9, line: 2, column: 1},
				Ident{start: 10, end: 10, line: 2, column: 2},
				Irq{start: 12, end: 14, line: 2, column: 4},
				Newline{start: 15, end: 15, line: 2, column: 7},
				IndentInc{start: 16, end: 17, line: 3, column: 1},
				AddEnable{start: 18, end: 27, line: 3, column: 3},
				Ass{start: 29, end: 29, line: 3, column: 14},
				Bool{start: 31, end: 34, line: 3, column: 16},
				Eof{start: 35, end: 35, line: 3, column: 20},
			},
		},
		{
			16,
			"type t static\n\twidth=7\n\nMain bus",
			[]Token{
				Type{start: 0, end: 3, line: 1, column: 1},
				Ident{start: 5, end: 5, line: 1, column: 6},
				Static{start: 7, end: 12, line: 1, column: 8},
				Newline{start: 13, end: 13, line: 1, column: 14},
				IndentInc{start: 14, end: 14, line: 2, column: 1},
				Width{start: 15, end: 19, line: 2, column: 2},
				Ass{start: 20, end: 20, line: 2, column: 7},
				Int{start: 21, end: 21, line: 2, column: 8},
				Newline{start: 22, end: 22, line: 2, column: 9},
				Newline{start: 23, end: 23, line: 3, column: 1},
				IndentDec{start: 23, end: 23, line: 3, column: 1},
				Ident{start: 24, end: 27, line: 4, column: 1},
				Bus{start: 29, end: 31, line: 4, column: 6},
				Eof{start: 32, end: 32, line: 4, column: 9},
			},
		},
		{
			17,
			"Main bus\n\t# Comment\n\tc config\n\t\twidth = 6\n\t# Comment 2\n\ts stream",
			[]Token{
				Ident{start: 0, end: 3, line: 1, column: 1},
				Bus{start: 5, end: 7, line: 1, column: 6},
				Newline{start: 8, end: 8, line: 1, column: 9},
				IndentInc{start: 9, end: 9, line: 2, column: 1},
				Comment{start: 10, end: 18, line: 2, column: 2},
				Newline{start: 19, end: 19, line: 2, column: 11},
				Ident{start: 21, end: 21, line: 3, column: 2},
				Config{start: 23, end: 28, line: 3, column: 4},
				Newline{start: 29, end: 29, line: 3, column: 10},
				IndentInc{start: 30, end: 31, line: 4, column: 1},
				Width{start: 32, end: 36, line: 4, column: 3},
				Ass{start: 38, end: 38, line: 4, column: 9},
				Int{start: 40, end: 40, line: 4, column: 11},
				Newline{start: 41, end: 41, line: 4, column: 12},
				IndentDec{start: 42, end: 42, line: 5, column: 1},
				Comment{start: 43, end: 53, line: 5, column: 2},
				Newline{start: 54, end: 54, line: 5, column: 13},
				Ident{start: 56, end: 56, line: 6, column: 2},
				Stream{start: 58, end: 63, line: 6, column: 4},
				Eof{start: 64, end: 64, line: 6, column: 10},
			},
		},
		{
			18,
			"masters = -0",
			[]Token{
				Masters{start: 0, end: 6, line: 1, column: 1},
				Ass{start: 8, end: 8, line: 1, column: 9},
				Sub{start: 10, end: 10, line: 1, column: 11},
				Int{start: 11, end: 11, line: 1, column: 12},
				Eof{start: 12, end: 12, line: 1, column: 13},
			},
		},
		{
			19,
			"size = a-b",
			[]Token{
				Size{start: 0, end: 3, line: 1, column: 1},
				Ass{start: 5, end: 5, line: 1, column: 6},
				Ident{start: 7, end: 7, line: 1, column: 8},
				Sub{start: 8, end: 8, line: 1, column: 9},
				Ident{start: 9, end: 9, line: 1, column: 10},
				Eof{start: 10, end: 10, line: 1, column: 11},
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
