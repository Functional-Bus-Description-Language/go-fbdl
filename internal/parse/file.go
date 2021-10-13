package parse

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
	Symbols map[string]Symbol
	Imports map[string]Import
}

func (f *File) AddSymbol(s Symbol) error {
	name := s.Name()

	if _, ok := f.Symbols[name]; ok {
		msg := `symbol '%s' defined at least twice in file, first occurence line %d, second line %d`
		return fmt.Errorf(msg, name, f.Symbols[name].LineNumber(), s.LineNumber())
	}
	f.Symbols[name] = s
	s.SetFile(f)

	return nil
}

func (f *File) GetSymbol(s string) (Symbol, error) {
	if sym, ok := f.Symbols[s]; ok {
		return sym, nil
	}

	return f.Pkg.GetSymbol(s)
}
