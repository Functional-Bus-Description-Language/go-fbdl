package parse

import (
	"fmt"
	"path"
	_ "strings"
	"sync"
)

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

	return nil
}

type Package struct {
	Name         string
	Path         string
	addFileMutex sync.Mutex
	Files        []File
	addSymbolMutex sync.Mutex
	Symbols      map[string]Symbol
}

func (p *Package) AddFile(f File) {
	p.addFileMutex.Lock()
	p.Files = append(p.Files, f)
	p.addFileMutex.Unlock()
}

func (p *Package) addSymbol(s Symbol) error {
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

type Packages map[string][]*Package

func GetName(path_ string) string {
	name := path.Base(path_)

	if name[:3] == "fbd-" {
		name = name[4:]
	}
	fmt.Println("name")

	return name
}
