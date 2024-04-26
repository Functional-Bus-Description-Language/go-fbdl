package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"reflect"
	"testing"
)

func TestBuildSingleConst(t *testing.T) {
	toks, _ := tok.Parse([]byte("const A = 15"), "")
	want := Const{
		Name:  toks[1].(tok.Ident),
		Value: Int{toks[3].(tok.Int)},
	}
	c := context{}
	got, err := buildSingleConst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 4 {
		t.Fatalf("c.i = %d", c.i)
	}
	if !reflect.DeepEqual(got[0], want) {
		t.Fatalf("got: %+v, want %+v", got[0], want)
	}
}

func TestBuildMultiConst(t *testing.T) {
	toks, _ := tok.Parse([]byte(`const
	A = 1
	B = 2 # Inline comment
	# Doc comment
	C = 3.14

	D = false`),
		"",
	)
	want := []Const{
		Const{Name: toks[3].(tok.Ident), Value: Int{toks[5].(tok.Int)}},
		Const{Name: toks[7].(tok.Ident), Value: Int{toks[9].(tok.Int)}},
		Const{
			Doc:   Doc{Lines: []tok.Comment{toks[11].(tok.Comment)}},
			Name:  toks[13].(tok.Ident),
			Value: Real{toks[15].(tok.Real)},
		},
		Const{Name: toks[17].(tok.Ident), Value: Bool{toks[19].(tok.Bool)}},
	}
	c := context{}
	got, err := buildMultiConst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 20 {
		t.Fatalf("c.i = %d", c.i)
	}

	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Fatalf("i: %d\ngot:\n%+v,\nwant\n%+v", i, got[i], want[i])
		}
	}
}
