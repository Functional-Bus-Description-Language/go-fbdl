package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"testing"
)

func TestBuildInstSingleLine(t *testing.T) {
	toks, _ := token.Parse([]byte("S [5]status; atomic = false; width = 10"))
	want := Inst{
		Name:  toks[0].(token.Ident),
		Count: Int{toks[2].(token.Int)},
		Type:  toks[4].(token.Status),
		Body: Body{
			Props: []Prop{
				Prop{toks[6].(token.Atomic), Bool{toks[8].(token.Bool)}},
				Prop{toks[10].(token.Width), Int{toks[12].(token.Int)}},
			},
		},
	}

	c := ctx{}
	got, err := buildInst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}
	if c.i != 13 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !got.eq(want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}

func TestBuildInstMultiLine(t *testing.T) {
	toks, _ := token.Parse([]byte(`B pkg.block_t(1, PI = 3.14)
	masters = 2; reset = "Sync"
	const FOO = true
	C config
		range = 8`),
	)
	want := Inst{
		Name: toks[0].(token.Ident),
		Type: toks[1].(token.QualIdent),
		Args: []Arg{
			Arg{Value: Int{toks[3].(token.Int)}},
			Arg{toks[5].(token.Ident), Real{toks[7].(token.Real)}},
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

	c := ctx{}
	got, err := buildInst(toks, &c)
	if err != nil {
		t.Fatalf("err != nil: %v", err)
	}

	if c.i != 31 {
		t.Fatalf("c.i = %d", c.i)
	}

	if !got.eq(want) {
		t.Fatalf("\ngot:\n%+v,\nwant\n%+v", got, want)
	}
}
