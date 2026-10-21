package token

import "testing"

func TestSpanText(t *testing.T) {
	for _, test := range []struct {
		name       string
		start, end int
		want       string
	}{
		{"first byte", 0, 1, "a"},
		{"middle byte", 1, 2, "b"},
		{"whole source", 0, 3, "abc"},
		{"empty span", 1, 1, ""},
		{"end of source", 3, 3, ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			tok := New[Ident](test.start, test.end)
			if got := tok.Text([]byte("abc")); got != test.want {
				t.Fatalf("Text() = %q, want %q", got, test.want)
			}
		})
	}
}

func TestParseEofSpan(t *testing.T) {
	for _, src := range []string{"", "a", "a\n"} {
		t.Run(src, func(t *testing.T) {
			toks, err := Parse([]byte(src))
			if err != nil {
				t.Fatal(err)
			}
			eof := toks[len(toks)-1].(Eof)
			if eof.Start() != len(src) || eof.End() != len(src) {
				t.Fatalf("EOF span = %+v, want [%d, %d)", eof, len(src), len(src))
			}
		})
	}
}

func TestParseErrorSpan(t *testing.T) {
	for _, test := range []struct {
		src        string
		start, end int
	}{
		{"?", 0, 1},
		{"a   ", 1, 4},
		{"a   \n", 1, 4},
		{"   ", 0, 3},
		{"0b", 0, 2},
		{"0o", 0, 2},
		{"0x", 0, 2},
		{"1e+", 0, 3},
		{"pkg._sym", 0, 8},
		{"a.b.c", 0, 5},
		{"\"abc", 0, 4},
		{"b\"01", 0, 4},
		{"o\"07", 0, 4},
		{"x\"af", 0, 4},
	} {
		t.Run(test.src, func(t *testing.T) {
			_, err := Parse([]byte(test.src))
			if err == nil {
				t.Fatal("expected a parse error")
			}
			tok := err.(Error).Toks[0]
			if tok.Start() != test.start || tok.End() != test.end {
				t.Fatalf("error span = %+v, want [%d, %d)", tok, test.start, test.end)
			}
		})
	}
}
