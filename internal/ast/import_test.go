package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func TestBuildSingleImport(t *testing.T) {
	toks, _ := token.Parse([]byte(`import "some/path"`), "")
	want := Import{
		Name: nil,
		Path: toks[1].(token.String),
	}
	ctx := context{toks: toks}
	got, err := buildSingleImport(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 2 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}

	toks, _ = token.Parse([]byte(`import pkg "path"`), "")
	want = Import{
		Name: toks[1].(token.Ident),
		Path: toks[2].(token.String),
	}
	ctx = context{toks: toks}
	got, err = buildSingleImport(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 3 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}
}

func TestBuildMultiImport(t *testing.T) {
	toks, _ := token.Parse([]byte(`import
  "path1"
  pkg "path2"

  "path3"`),
		"",
	)
	want := []Import{
		Import{Path: toks[3].(token.String)},
		Import{Name: toks[5].(token.Ident), Path: toks[6].(token.String)},
		Import{Path: toks[8].(token.String)},
	}

	ctx := context{toks: toks}
	got, err := buildMultiImport(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 9 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}

	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Fatalf("i: %d\ngot:\n%+v,\nwant\n%+v", i, got[i], want[i])
		}
	}
}
