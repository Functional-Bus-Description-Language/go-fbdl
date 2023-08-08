package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildSingleConst(t *testing.T) {
	toks, _ := token.Parse([]byte("const A = 15"))
	want := Const{
		Name:  toks[1].(token.Ident),
		Value: Int{toks[3].(token.Int)},
	}
	c := ctx{}
	got, err := buildSingleConst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 4 {
		t.Fatalf("c.i = %d", c.i)
	}
	if !got[0].eq(want) {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}
}

func TestBuildMultiConst(t *testing.T) {
	toks, _ := token.Parse([]byte(`const
	A = 1
	B = 2 # Inline comment
	# Doc comment
	C = 3.14

	D = false`),
	)
	want := []Const{
		Const{Name: toks[3].(token.Ident), Value: Int{toks[5].(token.Int)}},
		Const{Name: toks[7].(token.Ident), Value: Int{toks[9].(token.Int)}},
		Const{
			Doc:   Doc{Lines: []token.Comment{toks[11].(token.Comment)}},
			Name:  toks[13].(token.Ident),
			Value: Real{toks[15].(token.Real)},
		},
		Const{Name: toks[17].(token.Ident), Value: Bool{toks[19].(token.Bool)}},
	}
	c := ctx{}
	got, err := buildMultiConst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 20 {
		t.Fatalf("c.i = %d", c.i)
	}

	for i := range want {
		if !got[i].eq(want[i]) {
			t.Fatalf("i: %d\ngot:\n%+v,\nwant\n%+v", i, got[i], want[i])
		}
	}
}

func TestBuildInstSingleLine(t *testing.T) {
	toks, _ := token.Parse([]byte("S [5]status; atomic = false; width = 10"))
	want := Inst{
		Name:  toks[0].(token.Ident),
		Count: Int{toks[2].(token.Int)},
		Type:  toks[4].(token.Status),
		Body: Body{
			Props: []Prop{
				Prop{toks[6].(token.Atomic), Bool{toks[8].(token.Bool)}},
				Prop{toks[10].(token.Width), Int{toks[12].(token.Int)}},
			},
		},
	}

	c := ctx{}
	got, err := buildInst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 13 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !got.eq(want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}

func TestBuildInstMultiLine(t *testing.T) {
	toks, _ := token.Parse([]byte(`B pkg.block_t(1, PI = 3.14)
	masters = 2; reset = "Sync"
	const FOO = true
	C config
		range = 8`),
	)
	want := Inst{
		Name: toks[0].(token.Ident),
		Type: toks[1].(token.QualIdent),
		Args: []Arg{
			Arg{Value: Int{toks[3].(token.Int)}},
			Arg{toks[5].(token.Ident), Real{toks[7].(token.Real)}},
		},
		Body: Body{
			Consts: []Const{Const{Name: toks[20].(token.Ident), Value: Bool{toks[22].(token.Bool)}}},
			Insts: []Inst{
				Inst{
					Name: toks[24].(token.Ident),
					Type: toks[25].(token.Config),
					Body: Body{
						Props: []Prop{
							Prop{toks[28].(token.Range), Int{toks[30].(token.Int)}},
						},
					},
				},
			},
			Props: []Prop{
				Prop{toks[11].(token.Masters), Int{toks[13].(token.Int)}},
				Prop{toks[15].(token.Reset), String{toks[17].(token.String)}},
			},
		},
	}

	c := ctx{}
	got, err := buildInst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}

	if c.i != 31 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !got.eq(want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}

func TestBuildError(t *testing.T) {
	var tests = []struct {
		idx int // Test index, useful for navigation
		src string
		err error
	}{
		{
			0,
			"import *",
			fmt.Errorf("1:8: unexpected *, expected identifier, string or newline"),
		},
		{
			1,
			"import name +",
			fmt.Errorf("1:13: unexpected +, expected string"),
		},
		{
			2,
			"const A = ]",
			fmt.Errorf("1:11: unexpected ], expected expression"),
		},
		{
			3,
			"const A = foo(1 true)",
			fmt.Errorf("1:17: unexpected bool, expected , or )"),
		},
		{
			4,
			"const A = foo(, 1)",
			fmt.Errorf("1:15: unexpected ,, expected expression"),
		},
		{
			5,
			"const A = (a + b c",
			fmt.Errorf("1:18: unexpected identifier, expected )"),
		},
		{
			6,
			"const A = [1, 2, 3 4]",
			fmt.Errorf("1:20: unexpected integer, expected , or ]"),
		},
		{
			7,
			"const A = [, 1]",
			fmt.Errorf("1:12: unexpected ,, expected expression"),
		},
		{
			8,
			"const\n\tA 12",
			fmt.Errorf("2:4: unexpected integer, expected ="),
		},
		{
			9,
			"const\n\t2.24 = A",
			fmt.Errorf("2:2: unexpected real, expected identifier"),
		},
		{
			10,
			"const\nA = 2",
			fmt.Errorf("2:1: unexpected identifier, expected indent or newline"),
		},
		{
			11,
			"C type_t()",
			fmt.Errorf("1:9: empty argument list"),
		},
		{
			12,
			"C [3;config",
			fmt.Errorf("1:5: unexpected ;, expected ]"),
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		stream, err := token.Parse([]byte(test.src))
		if err != nil {
			t.Fatalf("%d: token.Parse: %v, expected nil", i, err)
		}

		_, err = Build(stream)
		if err == nil {
			t.Fatalf("%d: err == nil, expected != nil", i)
		}

		if err.Error() != test.err.Error() {
			t.Fatalf("%d:\n got: %v\nwant: %v", i, err, test.err)
		}
	}
}
