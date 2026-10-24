package ast

import (
	"reflect"
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func TestBuildParamListKeywordNames(t *testing.T) {
	for _, test := range []struct {
		src  string
		want []Param
	}{
		{
			"(bus, width)",
			[]Param{
				{Name: token.New[token.Ident](1, 4)},
				{Name: token.New[token.Ident](6, 11)},
			},
		},
		{
			"(bus = 1, width = bus)",
			[]Param{
				{Name: token.New[token.Ident](1, 4), Value: Int{token.New[token.Int](7, 8)}},
				{Name: token.New[token.Ident](10, 15), Value: Ident{token.New[token.Ident](18, 21)}},
			},
		},
		{
			"(width, bus = 1)",
			[]Param{
				{Name: token.New[token.Ident](1, 6)},
				{Name: token.New[token.Ident](8, 11), Value: Int{token.New[token.Int](14, 15)}},
			},
		},
	} {
		t.Run(test.src, func(t *testing.T) {
			toks, err := token.Parse([]byte(test.src))
			if err != nil {
				t.Fatalf(
					"Input:\n%s\n\ntoken.Parse: %v",
					test.src, err,
				)
			}
			tokens := tokenStream{toks: toks}
			got, err := buildParamList(&tokens)
			if err != nil {
				t.Fatalf(
					"Input:\n%s\n\nbuildParamList: %v",
					test.src, err,
				)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf(
					"Input:\n%s\n\ngot parameters %+v, want %+v",
					test.src, got, test.want,
				)
			}
		})
	}
}
