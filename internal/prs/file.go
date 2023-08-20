package prs

import (
	"fmt"
	"strings"
)

type File struct {
	Path    string
	Pkg     *Package
	Symbols SymbolContainer
	Imports map[string]Import
}

func (f *File) AddSymbol(s Symbol) error {
	name := s.Name()

	if !f.Symbols.Add(s) {
		msg := `%d:%d: redefinition of symbol '%s', first definition line %d column %d`
		first, _ := f.Symbols.GetByName(name)
		return fmt.Errorf(msg, s.Line(), s.Col(), name, first.Line(), first.Col())
	}
	s.setFile(f)

	return nil
}

func (f *File) GetSymbol(name string, kind SymbolKind) (Symbol, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := f.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Pkg.GetSymbol(symName, kind)
	}

	sym, ok := f.Symbols.Get(name, kind)
	if ok {
		return sym, nil
	}

	return f.Pkg.GetSymbol(name, kind)
}
