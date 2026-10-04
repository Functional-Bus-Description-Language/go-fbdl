package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func TestBuildSingleConst(t *testing.T) {
	toks, _ := token.Parse([]byte("const A = 15"), "")
	want := Const{
		Name:  toks[1].(token.Ident),
		Value: Int{toks[3].(token.Int)},
	}
	ctx := context{toks: toks}
	got, err := buildSingleConst(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 4 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}
}

func TestBuildMultiConst(t *testing.T) {
	toks, _ := token.Parse([]byte(`const
  A = 1
  B = 2 # Inline comment
  # Doc comment
  C = 3.14

  D = false`),
		"",
	)
	want := []Const{
		Const{Name: toks[3].(token.Ident), Value: Int{toks[5].(token.Int)}},
		Const{Name: toks[7].(token.Ident), Value: Int{toks[9].(token.Int)}},
		Const{
			Doc:   Doc{Lines: []token.Comment{toks[11].(token.Comment)}},
			Name:  toks[13].(token.Ident),
			Value: Float{toks[15].(token.Float)},
		},
		Const{Name: toks[18].(token.Ident), Value: Bool{toks[20].(token.Bool)}},
	}
	ctx := context{toks: toks}
	got, err := buildMultiConst(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 21 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}

	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Fatalf("i: %d\ngot:\n%+v,\nwant\n%+v", i, got[i], want[i])
		}
	}
}

func TestBuildMultiConstDocs(t *testing.T) {
	src := []byte(`const
  # Doc A
  A = 1

  # Doc B
  B = 2

  # Not a Doc

  C = 3
`)
	toks, err := token.Parse(src, "")
	if err != nil {
		t.Fatal(err)
	}
	file, err := Build(toks)
	if err != nil {
		t.Fatal(err)
	}
	if len(file.Consts) != 3 {
		t.Fatalf("got %d constants, want 3", len(file.Consts))
	}
	if got := file.Consts[0].Doc.Text(src); got != "Doc A" {
		t.Errorf("A: got doc %q, want %q", got, "Doc A")
	}
	if got := file.Consts[1].Doc.Text(src); got != "Doc B" {
		t.Errorf("B: got doc %q, want %q", got, "Doc B")
	}
	if !file.Consts[2].Doc.isEmpty() {
		t.Errorf("C: got doc %q, want no doc", file.Consts[2].Doc.Text(src))
	}
}
