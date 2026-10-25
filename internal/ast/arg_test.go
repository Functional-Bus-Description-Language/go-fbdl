package ast

import (
	"reflect"
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func TestBuildArgListKeywordNames(t *testing.T) {
	for _, test := range []struct {
		src  string
		want []Arg
	}{
		{
			"(bus = 1, width = bus)",
			[]Arg{
				{
					Name:          token.New[token.Ident](1, 4),
					Value:         Int{token.New[token.Int](7, 8)},
					ValueFirstTok: token.New[token.Int](7, 8),
				},
				{
					Name:          token.New[token.Ident](10, 15),
					Value:         Ident{token.New[token.Ident](18, 21)},
					ValueFirstTok: token.New[token.Bus](18, 21),
				},
			},
		},
		{
			"(width, bus = 1)",
			[]Arg{
				{
					Value:         Ident{token.New[token.Ident](1, 6)},
					ValueFirstTok: token.New[token.Width](1, 6),
				},
				{
					Name:          token.New[token.Ident](8, 11),
					Value:         Int{token.New[token.Int](14, 15)},
					ValueFirstTok: token.New[token.Int](14, 15),
				},
			},
		},
		{
			"(value = 1)",
			[]Arg{
				{
					Name:          token.New[token.Ident](1, 6),
					Value:         Int{token.New[token.Int](9, 10)},
					ValueFirstTok: token.New[token.Int](9, 10),
				},
			},
		},
	} {
		t.Run(test.src, func(t *testing.T) {
			toks, err := token.Parse([]byte(test.src))
			if err != nil {
				t.Fatal(err)
			}
			tokens := tokenStream{toks: toks}
			got, err := buildArgList(&tokens)
			if err != nil {
				t.Fatal(err)
			}
			want := ArgList{
				LParen: token.New[token.LParen](0, 1),
				Args:   test.want,
				RParen: token.New[token.RParen](len(test.src)-1, len(test.src)),
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got argument list %+v, want %+v", got, want)
			}
			if _, ok := tokens.peek(0).(token.Eof); !ok {
				t.Fatalf("unexpected token after argument list: %T", tokens.peek(0))
			}
		})
	}
}
