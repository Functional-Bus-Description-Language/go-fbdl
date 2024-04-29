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

type ArgList struct {
	LeftParen  tok.LeftParen
	Args       []Arg
	RightParen tok.RightParen
}

func (al ArgList) Len() int {
	return len(al.Args)
}

func buildArgList(astArgList ast.ArgList, src []byte, scope Scope) (ArgList, error) {
	if astArgList.Len() == 0 {
		return ArgList{}, nil
	}

	argList := ArgList{
		LeftParen:  astArgList.LeftParen,
		Args:       make([]Arg, 0, astArgList.Len()),
		RightParen: astArgList.RightParen,
	}
	names := make(map[string]bool)

	for _, aal := range astArgList.Args {
		arg := Arg{}

		if aal.Name != nil {
			name := tok.Text(aal.Name, src)
			if names[name] {
				return argList, tok.Error{
					Msg:  fmt.Sprintf("reassignment to '%s' argument", name),
					Toks: []tok.Token{aal.Name},
				}
			}
			names[name] = true
			arg.Name = name
		}

		val, err := MakeExpr(aal.Value, src, scope)
		if err != nil {
			return argList, err
		}
		arg.Value = val

		argList.Args = append(argList.Args, arg)
	}

	// Check whether arguments without name precede arguments with name.
	withName := false
	for i, arg := range argList.Args {
		if withName && arg.Name == "" {
			return argList, tok.Error{
				Msg:  "positional argument follows keyword argument",
				Toks: []tok.Token{astArgList.Args[i].ValueFirstTok},
			}
		}

		if arg.Name != "" {
			withName = true
		}
	}

	return argList, nil
}
