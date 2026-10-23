package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func TestBuildSingleConst(t *testing.T) {
	var tests = []struct {
		src     string
		want    []Const
		wantIdx int
	}{
		{
			"const A = 15",
			[]Const{{
				Name:  token.New[token.Ident](6, 7),
				Value: Int{token.New[token.Int](10, 12)},
			}},
			4,
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
		if tokStream.idx != test.wantIdx {
			t.Fatalf(
				"Input:\n%s\n\ntokStream.idx = %d, want %d",
				test.src, tokStream.idx, test.wantIdx,
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
		src     string
		want    []Const
		wantIdx int
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
			21,
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
		if tokStream.idx != test.wantIdx {
			t.Fatalf(
				"Input:\n%s\n\ntokStream.idx = %d, want %d",
				test.src, tokStream.idx, test.wantIdx,
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
