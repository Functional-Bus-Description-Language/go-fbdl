package prs

import (
	"fmt"
	"strings"
)

type Import struct {
	Path       string
	ImportName string
	Package    *Package
}

type File struct {
	Path    string
	Pkg     *Package
	Symbols SymbolContainer
	Imports map[string]Import
}

func (f *File) AddSymbol(s Symbol) error {
	name := s.Name()

	if !f.Symbols.Add(s) {
		msg := `line %d: symbol '%s' defined at least twice in file, first occurence line %d`
		first, _ := f.Symbols.Get(name)
		return fmt.Errorf(msg, s.LineNum(), name, first.LineNum())
	}
	s.SetFile(f)

	return nil
}

func (f *File) GetSymbol(name string) (Symbol, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := f.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Package.GetSymbol(symName)
	}

	sym, ok := f.Symbols.Get(name)
	if ok {
		return sym, nil
	}

	return f.Pkg.GetSymbol(name)
}
