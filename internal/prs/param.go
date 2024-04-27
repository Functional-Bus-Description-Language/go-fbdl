package prs

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Param struct represents parameter in the type definition parameter list,
// not the 'param' functionality.
type Param struct {
	Name      string
	DfltValue Expr
}

func buildParamList(astParams []ast.Parameter, src []byte, scope Scope) ([]Param, error) {
	if len(astParams) == 0 {
		return nil, nil
	}

	params := []Param{}
	names := make(map[string]bool)

	for _, ap := range astParams {
		p := Param{}

		name := tok.Text(ap.Name, src)
		if names[name] {
			return nil, tok.Error{
				Tok: ap.Name,
				Msg: fmt.Sprintf("redeclaration of '%s' parameter", name),
			}
		}
		names[name] = true
		p.Name = name

		if ap.Value != nil {
			v, err := MakeExpr(ap.Value, src, scope)
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
			return nil, tok.Error{
				Tok: astParams[i].Name,
				Msg: "parameters without default value must precede the ones with default value",
			}
		}

		if p.DfltValue != nil {
			withDflt = true
		}
	}

	return params, nil
}
