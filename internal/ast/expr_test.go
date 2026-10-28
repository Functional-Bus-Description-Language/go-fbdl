package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"reflect"
	"strings"
	"testing"
)

func TestBuildExpr(t *testing.T) {
	var tests = []struct {
		src  string
		want Expr
	}{
		{
			"id",
			Ident{Name: token.New[token.Ident](0, 2)},
		},
		{
			"-abc",
			UnaryExpr{
				Op: token.New[token.Sub](0, 1), X: Ident{Name: token.New[token.Ident](1, 4)},
			},
		},
		{
			"+ 10",
			UnaryExpr{
				Op: token.New[token.Add](0, 1), X: Int{token.New[token.Int](2, 4)},
			},
		},
		{
			"-1 & 1",
			BinaryExpr{
				X:  UnaryExpr{Op: token.New[token.Sub](0, 1), X: Int{token.New[token.Int](1, 2)}},
				Op: token.New[token.BitAnd](3, 4),
				Y:  Int{token.New[token.Int](5, 6)},
			},
		},
		{
			"+2 * 3",
			BinaryExpr{
				X:  UnaryExpr{Op: token.New[token.Add](0, 1), X: Int{token.New[token.Int](1, 2)}},
				Op: token.New[token.Mul](3, 4),
				Y:  Int{token.New[token.Int](5, 6)},
			},
		},
		{
			"1 & -1",
			BinaryExpr{
				X:  Int{token.New[token.Int](0, 1)},
				Op: token.New[token.BitAnd](2, 3),
				Y:  UnaryExpr{Op: token.New[token.Sub](4, 5), X: Int{token.New[token.Int](5, 6)}},
			},
		},
		{
			"-(1 & 1)",
			UnaryExpr{
				Op: token.New[token.Sub](0, 1),
				X: ParenExpr{
					LParen: token.New[token.LParen](1, 2),
					X: BinaryExpr{
						X:  Int{token.New[token.Int](2, 3)},
						Op: token.New[token.BitAnd](4, 5),
						Y:  Int{token.New[token.Int](6, 7)},
					},
					RParen: token.New[token.RParen](7, 8),
				},
			},
		},
		{
			"!true",
			UnaryExpr{
				Op: token.New[token.Neg](0, 1), X: Bool{token.New[token.Bool](1, 5)},
			},
		},
		{
			"!!false",
			UnaryExpr{
				Op: token.New[token.Neg](0, 1),
				X:  UnaryExpr{Op: token.New[token.Neg](1, 2), X: Bool{token.New[token.Bool](2, 7)}},
			},
		},
		{
			"!true == false",
			BinaryExpr{
				X:  UnaryExpr{Op: token.New[token.Neg](0, 1), X: Bool{token.New[token.Bool](1, 5)}},
				Op: token.New[token.Eq](6, 8),
				Y:  Bool{token.New[token.Bool](9, 14)},
			},
		},
		{
			"true && !false",
			BinaryExpr{
				X:  Bool{token.New[token.Bool](0, 4)},
				Op: token.New[token.And](5, 7),
				Y:  UnaryExpr{Op: token.New[token.Neg](8, 9), X: Bool{token.New[token.Bool](9, 14)}},
			},
		},
		{
			"(a >> b)",
			ParenExpr{
				LParen: token.New[token.LParen](0, 1),
				X: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](1, 2)},
					Op: token.New[token.RShift](3, 5),
					Y:  Ident{Name: token.New[token.Ident](6, 7)},
				},
				RParen: token.New[token.RParen](7, 8),
			},
		},
		{
			"floor(v)",
			Call{
				Name: token.New[token.Ident](0, 5),
				Args: []Expr{
					Ident{Name: token.New[token.Ident](6, 7)},
				},
			},
		},
		{
			"pkg.foo(1)",
			Call{
				Name: token.New[token.QualIdent](0, 7),
				Args: []Expr{
					Int{token.New[token.Int](8, 9)},
				},
			},
		},
		{
			"foo(12.35, true)",
			Call{
				Name: token.New[token.Ident](0, 3),
				Args: []Expr{
					Float{token.New[token.Float](4, 9)},
					Bool{token.New[token.Bool](11, 15)},
				},
			},
		},
		{
			"A + 1",
			BinaryExpr{
				X: Ident{Name: token.New[token.Ident](0, 1)}, Op: token.New[token.Add](2, 3), Y: Int{token.New[token.Int](4, 5)},
			},
		},
		{
			"A + B * C",
			BinaryExpr{
				X:  Ident{Name: token.New[token.Ident](0, 1)},
				Op: token.New[token.Add](2, 3),
				Y: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](4, 5)},
					Op: token.New[token.Mul](6, 7),
					Y:  Ident{Name: token.New[token.Ident](8, 9)},
				},
			},
		},
		{
			"A * B - C",
			BinaryExpr{
				X: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](0, 1)},
					Op: token.New[token.Mul](2, 3),
					Y:  Ident{Name: token.New[token.Ident](4, 5)},
				},
				Op: token.New[token.Sub](6, 7),
				Y:  Ident{Name: token.New[token.Ident](8, 9)},
			},
		},
		{
			"2 ** 3 ** 2",
			BinaryExpr{
				X:  Int{token.New[token.Int](0, 1)},
				Op: token.New[token.Exp](2, 4),
				Y: BinaryExpr{
					X:  Int{token.New[token.Int](5, 6)},
					Op: token.New[token.Exp](7, 9),
					Y:  Int{token.New[token.Int](10, 11)},
				},
			},
		},
		{
			"A ** B + C / D",
			BinaryExpr{
				X: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](0, 1)},
					Op: token.New[token.Exp](2, 4),
					Y:  Ident{Name: token.New[token.Ident](5, 6)},
				},
				Op: token.New[token.Add](7, 8),
				Y: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](9, 10)},
					Op: token.New[token.Div](11, 12),
					Y:  Ident{Name: token.New[token.Ident](13, 14)},
				},
			},
		},
		{
			"A * (B + C) / D",
			BinaryExpr{
				X: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](0, 1)},
					Op: token.New[token.Mul](2, 3),
					Y: ParenExpr{
						LParen: token.New[token.LParen](4, 5),
						X: BinaryExpr{
							X:  Ident{Name: token.New[token.Ident](5, 6)},
							Op: token.New[token.Add](7, 8),
							Y:  Ident{Name: token.New[token.Ident](9, 10)},
						},
						RParen: token.New[token.RParen](10, 11),
					},
				},
				Op: token.New[token.Div](12, 13),
				Y:  Ident{Name: token.New[token.Ident](14, 15)},
			},
		},
		{
			"A % B == D || false",
			BinaryExpr{
				X: BinaryExpr{
					X: BinaryExpr{
						X:  Ident{Name: token.New[token.Ident](0, 1)},
						Op: token.New[token.Rem](2, 3),
						Y:  Ident{Name: token.New[token.Ident](4, 5)},
					},
					Op: token.New[token.Eq](6, 8),
					Y:  Ident{Name: token.New[token.Ident](9, 10)},
				},
				Op: token.New[token.Or](11, 13),
				Y:  Bool{token.New[token.Bool](14, 19)},
			},
		},
		{
			"0:18",
			BinaryExpr{
				X: Int{X: token.New[token.Int](0, 1)}, Op: token.New[token.Colon](1, 2), Y: Int{token.New[token.Int](2, 4)},
			},
		},
		{
			"-1:0",
			BinaryExpr{
				X:  UnaryExpr{Op: token.New[token.Sub](0, 1), X: Int{X: token.New[token.Int](1, 2)}},
				Op: token.New[token.Colon](2, 3),
				Y:  Int{token.New[token.Int](3, 4)},
			},
		},
		{
			"bus + atomic",
			BinaryExpr{
				X:  Ident{Name: token.New[token.Ident](0, 3)},
				Op: token.New[token.Add](4, 5),
				Y:  Ident{Name: token.New[token.Ident](6, 12)},
			},
		},
		{
			"init-value",
			BinaryExpr{
				X:  Ident{Name: token.New[token.Ident](0, 4)},
				Op: token.New[token.Sub](4, 5),
				Y:  Ident{Name: token.New[token.Ident](5, 10)},
			},
		},
		{
			"reset-value(1)",
			BinaryExpr{
				X:  Ident{Name: token.New[token.Ident](0, 5)},
				Op: token.New[token.Sub](5, 6),
				Y: Call{
					Name: token.New[token.Ident](6, 11),
					Args: []Expr{Int{token.New[token.Int](12, 13)}},
				},
			},
		},
		{
			"-init-value",
			BinaryExpr{
				X: UnaryExpr{
					Op: token.New[token.Sub](0, 1),
					X:  Ident{Name: token.New[token.Ident](1, 5)},
				},
				Op: token.New[token.Sub](5, 6),
				Y:  Ident{Name: token.New[token.Ident](6, 11)},
			},
		},
		{
			"init-value + reset-value",
			BinaryExpr{
				X: BinaryExpr{
					X: BinaryExpr{
						X:  Ident{Name: token.New[token.Ident](0, 4)},
						Op: token.New[token.Sub](4, 5),
						Y:  Ident{Name: token.New[token.Ident](5, 10)},
					},
					Op: token.New[token.Add](11, 12),
					Y:  Ident{Name: token.New[token.Ident](13, 18)},
				},
				Op: token.New[token.Sub](18, 19),
				Y:  Ident{Name: token.New[token.Ident](19, 24)},
			},
		},
		{
			"enable-init-value",
			BinaryExpr{
				X: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](0, 6)},
					Op: token.New[token.Sub](6, 7),
					Y:  Ident{Name: token.New[token.Ident](7, 11)},
				},
				Op: token.New[token.Sub](11, 12),
				Y:  Ident{Name: token.New[token.Ident](12, 17)},
			},
		},
		{
			"width - init-value",
			BinaryExpr{
				X: BinaryExpr{
					X:  Ident{Name: token.New[token.Ident](0, 5)},
					Op: token.New[token.Sub](6, 7),
					Y:  Ident{Name: token.New[token.Ident](8, 12)},
				},
				Op: token.New[token.Sub](12, 13),
				Y:  Ident{Name: token.New[token.Ident](13, 18)},
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
		got, err := buildExpr(&tokStream, nil)
		if err != nil {
			t.Fatalf(
				"Input:\n%s\n\nbuildExpr: %v",
				test.src, err,
			)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf(
				"Input:\n%s\n\ngot: %+v\nwant: %+v",
				test.src, got, test.want,
			)
		}
		if _, ok := tokStream.peek(0).(token.Eof); !ok {
			t.Fatalf("unexpected token after expression %q: %T", test.src, tokStream.peek(0))
		}
	}
}

func TestBuildExprError(t *testing.T) {
	for _, src := range []string{"const A = !", "const A = true ! false", "const A = init-value +"} {
		t.Run(src, func(t *testing.T) {
			toks, err := token.Parse([]byte(src))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := Build(toks); err == nil {
				t.Fatal("expected an error")
			}
		})
	}
}

func BenchmarkBuildHyphenProperties(b *testing.B) {
	for _, count := range []int{2000, 4000, 8000} {
		b.Run(fmt.Sprint(count), func(b *testing.B) {
			var src strings.Builder
			for i := range count {
				fmt.Fprintf(&src, "const A%d = reset-value\n", i)
			}
			toks, err := token.Parse([]byte(src.String()))
			if err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for range b.N {
				if _, err := Build(toks); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
