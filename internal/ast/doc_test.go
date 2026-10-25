package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestDoc(t *testing.T) {
	src := `# Line 1
#Line 2
#
# Line 4`
	toks, _ := token.Parse([]byte(src))

	ctx := tokenStream{toks: toks}
	doc := buildDoc(&ctx)
	got := doc.Text([]byte(src))
	want := `Line 1
Line 2

Line 4`
	if got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestDocTextShortComments(t *testing.T) {
	for _, test := range []struct {
		src  string
		want string
	}{
		{"#x", "x"},
		{"# x", "x"},
		{"#", ""},
		{"# ", ""},
		{"#x\n#\n# y", "x\n\ny"},
	} {
		t.Run(test.src, func(t *testing.T) {
			src := []byte(test.src)
			toks, err := token.Parse(src)
			if err != nil {
				t.Fatal(err)
			}
			tokens := tokenStream{toks: toks}
			doc := buildDoc(&tokens)
			if got := doc.Text(src); got != test.want {
				t.Fatalf("got doc %q, want %q", got, test.want)
			}
		})
	}
}
