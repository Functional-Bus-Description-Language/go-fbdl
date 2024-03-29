package prs

import (
	"fmt"
	"strings"
)

type File struct {
	Path    string
	Pkg     *Package
	Imports map[string]Import
	symbolContainer
}

func (f *File) GetConst(name string) (*Const, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := f.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Pkg.GetConst(symName)
	}

	return f.Pkg.GetConst(name)
}

func (f *File) GetInst(name string) (*Inst, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := f.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Pkg.GetInst(symName)
	}

	return f.Pkg.GetInst(name)
}

func (f *File) GetType(name string) (*Type, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := f.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Pkg.GetType(symName)
	}

	return f.Pkg.GetType(name)
}
