package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildIdent(t *testing.T) {
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

func TestBuildUnaryExpr(t *testing.T) {
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

func TestBuildCallExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("floor(v)"))
	want := CallExpr{
		Name:   toks[0].(token.Ident),
		Lparen: toks[1].(token.LeftParen),
		Args: []Expr{
			Ident{Name: toks[2].(token.Ident)},
		},
		Rparen: toks[3].(token.RightParen),
	}

	i, got, err := buildExpr(toks, 0)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if i != 4 {
		t.Fatalf("i = %d", i)
	}
	if !want.eq(got.(CallExpr)) {
		t.Fatalf("\ngot:  %+v\nwant: %+v", got, want)
	}

	toks, _ = token.Parse([]byte("foo(12.35, true)"))
	want = CallExpr{
		Name:   toks[0].(token.Ident),
		Lparen: toks[1].(token.LeftParen),
		Args: []Expr{
			Real{Val: toks[2].(token.Real)},
			Bool{Val: toks[4].(token.Bool)},
		},
		Rparen: toks[5].(token.RightParen),
	}

	i, got, err = buildExpr(toks, 0)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if i != 6 {
		t.Fatalf("i = %d", i)
	}
	if !want.eq(got.(CallExpr)) {
		t.Fatalf("\ngot:  %+v\nwant: %+v", got, want)
	}
}
