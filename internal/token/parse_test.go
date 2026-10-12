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
				Newline{span{start: 0, end: 0}},
				Newline{span{start: 1, end: 1}},
				Eof{span{start: 2, end: 2}},
			},
		},
		{
			1,
			"# Comment line",
			[]Token{
				Comment{span{start: 0, end: 13}},
				Eof{span{start: 14, end: 14}},
			},
		},
		{
			2,
			"# Comment\n",
			[]Token{
				Comment{span{start: 0, end: 8}},
				Newline{span{start: 9, end: 9}},
				Eof{span{start: 10, end: 10}},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			[]Token{
				Comment{span{start: 0, end: 10}},
				Newline{span{start: 11, end: 11}},
				Comment{span{start: 12, end: 22}},
				Eof{span{start: 23, end: 23}},
			},
		},
		{
			4,
			"const A = true",
			[]Token{
				Const{span{start: 0, end: 4}},
				Ident{span{start: 6, end: 6}},
				Ass{span{start: 8, end: 8}},
				Bool{span{start: 10, end: 13}},
				Eof{span{start: 14, end: 14}},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			[]Token{
				Ident{span{start: 0, end: 2}},
				Mask{span{start: 4, end: 7}},
				Semicolon{span{start: 8, end: 8}},
				Atomic{span{start: 10, end: 15}},
				Ass{span{start: 17, end: 17}},
				Bool{span{start: 19, end: 23}},
				Eof{span{start: 24, end: 24}},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Irq{span{start: 2, end: 4}},
				Semicolon{span{start: 5, end: 5}},
				AddEnable{span{start: 7, end: 16}},
				Ass{span{start: 18, end: 18}},
				Bool{span{start: 20, end: 23}},
				Eof{span{start: 24, end: 24}},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			[]Token{
				Type{span{start: 0, end: 3}},
				Ident{span{start: 5, end: 9}},
				LParen{span{start: 10, end: 10}},
				Ident{span{start: 11, end: 11}},
				Ass{span{start: 13, end: 13}},
				Int{span{start: 15, end: 16}},
				RParen{span{start: 17, end: 17}},
				Config{span{start: 19, end: 24}},
				Semicolon{span{start: 25, end: 25}},
				Width{span{start: 27, end: 31}},
				Ass{span{start: 33, end: 33}},
				Ident{span{start: 35, end: 35}},
				Eof{span{start: 36, end: 36}},
			},
		},
		{
			8,
			"s static; init-value = x\"FFFFFFFF\"",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Static{span{start: 2, end: 7}},
				Semicolon{span{start: 8, end: 8}},
				InitValue{span{start: 10, end: 19}},
				Ass{span{start: 21, end: 21}},
				BitString{span{start: 23, end: 33}},
				Eof{span{start: 34, end: 34}},
			},
		},
		{
			9,
			"import foo \"path\"",
			[]Token{
				Import{span{start: 0, end: 5}},
				Ident{span{start: 7, end: 9}},
				String{span{start: 11, end: 16}},
				Eof{span{start: 17, end: 17}},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			[]Token{
				Const{span{start: 0, end: 4}},
				Ident{span{start: 6, end: 6}},
				Ass{span{start: 8, end: 8}},
				Int{span{start: 10, end: 10}},
				Exp{span{start: 11, end: 12}},
				Int{span{start: 13, end: 13}},
				Sub{span{start: 15, end: 15}},
				Int{span{start: 17, end: 17}},
				Eof{span{start: 18, end: 18}},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			[]Token{
				Const{span{start: 0, end: 4}},
				Ident{span{start: 6, end: 7}},
				Ass{span{start: 9, end: 9}},
				Int{span{start: 11, end: 13}},
				LShift{span{start: 15, end: 16}},
				Int{span{start: 18, end: 20}},
				Eof{span{start: 21, end: 21}},
			},
		},
		{
			12,
			"p proc; delay=10'ns",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Proc{span{start: 2, end: 5}},
				Semicolon{span{start: 6, end: 6}},
				Delay{span{start: 8, end: 12}},
				Ass{span{start: 13, end: 13}},
				Time{span{start: 14, end: 18}},
				Eof{span{start: 19, end: 19}},
			},
		},
		{
			13,
			"b [a&&true]block",
			[]Token{
				Ident{span{start: 0, end: 0}},
				LBracket{span{start: 2, end: 2}},
				Ident{span{start: 3, end: 3}},
				And{span{start: 4, end: 5}},
				Bool{span{start: 6, end: 9}},
				RBracket{span{start: 10, end: 10}},
				Block{span{start: 11, end: 15}},
				Eof{span{start: 16, end: 16}},
			},
		},
		{
			14,
			"const C_1 = 0xaf| 0x11",
			[]Token{
				Const{span{start: 0, end: 4}},
				Ident{span{start: 6, end: 8}},
				Ass{span{start: 10, end: 10}},
				Int{span{start: 12, end: 15}},
				BitOr{span{start: 16, end: 16}},
				Int{span{start: 18, end: 21}},
				Eof{span{start: 22, end: 22}},
			},
		},
		{
			15,
			"Main bus\n  i irq\n    add-enable = true",
			[]Token{
				Ident{span{start: 0, end: 3}},
				Bus{span{start: 5, end: 7}},
				Newline{span{start: 8, end: 8}},
				Indent{span{start: 9, end: 10}},
				Ident{span{start: 11, end: 11}},
				Irq{span{start: 13, end: 15}},
				Newline{span{start: 16, end: 16}},
				Indent{span{start: 17, end: 20}},
				AddEnable{span{start: 21, end: 30}},
				Ass{span{start: 32, end: 32}},
				Bool{span{start: 34, end: 37}},
				Eof{span{start: 38, end: 38}},
			},
		},
		{
			16,
			"type t static\n  width=7\n\nMain bus",
			[]Token{
				Type{span{start: 0, end: 3}},
				Ident{span{start: 5, end: 5}},
				Static{span{start: 7, end: 12}},
				Newline{span{start: 13, end: 13}},
				Indent{span{start: 14, end: 15}},
				Width{span{start: 16, end: 20}},
				Ass{span{start: 21, end: 21}},
				Int{span{start: 22, end: 22}},
				Newline{span{start: 23, end: 23}},
				Newline{span{start: 24, end: 24}},
				Dedent{span{start: 25, end: 25}},
				Ident{span{start: 25, end: 28}},
				Bus{span{start: 30, end: 32}},
				Eof{span{start: 33, end: 33}},
			},
		},
		{
			17,
			"Main bus\n  # Comment\n  c config\n    width = 6\n  # Comment 2\n  s stream",
			[]Token{
				Ident{span{start: 0, end: 3}},
				Bus{span{start: 5, end: 7}},
				Newline{span{start: 8, end: 8}},
				Indent{span{start: 9, end: 10}},
				Comment{span{start: 11, end: 19}},
				Newline{span{start: 20, end: 20}},
				Ident{span{start: 23, end: 23}},
				Config{span{start: 25, end: 30}},
				Newline{span{start: 31, end: 31}},
				Indent{span{start: 32, end: 35}},
				Width{span{start: 36, end: 40}},
				Ass{span{start: 42, end: 42}},
				Int{span{start: 44, end: 44}},
				Newline{span{start: 45, end: 45}},
				Dedent{span{start: 46, end: 47}},
				Comment{span{start: 48, end: 58}},
				Newline{span{start: 59, end: 59}},
				Ident{span{start: 62, end: 62}},
				Stream{span{start: 64, end: 69}},
				Eof{span{start: 70, end: 70}},
			},
		},
		{
			18,
			"\nmasters = -0",
			[]Token{
				Newline{span{start: 0, end: 0}},
				Masters{span{start: 1, end: 7}},
				Ass{span{start: 9, end: 9}},
				Sub{span{start: 11, end: 11}},
				Int{span{start: 12, end: 12}},
				Eof{span{start: 13, end: 13}},
			},
		},
		{
			19,
			"size = a-b",
			[]Token{
				Size{span{start: 0, end: 3}},
				Ass{span{start: 5, end: 5}},
				Ident{span{start: 7, end: 7}},
				Sub{span{start: 8, end: 8}},
				Ident{span{start: 9, end: 9}},
				Eof{span{start: 10, end: 10}},
			},
		},
		{
			20,
			"size = init-value",
			[]Token{
				Size{span{start: 0, end: 3}},
				Ass{span{start: 5, end: 5}},
				Ident{span{start: 7, end: 10}},
				Sub{span{start: 11, end: 11}},
				Ident{span{start: 12, end: 16}},
				Eof{span{start: 17, end: 17}},
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
				Const{span{start: 0, end: 4}},
				Newline{span{start: 5, end: 5}},
				Indent{span{start: 6, end: 7}},
				Ident{span{start: 8, end: 8}},
				Ass{span{start: 10, end: 10}},
				Int{span{start: 12, end: 12}},
				Newline{span{start: 13, end: 13}},
				Ident{span{start: 16, end: 16}},
				Ass{span{start: 18, end: 18}},
				Int{span{start: 20, end: 20}},
				Newline{span{start: 38, end: 38}},
				Comment{span{start: 41, end: 53}},
				Newline{span{start: 54, end: 54}},
				Ident{span{start: 57, end: 57}},
				Ass{span{start: 59, end: 59}},
				Float{span{start: 61, end: 64}},
				Eof{span{start: 65, end: 65}},
			},
		},
		{
			22,
			"abc.Def",
			[]Token{
				QualIdent{span{start: 0, end: 6}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			23,
			"a-b.C-d.E",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Sub{span{start: 1, end: 1}},
				QualIdent{span{start: 2, end: 4}},
				Sub{span{start: 5, end: 5}},
				QualIdent{span{start: 6, end: 8}},
				Eof{span{start: 9, end: 9}},
			},
		},
		{
			24,
			"range = 1:9",
			[]Token{
				Range{span{start: 0, end: 4}},
				Ass{span{start: 6, end: 6}},
				Int{span{start: 8, end: 8}},
				Colon{span{start: 9, end: 9}},
				Int{span{start: 10, end: 10}},
				Eof{span{start: 11, end: 11}},
			},
		},
		{
			25,
			"a-foo-ba",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Sub{span{start: 1, end: 1}},
				Ident{span{start: 2, end: 4}},
				Sub{span{start: 5, end: 5}},
				Ident{span{start: 6, end: 7}},
				Eof{span{start: 8, end: 8}},
			},
		},
		{
			26,
			"a-b.C-d.E",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Sub{span{start: 1, end: 1}},
				QualIdent{span{start: 2, end: 4}},
				Sub{span{start: 5, end: 5}},
				QualIdent{span{start: 6, end: 8}},
				Eof{span{start: 9, end: 9}},
			},
		},
		{
			27,
			"pkg.value--other.Value-",
			[]Token{
				QualIdent{span{start: 0, end: 8}},
				Sub{span{start: 9, end: 9}},
				Sub{span{start: 10, end: 10}},
				QualIdent{span{start: 11, end: 21}},
				Sub{span{start: 22, end: 22}},
				Eof{span{start: 23, end: 23}},
			},
		},
		{
			28,
			"x^y",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Xor{span{start: 1, end: 1}},
				Ident{span{start: 2, end: 2}},
				Eof{span{start: 3, end: 3}},
			},
		},
		{
			29,
			"x-",
			[]Token{
				Ident{span{start: 0, end: 0}},
				Sub{span{start: 1, end: 1}},
				Eof{span{start: 2, end: 2}},
			},
		},
		{
			30,
			"enable-init-value = 0",
			[]Token{
				EnableInitValue{span{start: 0, end: 16}},
				Ass{span{start: 18, end: 18}},
				Int{span{start: 20, end: 20}},
				Eof{span{start: 21, end: 21}},
			},
		},
		{
			31,
			"1e-3",
			[]Token{
				Float{span{start: 0, end: 3}},
				Eof{span{start: 4, end: 4}},
			},
		},
		{
			32,
			"1e+3",
			[]Token{
				Float{span{start: 0, end: 3}},
				Eof{span{start: 4, end: 4}},
			},
		},
		{
			33,
			"1E-12",
			[]Token{
				Float{span{start: 0, end: 4}},
				Eof{span{start: 5, end: 5}},
			},
		},
		{
			34,
			"1E+12",
			[]Token{
				Float{span{start: 0, end: 4}},
				Eof{span{start: 5, end: 5}},
			},
		},
		{
			35,
			"1.25e-3",
			[]Token{
				Float{span{start: 0, end: 6}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			36,
			"1.25E+3",
			[]Token{
				Float{span{start: 0, end: 6}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			37,
			"1e3",
			[]Token{
				Float{span{start: 0, end: 2}},
				Eof{span{start: 3, end: 3}},
			},
		},
		{
			38,
			"1e+0",
			[]Token{
				Float{span{start: 0, end: 3}},
				Eof{span{start: 4, end: 4}},
			},
		},
		{
			39,
			"1e-3-2",
			[]Token{
				Float{span{start: 0, end: 3}},
				Sub{span{start: 4, end: 4}},
				Int{span{start: 5, end: 5}},
				Eof{span{start: 6, end: 6}},
			},
		},
		{
			40,
			"1e+3+2",
			[]Token{
				Float{span{start: 0, end: 3}},
				Add{span{start: 4, end: 4}},
				Int{span{start: 5, end: 5}},
				Eof{span{start: 6, end: 6}},
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
			gotType := reflect.TypeOf(got[j])
			wantType := reflect.TypeOf(tok)
			if gotType != wantType ||
				got[j].Start() != tok.Start() ||
				got[j].End() != tok.End() {
				t.Fatalf(
					"\nTest: %d\n\nCode:\n%s\n\nToken: %d\ngot type: %v\ngot: %+v\nwant type %v\nwant: %+v",
					i, test.src, j, gotType, got[j], wantType, tok,
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
			"unterminated string literal",
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
		{
			22,
			"1e",
			"missing digits in number exponent",
		},
		{
			23,
			"1e+",
			"missing digits in number exponent",
		},
		{
			24,
			"1e-",
			"missing digits in number exponent",
		},
		{
			25,
			"1E+ ",
			"missing digits in number exponent",
		},
		{
			26,
			"1e-\n",
			"missing digits in number exponent",
		},
		{
			27,
			"1e+-3",
			"missing digits in number exponent",
		},
		{
			28,
			"1e--3",
			"missing digits in number exponent",
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

		tokErr := err.(Error)
		if tokErr.Msg != test.err {
			t.Fatalf("\nTest %d:\n\ngot:\n%v\n\nwant:\n%v", i, tokErr.Msg, test.err)
		}
	}
}
