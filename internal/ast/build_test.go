package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"testing"
)

func TestBuildError(t *testing.T) {
	var tests = []struct {
		idx int // Test index, useful for navigation
		src string
		err string
	}{
		{
			0,
			"import *",
			"unexpected '*', expected identifier, string or newline",
		},
		{
			1,
			"import name +",
			"unexpected '+', expected string",
		},
		{
			2,
			"const A = ]",
			"unexpected ']', expected expression",
		},
		{
			3,
			"const A = foo(1 true)",
			"unexpected bool, expected ',' or ')'",
		},
		{
			4,
			"const A = foo(, 1)",
			"unexpected ',', expected expression",
		},
		{
			5,
			"const A = (a + b c",
			"unexpected identifier, expected ')'",
		},
		{
			6,
			"const A = [1, 2, 3 4]",
			"unexpected integer, expected ',' or ']'",
		},
		{
			7,
			"const A = [, 1]",
			"unexpected ',', expected expression",
		},
		{
			8,
			"const\n\tA 12",
			"unexpected integer, expected '='",
		},
		{
			9,
			"const\n\t2.24 = A",
			"unexpected float, expected identifier",
		},
		{
			10,
			"const\nA = 2",
			"unexpected identifier, expected indent or newline",
		},
		{
			11,
			"C type_t()",
			"empty argument list",
		},
		{
			12,
			"C [3;config",
			"unexpected ';', expected ']'",
		},
		{
			13,
			"type type_t() config",
			"empty parameter list",
		},
		{
			14,
			"type t(,) config",
			"unexpected ',', expected identifier",
		},
		{
			15,
			"type t(a b) config",
			"unexpected identifier, expected '=', ')' or ','",
		},
		{
			16,
			"type t(a = 1;) static",
			"unexpected ';', expected ',' or ')'",
		},
		{
			17,
			"type 1 status",
			"unexpected integer, expected identifier",
		},
		{
			18,
			"type a [1,status",
			"unexpected ',', expected ']'",
		},
		{
			19,
			"import\n1",
			"unexpected integer, expected indent increase",
		},
		{
			20,
			"import\n\tabc 1",
			"unexpected integer, expected string",
		},
		{
			21,
			"import\n\t1",
			"unexpected integer, expected identifier or string",
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		_, err := Build([]byte(test.src), "")
		if err == nil {
			t.Fatalf("%d: err == nil, expected != nil", i)
		}

		tokErr := err.(tok.Error)
		if tokErr.Msg != test.err {
			t.Fatalf("\nTest %d:\n\ngot:\n%v\nwant:\n%v", i, tokErr.Msg, test.err)
		}
	}
}
