package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func checkExpr(ctx tokenStream, i int, got Expr, want Expr, err error) error {
	if err != nil {
		return err
	}

	errMsg := "context.i = %d, i = %d\n\ngot:  %+v\nwant: %+v"
	switch want := want.(type) {
	case Call:
		if !reflect.DeepEqual(got, want) {
			return fmt.Errorf(errMsg, ctx.idx, i, got, want)
		}
	default:
		if !reflect.DeepEqual(got, want) {
			return fmt.Errorf(errMsg, ctx.idx, i, got, want)
		}
	}

	return nil
}

func TestBuildIdent(t *testing.T) {
	toks, _ := token.Parse([]byte("id"))
	want := Ident{Name: toks[0]}
	ctx := tokenStream{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 1, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildUnaryExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("-abc"))
	want := UnaryExpr{
		Op: toks[0], X: Ident{Name: toks[1]},
	}
	ctx := tokenStream{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("+ 10"))
	want = UnaryExpr{
		Op: toks[0], X: Int{toks[1].(token.Int)},
	}
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildNegation(t *testing.T) {
	tests := []struct {
		src  string
		want Expr
	}{
		{
			src: "!true",
			want: UnaryExpr{
				Op: token.New[token.Neg](0, 1), X: Bool{token.New[token.Bool](1, 5)},
			},
		},
		{
			src: "!!false",
			want: UnaryExpr{
				Op: token.New[token.Neg](0, 1),
				X:  UnaryExpr{Op: token.New[token.Neg](1, 2), X: Bool{token.New[token.Bool](2, 7)}},
			},
		},
		{
			src: "!true == false",
			want: BinaryExpr{
				X:  UnaryExpr{Op: token.New[token.Neg](0, 1), X: Bool{token.New[token.Bool](1, 5)}},
				Op: token.New[token.Eq](6, 8),
				Y:  Bool{token.New[token.Bool](9, 14)},
			},
		},
		{
			src: "true && !false",
			want: BinaryExpr{
				X:  Bool{token.New[token.Bool](0, 4)},
				Op: token.New[token.And](5, 7),
				Y:  UnaryExpr{Op: token.New[token.Neg](8, 9), X: Bool{token.New[token.Bool](9, 14)}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			toks, err := token.Parse([]byte(test.src))
			if err != nil {
				t.Fatal(err)
			}
			ctx := tokenStream{toks: toks}
			got, err := buildExpr(&ctx, nil)
			if err := checkExpr(ctx, len(toks)-1, got, test.want, err); err != nil {
				t.Fatal(err)
			}
			if ctx.idx != len(toks)-1 {
				t.Fatalf("consumed %d tokens, want %d", ctx.idx, len(toks)-1)
			}
		})
	}
}

func TestBuildInvalidNegation(t *testing.T) {
	for _, src := range []string{"const A = !", "const A = true ! false"} {
		t.Run(src, func(t *testing.T) {
			toks, err := token.Parse([]byte(src))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := Build(toks); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}

func TestBuildParenExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("(a >> b)"))
	want := ParenExpr{
		LParen: toks[0].(token.LParen),
		X: BinaryExpr{
			X:  Ident{Name: toks[1]},
			Op: toks[2].(token.Operator),
			Y:  Ident{Name: toks[3]},
		},
		RParen: toks[4].(token.RParen),
	}
	ctx := tokenStream{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
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
	ctx := tokenStream{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 4, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("foo(12.35, true)"))
	want = Call{
		Name: toks[0].(token.Ident),
		Args: []Expr{
			Float{toks[2].(token.Float)},
			Bool{toks[4].(token.Bool)},
		},
	}
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 6, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildBinaryExpr(t *testing.T) {
	toks, _ := token.Parse([]byte("A + 1"))
	want := BinaryExpr{
		X: Ident{Name: toks[0]}, Op: toks[1].(token.Operator), Y: Int{toks[2].(token.Int)},
	}
	ctx := tokenStream{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
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
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
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
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
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
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("A * (B + C) / D"))
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(token.Operator),
			Y: ParenExpr{
				LParen: toks[2].(token.LParen),
				X: BinaryExpr{
					X:  Ident{Name: toks[3]},
					Op: toks[4].(token.Operator),
					Y:  Ident{Name: toks[5]},
				},
				RParen: toks[6].(token.RParen),
			},
		},
		Op: toks[7].(token.Operator),
		Y:  Ident{Name: toks[8]},
	}
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 9, got, want, err)
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
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("0:18"))
	want = BinaryExpr{
		X: Int{X: toks[0].(token.Int)}, Op: toks[1].(token.Colon), Y: Int{toks[2].(token.Int)},
	}
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = token.Parse([]byte("-1:0"))
	want = BinaryExpr{
		X:  UnaryExpr{Op: toks[0], X: Int{X: toks[1].(token.Int)}},
		Op: toks[2].(token.Operator),
		Y:  Int{toks[3].(token.Int)},
	}
	ctx = tokenStream{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
