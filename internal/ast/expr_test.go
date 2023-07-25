package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func checkExpr(i int, wantI int, got Expr, want Expr, err error) error {
	if err != nil {
		return err
	}

	errMsg := "i = %d, expected i = %d\n\ngot:  %+v\nwant: %+v"
	switch want := want.(type) {
	case CallExpr:
		if i != wantI || !want.eq(got.(CallExpr)) {
			return fmt.Errorf(errMsg, i, wantI, got, want)
		}
	default:
		if i != wantI || got != want {
			return fmt.Errorf(errMsg, i, wantI, got, want)
		}
	}

	return nil
}

func TestBuildIdent(t *testing.T) {
	toks, _ := token.Parse([]byte("id"))
	want := Ident{Name: toks[0]}
	i, got, err := buildExpr(toks, 0, nil)
	err = checkExpr(i, 1, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildUnaryExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("-abc"))
	want := UnaryExpr{
		Op: toks[0], X: Ident{Name: toks[1]},
	}
	i, got, err := buildExpr(toks, 0, nil)
	err = checkExpr(i, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("+ 10"))
	want = UnaryExpr{
		Op: toks[0], X: Int{Val: toks[1].(token.Int)},
	}
	i, got, err = buildExpr(toks, 0, nil)
	err = checkExpr(i, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildParenExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("(a >> b)"))
	want := ParenExpr{
		Lparen: toks[0].(token.LeftParen),
		X: BinaryExpr{
			X:  Ident{Name: toks[1]},
			Op: toks[2].(token.Operator),
			Y:  Ident{Name: toks[3]},
		},
		Rparen: toks[4].(token.RightParen),
	}
	i, got, err := buildExpr(toks, 0, nil)
	err = checkExpr(i, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
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
	i, got, err := buildExpr(toks, 0, nil)
	err = checkExpr(i, 4, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
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
	i, got, err = buildExpr(toks, 0, nil)
	err = checkExpr(i, 6, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildBinaryExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("A + 1"))
	want := BinaryExpr{
		X: Ident{Name: toks[0]}, Op: toks[1].(token.Operator), Y: Int{Val: toks[2].(token.Int)},
	}
	i, got, err := buildExpr(toks, 0, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if i != 3 {
		t.Fatalf("i = %d", i)
	}
	if got != want {
		t.Fatalf("\ngot:  %+v\nwant: %+v", got, want)
	}
	toks, _ = token.Parse([]byte("A + B * C"))
	want = BinaryExpr{
		X:  Ident{Name: toks[0]},
		Op: toks[1].(token.Operator),
		Y: BinaryExpr{
			X:  Ident{Name: toks[2]},
			Op: toks[3].(token.Operator),
			Y:  Ident{Name: toks[4]},
		},
	}
	i, got, err = buildExpr(toks, 0, nil)
	err = checkExpr(i, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("A * B - C"))
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(token.Operator),
			Y:  Ident{Name: toks[2]},
		},
		Op: toks[3].(token.Operator),
		Y:  Ident{Name: toks[4]},
	}
	i, got, err = buildExpr(toks, 0, nil)
	err = checkExpr(i, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("A ** B + C / D"))
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(token.Operator),
			Y:  Ident{Name: toks[2]},
		},
		Op: toks[3].(token.Operator),
		Y: BinaryExpr{
			X:  Ident{Name: toks[4]},
			Op: toks[5].(token.Operator),
			Y:  Ident{Name: toks[6]},
		},
	}
	i, got, err = buildExpr(toks, 0, nil)
	err = checkExpr(i, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("A * (B + C) / D"))
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(token.Operator),
			Y: ParenExpr{
				Lparen: toks[2].(token.LeftParen),
				X: BinaryExpr{
					X:  Ident{Name: toks[3]},
					Op: toks[4].(token.Operator),
					Y:  Ident{Name: toks[5]},
				},
				Rparen: toks[6].(token.RightParen),
			},
		},
		Op: toks[7].(token.Operator),
		Y:  Ident{Name: toks[8]},
	}
	i, got, err = buildExpr(toks, 0, nil)
	err = checkExpr(i, 9, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("A % B == D || false"))
	want = BinaryExpr{
		X: BinaryExpr{
			X: BinaryExpr{
				X:  Ident{Name: toks[0]},
				Op: toks[1].(token.Operator),
				Y:  Ident{Name: toks[2]},
			},
			Op: toks[3].(token.Operator),
			Y:  Ident{Name: toks[4]},
		},
		Op: toks[5].(token.Operator),
		Y:  Bool{Val: toks[6].(token.Bool)},
	}
	i, got, err = buildExpr(toks, 0, nil)
	err = checkExpr(i, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
