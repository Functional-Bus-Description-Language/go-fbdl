package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

func resolveArgLists(packages prs.Packages) error {
	for _, pkgs := range packages {
		for _, pkg := range pkgs {
			err := resolveArgListsInSymbols(pkg.Symbols())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func resolveArgListsInSymbols(syms []prs.Symbol) error {
	for _, s := range syms {
		f, ok := s.(prs.Functionality)
		if !ok {
			continue
		}

		if !util.IsBaseType(f.Type()) {
			resolvedArgs, err := resolveArgs(f)
			if err != nil {
				return fmt.Errorf(
					"%s:%d:%d: %v",
					s.File().Path, s.Line(), s.Col(), err,
				)
			}

			f.SetResolvedArgs(resolvedArgs)
		}

		if len(f.Symbols()) > 0 {
			return resolveArgListsInSymbols(f.Symbols())
		}
	}

	return nil
}

func resolveArgs(symbol prs.Functionality) (map[string]prs.Expr, error) {
	var err error
	args := symbol.Args()
	resolvedArgs := make(map[string]prs.Expr)
	inPositionalArgs := true

	typeSymbol, err := symbol.GetType(symbol.Type())
	if err != nil {
		return nil, err
	}

	params := typeSymbol.Params()

	var argName string
	var argHasName bool
	for i, p := range params {
		if inPositionalArgs {
			if i < len(args) {
				argHasName = args[i].Name != ""
				argName = args[i].Name
			} else {
				inPositionalArgs = false
				argHasName = false
			}

			if argHasName {
				inPositionalArgs = false
				if argName == p.Name {
					resolvedArgs[p.Name] = args[i].Value
				} else {
					found := false
					for _, ar := range args {
						if ar.Name == p.Name {
							resolvedArgs[p.Name] = ar.Value
							found = true
						}
					}
					if !found {
						resolvedArgs[p.Name] = p.DfltValue
					}
				}
			} else {
				if i < len(args) {
					resolvedArgs[p.Name] = args[i].Value
				} else {
					resolvedArgs[p.Name] = p.DfltValue
				}
			}
		} else {
			found := false
			for _, ar := range args {
				if ar.Name == p.Name {
					resolvedArgs[p.Name] = ar.Value
					found = true
				}
			}
			if !found {
				resolvedArgs[p.Name] = p.DfltValue
			}
		}
	}

	return resolvedArgs, nil
}
