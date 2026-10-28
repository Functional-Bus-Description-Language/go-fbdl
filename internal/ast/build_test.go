package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildError(t *testing.T) {
	var tests = []struct {
		idx int // Test index, useful for navigation
		src string
		err string
	}{
		{
			0,
			"import *",
			"unexpected '*', expected identifier, string or newline",
		},
		{
			1,
			"import name +",
			"unexpected '+', expected string",
		},
		{
			2,
			"const A = ]",
			"unexpected ']', expected expression",
		},
		{
			3,
			"const A = foo(1 true)",
			"unexpected bool, expected ',' or ')'",
		},
		{
			4,
			"const A = foo(, 1)",
			"unexpected ',', expected expression",
		},
		{
			5,
			"const A = (a + b c",
			"unexpected identifier, expected ')'",
		},
		{
			6,
			"const A = [1, 2, 3 4]",
			"unexpected integer, expected ',' or ']'",
		},
		{
			7,
			"const A = [, 1]",
			"unexpected ',', expected expression",
		},
		{
			8,
			"const\n  A 12",
			"unexpected integer, expected '='",
		},
		{
			9,
			"const\n  2.24 = A",
			"unexpected float, expected identifier",
		},
		{
			10,
			"const\nA = 2",
			"unexpected identifier, expected indent or newline",
		},
		{
			11,
			"C type_t()",
			"empty argument list",
		},
		{
			12,
			"C [3;config",
			"unexpected ';', expected ']'",
		},
		{
			13,
			"type type_t() config",
			"empty parameter list",
		},
		{
			14,
			"type t(,) config",
			"unexpected ',', expected identifier",
		},
		{
			15,
			"type t(a b) config",
			"unexpected identifier, expected '=', ')' or ','",
		},
		{
			16,
			"type t(a = 1;) static",
			"unexpected ';', expected ',' or ')'",
		},
		{
			17,
			"type 1 status",
			"unexpected integer, expected identifier",
		},
		{
			18,
			"type a [1,status",
			"unexpected ',', expected ']'",
		},
		{
			19,
			"import\n1",
			"unexpected integer, expected indent increase",
		},
		{
			20,
			"import\n  abc 1",
			"unexpected integer, expected string",
		},
		{
			21,
			"import\n  1",
			"unexpected integer, expected identifier or string",
		},
		{
			22,
			"const a = 1 const b = 2",
			"unexpected 'const', expected newline or end of file",
		},
		{
			23,
			"const\n  a = 1 b = 2",
			"unexpected identifier, expected newline or end of file",
		},
		{
			24,
			"top bus\n  const a = 1 const b = 2",
			"unexpected 'const', expected newline or end of file",
		},
		{
			25,
			"const a = 1 type T config",
			"unexpected 'type', expected newline or end of file",
		},
		{
			26,
			`import "a" const b = 2`,
			"unexpected 'const', expected newline or end of file",
		},
		{
			27,
			`import "a" import "b"`,
			"unexpected 'import', expected newline or end of file",
		},
		{
			28,
			"import\n  \"a\" \"b\"",
			"unexpected string, expected newline or end of file",
		},
		{
			29,
			"import\n  pkg \"a\" other \"b\"",
			"unexpected identifier, expected newline or end of file",
		},
	}

	for i, test := range tests {
		if i != test.idx {
			t.Fatalf("Invalid test index %d, expected %d", test.idx, i)
		}

		toks, err := token.Parse([]byte(test.src))
		if err != nil {
			t.Fatalf("%d: tokenize source: %v", i, err)
		}

		_, err = Build(toks)
		if err == nil {
			t.Fatalf("%d: err == nil, expected != nil", i)
		}

		tokErr := err.(token.Error)
		if tokErr.Msg != test.err {
			t.Fatalf("\nTest %d:\n\ngot:\n%v\nwant:\n%v", i, tokErr.Msg, test.err)
		}
	}
}

