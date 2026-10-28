package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func TestBuildInstSingleLine(t *testing.T) {
	for _, test := range []struct {
		src  string
		want Inst
	}{
		{
			"S [5]status; atomic = false; width = 10",
			Inst{
				Name:  token.New[token.Ident](0, 1),
				Count: Int{token.New[token.Int](3, 4)},
				Type:  token.New[token.Status](5, 11),
				Body: Body{
					Props: []Prop{
						{token.New[token.Atomic](13, 19), Bool{token.New[token.Bool](22, 27)}},
						{token.New[token.Width](29, 34), Int{token.New[token.Int](37, 39)}},
					},
				},
			},
		},
		{
			"S width",
			Inst{
				Name: token.New[token.Ident](0, 1),
				Type: token.New[token.Ident](2, 7),
			},
		},
		{
			"S atomic",
			Inst{
				Name: token.New[token.Ident](0, 1),
				Type: token.New[token.Ident](2, 8),
			},
		},
	} {
		t.Run(test.src, func(t *testing.T) {
			toks, err := token.Parse([]byte(test.src))
			if err != nil {
				t.Fatal(err)
			}
			ctx := tokenStream{toks: toks}
			got, err := buildInst(&ctx)
			if err != nil {
				t.Fatal(err)
			}
			if ctx.idx != len(toks)-1 {
				t.Fatalf("ctx.idx = %d, want %d", ctx.idx, len(toks)-1)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, test.want)
			}
		})
	}
}

func TestBuildInstMultiLine(t *testing.T) {
	toks, _ := token.Parse([]byte(`B pkg.Block_t(1, PI = 3.14)
  masters = 2; reset = "Sync"
  const FOO = true
  C config
    range = 8`))
	want := Inst{
		Name: toks[0].(token.Ident),
		Type: toks[1].(token.QualIdent),
		ArgList: ArgList{
			LParen: toks[2].(token.LParen),
			Args: []Arg{
				Arg{nil, Int{toks[3].(token.Int)}, toks[3].(token.Int)},
				Arg{toks[5].(token.Ident), Float{toks[7].(token.Float)}, toks[7].(token.Float)},
			},
			RParen: toks[8].(token.RParen),
		},
		Body: Body{
			Consts: []Const{Const{Name: toks[20].(token.Ident), Value: Bool{toks[22].(token.Bool)}}},
			Insts: []Inst{
				Inst{
					Name: toks[24].(token.Ident),
					Type: toks[25].(token.Config),
					Body: Body{
						Props: []Prop{
							Prop{toks[28].(token.Range), Int{toks[30].(token.Int)}},
						},
					},
				},
			},
			Props: []Prop{
				Prop{toks[11].(token.Masters), Int{toks[13].(token.Int)}},
				Prop{toks[15].(token.Reset), String{toks[17].(token.String)}},
			},
		},
	}

	ctx := tokenStream{toks: toks}
	got, err := buildInst(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}

	if ctx.idx != 31 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}
