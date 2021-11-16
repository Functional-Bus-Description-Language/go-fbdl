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
	Symbols        SymbolContainer
}

func (p *Package) AddFile(f *File) {
	p.addFileMutex.Lock()
	p.Files = append(p.Files, f)
	p.addFileMutex.Unlock()
}

func (p *Package) AddSymbol(s Symbol) error {
	p.addSymbolMutex.Lock()
	defer p.addSymbolMutex.Unlock()

	if !p.Symbols.Add(s) {
		msg := `symbol '%s' defined at least twice in package '%s', first occurence line %d, second line %d`
		first, _ := p.Symbols.Get(s.Name())
		return fmt.Errorf(msg, s.Name(), p.Name, first.LineNumber(), s.LineNumber())
	}

	return nil
}

func (p *Package) GetSymbol(name string) (Symbol, error) {
	sym, ok := p.Symbols.Get(name)
	if ok {
		return sym, nil
	}

	return nil, fmt.Errorf("symbol '%s' not found in package '%s'", name, p.Name)
}

func GetName(path_ string) string {
	name := path.Base(path_)

	if name[:3] == "fbd-" {
		name = name[4:]
	}
	fmt.Println("name")

	return name
}