func TestBuildDeclarationLineEnds(t *testing.T) {
	for _, test := range []struct {
		src     string
		consts  int
		imports int
	}{
		{"const a = 1\nconst b = 2", 2, 0},
		{"const a = reset-value\nconst b = init-value\nconst c = bus + width", 3, 0},
		{"const a = enable-init-value\nmain bus\n  reset-value = a", 1, 0},
		{"const a = 1 # Inline comment\nconst b = 2\n", 2, 0},
		{"const\n  a = 1\n  b = 2\nconst c = 3", 3, 0},
		{"import \"a\"\nimport pkg \"b\"", 0, 2},
		{"import\n  \"a\"\n  pkg \"b\"\nconst c = 3", 1, 2},
		{"import\n\n  \"a\"\n  pkg \"b\"\nconst c = 3", 1, 2},
		{"import\n\n\n  \"a\"\n  pkg \"b\"\nconst c = 3", 1, 2},
		{"import\n  \"a\" # Inline comment\n\n  pkg \"b\"", 0, 2},
	} {
		t.Run(test.src, func(t *testing.T) {
			toks, err := token.Parse([]byte(test.src))
			if err != nil {
				t.Fatal(err)
			}
			file, err := Build(toks)
			if err != nil {
				t.Fatal(err)
			}
			if len(file.Consts) != test.consts || len(file.Imports) != test.imports {
				t.Fatalf("got %d constants and %d imports, want %d constants and %d imports",
					len(file.Consts), len(file.Imports), test.consts, test.imports)
			}
		})
	}
}

func TestBuildDocDoesNotLeak(t *testing.T) {
	for _, test := range []struct {
		src       string
		constDocs []string
		instDocs  []string
		typeDocs  []string
	}{
		{
			"# Documentation for a\nconst a = 1\n\nconst b = 2",
			[]string{"Documentation for a", ""}, nil, nil,
		},
		{
			"# Documentation for a\nconst a = 1\nconst b = 2",
			[]string{"Documentation for a", ""}, nil, nil,
		},
		{
			"# Group documentation\nconst\n  a = 1\n  b = 2\nconst c = 3",
			[]string{"Group documentation", "", ""}, nil, nil,
		},
		{
			"# Instance documentation\na config\nb config",
			nil, []string{"Instance documentation", ""}, nil,
		},
		{
			"# Instance documentation\na config\n\n# Documentation for b\nb config",
			nil, []string{"Instance documentation", "Documentation for b"}, nil,
		},
		{
			"a config\n\n\n",
			nil, []string{""}, nil,
		},
		{
			"# Instance documentation\na config\n  width = 1\nconst b = 2",
			[]string{""}, []string{"Instance documentation"}, nil,
		},
		{
			"# Import comment\nimport \"a\"\nconst b = 2",
			[]string{""}, nil, nil,
		},
		{
			"# Type comment\ntype T config\nconst b = 2",
			[]string{""}, nil, []string{"Type comment"},
		},
		{
			"# Type comment\ntype T config; width = 1\ntype U config",
			nil, nil, []string{"Type comment", ""},
		},
		{
			"# First type\ntype T config\n# Second type\ntype U config",
			nil, nil, []string{"First type", "Second type"},
		},
		{
			"# Detached type comment\n\ntype T config",
			nil, nil, []string{""},
		},
		{
			"# Detached comment\n\nconst a = 1",
			[]string{""}, nil, nil,
		},
		{
			"# Documentation for a\nconst a = 1\n# Documentation for b\nconst b = 2",
			[]string{"Documentation for a", "Documentation for b"}, nil, nil,
		},
	} {
		t.Run(test.src, func(t *testing.T) {
			src := []byte(test.src)
			toks, err := token.Parse(src)
			if err != nil {
				t.Fatal(err)
			}
			file, err := Build(toks)
			if err != nil {
				t.Fatal(err)
			}
			if len(file.Consts) != len(test.constDocs) || len(file.Insts) != len(test.instDocs) {
				t.Fatalf("got %d constants and %d instances, want %d constants and %d instances",
					len(file.Consts), len(file.Insts), len(test.constDocs), len(test.instDocs))
			}
			if len(file.Types) != len(test.typeDocs) {
				t.Fatalf("got %d types, want %d", len(file.Types), len(test.typeDocs))
			}
			for i, con := range file.Consts {
				if got := con.Doc.Text(src); got != test.constDocs[i] {
					t.Errorf("constant %s: got doc %q, want %q", con.Name.Text(src), got, test.constDocs[i])
				}
			}
			for i, inst := range file.Insts {
				if got := inst.Doc.Text(src); got != test.instDocs[i] {
					t.Errorf("instance %s: got doc %q, want %q", inst.Name.Text(src), got, test.instDocs[i])
				}
			}
			for i, typ := range file.Types {
				if got := typ.Doc.Text(src); got != test.typeDocs[i] {
					t.Errorf("type %s: got doc %q, want %q", typ.Name.Text(src), got, test.typeDocs[i])
				}
			}
		})
	}
}
