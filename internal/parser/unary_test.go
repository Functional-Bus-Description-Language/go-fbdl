package parser

import (
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

func TestUnaryExpression(t *testing.T) {
	tests := []struct {
		src  string
		want types.Value
		err  string
	}{
		{src: "!true", want: types.Bool(false)},
		{src: "!false", want: types.Bool(true)},
		{src: "!!true", want: types.Bool(true)},
		{src: "!!!false", want: types.Bool(true)},
		{src: "+1", want: types.Int(1)},
		{src: "-1", want: types.Int(-1)},
		{src: "-3 >> 1", want: types.Int(-2)},
		{src: "-+3 >> 1", want: types.Int(-2)},
		{src: "--3 >> 1", want: types.Int(1)},
		{src: "+2 * 3", want: types.Int(6)},
		{src: "-2 ** 2", want: types.Int(4)},
		{src: "1 + 2 * 3", want: types.Int(7)},
		{src: "1 - 2 * 3", want: types.Int(-5)},
		{src: "!1", err: "operand of negation must be of type bool, current type integer"},
		{src: "!1.5", err: "operand of negation must be of type bool, current type float"},
		{src: "!\"text\"", err: "operand of negation must be of type bool, current type string"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			src := []byte("const A = " + test.src)
			toks, err := token.Parse(src)
			if err != nil {
				t.Fatal(err)
			}
			astFile, err := ast.Build(toks)
			if err != nil {
				t.Fatal(err)
			}
			expr, err := MakeExpr(astFile.Consts[0].Value, &File{Src: src}, nil)
			if err != nil {
				t.Fatal(err)
			}
			got, err := expr.Eval()
			if test.err != "" {
				gotErr, ok := err.(Error)
				if !ok || gotErr.err.Msg != test.err {
					t.Fatalf("got error %v, want %q", err, test.err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Fatalf("got %v, want %v", got, test.want)
			}
		})
	}
}
