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

func buildParamList(astParams []ast.Param, src []byte, parent Searchable) ([]Param, error) {
	if len(astParams) == 0 {
		return nil, nil
	}

	params := []Param{}
	names := make(map[string]bool)

	for _, ap := range astParams {
		p := Param{}

		name := tok.Text(ap.Name, src)
		if names[name] {
			return nil, fmt.Errorf(
				"%s: redeclaration of '%s' parameter",
				tok.Loc(ap.Name), name,
			)
		}
		names[name] = true
		p.Name = name

		if ap.Value != nil {
			v, err := MakeExpr(ap.Value, src, parent)
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
			return nil, fmt.Errorf(
				"%s: parameters without default value must precede the ones with default value",
				tok.Loc(astParams[i].Name),
			)
		}

		if p.DfltValue != nil {
			withDflt = true
		}
	}

	return params, nil
}
