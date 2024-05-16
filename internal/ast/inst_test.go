package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"reflect"
	"testing"
)

func TestBuildInstSingleLine(t *testing.T) {
	toks, _ := tok.Parse([]byte("S [5]status; atomic = false; width = 10"), "")
	want := Inst{
		Name:  toks[0].(tok.Ident),
		Count: Int{toks[2].(tok.Int)},
		Type:  toks[4].(tok.Status),
		Body: Body{
			Props: []Prop{
				Prop{toks[6].(tok.Atomic), Bool{toks[8].(tok.Bool)}},
				Prop{toks[10].(tok.Width), Int{toks[12].(tok.Int)}},
			},
		},
	}

	ctx := context{toks: toks}
	got, err := buildInst(&ctx)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if ctx.idx != 13 {
		t.Fatalf("ctx.idx = %d", ctx.idx)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}

func TestBuildInstMultiLine(t *testing.T) {
	toks, _ := tok.Parse([]byte(`B pkg.Block_t(1, PI = 3.14)
	masters = 2; reset = "Sync"
	const FOO = true
	C config
		range = 8`),
		"",
	)
	want := Inst{
		Name: toks[0].(tok.Ident),
		Type: toks[1].(tok.QualIdent),
		ArgList: ArgList{
			LParen: toks[2].(tok.LParen),
			Args: []Arg{
				Arg{nil, Int{toks[3].(tok.Int)}, toks[3].(tok.Int)},
				Arg{toks[5].(tok.Ident), Float{toks[7].(tok.Float)}, toks[7].(tok.Float)},
			},
			RParen: toks[8].(tok.RParen),
		},
		Body: Body{
			Consts: []Const{Const{Name: toks[20].(tok.Ident), Value: Bool{toks[22].(tok.Bool)}}},
			Insts: []Inst{
				Inst{
					Name: toks[24].(tok.Ident),
					Type: toks[25].(tok.Config),
					Body: Body{
						Props: []Prop{
							Prop{toks[28].(tok.Range), Int{toks[30].(tok.Int)}},
						},
					},
				},
			},
			Props: []Prop{
				Prop{toks[11].(tok.Masters), Int{toks[13].(tok.Int)}},
				Prop{toks[15].(tok.Reset), String{toks[17].(tok.String)}},
			},
		},
	}

	ctx := context{toks: toks}
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
