package prs

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
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
	symbolContainer
}

func (p *Package) AddFile(f *File) {
	p.filesMutex.Lock()
	p.Files = append(p.Files, f)
	p.filesMutex.Unlock()
}

func (p *Package) addConst(c *Const) error {
	p.symbolsMutex.Lock()
	defer p.symbolsMutex.Unlock()

	if !p.symbolContainer.addConst(c) {
		msg := fmt.Sprintf(
			"redefinition of constant '%s' in package '%s'", c.name, p.Name,
		)
		first, _ := p.symbolContainer.GetConst(c.name)
		return tok.Error{Msg: msg, Toks: []tok.Token{c.tok, first.tok}}
	}

	return nil
}

func (p *Package) addInst(i *Inst) error {
	p.symbolsMutex.Lock()
	defer p.symbolsMutex.Unlock()

	if !p.symbolContainer.addInst(i) {
		msg := `reinstantiation of '%s' in package '%s'
  %s:%d:%d
  %s:%d:%d`
		first, _ := p.symbolContainer.GetConst(i.name)
		return fmt.Errorf(
			msg, i.name, p.Name,
			first.File().Path, first.Line(), first.Col(),
			i.File().Path, i.Line(), i.Col(),
		)
	}

	return nil
}

func (p *Package) addType(t *Type) error {
	p.symbolsMutex.Lock()
	defer p.symbolsMutex.Unlock()

	if !p.symbolContainer.addType(t) {
		msg := fmt.Sprintf(
			"redefinition of type '%s' in package '%s'", t.name, p.Name,
		)
		first, _ := p.symbolContainer.GetType(t.name)
		return tok.Error{Msg: msg, Toks: []tok.Token{t.tok, first.tok}}
	}

	return nil
}

func (p *Package) GetConst(name string) (*Const, error) {
	sym, ok := p.symbolContainer.GetConst(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf("constant '%s' not found in package '%s'", name, p.Name)
}

func (p *Package) GetInst(name string) (*Inst, error) {
	sym, ok := p.symbolContainer.GetInst(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf("instantiation '%s' not found in package '%s'", name, p.Name)
}

func (p *Package) GetType(name string) (*Type, error) {
	sym, ok := p.symbolContainer.GetType(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf("type '%s' not found in package '%s'", name, p.Name)
}
