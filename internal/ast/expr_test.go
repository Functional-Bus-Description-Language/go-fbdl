package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestIdentExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("id"))
	want := Ident{Name: toks[0]}

	i, got, err := buildExpr(toks, 0)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if i != 1 {
		t.Fatalf("i = %d", i)
	}
	if got != want {
		t.Fatalf("\ngot:  %+v\nwant: %+v", got, want)
	}
}

func TestUnaryExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("-abc"))
	want := UnaryExpr{
		Op: toks[0], X: Ident{Name: toks[1]},
	}

	i, got, err := buildExpr(toks, 0)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if i != 2 {
		t.Fatalf("i = %d", i)
	}
	if got != want {
		t.Fatalf("\ngot:  %+v\nwant: %+v", got, want)
	}

	toks, _ = token.Parse([]byte("+ 10"))
	want = UnaryExpr{
		Op: toks[0], X: Int{Val: toks[1].(token.Int)},
	}

	i, got, err = buildExpr(toks, 0)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if i != 2 {
		t.Fatalf("i = %d", i)
	}
	if got != want {
		t.Fatalf("\ngot:  %+v\nwant: %+v", got, want)
	}
}
