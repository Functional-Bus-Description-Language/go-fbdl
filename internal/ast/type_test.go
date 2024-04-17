package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"reflect"
	"testing"
)

func TestBuildTypeSingleLine(t *testing.T) {
	toks, _ := tok.Parse([]byte("type foo_t(W=1) [8]config; width = W"))
	want := Type{
		Name:   toks[1].(tok.Ident),
		Params: []Param{Param{toks[3].(tok.Ident), Int{toks[5].(tok.Int)}}},
		Count:  Int{toks[8].(tok.Int)},
		Type:   toks[10].(tok.Config),
		Body: Body{
			Props: []Property{Property{toks[12].(tok.Width), Ident{toks[14].(tok.Ident)}}},
		},
	}

	c := context{}
	got, err := buildType(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 15 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}

func TestBuildTypeMultiLine(t *testing.T) {
	toks, _ := tok.Parse([]byte(`type foo_t bar(1, N = 2)
	const A = "a"
	init-value = A
	type cfg_t config`),
	)
	want := Type{
		Name: toks[1].(tok.Ident),
		Type: toks[2].(tok.Ident),
		Args: []Arg{
			Arg{nil, Int{toks[4].(tok.Int)}, toks[4].(tok.Int)},
			Arg{toks[6].(tok.Ident), Int{toks[8].(tok.Int)}, toks[8].(tok.Int)},
		},
		Body: Body{
			Consts: []Const{
				Const{Name: toks[13].(tok.Ident), Value: String{toks[15].(tok.String)}},
			},
			Props: []Property{
				Property{
					Name:  toks[17].(tok.InitValue),
					Value: Ident{toks[19].(tok.Ident)},
				},
			},
			Types: []Type{Type{Name: toks[22].(tok.Ident), Type: toks[23].(tok.Config)}},
		},
	}

	c := context{}
	got, err := buildType(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 24 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}
