package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

func resolveArgumentLists(packages prs.Packages) error {
	for name, pkgs := range packages {
		for _, pkg := range pkgs {
			err := resolveArgumentListsInSymbols(pkg.Symbols)
			if err != nil {
				return fmt.Errorf("package '%s': %v", name, err)
			}
		}
	}

	return nil
}

func resolveArgumentListsInSymbols(symbols prs.SymbolContainer) error {
	for _, s := range symbols {
		name := s.Name()
		e, ok := s.(prs.Element)
		if !ok {
			continue
		}

		if !util.IsBaseType(e.Type()) {
			resolvedArgs, err := resolveArguments(e)
			if err != nil {
				return fmt.Errorf("cannot resolve argument list for symbol '%s': %v", name, err)
			}

			e.SetResolvedArgs(resolvedArgs)
		}

		if len(e.Symbols()) > 0 {
			return resolveArgumentListsInSymbols(e.Symbols())
		}
	}

	return nil
}

func resolveArguments(symbol prs.Element) (map[string]prs.Expression, error) {
	var err error
	args := symbol.Args()
	resolvedArgs := make(map[string]prs.Expression)
	inPositionalArgs := true

	typeSymbol, err := symbol.GetSymbol(symbol.Type())
	if err != nil {
		return nil, fmt.Errorf("cannot get symbol '%s' for element type: %v", symbol.Type(), err)
	}

	params := typeSymbol.(prs.Element).Params()

	var argName string
	var argHasName bool
	for i, p := range params {
		if inPositionalArgs {
			if i < len(args) {
				argHasName = args[i].HasName
				argName = args[i].Name
			} else {
				inPositionalArgs = false
				argHasName = false
			}

			if argHasName == true {
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
						resolvedArgs[p.Name] = p.DefaultValue
					}
				}
			} else {
				if i < len(args) {
					resolvedArgs[p.Name] = args[i].Value
				} else {
					resolvedArgs[p.Name] = p.DefaultValue
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
				resolvedArgs[p.Name] = p.DefaultValue
			}
		}
	}

	return resolvedArgs, nil
}
