package token

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
				Newline{position{start: 0, end: 1}},
				Eof{position{start: 2, end: 2}},
			},
		},
		{
			1,
			"# Comment line",
			[]Token{
				Comment{position{start: 0, end: 13}},
				Eof{position{start: 14, end: 14}},
			},
		},
		{
			2,
			"# Comment\n",
			[]Token{
				Comment{position{start: 0, end: 8}},
				Newline{position{start: 9, end: 9}},
				Eof{position{start: 10, end: 10}},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			[]Token{
				Comment{position{start: 0, end: 10}},
				Newline{position{start: 11, end: 11}},
				Comment{position{start: 12, end: 22}},
				Eof{position{start: 23, end: 23}},
			},
		},
		{
			4,
			"const A = true",
			[]Token{
				Const{position{start: 0, end: 4}},
				Ident{position{start: 6, end: 6}},
				Ass{position{start: 8, end: 8}},
				Bool{position{start: 10, end: 13}},
				Eof{position{start: 14, end: 14}},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			[]Token{
				Ident{position{start: 0, end: 2}},
				Mask{position{start: 4, end: 7}},
				Semicolon{position{start: 8, end: 8}},
				Atomic{position{start: 10, end: 15}},
				Ass{position{start: 17, end: 17}},
				Bool{position{start: 19, end: 23}},
				Eof{position{start: 24, end: 24}},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			[]Token{
				Ident{position{start: 0, end: 0}},
				Irq{position{start: 2, end: 4}},
				Semicolon{position{start: 5, end: 5}},
				AddEnable{position{start: 7, end: 16}},
				Ass{position{start: 18, end: 18}},
				Bool{position{start: 20, end: 23}},
				Eof{position{start: 24, end: 24}},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			[]Token{
				Type{position{start: 0, end: 3}},
				Ident{position{start: 5, end: 9}},
				LParen{position{start: 10, end: 10}},
				Ident{position{start: 11, end: 11}},
				Ass{position{start: 13, end: 13}},
				Int{position{start: 15, end: 16}},
				RParen{position{start: 17, end: 17}},
				Config{position{start: 19, end: 24}},
				Semicolon{position{start: 25, end: 25}},
				Width{position{start: 27, end: 31}},
				Ass{position{start: 33, end: 33}},
				Ident{position{start: 35, end: 35}},
				Eof{position{start: 36, end: 36}},
			},
		},
		{
			8,
			"s static; init-value = x\"FFFFFFFF\"",
			[]Token{
				Ident{position{start: 0, end: 0}},
				Static{position{start: 2, end: 7}},
				Semicolon{position{start: 8, end: 8}},
				InitValue{position{start: 10, end: 19}},
				Ass{position{start: 21, end: 21}},
				BitString{position{start: 23, end: 33}},
				Eof{position{start: 34, end: 34}},
			},
		},
		{
			9,
			"import foo \"path\"",
			[]Token{
				Import{position{start: 0, end: 5}},
				Ident{position{start: 7, end: 9}},
				String{position{start: 11, end: 16}},
				Eof{position{start: 17, end: 17}},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			[]Token{
				Const{position{start: 0, end: 4}},
				Ident{position{start: 6, end: 6}},
				Ass{position{start: 8, end: 8}},
				Int{position{start: 10, end: 10}},
				Exp{position{start: 11, end: 12}},
				Int{position{start: 13, end: 13}},
				Sub{position{start: 15, end: 15}},
				Int{position{start: 17, end: 17}},
				Eof{position{start: 18, end: 18}},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			[]Token{
				Const{position{start: 0, end: 4}},
				Ident{position{start: 6, end: 7}},
				Ass{position{start: 9, end: 9}},
				Int{position{start: 11, end: 13}},
				LShift{position{start: 15, end: 16}},
				Int{position{start: 18, end: 20}},
				Eof{position{start: 21, end: 21}},
			},
		},
		{
			12,
			"p proc; delay=10 ns",
			[]Token{
				Ident{position{start: 0, end: 0}},
				Proc{position{start: 2, end: 5}},
				Semicolon{position{start: 6, end: 6}},
				Delay{position{start: 8, end: 12}},
				Ass{position{start: 13, end: 13}},
				Time{position{start: 14, end: 18}},
				Eof{position{start: 19, end: 19}},
			},
		},
		{
			13,
			"b [a&&true]block",
			[]Token{
				Ident{position{start: 0, end: 0}},
				LBracket{position{start: 2, end: 2}},
				Ident{position{start: 3, end: 3}},
				And{position{start: 4, end: 5}},
				Bool{position{start: 6, end: 9}},
				RBracket{position{start: 10, end: 10}},
				Block{position{start: 11, end: 15}},
				Eof{position{start: 16, end: 16}},
			},
		},
		{
			14,
			"const C_1 = 0xaf| 0x11",
			[]Token{
				Const{position{start: 0, end: 4}},
				Ident{position{start: 6, end: 8}},
				Ass{position{start: 10, end: 10}},
				Int{position{start: 12, end: 15}},
				BitOr{position{start: 16, end: 16}},
				Int{position{start: 18, end: 21}},
				Eof{position{start: 22, end: 22}},
			},
		},
		{
			15,
			"Main bus\n  i irq\n    add-enable = true",
			[]Token{
				Ident{position{start: 0, end: 3}},
				Bus{position{start: 5, end: 7}},
				Newline{position{start: 8, end: 8}},
				Indent{position{start: 9, end: 10}},
				Ident{position{start: 11, end: 11}},
				Irq{position{start: 13, end: 15}},
				Newline{position{start: 16, end: 16}},
				Indent{position{start: 17, end: 20}},
				AddEnable{position{start: 21, end: 30}},
				Ass{position{start: 32, end: 32}},
				Bool{position{start: 34, end: 37}},
				Eof{position{start: 38, end: 38}},
			},
		},
		{
			16,
			"type t static\n  width=7\n\nMain bus",
			[]Token{
				Type{position{start: 0, end: 3}},
				Ident{position{start: 5, end: 5}},
				Static{position{start: 7, end: 12}},
				Newline{position{start: 13, end: 13}},
				Indent{position{start: 14, end: 15}},
				Width{position{start: 16, end: 20}},
				Ass{position{start: 21, end: 21}},
				Int{position{start: 22, end: 22}},
				Newline{position{start: 23, end: 24}},
				Dedent{position{start: 25, end: 25}},
				Ident{position{start: 25, end: 28}},
				Bus{position{start: 30, end: 32}},
				Eof{position{start: 33, end: 33}},
			},
		},
		{
			17,
			"Main bus\n  # Comment\n  c config\n    width = 6\n  # Comment 2\n  s stream",
			[]Token{
				Ident{position{start: 0, end: 3}},
				Bus{position{start: 5, end: 7}},
				Newline{position{start: 8, end: 8}},
				Indent{position{start: 9, end: 10}},
				Comment{position{start: 11, end: 19}},
				Newline{position{start: 20, end: 20}},
				Ident{position{start: 23, end: 23}},
				Config{position{start: 25, end: 30}},
				Newline{position{start: 31, end: 31}},
				Indent{position{start: 32, end: 35}},
				Width{position{start: 36, end: 40}},
				Ass{position{start: 42, end: 42}},
				Int{position{start: 44, end: 44}},
				Newline{position{start: 45, end: 45}},
				Dedent{position{start: 46, end: 47}},
				Comment{position{start: 48, end: 58}},
				Newline{position{start: 59, end: 59}},
				Ident{position{start: 62, end: 62}},
				Stream{position{start: 64, end: 69}},
				Eof{position{start: 70, end: 70}},
			},
		},
		{
			18,
			"masters = -0",
			[]Token{
				Masters{position{start: 0, end: 6}},
				Ass{position{start: 8, end: 8}},
				Sub{position{start: 10, end: 10}},
				Int{position{start: 11, end: 11}},
				Eof{position{start: 12, end: 12}},
			},
		},
		{
			19,
			"size = a-b",
			[]Token{
				Size{position{start: 0, end: 3}},
				Ass{position{start: 5, end: 5}},
				Ident{position{start: 7, end: 7}},
				Sub{position{start: 8, end: 8}},
				Ident{position{start: 9, end: 9}},
				Eof{position{start: 10, end: 10}},
			},
		},
		{
			20,
			"size = init-value",
			[]Token{
				Size{position{start: 0, end: 3}},
				Ass{position{start: 5, end: 5}},
				Ident{position{start: 7, end: 10}},
				Sub{position{start: 11, end: 11}},
				Ident{position{start: 12, end: 16}},
				Eof{position{start: 17, end: 17}},
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
				Const{position{start: 0, end: 4}},
				Newline{position{start: 5, end: 5}},
				Indent{position{start: 6, end: 7}},
				Ident{position{start: 8, end: 8}},
				Ass{position{start: 10, end: 10}},
				Int{position{start: 12, end: 12}},
				Newline{position{start: 13, end: 13}},
				Ident{position{start: 16, end: 16}},
				Ass{position{start: 18, end: 18}},
				Int{position{start: 20, end: 20}},
				Newline{position{start: 38, end: 38}},
				Comment{position{start: 41, end: 53}},
				Newline{position{start: 54, end: 54}},
				Ident{position{start: 57, end: 57}},
				Ass{position{start: 59, end: 59}},
				Float{position{start: 61, end: 64}},
				Eof{position{start: 65, end: 65}},
			},
		},
		{
			22,
			"abc.Def",
			[]Token{
				QualIdent{position{start: 0, end: 6}},
				Eof{position{start: 7, end: 7}},
			},
		},
		{
			23,
			"a-b.C-d.E",
			[]Token{
				Ident{position{start: 0, end: 0}},
				Sub{position{start: 1, end: 1}},
				QualIdent{position{start: 2, end: 4}},
				Sub{position{start: 5, end: 5}},
				QualIdent{position{start: 6, end: 8}},
				Eof{position{start: 9, end: 9}},
			},
		},
		{
			24,
			"range = 1:9",
			[]Token{
				Range{position{start: 0, end: 4}},
				Ass{position{start: 6, end: 6}},
				Int{position{start: 8, end: 8}},
				Colon{position{start: 9, end: 9}},
				Int{position{start: 10, end: 10}},
				Eof{position{start: 11, end: 11}},
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
				got[j].End() != tok.End() {
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
			"odd number (1) of spaces in indent, expected even number",
		},
		{
			1,
			";\n",
			"extra ';' at line end",
		},
		{
			2,
			" ;\n",
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
			"tab character '\\t' allowed only in comments, use spaces",
		},
		{
			16,
			"; \n",
			"extra space at line end",
		},
		{
			17,
			"Main bus\n\t c config",
			"tab character '\\t' allowed only in comments, use spaces",
		},
		{
			18,
			"Main bus\n    c config",
			"multi indent increase, previous indent 0 , current indent 2",
		},
		{
			19,
			"pkg._sym",
			"symbol name in qualified identifier must start with letter",
		},
		{
			20,
			"a-b._c",
			"symbol name in qualified identifier must start with letter",
		},
		{
			21,
			"pkg.3c",
			"symbol name in qualified identifier must start with letter",
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
