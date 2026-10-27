package ast

import (
	"reflect"
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func TestBuildSingleConst(t *testing.T) {
	var tests = []struct {
		src  string
		want []Const
	}{
		{
			`const A = 15`,
			[]Const{{
				Name:  token.New[token.Ident](6, 7),
				Value: Int{token.New[token.Int](10, 12)},
			}},
		},
		{
			`const bus = "ab"`,
			[]Const{{
				Name:  token.New[token.Ident](6, 9),
				Value: String{token.New[token.String](12, 16)},
			}},
		},
		{
			`const width = 1`,
			[]Const{{
				Name:  token.New[token.Ident](6, 11),
				Value: Int{token.New[token.Int](14, 15)},
			}},
		},
	}

	for _, test := range tests {
		toks, err := token.Parse([]byte(test.src))
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\ntoken.Parse: %v",
				test.src, err,
			)
		}
		tokStream := tokenStream{toks: toks}
		tokStream.next()
		got, err := buildSingleConst(&tokStream)
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\nbuildSingleConst: %v",
				test.src, err,
			)
		}
		if len(got) != len(test.want) {
			t.Fatalf(
				"Input:\n%s\n\ngot %d constants, want %d",
				test.src, len(got), len(test.want),
			)
		}
		for j, want := range test.want {
			if !reflect.DeepEqual(got[j], want) {
				t.Fatalf(
					"Input:\n%s\n\nConstant: %d\ngot: %+v\nwant: %+v",
					test.src, j, got[j], want,
				)
			}
		}
	}
}

func TestBuildMultiConst(t *testing.T) {
	var tests = []struct {
		src  string
		want []Const
	}{
		{
			`const
  A = 1
  B = 2 # Inline comment
  # Doc comment
  C = 3.14

  D = false`,
			[]Const{
				{
					Name:  token.New[token.Ident](8, 9),
					Value: Int{token.New[token.Int](12, 13)},
				},
				{
					Name:  token.New[token.Ident](16, 17),
					Value: Int{token.New[token.Int](20, 21)},
				},
				{
					Doc:   Doc{Lines: []token.Comment{token.New[token.Comment](41, 54)}},
					Name:  token.New[token.Ident](57, 58),
					Value: Float{token.New[token.Float](61, 65)},
				},
				{
					Name:  token.New[token.Ident](69, 70),
					Value: Bool{token.New[token.Bool](73, 78)},
				},
			},
		},
		{
			`const
  bus = "ab"
  width = 1`,
			[]Const{
				{
					Name:  token.New[token.Ident](8, 11),
					Value: String{token.New[token.String](14, 18)},
				},
				{
					Name:  token.New[token.Ident](21, 26),
					Value: Int{token.New[token.Int](29, 30)},
				},
			},
		},
	}

	for _, test := range tests {
		toks, err := token.Parse([]byte(test.src))
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\ntoken.Parse: %v",
				test.src, err,
			)
		}
		tokStream := tokenStream{toks: toks}
		tokStream.next()
		got, err := buildMultiConst(&tokStream)
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\nbuildMultiConst: %v",
				test.src, err,
			)
		}
		if len(got) != len(test.want) {
			t.Fatalf(
				"Input:\n%s\n\ngot %d constants, want %d",
				test.src, len(got), len(test.want),
			)
		}
		for j, want := range test.want {
			if !reflect.DeepEqual(got[j], want) {
				t.Fatalf(
					"Input:\n%s\n\nConstant: %d\ngot: %+v\nwant: %+v",
					test.src, j, got[j], want,
				)
			}
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
	toks, err := token.Parse(src)
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

func TestBuildConstGroupDocFallback(t *testing.T) {
	src := []byte(`# Group documentation
const
  # Specific documentation
  A = 1
  # Documentation for B
  B = 2
# Fallback documentation
const
  C = 3
  D = 4
const E = 5

main bus
  # Group documentation
  const
    # Specific documentation
    A = 1
    # Documentation for B
    B = 2
  # Fallback documentation
  const
    C = 3
    D = 4
  const E = 5

type Container block
  # Group documentation
  const
    # Specific documentation
    A = 1
    # Documentation for B
    B = 2
  # Fallback documentation
  const
    C = 3
    D = 4
  const E = 5`)
	toks, err := token.Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	file, err := Build(toks)
	if err != nil {
		t.Fatal(err)
	}
	if len(file.Insts) != 1 || len(file.Types) != 1 {
		t.Fatal("expected one instance and one type")
	}

	wantDocs := []string{
		"Specific documentation", // A keeps its own documentation.
		"Documentation for B",    // B keeps its own documentation too.
		"Fallback documentation", // C uses its group's documentation.
		"",                       // D does not inherit the group documentation.
		"",                       // E is outside the group.
	}
	for _, test := range []struct {
		name   string
		consts []Const
	}{
		{"file", file.Consts},
		{"instance", file.Insts[0].Body.Consts},
		{"type", file.Types[0].Body.Consts},
	} {
		t.Run(test.name, func(t *testing.T) {
			if len(test.consts) != len(wantDocs) {
				t.Fatalf("got %d constants, want %d", len(test.consts), len(wantDocs))
			}
			for i, con := range test.consts {
				if got := con.Doc.Text(src); got != wantDocs[i] {
					t.Errorf("constant %s: got doc %q, want %q", con.Name.Text(src), got, wantDocs[i])
				}
			}
		})
	}
}

func TestBuildConstKeywordNames(t *testing.T) {
	for _, test := range []struct {
		src  string
		want []string
	}{
		{`const bus = "ab"`, []string{"bus"}},
		{`const width = 1`, []string{"width"}},
		{"const\n  bus = \"ab\"\n  width = 1", []string{"bus", "width"}},
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
			var got []string
			for _, con := range file.Consts {
				got = append(got, con.Name.Text([]byte(test.src)))
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("got names %v, want %v", got, test.want)
			}
		})
	}
}
