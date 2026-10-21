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
				Newline{span{start: 0, end: 1}},
				Newline{span{start: 1, end: 2}},
				Eof{span{start: 2, end: 2}},
			},
		},
		{
			1,
			"# Comment line",
			[]Token{
				Comment{span{start: 0, end: 14}},
				Eof{span{start: 14, end: 14}},
			},
		},
		{
			2,
			"# Comment\n",
			[]Token{
				Comment{span{start: 0, end: 9}},
				Newline{span{start: 9, end: 10}},
				Eof{span{start: 10, end: 10}},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			[]Token{
				Comment{span{start: 0, end: 11}},
				Newline{span{start: 11, end: 12}},
				Comment{span{start: 12, end: 23}},
				Eof{span{start: 23, end: 23}},
			},
		},
		{
			4,
			"const A = true",
			[]Token{
				Const{span{start: 0, end: 5}},
				Ident{span{start: 6, end: 7}},
				Ass{span{start: 8, end: 9}},
				Bool{span{start: 10, end: 14}},
				Eof{span{start: 14, end: 14}},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			[]Token{
				Ident{span{start: 0, end: 3}},
				Mask{span{start: 4, end: 8}},
				Semicolon{span{start: 8, end: 9}},
				Atomic{span{start: 10, end: 16}},
				Ass{span{start: 17, end: 18}},
				Bool{span{start: 19, end: 24}},
				Eof{span{start: 24, end: 24}},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Irq{span{start: 2, end: 5}},
				Semicolon{span{start: 5, end: 6}},
				AddEnable{span{start: 7, end: 17}},
				Ass{span{start: 18, end: 19}},
				Bool{span{start: 20, end: 24}},
				Eof{span{start: 24, end: 24}},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			[]Token{
				Type{span{start: 0, end: 4}},
				Ident{span{start: 5, end: 10}},
				LParen{span{start: 10, end: 11}},
				Ident{span{start: 11, end: 12}},
				Ass{span{start: 13, end: 14}},
				Int{span{start: 15, end: 17}},
				RParen{span{start: 17, end: 18}},
				Config{span{start: 19, end: 25}},
				Semicolon{span{start: 25, end: 26}},
				Width{span{start: 27, end: 32}},
				Ass{span{start: 33, end: 34}},
				Ident{span{start: 35, end: 36}},
				Eof{span{start: 36, end: 36}},
			},
		},
		{
			8,
			"s static; init-value = x\"FFFFFFFF\"",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Static{span{start: 2, end: 8}},
				Semicolon{span{start: 8, end: 9}},
				InitValue{span{start: 10, end: 20}},
				Ass{span{start: 21, end: 22}},
				BitString{span{start: 23, end: 34}},
				Eof{span{start: 34, end: 34}},
			},
		},
		{
			9,
			"import foo \"path\"",
			[]Token{
				Import{span{start: 0, end: 6}},
				Ident{span{start: 7, end: 10}},
				String{span{start: 11, end: 17}},
				Eof{span{start: 17, end: 17}},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			[]Token{
				Const{span{start: 0, end: 5}},
				Ident{span{start: 6, end: 7}},
				Ass{span{start: 8, end: 9}},
				Int{span{start: 10, end: 11}},
				Exp{span{start: 11, end: 13}},
				Int{span{start: 13, end: 14}},
				Sub{span{start: 15, end: 16}},
				Int{span{start: 17, end: 18}},
				Eof{span{start: 18, end: 18}},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			[]Token{
				Const{span{start: 0, end: 5}},
				Ident{span{start: 6, end: 8}},
				Ass{span{start: 9, end: 10}},
				Int{span{start: 11, end: 14}},
				LShift{span{start: 15, end: 17}},
				Int{span{start: 18, end: 21}},
				Eof{span{start: 21, end: 21}},
			},
		},
		{
			12,
			"p proc; delay=10'ns",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Proc{span{start: 2, end: 6}},
				Semicolon{span{start: 6, end: 7}},
				Delay{span{start: 8, end: 13}},
				Ass{span{start: 13, end: 14}},
				Time{span{start: 14, end: 19}},
				Eof{span{start: 19, end: 19}},
			},
		},
		{
			13,
			"b [a&&true]block",
			[]Token{
				Ident{span{start: 0, end: 1}},
				LBracket{span{start: 2, end: 3}},
				Ident{span{start: 3, end: 4}},
				And{span{start: 4, end: 6}},
				Bool{span{start: 6, end: 10}},
				RBracket{span{start: 10, end: 11}},
				Block{span{start: 11, end: 16}},
				Eof{span{start: 16, end: 16}},
			},
		},
		{
			14,
			"const C_1 = 0xaf| 0x11",
			[]Token{
				Const{span{start: 0, end: 5}},
				Ident{span{start: 6, end: 9}},
				Ass{span{start: 10, end: 11}},
				Int{span{start: 12, end: 16}},
				BitOr{span{start: 16, end: 17}},
				Int{span{start: 18, end: 22}},
				Eof{span{start: 22, end: 22}},
			},
		},
		{
			15,
			"Main bus\n  i irq\n    add-enable = true",
			[]Token{
				Ident{span{start: 0, end: 4}},
				Bus{span{start: 5, end: 8}},
				Newline{span{start: 8, end: 9}},
				Indent{span{start: 9, end: 11}},
				Ident{span{start: 11, end: 12}},
				Irq{span{start: 13, end: 16}},
				Newline{span{start: 16, end: 17}},
				Indent{span{start: 17, end: 21}},
				AddEnable{span{start: 21, end: 31}},
				Ass{span{start: 32, end: 33}},
				Bool{span{start: 34, end: 38}},
				Eof{span{start: 38, end: 38}},
			},
		},
		{
			16,
			"type t static\n  width=7\n\nMain bus",
			[]Token{
				Type{span{start: 0, end: 4}},
				Ident{span{start: 5, end: 6}},
				Static{span{start: 7, end: 13}},
				Newline{span{start: 13, end: 14}},
				Indent{span{start: 14, end: 16}},
				Width{span{start: 16, end: 21}},
				Ass{span{start: 21, end: 22}},
				Int{span{start: 22, end: 23}},
				Newline{span{start: 23, end: 24}},
				Newline{span{start: 24, end: 25}},
				Dedent{span{start: 25, end: 26}},
				Ident{span{start: 25, end: 29}},
				Bus{span{start: 30, end: 33}},
				Eof{span{start: 33, end: 33}},
			},
		},
		{
			17,
			"Main bus\n  # Comment\n  c config\n    width = 6\n  # Comment 2\n  s stream",
			[]Token{
				Ident{span{start: 0, end: 4}},
				Bus{span{start: 5, end: 8}},
				Newline{span{start: 8, end: 9}},
				Indent{span{start: 9, end: 11}},
				Comment{span{start: 11, end: 20}},
				Newline{span{start: 20, end: 21}},
				Ident{span{start: 23, end: 24}},
				Config{span{start: 25, end: 31}},
				Newline{span{start: 31, end: 32}},
				Indent{span{start: 32, end: 36}},
				Width{span{start: 36, end: 41}},
				Ass{span{start: 42, end: 43}},
				Int{span{start: 44, end: 45}},
				Newline{span{start: 45, end: 46}},
				Dedent{span{start: 46, end: 48}},
				Comment{span{start: 48, end: 59}},
				Newline{span{start: 59, end: 60}},
				Ident{span{start: 62, end: 63}},
				Stream{span{start: 64, end: 70}},
				Eof{span{start: 70, end: 70}},
			},
		},
		{
			18,
			"\nmasters = -0",
			[]Token{
				Newline{span{start: 0, end: 1}},
				Masters{span{start: 1, end: 8}},
				Ass{span{start: 9, end: 10}},
				Sub{span{start: 11, end: 12}},
				Int{span{start: 12, end: 13}},
				Eof{span{start: 13, end: 13}},
			},
		},
		{
			19,
			"size = a-b",
			[]Token{
				Size{span{start: 0, end: 4}},
				Ass{span{start: 5, end: 6}},
				Ident{span{start: 7, end: 8}},
				Sub{span{start: 8, end: 9}},
				Ident{span{start: 9, end: 10}},
				Eof{span{start: 10, end: 10}},
			},
		},
		{
			20,
			"size = init-value",
			[]Token{
				Size{span{start: 0, end: 4}},
				Ass{span{start: 5, end: 6}},
				Ident{span{start: 7, end: 11}},
				Sub{span{start: 11, end: 12}},
				Ident{span{start: 12, end: 17}},
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
				Const{span{start: 0, end: 5}},
				Newline{span{start: 5, end: 6}},
				Indent{span{start: 6, end: 8}},
				Ident{span{start: 8, end: 9}},
				Ass{span{start: 10, end: 11}},
				Int{span{start: 12, end: 13}},
				Newline{span{start: 13, end: 14}},
				Ident{span{start: 16, end: 17}},
				Ass{span{start: 18, end: 19}},
				Int{span{start: 20, end: 21}},
				Newline{span{start: 38, end: 39}},
				Comment{span{start: 41, end: 54}},
				Newline{span{start: 54, end: 55}},
				Ident{span{start: 57, end: 58}},
				Ass{span{start: 59, end: 60}},
				Float{span{start: 61, end: 65}},
				Eof{span{start: 65, end: 65}},
			},
		},
		{
			22,
			"abc.Def",
			[]Token{
				QualIdent{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			23,
			"a-b.C-d.E",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Sub{span{start: 1, end: 2}},
				QualIdent{span{start: 2, end: 5}},
				Sub{span{start: 5, end: 6}},
				QualIdent{span{start: 6, end: 9}},
				Eof{span{start: 9, end: 9}},
			},
		},
		{
			24,
			"range = 1:9",
			[]Token{
				Range{span{start: 0, end: 5}},
				Ass{span{start: 6, end: 7}},
				Int{span{start: 8, end: 9}},
				Colon{span{start: 9, end: 10}},
				Int{span{start: 10, end: 11}},
				Eof{span{start: 11, end: 11}},
			},
		},
		{
			25,
			"a-foo-ba",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Sub{span{start: 1, end: 2}},
				Ident{span{start: 2, end: 5}},
				Sub{span{start: 5, end: 6}},
				Ident{span{start: 6, end: 8}},
				Eof{span{start: 8, end: 8}},
			},
		},
		{
			26,
			"a-b.C-d.E",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Sub{span{start: 1, end: 2}},
				QualIdent{span{start: 2, end: 5}},
				Sub{span{start: 5, end: 6}},
				QualIdent{span{start: 6, end: 9}},
				Eof{span{start: 9, end: 9}},
			},
		},
		{
			27,
			"pkg.value--other.Value-",
			[]Token{
				QualIdent{span{start: 0, end: 9}},
				Sub{span{start: 9, end: 10}},
				Sub{span{start: 10, end: 11}},
				QualIdent{span{start: 11, end: 22}},
				Sub{span{start: 22, end: 23}},
				Eof{span{start: 23, end: 23}},
			},
		},
		{
			28,
			"x^y",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Xor{span{start: 1, end: 2}},
				Ident{span{start: 2, end: 3}},
				Eof{span{start: 3, end: 3}},
			},
		},
		{
			29,
			"x-",
			[]Token{
				Ident{span{start: 0, end: 1}},
				Sub{span{start: 1, end: 2}},
				Eof{span{start: 2, end: 2}},
			},
		},
		{
			30,
			"enable-init-value = 0",
			[]Token{
				EnableInitValue{span{start: 0, end: 17}},
				Ass{span{start: 18, end: 19}},
				Int{span{start: 20, end: 21}},
				Eof{span{start: 21, end: 21}},
			},
		},
		{
			31,
			"1e-3",
			[]Token{
				Float{span{start: 0, end: 4}},
				Eof{span{start: 4, end: 4}},
			},
		},
		{
			32,
			"1e+3",
			[]Token{
				Float{span{start: 0, end: 4}},
				Eof{span{start: 4, end: 4}},
			},
		},
		{
			33,
			"1E-12",
			[]Token{
				Float{span{start: 0, end: 5}},
				Eof{span{start: 5, end: 5}},
			},
		},
		{
			34,
			"1E+12",
			[]Token{
				Float{span{start: 0, end: 5}},
				Eof{span{start: 5, end: 5}},
			},
		},
		{
			35,
			"1.25e-3",
			[]Token{
				Float{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			36,
			"1.25E+3",
			[]Token{
				Float{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			37,
			"1e3",
			[]Token{
				Float{span{start: 0, end: 3}},
				Eof{span{start: 3, end: 3}},
			},
		},
		{
			38,
			"1e+0",
			[]Token{
				Float{span{start: 0, end: 4}},
				Eof{span{start: 4, end: 4}},
			},
		},
		{
			39,
			"1e-3-2",
			[]Token{
				Float{span{start: 0, end: 4}},
				Sub{span{start: 4, end: 5}},
				Int{span{start: 5, end: 6}},
				Eof{span{start: 6, end: 6}},
			},
		},
		{
			40,
			"1e+3+2",
			[]Token{
				Float{span{start: 0, end: 4}},
				Add{span{start: 4, end: 5}},
				Int{span{start: 5, end: 6}},
				Eof{span{start: 6, end: 6}},
			},
		},
		{
			41,
			"config config",
			[]Token{
				Ident{span{start: 0, end: 6}},
				Config{span{start: 7, end: 13}},
				Eof{span{start: 13, end: 13}},
			},
		},
		{
			42,
			"\nconfig config",
			[]Token{
				Newline{span{start: 0, end: 1}},
				Ident{span{start: 1, end: 7}},
				Config{span{start: 8, end: 14}},
				Eof{span{start: 14, end: 14}},
			},
		},
		{
			43,
			"foo-1",
			[]Token{
				Ident{span{start: 0, end: 3}},
				Sub{span{start: 3, end: 4}},
				Int{span{start: 4, end: 5}},
				Eof{span{start: 5, end: 5}},
			},
		},
		{
			44,
			"foo-1e-1",
			[]Token{
				Ident{span{start: 0, end: 3}},
				Sub{span{start: 3, end: 4}},
				Float{span{start: 4, end: 8}},
				Eof{span{start: 8, end: 8}},
			},
		},
		{
			45,
			"1^2",
			[]Token{
				Int{span{start: 0, end: 1}},
				Xor{span{start: 1, end: 2}},
				Int{span{start: 2, end: 3}},
				Eof{span{start: 3, end: 3}},
			},
		},
		{
			46,
			"1!=2",
			[]Token{
				Int{span{start: 0, end: 1}},
				Neq{span{start: 1, end: 3}},
				Int{span{start: 3, end: 4}},
				Eof{span{start: 4, end: 4}},
			},
		},
		{
			47,
			"1#Comment",
			[]Token{
				Int{span{start: 0, end: 1}},
				Eof{span{start: 9, end: 9}},
			},
		},
		{
			48,
			`b"hHlL"`,
			[]Token{
				BitString{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			49,
			`B"hHlL"`,
			[]Token{
				BitString{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			50,
			`o"hHlL"`,
			[]Token{
				BitString{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			51,
			`O"hHlL"`,
			[]Token{
				BitString{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			52,
			`x"hHlL"`,
			[]Token{
				BitString{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
			},
		},
		{
			53,
			`X"hHlL"`,
			[]Token{
				BitString{span{start: 0, end: 7}},
				Eof{span{start: 7, end: 7}},
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
			"extra space at line end",
		},
		{
			1,
			";\n",
			"extra ';' at line end",
		},
		{
			2,
			"f;\n",
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
		{
			29,
			"0b",
			"missing digit after bin prefix",
		},
		{
			30,
			"0o",
			"missing digit after octal prefix",
		},
		{
			31,
			"0x",
			"missing digit after hex prefix",
		},
		{
			32,
			"0B",
			"missing digit after bin prefix",
		},
		{
			33,
			"0O",
			"missing digit after octal prefix",
		},
		{
			34,
			"0X",
			"missing digit after hex prefix",
		},
		{
			35,
			"0b ",
			"missing digit after bin prefix",
		},
		{
			36,
			"0o\n",
			"missing digit after octal prefix",
		},
		{
			37,
			"0x\t",
			"missing digit after hex prefix",
		},
		{
			38,
			"0b+1",
			"missing digit after bin prefix",
		},
		{
			39,
			"0o)",
			"missing digit after octal prefix",
		},
		{
			40,
			"0x#Comment",
			"missing digit after hex prefix",
		},
		{
			41,
			"a-.b",
			"missing package in qualified identifier",
		},
		{
			42,
			"a--.b",
			"missing package in qualified identifier",
		},
		{
			43,
			"pkg.a-.b",
			"missing package in qualified identifier",
		},
		{
			44,
			"const X = a-.b-c",
			"missing package in qualified identifier",
		},
		{
			45,
			"const X = 1 ",
			"extra space at line end",
		},
		{
			46,
			"const X = 1   ",
			"extra 3 spaces at line end",
		},
		{
			47,
			" ",
			"extra space at line end",
		},
		{
			48,
			"  ",
			"extra 2 spaces at line end",
		},
		{
			49,
			"\n  ",
			"extra 2 spaces at line end",
		},
		{
			50,
			"Main bus\n  width = 32 ",
			"extra space at line end",
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
