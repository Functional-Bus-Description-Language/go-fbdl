package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"testing"
)

func TestBuildTypeSingleLine(t *testing.T) {
	var tests = []struct {
		src  string
		want Type
	}{
		{
			"type foo_t(W=1) [8]config; width = W",
			Type{
				Name:   token.New[token.Ident](5, 10),
				Params: []Param{{token.New[token.Ident](11, 12), Int{token.New[token.Int](13, 14)}}},
				Count:  Int{token.New[token.Int](17, 18)},
				Type:   token.New[token.Config](19, 25),
				Body: Body{
					Props: []Prop{{token.New[token.Width](27, 32), Ident{token.New[token.Ident](35, 36)}}},
				},
			},
		},
	}

	for _, test := range tests {
		toks, err := token.Parse([]byte(test.src))
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\ntoken.Parse: %v",
				test.src, err,
			)
		}
		tokStream := tokenStream{toks: toks}
		got, err := buildType(&tokStream)
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\nbuildType: %v",
				test.src, err,
			)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf(
				"Input:\n%s\n\ngot: %+v\nwant: %+v",
				test.src, got, test.want,
			)
		}
	}
}

func TestBuildTypeMultiLine(t *testing.T) {
	var tests = []struct {
		src  string
		want Type
	}{
		{
			`type foo_t bar(1, N = 2)
  const A = "a"
  init-value = A
  type cfg_t config`,
			Type{
				Name: token.New[token.Ident](5, 10),
				Type: token.New[token.Ident](11, 14),
				Args: ArgList{
					LParen: token.New[token.LParen](14, 15),
					Args: []Arg{
						{nil, Int{token.New[token.Int](15, 16)}, token.New[token.Int](15, 16)},
						{token.New[token.Ident](18, 19), Int{token.New[token.Int](22, 23)}, token.New[token.Int](22, 23)},
					},
					RParen: token.New[token.RParen](23, 24),
				},
				Body: Body{
					Consts: []Const{
						{Name: token.New[token.Ident](33, 34), Value: String{token.New[token.String](37, 40)}},
					},
					Props: []Prop{
						{
							Name:  token.New[token.InitValue](43, 53),
							Value: Ident{token.New[token.Ident](56, 57)},
						},
					},
					Types: []Type{{Name: token.New[token.Ident](65, 70), Type: token.New[token.Config](71, 77)}},
				},
			},
		},
	}

	for _, test := range tests {
		toks, err := token.Parse([]byte(test.src))
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\ntoken.Parse: %v",
				test.src, err,
			)
		}
		tokStream := tokenStream{toks: toks}
		got, err := buildType(&tokStream)
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\nbuildType: %v",
				test.src, err,
			)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf(
				"Input:\n%s\n\ngot: %+v\nwant: %+v",
				test.src, got, test.want,
			)
		}
	}
}
