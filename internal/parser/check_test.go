package parser

import (
	"strings"
	"testing"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func TestPropConflict(t *testing.T) {
	for _, declaration := range []string{"cfg config", "type cfg_t config"} {
		for _, names := range [][2]string{{"width", "range"}, {"range", "width"}} {
			t.Run(declaration+"/"+names[0], func(t *testing.T) {
				file := &File{
					Path: "conflict.fbd",
					Src:  []byte(declaration + "\n  " + names[0] + " = 2\n  " + names[1] + " = 3\n"),
				}
				toks, err := token.Parse(file.Src)
				if err != nil {
					t.Fatal(err)
				}
				astFile, err := ast.Build(toks)
				if err != nil {
					t.Fatal(err)
				}
				var props []ast.Prop
				if len(astFile.Types) > 0 {
					props = astFile.Types[0].Body.Props
					_, err = buildType(astFile.Types[0], file)
				} else {
					props = astFile.Insts[0].Body.Props
					_, err = buildInst(astFile.Insts[0], file)
				}
				if err == nil {
					t.Fatal("expected property conflict")
				}
				conflictError, ok := err.(*token.Error)
				if !ok {
					t.Fatalf("expected token.Error from builder, got %T", err)
				}
				err = Error{*conflictError, file}
				want := "error: cannot set '" + names[1] + "' property, because '" + names[0] + "' property is already set\n" +
					"conflict.fbd +3:3\n   |\n 3 |   " + names[1] + " = 3\n   |   ^^^^^\n" +
					"conflict.fbd +2:3\n   |\n 2 |   " + names[0] + " = 2\n   |   ^^^^^\n"
				if got := err.Error(); got != want {
					t.Fatalf("got %q, want %q", got, want)
				}

				first := Prop{Name: names[0], NameTok: props[0].Name}
				second := Prop{Name: names[1], NameTok: props[1].Name}
				conflict := checkPropConflict(second, PropContainer{first})
				if conflict == nil || len(conflict.Toks) != 2 || conflict.Toks[0] != second.NameTok || conflict.Toks[1] != first.NameTok {
					t.Fatalf("expected both property tokens, got %+v", conflict)
				}
				if strings.Contains(conflict.Error(), "line") {
					t.Fatal("conflict detection should not format source locations")
				}
				if checkPropConflict(second, nil) != nil || checkPropConflict(first, PropContainer{first}) != nil {
					t.Fatal("nonconflicting properties should not return an error")
				}
			})
		}
	}
}
