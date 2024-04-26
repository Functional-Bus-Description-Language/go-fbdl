package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"testing"
)

func TestDoc(t *testing.T) {
	src := `# Line 1
#Line 2
#
# Line 4`
	toks, _ := tok.Parse([]byte(src), "")

	c := context{}
	doc := buildDoc(toks, &c)
	got := doc.Text([]byte(src))
	want := `Line 1
Line 2

Line 4`
	if got != want {
		t.Fatalf("\ngot:\n%s\nwant:\n%s", got, want)
	}
}
