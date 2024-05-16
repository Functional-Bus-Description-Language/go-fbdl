package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"reflect"
	"testing"
)

func checkExpr(ctx context, i int, got Expr, want Expr, err error) error {
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
	toks, _ := tok.Parse([]byte("id"), "")
	want := Ident{Name: toks[0]}
	ctx := context{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 1, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildUnaryExpr(t *testing.T) {
	toks, _ := tok.Parse([]byte("-abc"), "")
	want := UnaryExpr{
		Op: toks[0], X: Ident{Name: toks[1]},
	}
	ctx := context{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("+ 10"), "")
	want = UnaryExpr{
		Op: toks[0], X: Int{toks[1].(tok.Int)},
	}
	ctx.idx = 0
	ctx.toks = toks
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 2, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildParenExpr(t *testing.T) {
	toks, _ := tok.Parse([]byte("(a >> b)"), "")
	want := ParenExpr{
		LParen: toks[0].(tok.LParen),
		X: BinaryExpr{
			X:  Ident{Name: toks[1]},
			Op: toks[2].(tok.Operator),
			Y:  Ident{Name: toks[3]},
		},
		RParen: toks[4].(tok.RParen),
	}
	ctx := context{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildCall(t *testing.T) {
	toks, _ := tok.Parse([]byte("floor(v)"), "")
	want := Call{
		Name: toks[0].(tok.Ident),
		Args: []Expr{
			Ident{Name: toks[2].(tok.Ident)},
		},
	}
	ctx := context{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 4, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("foo(12.35, true)"), "")
	want = Call{
		Name: toks[0].(tok.Ident),
		Args: []Expr{
			Real{toks[2].(tok.Real)},
			Bool{toks[4].(tok.Bool)},
		},
	}
	ctx.idx = 0
	ctx.toks = toks
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 6, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBuildBinaryExpr(t *testing.T) {
	toks, _ := tok.Parse([]byte("A + 1"), "")
	want := BinaryExpr{
		X: Ident{Name: toks[0]}, Op: toks[1].(tok.Operator), Y: Int{toks[2].(tok.Int)},
	}
	ctx := context{toks: toks}
	got, err := buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("A + B * C"), "")
	want = BinaryExpr{
		X:  Ident{Name: toks[0]},
		Op: toks[1].(tok.Operator),
		Y: BinaryExpr{
			X:  Ident{Name: toks[2]},
			Op: toks[3].(tok.Operator),
			Y:  Ident{Name: toks[4]},
		},
	}
	ctx.idx = 0
	ctx.toks = toks
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("A * B - C"), "")
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(tok.Operator),
			Y:  Ident{Name: toks[2]},
		},
		Op: toks[3].(tok.Operator),
		Y:  Ident{Name: toks[4]},
	}
	ctx.idx = 0
	ctx.toks = toks
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("A ** B + C / D"), "")
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(tok.Operator),
			Y:  Ident{Name: toks[2]},
		},
		Op: toks[3].(tok.Operator),
		Y: BinaryExpr{
			X:  Ident{Name: toks[4]},
			Op: toks[5].(tok.Operator),
			Y:  Ident{Name: toks[6]},
		},
	}
	ctx.idx = 0
	ctx.toks = toks
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("A * (B + C) / D"), "")
	want = BinaryExpr{
		X: BinaryExpr{
			X:  Ident{Name: toks[0]},
			Op: toks[1].(tok.Operator),
			Y: ParenExpr{
				LParen: toks[2].(tok.LParen),
				X: BinaryExpr{
					X:  Ident{Name: toks[3]},
					Op: toks[4].(tok.Operator),
					Y:  Ident{Name: toks[5]},
				},
				RParen: toks[6].(tok.RParen),
			},
		},
		Op: toks[7].(tok.Operator),
		Y:  Ident{Name: toks[8]},
	}
	ctx.idx = 0
	ctx.toks = toks
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 9, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("A % B == D || false"), "")
	want = BinaryExpr{
		X: BinaryExpr{
			X: BinaryExpr{
				X:  Ident{Name: toks[0]},
				Op: toks[1].(tok.Operator),
				Y:  Ident{Name: toks[2]},
			},
			Op: toks[3].(tok.Operator),
			Y:  Ident{Name: toks[4]},
		},
		Op: toks[5].(tok.Operator),
		Y:  Bool{toks[6].(tok.Bool)},
	}
	ctx.idx = 0
	ctx.toks = toks
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 7, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("0:18"), "")
	want = BinaryExpr{
		X: Int{X: toks[0].(tok.Int)}, Op: toks[1].(tok.Colon), Y: Int{toks[2].(tok.Int)},
	}
	ctx = context{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	toks, _ = tok.Parse([]byte("-1:0"), "")
	want = BinaryExpr{
		X:  UnaryExpr{Op: toks[0], X: Int{X: toks[1].(tok.Int)}},
		Op: toks[2].(tok.Operator),
		Y:  Int{toks[3].(tok.Int)},
	}
	ctx = context{toks: toks}
	got, err = buildExpr(&ctx, nil)
	err = checkExpr(ctx, 5, got, want, err)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
