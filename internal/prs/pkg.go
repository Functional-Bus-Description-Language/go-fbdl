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
	Name string
	Path string

	filesMutex sync.Mutex
	Files      []*File

	symbolsMutex sync.Mutex
	Symbols      SymbolContainer
}

func (p *Package) AddFile(f *File) {
	p.filesMutex.Lock()
	p.Files = append(p.Files, f)
	p.filesMutex.Unlock()
}

func (p *Package) AddSymbol(s Symbol) error {
	p.symbolsMutex.Lock()
	defer p.symbolsMutex.Unlock()

	if !p.Symbols.Add(s) {
		msg := `redefinition of symbol '%s' in package '%s'
  %s:%d:%d
  %s:%d:%d`
		first, _ := p.Symbols.GetByName(s.Name())
		return fmt.Errorf(
			msg, s.Name(), p.Name,
			first.File().Path, first.Line(), first.Col(),
			s.File().Path, s.Line(), s.Col(),
		)
	}

	return nil
}

func (p *Package) GetConst(name string) (*Const, error) {
	sym, ok := p.Symbols.GetConst(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf("constant '%s' not found in package '%s'", name, p.Name)
}

func (p *Package) GetInst(name string) (*Inst, error) {
	sym, ok := p.Symbols.GetInst(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf("instantiation '%s' not found in package '%s'", name, p.Name)
}

func (p *Package) GetType(name string) (*Type, error) {
	sym, ok := p.Symbols.GetType(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf("type '%s' not found in package '%s'", name, p.Name)
}

func GetName(path_ string) string {
	name := path.Base(path_)

	if name[:3] == "fbd-" {
		name = name[4:]
	}

	return name
}
