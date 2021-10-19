package prs

import (
	"fmt"
	"path"
	"sync"
)

type Packages map[string][]*Package

type Package struct {
	Name           string
	Path           string
	addFileMutex   sync.Mutex
	Files          []*File
	addSymbolMutex sync.Mutex
	Symbols        map[string]Symbol
}

func (p *Package) AddFile(f *File) {
	p.addFileMutex.Lock()
	p.Files = append(p.Files, f)
	p.addFileMutex.Unlock()
}

func (p *Package) AddSymbol(s Symbol) error {
	p.addSymbolMutex.Lock()
	defer p.addSymbolMutex.Unlock()

	name := s.Name()

	if _, ok := p.Symbols[name]; ok {
		msg := `symbol '%s' defined at least twice in package '%s', first occurence line %d, second line %d`
		return fmt.Errorf(msg, name, p.Name, p.Symbols[name].LineNumber(), s.LineNumber())
	}
	p.Symbols[name] = s

	return nil
}

func (p *Package) GetSymbol(s string) (Symbol, error) {
	if sym, ok := p.Symbols[s]; ok {
		return sym, nil
	}

	return nil, fmt.Errorf("symbol '%s' not found in package '%s'", s, p.Name)
}

func GetName(path_ string) string {
	name := path.Base(path_)

	if name[:3] == "fbd-" {
		name = name[4:]
	}
	fmt.Println("name")

	return name
}
