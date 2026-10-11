package parser

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func TestError(t *testing.T) {
	tests := []struct {
		src  string
		idx  int
		want string
	}{
		{"abc\n  name", 3, "test.fbd +2:3\n   |\n 2 |   name\n   |   ^^^^\n"},
		{"abc\n", 2, "test.fbd +2:1\n   |\n 2 | \n   | ^\n"},
		{"", 0, "test.fbd +1:1\n   |\n 1 | \n   | ^\n"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			src := []byte(test.src)
			toks, err := token.Parse(src)
			if err != nil {
				t.Fatal(err)
			}
			file := &File{Path: "test.fbd", Src: src}
			err = Error{token.Error{Msg: "invalid", Toks: []token.Token{toks[test.idx]}}, file}
			want := "error: invalid\n" + test.want
			if got := err.Error(); got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
			var tokErr token.Error
			if !errors.As(err, &tokErr) || tokErr.Msg != "invalid" {
				t.Fatal("token error cannot be unwrapped")
			}
		})
	}
}

func TestPackageErrorSources(t *testing.T) {
	pkg := &Package{Name: "test"}
	for i, src := range []string{"const A = 1", "\nconst A = 2"} {
		file := &File{Path: fmt.Sprintf("%d.fbd", i), Src: []byte(src), Pkg: pkg}
		toks, err := token.Parse(file.Src)
		if err != nil {
			t.Fatal(err)
		}
		astFile, err := ast.Build(toks)
		if err != nil {
			t.Fatal(err)
		}
		consts, err := buildConsts(astFile.Consts, file, file)
		if err != nil {
			t.Fatal(err)
		}
		err = pkg.addConst(consts[0])
		if i == 0 {
			if err != nil {
				t.Fatal(err)
			}
			continue
		}
		if err == nil {
			t.Fatal("expected a redefinition error")
		}
		if _, ok := err.(token.Error); !ok {
			t.Fatalf("expected token.Error from package, got %T", err)
		}
		err = Error{err.(token.Error), file}
		for _, want := range []string{"1.fbd +2:7", "first definition 0.fbd:1:7", "const A = 2"} {
			if !strings.Contains(err.Error(), want) {
				t.Fatalf("missing %q in %q", want, err.Error())
			}
		}
	}
}

func TestExpressionErrorSource(t *testing.T) {
	file := &File{Path: "definition.fbd", Src: []byte("const A = 1 << -1")}
	toks, err := token.Parse(file.Src)
	if err != nil {
		t.Fatal(err)
	}
	astFile, err := ast.Build(toks)
	if err != nil {
		t.Fatal(err)
	}
	consts, err := buildConsts(astFile.Consts, file, file)
	if err != nil {
		t.Fatal(err)
	}
	_, err = consts[0].Value.Eval()
	if _, ok := err.(Error); !ok {
		t.Fatalf("expected parser.Error from evaluation, got %T", err)
	}
	if !strings.Contains(err.Error(), "definition.fbd +1:13") {
		t.Fatalf("expected expression error with defining file, got %v", err)
	}
}

func TestPublicExpressionErrors(t *testing.T) {
	for _, src := range []string{"const A = unknown(1)", "const A = 99999999999999999999999999999", "const A = 1 << -1", "const A = [1 << -1]", "const A = -true", "const A = bool(1 << -1)"} {
		t.Run(src, func(t *testing.T) {
			file := &File{Path: "expression.fbd", Src: []byte(src)}
			toks, err := token.Parse(file.Src)
			if err != nil {
				t.Fatal(err)
			}
			astFile, err := ast.Build(toks)
			if err != nil {
				t.Fatal(err)
			}
			expr, err := MakeExpr(astFile.Consts[0].Value, file, file)
			if err == nil {
				_, err = expr.Eval()
			}
			parserErr, ok := err.(Error)
			if !ok || parserErr.file != file || len(parserErr.err.Toks) == 0 {
				t.Fatalf("expected parser.Error with the defining file and token, got %#v", err)
			}
			if !strings.Contains(err.Error(), "expression.fbd +1:") {
				t.Fatalf("missing defining location in %q", err.Error())
			}
		})
	}
}

func TestPublicLookupErrors(t *testing.T) {
	pkg := &Package{Name: "test"}
	file := &File{Path: "test.fbd", Pkg: pkg}
	typ := &Type{symbol: symbol{file: file, scope: file}}
	inst := &Inst{symbol: symbol{file: file, scope: file}}
	for _, scope := range []Scope{pkg, file, typ, inst} {
		_, constErr := scope.GetConst("missing")
		_, typeErr := scope.GetType("missing")
		_, instErr := scope.GetInst("missing")
		for _, err := range []error{constErr, typeErr, instErr} {
			if _, ok := err.(Error); !ok {
				t.Fatalf("expected parser.Error from %T, got %T", scope, err)
			}
		}
	}
}

func TestParsePackagesErrors(t *testing.T) {
	for _, src := range []string{"main bus\n  cfg config\n    width = 2\n    range = 3", "const A = unknown(1)", "@", "import missing \"missing\"\nmain bus"} {
		t.Run(src, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "bus.fbd")
			if err := os.WriteFile(path, []byte(src), 0600); err != nil {
				t.Fatal(err)
			}
			err := ParsePackages(Packages{"main": {&Package{Name: "main", Path: path}}})
			parserErr, ok := err.(Error)
			if !ok || parserErr.file == nil || parserErr.file.Path != path {
				t.Fatalf("expected parser.Error with parsed file, got %#v", err)
			}
		})
	}
}

func TestReferencedExpressionErrorSource(t *testing.T) {
	pkg := &Package{Name: "test"}
	definition := &File{Path: "definition.fbd", Src: []byte("const A = 1 << -1"), Pkg: pkg}
	toks, err := token.Parse(definition.Src)
	if err != nil {
		t.Fatal(err)
	}
	astFile, err := ast.Build(toks)
	if err != nil {
		t.Fatal(err)
	}
	consts, err := buildConsts(astFile.Consts, definition, definition)
	if err != nil {
		t.Fatal(err)
	}
	if err := pkg.addConst(consts[0]); err != nil {
		t.Fatal(err)
	}

	caller := &File{Path: "caller.fbd", Src: []byte("A"), Pkg: pkg}
	toks, err = token.Parse(caller.Src)
	if err != nil {
		t.Fatal(err)
	}
	expr := MakeDeclaredIdentifier(ast.Ident{Name: toks[0]}, caller, caller)
	_, err = expr.Eval()
	parserErr, ok := err.(Error)
	if !ok || parserErr.file != definition {
		t.Fatalf("expected parser.Error from defining file, got %#v", err)
	}
	if !strings.Contains(err.Error(), "definition.fbd +1:13") {
		t.Fatalf("missing defining location in %q", err.Error())
	}
}
