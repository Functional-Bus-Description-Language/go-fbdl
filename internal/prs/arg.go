package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Arg struct represents an argument in the argument list.
type Arg struct {
	Name  string
	Value Expr
}

func buildArgList(astArgs []ast.Arg, src []byte, parent Searchable) ([]Arg, error) {
	if len(astArgs) == 0 {
		return nil, nil
	}

	args := []Arg{}
	names := make(map[string]bool)

	for _, aa := range astArgs {
		a := Arg{}

		if aa.Name != nil {
			name := tok.Text(aa.Name, src)
			if names[name] {
				return nil, tok.Error{
					Tok: aa.Name,
					Msg: fmt.Sprintf("reassignment to '%s' argument", name),
				}
			}
			names[name] = true
			a.Name = name
		}

		v, err := MakeExpr(aa.Value, src, parent)
		if err != nil {
			return nil, err
		}
		a.Value = v

		args = append(args, a)
	}

	// Check whether arguments without name precede arguments with name.
	withName := false
	for i, a := range args {
		if withName && a.Name == "" {
			return nil, tok.Error{
				Tok: astArgs[i].ValueFirstTok,
				Msg: "positional argument follows keyword argument",
			}
		}

		if a.Name != "" {
			withName = true
		}
	}

	return args, nil
}
