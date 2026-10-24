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
				Newline{span{0, 1}},
				Newline{span{1, 2}},
				Eof{span{2, 2}},
			},
		},
		{
			1,
			"# Comment line",
			[]Token{
				Comment{span{0, 14}},
				Eof{span{14, 14}},
			},
		},
		{
			2,
			"# Comment\n",
			[]Token{
				Comment{span{0, 9}},
				Newline{span{9, 10}},
				Eof{span{10, 10}},
			},
		},
		{
			3,
			"# Comment 1\n# Comment 2",
			[]Token{
				Comment{span{0, 11}},
				Newline{span{11, 12}},
				Comment{span{12, 23}},
				Eof{span{23, 23}},
			},
		},
		{
			4,
			"const A = true",
			[]Token{
				Const{span{0, 5}},
				Ident{span{6, 7}},
				Ass{span{8, 9}},
				Bool{span{10, 14}},
				Eof{span{14, 14}},
			},
		},
		{
			5,
			"foo mask; atomic = false",
			[]Token{
				Ident{span{0, 3}},
				Mask{span{4, 8}},
				Semicolon{span{8, 9}},
				Atomic{span{10, 16}},
				Ass{span{17, 18}},
				Bool{span{19, 24}},
				Eof{span{24, 24}},
			},
		},
		{
			6,
			"i irq; add-enable = true",
			[]Token{
				Ident{span{0, 1}},
				Irq{span{2, 5}},
				Semicolon{span{5, 6}},
				AddEnable{span{7, 17}},
				Ass{span{18, 19}},
				Bool{span{20, 24}},
				Eof{span{24, 24}},
			},
		},
		{
			7,
			"type cfg_t(w = 10) config; width = w",
			[]Token{
				Type{span{0, 4}},
				Ident{span{5, 10}},
				LParen{span{10, 11}},
				Ident{span{11, 12}},
				Ass{span{13, 14}},
				Int{span{15, 17}},
				RParen{span{17, 18}},
				Config{span{19, 25}},
				Semicolon{span{25, 26}},
				Width{span{27, 32}},
				Ass{span{33, 34}},
				Ident{span{35, 36}},
				Eof{span{36, 36}},
			},
		},
		{
			8,
			"s static; init-value = x\"FFFFFFFF\"",
			[]Token{
				Ident{span{0, 1}},
				Static{span{2, 8}},
				Semicolon{span{8, 9}},
				InitValue{span{10, 20}},
				Ass{span{21, 22}},
				BitString{span{23, 34}},
				Eof{span{34, 34}},
			},
		},
		{
			9,
			"import foo \"path\"",
			[]Token{
				Import{span{0, 6}},
				Ident{span{7, 10}},
				String{span{11, 17}},
				Eof{span{17, 17}},
			},
		},
		{
			10,
			"const A = 2**5 - 1",
			[]Token{
				Const{span{0, 5}},
				Ident{span{6, 7}},
				Ass{span{8, 9}},
				Int{span{10, 11}},
				Exp{span{11, 13}},
				Int{span{13, 14}},
				Sub{span{15, 16}},
				Int{span{17, 18}},
				Eof{span{18, 18}},
			},
		},
		{
			11,
			"const A1 = 0b1 << 0o3",
			[]Token{
				Const{span{0, 5}},
				Ident{span{6, 8}},
				Ass{span{9, 10}},
				Int{span{11, 14}},
				LShift{span{15, 17}},
				Int{span{18, 21}},
				Eof{span{21, 21}},
			},
		},
		{
			12,
			"p proc; delay=10'ns",
			[]Token{
				Ident{span{0, 1}},
				Proc{span{2, 6}},
				Semicolon{span{6, 7}},
				Delay{span{8, 13}},
				Ass{span{13, 14}},
				Time{span{14, 19}},
				Eof{span{19, 19}},
			},
		},
		{
			13,
			"b [a&&true]block",
			[]Token{
				Ident{span{0, 1}},
				LBracket{span{2, 3}},
				Ident{span{3, 4}},
				And{span{4, 6}},
				Bool{span{6, 10}},
				RBracket{span{10, 11}},
				Block{span{11, 16}},
				Eof{span{16, 16}},
			},
		},
		{
			14,
			"const C_1 = 0xaf| 0x11",
			[]Token{
				Const{span{0, 5}},
				Ident{span{6, 9}},
				Ass{span{10, 11}},
				Int{span{12, 16}},
				BitOr{span{16, 17}},
				Int{span{18, 22}},
				Eof{span{22, 22}},
			},
		},
		{
			15,
			"Main bus\n  i irq\n    add-enable = true",
			[]Token{
				Ident{span{0, 4}},
				Bus{span{5, 8}},
				Newline{span{8, 9}},
				Indent{span{9, 11}},
				Ident{span{11, 12}},
				Irq{span{13, 16}},
				Newline{span{16, 17}},
				Indent{span{17, 21}},
				AddEnable{span{21, 31}},
				Ass{span{32, 33}},
				Bool{span{34, 38}},
				Eof{span{38, 38}},
			},
		},
		{
			16,
			"type t static\n  width=7\n\nMain bus",
			[]Token{
				Type{span{0, 4}},
				Ident{span{5, 6}},
				Static{span{7, 13}},
				Newline{span{13, 14}},
				Indent{span{14, 16}},
				Width{span{16, 21}},
				Ass{span{21, 22}},
				Int{span{22, 23}},
				Newline{span{23, 24}},
				Newline{span{24, 25}},
				Dedent{span{25, 26}},
				Ident{span{25, 29}},
				Bus{span{30, 33}},
				Eof{span{33, 33}},
			},
		},
		{
			17,
			"Main bus\n  # Comment\n  c config\n    width = 6\n  # Comment 2\n  s stream",
			[]Token{
				Ident{span{0, 4}},
				Bus{span{5, 8}},
				Newline{span{8, 9}},
				Indent{span{9, 11}},
				Comment{span{11, 20}},
				Newline{span{20, 21}},
				Ident{span{23, 24}},
				Config{span{25, 31}},
				Newline{span{31, 32}},
				Indent{span{32, 36}},
				Width{span{36, 41}},
				Ass{span{42, 43}},
				Int{span{44, 45}},
				Newline{span{45, 46}},
				Dedent{span{46, 48}},
				Comment{span{48, 59}},
				Newline{span{59, 60}},
				Ident{span{62, 63}},
				Stream{span{64, 70}},
				Eof{span{70, 70}},
			},
		},
		{
			18,
			"\nmasters = -0",
			[]Token{
				Newline{span{0, 1}},
				Masters{span{1, 8}},
				Ass{span{9, 10}},
				Sub{span{11, 12}},
				Int{span{12, 13}},
				Eof{span{13, 13}},
			},
		},
		{
			19,
			"size = a-b",
			[]Token{
				Size{span{0, 4}},
				Ass{span{5, 6}},
				Ident{span{7, 8}},
				Sub{span{8, 9}},
				Ident{span{9, 10}},
				Eof{span{10, 10}},
			},
		},
		{
			20,
			"size = init-value",
			[]Token{
				Size{span{0, 4}},
				Ass{span{5, 6}},
				InitValue{span{7, 17}},
				Eof{span{17, 17}},
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
				Const{span{0, 5}},
				Newline{span{5, 6}},
				Indent{span{6, 8}},
				Ident{span{8, 9}},
				Ass{span{10, 11}},
				Int{span{12, 13}},
				Newline{span{13, 14}},
				Ident{span{16, 17}},
				Ass{span{18, 19}},
				Int{span{20, 21}},
				Newline{span{38, 39}},
				Comment{span{41, 54}},
				Newline{span{54, 55}},
				Ident{span{57, 58}},
				Ass{span{59, 60}},
				Float{span{61, 65}},
				Eof{span{65, 65}},
			},
		},
		{
			22,
			"abc.Def",
			[]Token{
				QualIdent{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			23,
			"a-b.C-d.E",
			[]Token{
				Ident{span{0, 1}},
				Sub{span{1, 2}},
				QualIdent{span{2, 5}},
				Sub{span{5, 6}},
				QualIdent{span{6, 9}},
				Eof{span{9, 9}},
			},
		},
		{
			24,
			"range = 1:9",
			[]Token{
				Range{span{0, 5}},
				Ass{span{6, 7}},
				Int{span{8, 9}},
				Colon{span{9, 10}},
				Int{span{10, 11}},
				Eof{span{11, 11}},
			},
		},
		{
			25,
			"a-foo-ba",
			[]Token{
				Ident{span{0, 1}},
				Sub{span{1, 2}},
				Ident{span{2, 5}},
				Sub{span{5, 6}},
				Ident{span{6, 8}},
				Eof{span{8, 8}},
			},
		},
		{
			26,
			"a-b.C-d.E",
			[]Token{
				Ident{span{0, 1}},
				Sub{span{1, 2}},
				QualIdent{span{2, 5}},
				Sub{span{5, 6}},
				QualIdent{span{6, 9}},
				Eof{span{9, 9}},
			},
		},
		{
			27,
			"pkg.value--other.Value-",
			[]Token{
				QualIdent{span{0, 9}},
				Sub{span{9, 10}},
				Sub{span{10, 11}},
				QualIdent{span{11, 22}},
				Sub{span{22, 23}},
				Eof{span{23, 23}},
			},
		},
		{
			28,
			"x^y",
			[]Token{
				Ident{span{0, 1}},
				Xor{span{1, 2}},
				Ident{span{2, 3}},
				Eof{span{3, 3}},
			},
		},
		{
			29,
			"x-",
			[]Token{
				Ident{span{0, 1}},
				Sub{span{1, 2}},
				Eof{span{2, 2}},
			},
		},
		{
			30,
			"enable-init-value = 0",
			[]Token{
				EnableInitValue{span{0, 17}},
				Ass{span{18, 19}},
				Int{span{20, 21}},
				Eof{span{21, 21}},
			},
		},
		{
			31,
			"1e-3",
			[]Token{
				Float{span{0, 4}},
				Eof{span{4, 4}},
			},
		},
		{
			32,
			"1e+3",
			[]Token{
				Float{span{0, 4}},
				Eof{span{4, 4}},
			},
		},
		{
			33,
			"1E-12",
			[]Token{
				Float{span{0, 5}},
				Eof{span{5, 5}},
			},
		},
		{
			34,
			"1E+12",
			[]Token{
				Float{span{0, 5}},
				Eof{span{5, 5}},
			},
		},
		{
			35,
			"1.25e-3",
			[]Token{
				Float{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			36,
			"1.25E+3",
			[]Token{
				Float{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			37,
			"1e3",
			[]Token{
				Float{span{0, 3}},
				Eof{span{3, 3}},
			},
		},
		{
			38,
			"1e+0",
			[]Token{
				Float{span{0, 4}},
				Eof{span{4, 4}},
			},
		},
		{
			39,
			"1e-3-2",
			[]Token{
				Float{span{0, 4}},
				Sub{span{4, 5}},
				Int{span{5, 6}},
				Eof{span{6, 6}},
			},
		},
		{
			40,
			"1e+3+2",
			[]Token{
				Float{span{0, 4}},
				Add{span{4, 5}},
				Int{span{5, 6}},
				Eof{span{6, 6}},
			},
		},
		{
			41,
			"config config",
			[]Token{
				Config{span{0, 6}},
				Config{span{7, 13}},
				Eof{span{13, 13}},
			},
		},
		{
			42,
			"\nconfig config",
			[]Token{
				Newline{span{0, 1}},
				Config{span{1, 7}},
				Config{span{8, 14}},
				Eof{span{14, 14}},
			},
		},
		{
			43,
			"foo-1",
			[]Token{
				Ident{span{0, 3}},
				Sub{span{3, 4}},
				Int{span{4, 5}},
				Eof{span{5, 5}},
			},
		},
		{
			44,
			"foo-1e-1",
			[]Token{
				Ident{span{0, 3}},
				Sub{span{3, 4}},
				Float{span{4, 8}},
				Eof{span{8, 8}},
			},
		},
		{
			45,
			"1^2",
			[]Token{
				Int{span{0, 1}},
				Xor{span{1, 2}},
				Int{span{2, 3}},
				Eof{span{3, 3}},
			},
		},
		{
			46,
			"1!=2",
			[]Token{
				Int{span{0, 1}},
				Neq{span{1, 3}},
				Int{span{3, 4}},
				Eof{span{4, 4}},
			},
		},
		{
			47,
			"1#Comment",
			[]Token{
				Int{span{0, 1}},
				Eof{span{9, 9}},
			},
		},
		{
			48,
			`b"hHlL"`,
			[]Token{
				BitString{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			49,
			`B"hHlL"`,
			[]Token{
				BitString{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			50,
			`o"hHlL"`,
			[]Token{
				BitString{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			51,
			`O"hHlL"`,
			[]Token{
				BitString{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			52,
			`x"hHlL"`,
			[]Token{
				BitString{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			53,
			`X"hHlL"`,
			[]Token{
				BitString{span{0, 7}},
				Eof{span{7, 7}},
			},
		},
		{
			54,
			`const config = atomic`,
			[]Token{
				Const{span{0, 5}},
				Config{span{6, 12}},
				Ass{span{13, 14}},
				Atomic{span{15, 21}},
				Eof{span{21, 21}},
			},
		},
		{
			55,
			`const bus = true-false`,
			[]Token{
				Const{span{0, 5}},
				Bus{span{6, 9}},
				Ass{span{10, 11}},
				Bool{span{12, 16}},
				Sub{span{16, 17}},
				Bool{span{17, 22}},
				Eof{span{22, 22}},
			},
		},
		{
			56,
			`const bus = true- false`,
			[]Token{
				Const{span{0, 5}},
				Bus{span{6, 9}},
				Ass{span{10, 11}},
				Bool{span{12, 16}},
				Sub{span{16, 17}},
				Bool{span{18, 23}},
				Eof{span{23, 23}},
			},
		},
		{
			57,
			`const c = a-b"01"`,
			[]Token{
				Const{span{0, 5}},
				Ident{span{6, 7}},
				Ass{span{8, 9}},
				Ident{span{10, 11}},
				Sub{span{11, 12}},
				BitString{span{12, 17}},
				Eof{span{17, 17}},
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
			"tab character '\\t' allowed only in comments and strings, use spaces",
		},
		{
			16,
			"; \n",
			"extra space at line end",
		},
		{
			17,
			"Main bus\n\t c config",
			"tab character '\\t' allowed only in comments and strings, use spaces",
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

func TestParseSubtractionWords(t *testing.T) {
	for _, test := range []struct {
		src  string
		want []Token
	}{
		{
			"config-width",
			[]Token{Config{span{0, 6}}, Sub{span{6, 7}}, Width{span{7, 12}}, Eof{span{12, 12}}},
		},
		{
			"true-pkg.value-false",
			[]Token{Bool{span{0, 4}}, Sub{span{4, 5}}, QualIdent{span{5, 14}}, Sub{span{14, 15}}, Bool{span{15, 20}}, Eof{span{20, 20}}},
		},
		{
			"true--false",
			[]Token{Bool{span{0, 4}}, Sub{span{4, 5}}, Sub{span{5, 6}}, Bool{span{6, 11}}, Eof{span{11, 11}}},
		},
	} {
		t.Run(test.src, func(t *testing.T) {
			got, err := Parse([]byte(test.src))
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("Parse() = %#v, want %#v", got, test.want)
			}
		})
	}
}

func TestParseSubtractionBitStrings(t *testing.T) {
	for _, literal := range []string{`b"01"`, `B"01"`, `o"07"`, `O"07"`, `x"af"`, `X"af"`} {
		for _, left := range []struct {
			src string
			tok Token
		}{
			{"a", Ident{span{0, 1}}},
			{"pkg.a", QualIdent{span{0, 5}}},
		} {
			src := left.src + "-" + literal
			t.Run(src, func(t *testing.T) {
				got, err := Parse([]byte(src))
				if err != nil {
					t.Fatal(err)
				}
				end := len(src)
				want := []Token{
					left.tok,
					Sub{span{len(left.src), len(left.src) + 1}},
					BitString{span{len(left.src) + 1, end}},
					Eof{span{end, end}},
				}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("Parse() = %#v, want %#v", got, want)
				}
			})
		}
	}
}
