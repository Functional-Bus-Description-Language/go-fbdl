package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestIdentExpr(t *testing.T) {
	toks, err := token.Parse([]byte("id"))
	if err != nil {
		t.Fatalf("%v", err)
	}
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
	toks, err := token.Parse([]byte("-abc"))
	if err != nil {
		t.Fatalf("%v", err)
	}
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
}
