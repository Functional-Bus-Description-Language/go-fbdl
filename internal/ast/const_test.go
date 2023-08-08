package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildSingleConst(t *testing.T) {
	toks, _ := token.Parse([]byte("const A = 15"))
	want := Const{
		Name:  toks[1].(token.Ident),
		Value: Int{toks[3].(token.Int)},
	}
	c := ctx{}
	got, err := buildSingleConst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 4 {
		t.Fatalf("c.i = %d", c.i)
	}
	if !got[0].eq(want) {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}
}

func TestBuildMultiConst(t *testing.T) {
	toks, _ := token.Parse([]byte(`const
	A = 1
	B = 2 # Inline comment
	# Doc comment
	C = 3.14

	D = false`),
	)
	want := []Const{
		Const{Name: toks[3].(token.Ident), Value: Int{toks[5].(token.Int)}},
		Const{Name: toks[7].(token.Ident), Value: Int{toks[9].(token.Int)}},
		Const{
			Doc:   Doc{Lines: []token.Comment{toks[11].(token.Comment)}},
			Name:  toks[13].(token.Ident),
			Value: Real{toks[15].(token.Real)},
		},
		Const{Name: toks[17].(token.Ident), Value: Bool{toks[19].(token.Bool)}},
	}
	c := ctx{}
	got, err := buildMultiConst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 20 {
		t.Fatalf("c.i = %d", c.i)
	}

	for i := range want {
		if !got[i].eq(want[i]) {
			t.Fatalf("i: %d\ngot:\n%+v,\nwant\n%+v", i, got[i], want[i])
		}
	}
}
