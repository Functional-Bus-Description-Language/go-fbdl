package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func TestBuildTypeSingleLine(t *testing.T) {
	toks, _ := token.Parse([]byte("type foo_t(W=1) [8]config; width = W"), "")
	want := Type{
		Name:   toks[1].(token.Ident),
		Params: []Param{Param{toks[3].(token.Ident), Int{toks[5].(token.Int)}}},
		Count:  Int{toks[8].(token.Int)},
		Type:   toks[10].(token.Config),
		Body: Body{
			Props: []Prop{Prop{toks[12].(token.Width), Ident{toks[14].(token.Ident)}}},
		},
	}

	ctx := tokenStream{toks: toks}
	got, err := buildType(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 15 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}

func TestBuildTypeMultiLine(t *testing.T) {
	toks, _ := token.Parse([]byte(`type foo_t bar(1, N = 2)
  const A = "a"
  init-value = A
  type cfg_t config`),
		"",
	)
	want := Type{
		Name: toks[1].(token.Ident),
		Type: toks[2].(token.Ident),
		Args: ArgList{
			LParen: toks[3].(token.LParen),
			Args: []Arg{
				Arg{nil, Int{toks[4].(token.Int)}, toks[4].(token.Int)},
				Arg{toks[6].(token.Ident), Int{toks[8].(token.Int)}, toks[8].(token.Int)},
			},
			RParen: toks[9].(token.RParen),
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
			Types: []Type{Type{Name: toks[22].(token.Ident), Type: toks[23].(token.Config)}},
		},
	}

	ctx := tokenStream{toks: toks}
	got, err := buildType(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 24 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}
