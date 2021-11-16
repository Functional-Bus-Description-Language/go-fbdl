package prs

import (
	"fmt"
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
		return fmt.Errorf(msg, s.LineNumber(), name, first.LineNumber())
	}
	s.SetFile(f)

	return nil
}

func (f *File) GetSymbol(name string) (Symbol, error) {
	sym, ok := f.Symbols.Get(name)
	if ok {
		return sym, nil
	}

	return f.Pkg.GetSymbol(name)
}
