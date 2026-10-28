package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func TestBuildSingleImport(t *testing.T) {
	for _, test := range []struct {
		src  string
		want Import
	}{
		{
			`import "some/path"`,
			Import{Path: token.New[token.String](7, 18)},
		},
		{
			`import pkg "path"`,
			Import{Name: token.New[token.Ident](7, 10), Path: token.New[token.String](11, 17)},
		},
		{
			`import bus "path"`,
			Import{Name: token.New[token.Ident](7, 10), Path: token.New[token.String](11, 17)},
		},
		{
			`import width "path"`,
			Import{Name: token.New[token.Ident](7, 12), Path: token.New[token.String](13, 19)},
		},
	} {
		t.Run(test.src, func(t *testing.T) {
			toks, err := token.Parse([]byte(test.src))
			if err != nil {
				t.Fatal(err)
			}
			ctx := tokenStream{toks: toks}
			got, err := buildImport(&ctx)
			if err != nil {
				t.Fatal(err)
			}
			if ctx.idx != len(toks)-1 {
				t.Fatalf("ctx.idx = %d, want %d", ctx.idx, len(toks)-1)
			}
			if !reflect.DeepEqual(got, []Import{test.want}) {
				t.Fatalf("got: %+v, want %+v", got, []Import{test.want})
			}
		})
	}
}

func TestBuildMultiImport(t *testing.T) {
	toks, err := token.Parse([]byte(`import
  "path1"
  pkg "path2"

  "path3"
  bus "path4"
  width "path5"`))
	if err != nil {
		t.Fatal(err)
	}
	want := []Import{
		Import{Path: toks[3].(token.String)},
		Import{Name: toks[5].(token.Ident), Path: toks[6].(token.String)},
		Import{Path: toks[9].(token.String)},
		Import{Name: token.New[token.Ident](44, 47), Path: toks[12].(token.String)},
		Import{Name: token.New[token.Ident](58, 63), Path: toks[15].(token.String)},
	}

	ctx := tokenStream{toks: toks}
	got, err := buildMultiImport(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != len(toks)-1 {
		t.Fatalf("ctx.idx = %d, want %d", ctx.idx, len(toks)-1)
	}
	if len(got) != len(want) {
		t.Fatalf("got %d imports, want %d", len(got), len(want))
	}

	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Fatalf("i: %d\ngot:\n%+v,\nwant\n%+v", i, got[i], want[i])
		}
	}
}

func TestBuildMultiImportIgnoresComments(t *testing.T) {
	src := []byte(`import
  # Before the first import
  # Another comment
  "path1" # Inline comment

  # Between imports
  pkg "path2"
  # After the last import
const A = 1
# Documentation for B
const B = 2`)
	toks, err := token.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	file, err := Build(toks)
	if err != nil {
		t.Fatal(err)
	}
	if len(file.Imports) != 2 {
		t.Fatalf("got %d imports, want 2", len(file.Imports))
	}
	if got := file.Imports[0].Path.Text(src); got != `"path1"` {
		t.Errorf("first import: got path %q, want %q", got, `"path1"`)
	}
	if file.Imports[0].Name != nil {
		t.Error("first import should have no alias")
	}
	if got := file.Imports[1].Path.Text(src); got != `"path2"` {
		t.Errorf("second import: got path %q, want %q", got, `"path2"`)
	}
	if file.Imports[1].Name == nil || file.Imports[1].Name.Text(src) != "pkg" {
		t.Error("second import should have alias pkg")
	}
	if len(file.Consts) != 2 {
		t.Fatalf("got %d constants, want 2", len(file.Consts))
	}
	if !file.Consts[0].Doc.isEmpty() {
		t.Errorf("A: got doc %q, want no doc", file.Consts[0].Doc.Text(src))
	}
	if got := file.Consts[1].Doc.Text(src); got != "Documentation for B" {
		t.Errorf("B: got doc %q, want %q", got, "Documentation for B")
	}
}

func TestBuildMultiImportIgnoresCommentAtEOF(t *testing.T) {
	src := []byte(`import
  "path"
  # Ignored comment without a final newline`)
	toks, err := token.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	file, err := Build(toks)
	if err != nil {
		t.Fatal(err)
	}
	if len(file.Imports) != 1 {
		t.Fatalf("got %d imports, want 1", len(file.Imports))
	}
	if got := file.Imports[0].Path.Text(src); got != `"path"` {
		t.Errorf("got path %q, want %q", got, `"path"`)
	}
}
