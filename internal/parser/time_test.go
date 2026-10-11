package parser

import (
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

func TestApostropheTimeLiteral(t *testing.T) {
	tests := []struct {
		src  string
		want types.Time
	}{
		{"10'ns", types.Time{Ns: 10}},
		{"10'us", types.Time{Ns: 10000}},
		{"10'ms", types.Time{Ns: 10000000}},
		{"10's", types.Time{S: 10}},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			src := []byte(test.src)
			toks, err := token.Parse(src, "time.fbd")
			if err != nil {
				t.Fatal(err)
			}
			tim, err := MakeTime(ast.Time{X: toks[0].(token.Time)}, src, nil)
			if err != nil {
				t.Fatal(err)
			}
			got, err := tim.Eval()
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Fatalf("got %+v, want %+v", got, test.want)
			}
		})
	}
}
