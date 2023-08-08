package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildTypeSingleLine(t *testing.T) {
	toks, _ := token.Parse([]byte("type foo_t(W=1) [8]config; width = W"))
	want := Type{
		Name:   toks[1].(token.Ident),
		Params: []Param{Param{toks[3].(token.Ident), Int{toks[5].(token.Int)}}},
		Count:  Int{toks[8].(token.Int)},
		Type:   toks[10].(token.Config),
		Body: Body{
			Props: []Prop{Prop{toks[12].(token.Width), Ident{toks[14].(token.Ident)}}},
		},
	}

	c := ctx{}
	got, err := buildType(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 15 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !got.eq(want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}

func TestBuildTypeMultiLine(t *testing.T) {
	toks, _ := token.Parse([]byte(`type foo_t bar(1, N = 2)
	const A = "a"
	init-value = A`),
	)
	want := Type{
		Name: toks[1].(token.Ident),
		Type: toks[2].(token.Ident),
		Args: []Arg{
			Arg{Value: Int{toks[4].(token.Int)}},
			Arg{toks[6].(token.Ident), Int{toks[8].(token.Int)}},
		},
		Body: Body{
			Consts: []Const{
				Const{Name: toks[13].(token.Ident), Value: String{toks[15].(token.String)}},
			},
			Props: []Prop{
				Prop{
					Name:  toks[17].(token.InitValue),
					Value: Ident{toks[19].(token.Ident)},
				},
			},
		},
	}

	c := ctx{}
	got, err := buildType(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 20 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !got.eq(want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}
