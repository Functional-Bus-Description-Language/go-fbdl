package parser

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
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

func (pkg *Package) AddFile(f *File) {
	pkg.filesMutex.Lock()
	pkg.Files = append(pkg.Files, f)
	pkg.filesMutex.Unlock()
}

func (pkg *Package) addConst(c *Const) error {
	pkg.symbolsMutex.Lock()
	defer pkg.symbolsMutex.Unlock()

	if !pkg.symbolContainer.addConst(c) {
		msg := fmt.Sprintf(
			"redefinition of constant '%s' in package '%s'", c.name, pkg.Name,
		)
		first, _ := pkg.symbolContainer.GetConst(c.name)
		msg += fmt.Sprintf("; first definition %s:%d:%d", first.File().Path, first.Line(), first.Col())
		return token.Error{Msg: msg, Toks: []token.Token{c.token}}
	}

	return nil
}

func (pkg *Package) addInst(i *Inst) error {
	pkg.symbolsMutex.Lock()
	defer pkg.symbolsMutex.Unlock()

	if !pkg.symbolContainer.addInst(i) {
		msg := `reinstantiation of '%s' in package '%s'
  %s:%d:%d
  %s:%d:%d`
		first, _ := pkg.symbolContainer.GetInst(i.name)
		return token.Error{Msg: fmt.Sprintf(
			msg, i.name, pkg.Name,
			first.File().Path, first.Line(), first.Col(),
			i.File().Path, i.Line(), i.Col(),
		)}
	}

	return nil
}

func (pkg *Package) addType(t *Type) error {
	pkg.symbolsMutex.Lock()
	defer pkg.symbolsMutex.Unlock()

	if !pkg.symbolContainer.addType(t) {
		msg := fmt.Sprintf(
			"redefinition of type '%s' in package '%s'", t.name, pkg.Name,
		)
		first, _ := pkg.symbolContainer.GetType(t.name)
		msg += fmt.Sprintf("; first definition %s:%d:%d", first.File().Path, first.Line(), first.Col())
		return token.Error{Msg: msg, Toks: []token.Token{t.token}}
	}

	return nil
}

func (pkg *Package) GetConst(name string) (*Const, error) {
	sym, ok := pkg.symbolContainer.GetConst(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf(
		"constant '%s' not found in package '%s'",
		name, pkg.Name,
	)
}

func (pkg *Package) GetInst(name string) (*Inst, error) {
	sym, ok := pkg.symbolContainer.GetInst(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf(
		"instantiation '%s' not found in package '%s'",
		name, pkg.Name,
	)
}

func (pkg *Package) GetType(name string) (*Type, error) {
	sym, ok := pkg.symbolContainer.GetType(name)
	if ok {
		return sym, nil
	}
	return nil, fmt.Errorf(
		"type '%s' not found in package '%s'",
		name, pkg.Name,
	)
}
