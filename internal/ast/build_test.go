package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildSingleConst(t *testing.T) {
	toks, _ := token.Parse([]byte("const A = 15"))
	want := SingleConst{
		Name: toks[1].(token.Ident),
		Expr: Int{toks[3].(token.Int)},
	}
	c := ctx{}
	got, err := buildConst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 4 {
		t.Fatalf("c.i = %d", c.i)
	}
	if got != want {
		t.Fatalf("got: %+v, want %+v", got, want)
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
