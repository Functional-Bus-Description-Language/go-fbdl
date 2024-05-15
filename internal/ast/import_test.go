package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"reflect"
	"testing"
)

func TestBuildSingleImport(t *testing.T) {
	toks, _ := tok.Parse([]byte(`import "some/path"`), "")
	want := Import{
		Name: nil,
		Path: toks[1].(tok.String),
	}
	ctx := context{}
	got, err := buildSingleImport(toks, &ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 2 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}

	toks, _ = tok.Parse([]byte(`import pkg "path"`), "")
	want = Import{
		Name: toks[1].(tok.Ident),
		Path: toks[2].(tok.String),
	}
	ctx = context{}
	got, err = buildSingleImport(toks, &ctx)
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
	toks, _ := tok.Parse([]byte(`import
	"path1"
	pkg "path2"

	"path3"`),
		"",
	)
	want := []Import{
		Import{Path: toks[3].(tok.String)},
		Import{Name: toks[5].(tok.Ident), Path: toks[6].(tok.String)},
		Import{Path: toks[8].(tok.String)},
	}

	ctx := context{}
	got, err := buildMultiImport(toks, &ctx)
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
