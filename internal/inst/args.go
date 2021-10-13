package inst

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parse"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/value"
	_ "log"
)

func resolveArgumentLists(packages parse.Packages) error {
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

func resolveArgumentListsInSymbols(symbols map[string]parse.Symbol) error {
	for name, s := range symbols {
		e, ok := s.(parse.Element)
		if !ok {
			continue
		}

		if util.IsBaseType(e.Type()) {
			continue
		}

		resolvedArgs, err := resolveArguments(e)
		if err != nil {
			return fmt.Errorf("cannot resolve argument list for symbol '%s': %v", name, err)
		}

		e.SetResolvedArgs(resolvedArgs)
	}

	return nil
}

func resolveArguments(symbol parse.Element) (map[string]value.Value, error) {
	var err error
	args := symbol.Args()
	resolvedArgs := make(map[string]value.Value)
	inPositionalArgs := true

	type_symbol, err := symbol.GetSymbol(symbol.Type())
	if err != nil {
		return nil, fmt.Errorf("cannot get symbol '%s'", symbol.Type())
	}

	params := type_symbol.(parse.Element).Params()

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
					v, err := args[i].Value.Value()
					if err != nil {
						return nil, fmt.Errorf("resolve arguments: %v", err)
					}
					resolvedArgs[p.Name] = v
				} else {
					found := false
					for _, ar := range args {
						if ar.Name == p.Name {
							v, err := ar.Value.Value()
							if err != nil {
								return nil, fmt.Errorf("resolve arguments: %v", err)
							}
							resolvedArgs[p.Name] = v
							found = true
						}
					}
					if !found {
						v, err := p.DefaultValue.Value()
						if err != nil {
							return nil, fmt.Errorf("resolve arguments: %v", err)
						}
						resolvedArgs[p.Name] = v
					}
				}
			} else {
				var v value.Value
				if i < len(args) {
					v, err = args[i].Value.Value()
				} else {
					v, err = p.DefaultValue.Value()
				}
				if err != nil {
					return nil, fmt.Errorf("resolve arguments: %v", err)
				}
				resolvedArgs[p.Name] = v
			}
		} else {
			found := false
			for _, ar := range args {
				if ar.Name == p.Name {
					v, err := ar.Value.Value()
					if err != nil {
						return nil, fmt.Errorf("resolve arguments: %v", err)
					}
					resolvedArgs[p.Name] = v
					found = true
				}
			}
			if !found {
				v, err := p.DefaultValue.Value()
				if err != nil {
					return nil, fmt.Errorf("resolve arguments: %v", err)
				}
				resolvedArgs[p.Name] = v
			}
		}
	}

	return resolvedArgs, nil
}
