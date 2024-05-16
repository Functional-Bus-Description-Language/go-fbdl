package tok

import (
	"reflect"
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
				Newline{position{start: 0, end: 1, line: 1, column: 1}},
				Eof{position{start: 2, end: 2, line: 3, column: 1}},
			},
		},
		{
			1,
			"# Comment line",
			[]Token{
				Comment{position{start: 0, end: 13, line: 1, column: 1}},
				Eof{position{start: 14, end: 14, line: 1, column: 15}},
			},
		},
		{
			2,
			"# Comment\n",
			[]Token{
				Comment{position{start: 0, end: 8, line: 1, column: 1}},
				Newline{position{start: 9, end: 9, line: 1, column: 10}},
				Eof{position{start: 10, end: 10, line: 2, column: 1}},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			[]Token{
				Comment{position{start: 0, end: 10, line: 1, column: 1}},
				Newline{position{start: 11, end: 11, line: 1, column: 12}},
				Comment{position{start: 12, end: 22, line: 2, column: 1}},
				Eof{position{start: 23, end: 23, line: 2, column: 12}},
			},
		},
		{
			4,
			"const A = true",
			[]Token{
				Const{position{start: 0, end: 4, line: 1, column: 1}},
				Ident{position{start: 6, end: 6, line: 1, column: 7}},
				Ass{position{start: 8, end: 8, line: 1, column: 9}},
				Bool{position{start: 10, end: 13, line: 1, column: 11}},
				Eof{position{start: 14, end: 14, line: 1, column: 15}},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			[]Token{
				Ident{position{start: 0, end: 2, line: 1, column: 1}},
				Mask{position{start: 4, end: 7, line: 1, column: 5}},
				Semicolon{position{start: 8, end: 8, line: 1, column: 9}},
				Atomic{position{start: 10, end: 15, line: 1, column: 11}},
				Ass{position{start: 17, end: 17, line: 1, column: 18}},
				Bool{position{start: 19, end: 23, line: 1, column: 20}},
				Eof{position{start: 24, end: 24, line: 1, column: 25}},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			[]Token{
				Ident{position{start: 0, end: 0, line: 1, column: 1}},
				Irq{position{start: 2, end: 4, line: 1, column: 3}},
				Semicolon{position{start: 5, end: 5, line: 1, column: 6}},
				AddEnable{position{start: 7, end: 16, line: 1, column: 8}},
				Ass{position{start: 18, end: 18, line: 1, column: 19}},
				Bool{position{start: 20, end: 23, line: 1, column: 21}},
				Eof{position{start: 24, end: 24, line: 1, column: 25}},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			[]Token{
				Type{position{start: 0, end: 3, line: 1, column: 1}},
				Ident{position{start: 5, end: 9, line: 1, column: 6}},
				LParen{position{start: 10, end: 10, line: 1, column: 11}},
				Ident{position{start: 11, end: 11, line: 1, column: 12}},
				Ass{position{start: 13, end: 13, line: 1, column: 14}},
				Int{position{start: 15, end: 16, line: 1, column: 16}},
				RParen{position{start: 17, end: 17, line: 1, column: 18}},
				Config{position{start: 19, end: 24, line: 1, column: 20}},
				Semicolon{position{start: 25, end: 25, line: 1, column: 26}},
				Width{position{start: 27, end: 31, line: 1, column: 28}},
				Ass{position{start: 33, end: 33, line: 1, column: 34}},
				Ident{position{start: 35, end: 35, line: 1, column: 36}},
				Eof{position{start: 36, end: 36, line: 1, column: 37}},
			},
		},
		{
			8,
			"s static; groups = [\"a\", \"b\"]",
			[]Token{
				Ident{position{start: 0, end: 0, line: 1, column: 1}},
				Static{position{start: 2, end: 7, line: 1, column: 3}},
				Semicolon{position{start: 8, end: 8, line: 1, column: 9}},
				Groups{position{start: 10, end: 15, line: 1, column: 11}},
				Ass{position{start: 17, end: 17, line: 1, column: 18}},
				LBracket{position{start: 19, end: 19, line: 1, column: 20}},
				String{position{start: 20, end: 22, line: 1, column: 21}},
				Comma{position{start: 23, end: 23, line: 1, column: 24}},
				String{position{start: 25, end: 27, line: 1, column: 26}},
				RBracket{position{start: 28, end: 28, line: 1, column: 29}},
				Eof{position{start: 29, end: 29, line: 1, column: 30}},
			},
		},
		{
			9,
			"import foo \"path\"",
			[]Token{
				Import{position{start: 0, end: 5, line: 1, column: 1}},
				Ident{position{start: 7, end: 9, line: 1, column: 8}},
				String{position{start: 11, end: 16, line: 1, column: 12}},
				Eof{position{start: 17, end: 17, line: 1, column: 18}},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			[]Token{
				Const{position{start: 0, end: 4, line: 1, column: 1}},
				Ident{position{start: 6, end: 6, line: 1, column: 7}},
				Ass{position{start: 8, end: 8, line: 1, column: 9}},
				Int{position{start: 10, end: 10, line: 1, column: 11}},
				Exp{position{start: 11, end: 12, line: 1, column: 12}},
				Int{position{start: 13, end: 13, line: 1, column: 14}},
				Sub{position{start: 15, end: 15, line: 1, column: 16}},
				Int{position{start: 17, end: 17, line: 1, column: 18}},
				Eof{position{start: 18, end: 18, line: 1, column: 19}},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			[]Token{
				Const{position{start: 0, end: 4, line: 1, column: 1}},
				Ident{position{start: 6, end: 7, line: 1, column: 7}},
				Ass{position{start: 9, end: 9, line: 1, column: 10}},
				Int{position{start: 11, end: 13, line: 1, column: 12}},
				LShift{position{start: 15, end: 16, line: 1, column: 16}},
				Int{position{start: 18, end: 20, line: 1, column: 19}},
				Eof{position{start: 21, end: 21, line: 1, column: 22}},
			},
		},
		{
			12,
			"p proc; delay=10 ns",
			[]Token{
				Ident{position{start: 0, end: 0, line: 1, column: 1}},
				Proc{position{start: 2, end: 5, line: 1, column: 3}},
				Semicolon{position{start: 6, end: 6, line: 1, column: 7}},
				Delay{position{start: 8, end: 12, line: 1, column: 9}},
				Ass{position{start: 13, end: 13, line: 1, column: 14}},
				Time{position{start: 14, end: 18, line: 1, column: 15}},
				Eof{position{start: 19, end: 19, line: 1, column: 20}},
			},
		},
		{
			13,
			"b [a&&true]block",
			[]Token{
				Ident{position{start: 0, end: 0, line: 1, column: 1}},
				LBracket{position{start: 2, end: 2, line: 1, column: 3}},
				Ident{position{start: 3, end: 3, line: 1, column: 4}},
				And{position{start: 4, end: 5, line: 1, column: 5}},
				Bool{position{start: 6, end: 9, line: 1, column: 7}},
				RBracket{position{start: 10, end: 10, line: 1, column: 11}},
				Block{position{start: 11, end: 15, line: 1, column: 12}},
				Eof{position{start: 16, end: 16, line: 1, column: 17}},
			},
		},
		{
			14,
			"const C_1 = 0xaf| 0x11",
			[]Token{
				Const{position{start: 0, end: 4, line: 1, column: 1}},
				Ident{position{start: 6, end: 8, line: 1, column: 7}},
				Ass{position{start: 10, end: 10, line: 1, column: 11}},
				Int{position{start: 12, end: 15, line: 1, column: 13}},
				BitOr{position{start: 16, end: 16, line: 1, column: 17}},
				Int{position{start: 18, end: 21, line: 1, column: 19}},
				Eof{position{start: 22, end: 22, line: 1, column: 23}},
			},
		},
		{
			15,
			"Main bus\n\ti irq\n\t\tadd-enable = true",
			[]Token{
				Ident{position{start: 0, end: 3, line: 1, column: 1}},
				Bus{position{start: 5, end: 7, line: 1, column: 6}},
				Newline{position{start: 8, end: 8, line: 1, column: 9}},
				Indent{position{start: 9, end: 9, line: 2, column: 1}},
				Ident{position{start: 10, end: 10, line: 2, column: 2}},
				Irq{position{start: 12, end: 14, line: 2, column: 4}},
				Newline{position{start: 15, end: 15, line: 2, column: 7}},
				Indent{position{start: 16, end: 17, line: 3, column: 1}},
				AddEnable{position{start: 18, end: 27, line: 3, column: 3}},
				Ass{position{start: 29, end: 29, line: 3, column: 14}},
				Bool{position{start: 31, end: 34, line: 3, column: 16}},
				Eof{position{start: 35, end: 35, line: 3, column: 20}},
			},
		},
		{
			16,
			"type t static\n\twidth=7\n\nMain bus",
			[]Token{
				Type{position{start: 0, end: 3, line: 1, column: 1}},
				Ident{position{start: 5, end: 5, line: 1, column: 6}},
				Static{position{start: 7, end: 12, line: 1, column: 8}},
				Newline{position{start: 13, end: 13, line: 1, column: 14}},
				Indent{position{start: 14, end: 14, line: 2, column: 1}},
				Width{position{start: 15, end: 19, line: 2, column: 2}},
				Ass{position{start: 20, end: 20, line: 2, column: 7}},
				Int{position{start: 21, end: 21, line: 2, column: 8}},
				Newline{position{start: 22, end: 23, line: 2, column: 9}},
				Dedent{position{start: 24, end: 24, line: 4, column: 1}},
				Ident{position{start: 24, end: 27, line: 4, column: 1}},
				Bus{position{start: 29, end: 31, line: 4, column: 6}},
				Eof{position{start: 32, end: 32, line: 4, column: 9}},
			},
		},
		{
			17,
			"Main bus\n\t# Comment\n\tc config\n\t\twidth = 6\n\t# Comment 2\n\ts stream",
			[]Token{
				Ident{position{start: 0, end: 3, line: 1, column: 1}},
				Bus{position{start: 5, end: 7, line: 1, column: 6}},
				Newline{position{start: 8, end: 8, line: 1, column: 9}},
				Indent{position{start: 9, end: 9, line: 2, column: 1}},
				Comment{position{start: 10, end: 18, line: 2, column: 2}},
				Newline{position{start: 19, end: 19, line: 2, column: 11}},
				Ident{position{start: 21, end: 21, line: 3, column: 2}},
				Config{position{start: 23, end: 28, line: 3, column: 4}},
				Newline{position{start: 29, end: 29, line: 3, column: 10}},
				Indent{position{start: 30, end: 31, line: 4, column: 1}},
				Width{position{start: 32, end: 36, line: 4, column: 3}},
				Ass{position{start: 38, end: 38, line: 4, column: 9}},
				Int{position{start: 40, end: 40, line: 4, column: 11}},
				Newline{position{start: 41, end: 41, line: 4, column: 12}},
				Dedent{position{start: 42, end: 42, line: 5, column: 1}},
				Comment{position{start: 43, end: 53, line: 5, column: 2}},
				Newline{position{start: 54, end: 54, line: 5, column: 13}},
				Ident{position{start: 56, end: 56, line: 6, column: 2}},
				Stream{position{start: 58, end: 63, line: 6, column: 4}},
				Eof{position{start: 64, end: 64, line: 6, column: 10}},
			},
		},
		{
			18,
			"masters = -0",
			[]Token{
				Masters{position{start: 0, end: 6, line: 1, column: 1}},
				Ass{position{start: 8, end: 8, line: 1, column: 9}},
				Sub{position{start: 10, end: 10, line: 1, column: 11}},
				Int{position{start: 11, end: 11, line: 1, column: 12}},
				Eof{position{start: 12, end: 12, line: 1, column: 13}},
			},
		},
		{
			19,
			"size = a-b",
			[]Token{
				Size{position{start: 0, end: 3, line: 1, column: 1}},
				Ass{position{start: 5, end: 5, line: 1, column: 6}},
				Ident{position{start: 7, end: 7, line: 1, column: 8}},
				Sub{position{start: 8, end: 8, line: 1, column: 9}},
				Ident{position{start: 9, end: 9, line: 1, column: 10}},
				Eof{position{start: 10, end: 10, line: 1, column: 11}},
			},
		},
		{
			20,
			"size = init-value",
			[]Token{
				Size{position{start: 0, end: 3, line: 1, column: 1}},
				Ass{position{start: 5, end: 5, line: 1, column: 6}},
				Ident{position{start: 7, end: 10, line: 1, column: 8}},
				Sub{position{start: 11, end: 11, line: 1, column: 12}},
				Ident{position{start: 12, end: 16, line: 1, column: 13}},
				Eof{position{start: 17, end: 17, line: 1, column: 18}},
			},
		},
		{
			21,
			`const
	A = 1
	B = 2 # Inline comment
	# Doc comment
	C = 3.14`,
			[]Token{
				Const{position{start: 0, end: 4, line: 1, column: 1}},
				Newline{position{start: 5, end: 5, line: 1, column: 6}},
				Indent{position{start: 6, end: 6, line: 2, column: 1}},
				Ident{position{start: 7, end: 7, line: 2, column: 2}},
				Ass{position{start: 9, end: 9, line: 2, column: 4}},
				Int{position{start: 11, end: 11, line: 2, column: 6}},
				Newline{position{start: 12, end: 12, line: 2, column: 7}},
				Ident{position{start: 14, end: 14, line: 3, column: 2}},
				Ass{position{start: 16, end: 16, line: 3, column: 4}},
				Int{position{start: 18, end: 18, line: 3, column: 6}},
				Newline{position{start: 36, end: 36, line: 3, column: 24}},
				Comment{position{start: 38, end: 50, line: 4, column: 2}},
				Newline{position{start: 51, end: 51, line: 4, column: 15}},
				Ident{position{start: 53, end: 53, line: 5, column: 2}},
				Ass{position{start: 55, end: 55, line: 5, column: 4}},
				Float{position{start: 57, end: 60, line: 5, column: 6}},
				Eof{position{start: 61, end: 61, line: 5, column: 10}},
			},
		},
		{
			22,
			"abc.Def",
			[]Token{
				QualIdent{position{start: 0, end: 6, line: 1, column: 1}},
				Eof{position{start: 7, end: 7, line: 1, column: 8}},
			},
		},
		{
			23,
			"a-b.C-d.E",
			[]Token{
				Ident{position{start: 0, end: 0, line: 1, column: 1}},
				Sub{position{start: 1, end: 1, line: 1, column: 2}},
				QualIdent{position{start: 2, end: 4, line: 1, column: 3}},
				Sub{position{start: 5, end: 5, line: 1, column: 6}},
				QualIdent{position{start: 6, end: 8, line: 1, column: 7}},
				Eof{position{start: 9, end: 9, line: 1, column: 10}},
			},
		},
		{
			24,
			"range = 1:9",
			[]Token{
				Range{position{start: 0, end: 4, line: 1, column: 1}},
				Ass{position{start: 6, end: 6, line: 1, column: 7}},
				Int{position{start: 8, end: 8, line: 1, column: 9}},
				Colon{position{start: 9, end: 9, line: 1, column: 10}},
				Int{position{start: 10, end: 10, line: 1, column: 11}},
				Eof{position{start: 11, end: 11, line: 1, column: 12}},
			},
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		got, err := Parse([]byte(test.src), "")
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
			if reflect.TypeOf(got[j]) != reflect.TypeOf(tok) ||
				got[j].Start() != tok.Start() ||
				got[j].End() != tok.End() ||
				got[j].Line() != tok.Line() ||
				got[j].Column() != tok.Column() {
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
		err string
	}{
		{
			0,
			"\n ",
			"space character ' ' not allowed for indent",
		},
		{
			1,
			";\n",
			"extra ';' at line end",
		},
		{
			2,
			" ; \n",
			"extra ';' at line end",
		},
		{
			3,
			";;",
			"redundant ';'",
		},
		{
			4,
			"b\"01-uUwWxXzZ3\"",
			"invalid character '3' in binary bit string",
		},
		{
			5,
			"B\"0",
			"unterminated binary bit string, probably missing '\"'",
		},
		{
			6,
			"o\"01234567-uUwWxXzZ8\"",
			"invalid character '8' in octal bit string",
		},
		{
			7,
			"O\"0",
			"unterminated octal bit string, probably missing '\"'",
		},
		{
			8,
			"x\"0123456789aAbBcCdDeEfF-uUwWxXzZ8g\"",
			"invalid character 'g' in hex bit string",
		},
		{
			9,
			"X\"0",
			"unterminated hex bit string, probably missing '\"'",
		},
		{
			10,
			",,",
			"redundant ','",
		},
		{
			11,
			"1.2.3",
			"second point character '.' in number",
		},
		{
			12,
			"1e2.",
			"point character '.' after exponent in number",
		},
		{
			13,
			"1e2d",
			"invalid character 'd' in number",
		},
		{
			14,
			"\n\"str",
			"unterminated string, probably missing '\"'",
		},
		{
			15,
			"\t",
			"tab character '\\t' not allowed for alignment",
		},
		{
			16,
			"Main bus\n\tc config;\twidth = 7",
			"tab character '\\t' not allowed for alignment",
		},
		{
			17,
			"Main bus\n\t c config",
			"space character ' ' right after tab character '\\t'",
		},
		{
			18,
			"Main bus\n\t\tc config",
			"multi indent increase",
		},
		{
			19,
			"pkg.sym",
			"symbol name in qualified identifier must start with upper case letter",
		},
		{
			20,
			"a-b.c",
			"symbol name in qualified identifier must start with upper case letter",
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		_, err := Parse([]byte(test.src), "")
		if err == nil {
			t.Fatalf("%d: err == nil, expected != nil", i)
		}

		tokErr := err.(Error)
		if tokErr.Msg != test.err {
			t.Fatalf("\nTest %d:\n\ngot:\n%v\n\nwant:\n%v", i, tokErr.Msg, test.err)
		}
	}
}
