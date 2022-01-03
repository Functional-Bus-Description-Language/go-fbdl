package prs

import (
	"fmt"
	"path"
	"strings"
	"sync"
)

type Packages map[string][]*Package

// GetMatching returns list of pointers to packages which Path contains path suffix.
func (packages Packages) GetMatching(path string) []*Package {
	ret := []*Package{}

	for _, pkgs := range packages {
		for _, pkg := range pkgs {
			if strings.HasSuffix(pkg.Path, path) {
				ret = append(ret, pkg)
			}
		}
	}

	return ret
}

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
		first, _ := p.Symbols.Get(s.Name(), s.Kind())
		return fmt.Errorf(msg, s.Name(), p.Name, first.LineNum(), s.LineNum())
	}

	return nil
}

func (p *Package) GetSymbol(name string, kind SymbolKind) (Symbol, error) {
	sym, ok := p.Symbols.Get(name, kind)
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

	return name
}
