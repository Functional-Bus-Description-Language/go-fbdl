package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func checkExpr(c ctx, i int, got Expr, want Expr, err error) error {
	if err != nil {
		return err
	}

	errMsg := "ctx.i = %d, i = %d\n\ngot:  %+v\nwant: %+v"
	switch want := want.(type) {
	case Call:
		if !want.eq(got.(Call)) {
			return fmt.Errorf(errMsg, c.i, i, got, want)
		}
	default:
		if got != want {
			return fmt.Errorf(errMsg, c.i, i, got, want)
		}
	}

	return nil
}

func TestBuildIdent(t *testing.T) {
	toks, _ := token.Parse([]byte("id"))
	want := Ident{Name: toks[0]}
	c := ctx{}
	got, err := buildExpr(toks, &c, nil)
	err = checkExpr(c, 1, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildUnaryExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("-abc"))
	want := UnaryExpr{
		Op: toks[0], X: Ident{Name: toks[1]},
	}
	c := ctx{}
	got, err := buildExpr(toks, &c, nil)
	err = checkExpr(c, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("+ 10"))
	want = UnaryExpr{
		Op: toks[0], X: Int{toks[1].(token.Int)},
	}
	c.i = 0
	got, err = buildExpr(toks, &c, nil)
	err = checkExpr(c, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildParenExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("(a >> b)"))
	want := ParenExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[1]},
			Op: toks[2].(token.Operator),
			Y:  Ident{Name: toks[3]},
		},
	}
	c := ctx{}
	got, err := buildExpr(toks, &c, nil)
	err = checkExpr(c, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildCall(t *testing.T) {
	toks, _ := token.Parse([]byte("floor(v)"))
	want := Call{
		Name: toks[0].(token.Ident),
		Args: []Expr{
			Ident{Name: toks[2].(token.Ident)},
		},
	}
	c := ctx{}
	got, err := buildExpr(toks, &c, nil)
	err = checkExpr(c, 4, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("foo(12.35, true)"))
	want = Call{
		Name: toks[0].(token.Ident),
		Args: []Expr{
			Real{toks[2].(token.Real)},
			Bool{toks[4].(token.Bool)},
		},
	}
	c.i = 0
	got, err = buildExpr(toks, &c, nil)
	err = checkExpr(c, 6, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildBinaryExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("A + 1"))
	want := BinaryExpr{
		X: Ident{Name: toks[0]}, Op: toks[1].(token.Operator), Y: Int{toks[2].(token.Int)},
	}
	c := ctx{}
	got, err := buildExpr(toks, &c, nil)
	if err != nil {
		t.Fatalf("%v", err)
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
	c.i = 0
	got, err = buildExpr(toks, &c, nil)
	err = checkExpr(c, 5, got, want, err)
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
	c.i = 0
	got, err = buildExpr(toks, &c, nil)
	err = checkExpr(c, 5, got, want, err)
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
	c.i = 0
	got, err = buildExpr(toks, &c, nil)
	err = checkExpr(c, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("A * (B + C) / D"))
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(token.Operator),
			Y: ParenExpr{
				X: BinaryExpr{
					X:  Ident{Name: toks[3]},
					Op: toks[4].(token.Operator),
					Y:  Ident{Name: toks[5]},
				},
			},
		},
		Op: toks[7].(token.Operator),
		Y:  Ident{Name: toks[8]},
	}
	c.i = 0
	got, err = buildExpr(toks, &c, nil)
	err = checkExpr(c, 9, got, want, err)
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
		Y:  Bool{toks[6].(token.Bool)},
	}
	c.i = 0
	got, err = buildExpr(toks, &c, nil)
	err = checkExpr(c, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
