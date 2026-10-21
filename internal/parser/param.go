package parser

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Param struct represents parameter in the type definition parameter list,
// not the 'param' functionality.
type Param struct {
	Name      string
	DfltValue Expr
}

func buildParamList(astParams []ast.Param, file *File, scope Scope) ([]Param, error) {
	src := file.Src
	if len(astParams) == 0 {
		return nil, nil
	}

	params := []Param{}
	names := make(map[string]bool)

	for _, ap := range astParams {
		p := Param{}

		name := ap.Name.Text(src)
		if names[name] {
			return nil, token.Error{
				Msg:  fmt.Sprintf("redeclaration of '%s' parameter", name),
				Toks: []token.Token{ap.Name},
			}
		}
		names[name] = true
		p.Name = name

		if ap.Value != nil {
			v, err := MakeExpr(ap.Value, file, scope)
			if err != nil {
				return nil, err
			}
			p.DfltValue = v
		}

		params = append(params, p)
	}

	// Check whether parameters without default value precede parameters with default value.
	withDflt := false
	for i, p := range params {
		if withDflt && p.DfltValue == nil {
			return nil, token.Error{
				Msg:  "parameters without default value must precede the ones with default value",
				Toks: []token.Token{astParams[i].Name},
			}
		}

		if p.DfltValue != nil {
			withDflt = true
		}
	}

	return params, nil
}
