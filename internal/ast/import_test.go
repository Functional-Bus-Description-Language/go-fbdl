package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildSingleImport(t *testing.T) {
	toks, _ := token.Parse([]byte(`import pkg "path"`))
	want := Import{
		Name: toks[1].(token.Ident),
		Path: toks[2].(token.String),
	}
	c := ctx{}
	got, err := buildSingleImport(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 3 {
		t.Fatalf("c.i = %d", c.i)
	}
	if got[0] != want {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}
}

func TestBuildMultiImport(t *testing.T) {
	toks, _ := token.Parse([]byte(`import
	"path1"
	pkg "path2"

	"path3"`),
	)
	want := []Import{
		Import{Path: toks[3].(token.String)},
		Import{Name: toks[5].(token.Ident), Path: toks[6].(token.String)},
		Import{Path: toks[8].(token.String)},
	}

	c := ctx{}
	got, err := buildMultiImport(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 9 {
		t.Fatalf("c.i = %d", c.i)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("i: %d\ngot:\n%+v,\nwant\n%+v", i, got[i], want[i])
		}
	}
}
