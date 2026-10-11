package parser

import (
	"fmt"
	"strings"
)

type File struct {
	Path    string
	Src     []byte
	Pkg     *Package
	Imports map[string]Import
	symbolContainer
}

func (file *File) GetConst(name string) (*Const, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := file.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Pkg.GetConst(symName)
	}

	return file.Pkg.GetConst(name)
}

func (file *File) GetInst(name string) (*Inst, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := file.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Pkg.GetInst(symName)
	}

	return file.Pkg.GetInst(name)
}

func (file *File) GetType(name string) (*Type, error) {
	if strings.Contains(name, ".") {
		parts := strings.Split(name, ".")
		pkgName := parts[0]
		symName := parts[1]

		pkg, ok := file.Imports[pkgName]
		if !ok {
			return nil, fmt.Errorf("package '%s' is not imported", pkgName)
		}

		return pkg.Pkg.GetType(symName)
	}

	return file.Pkg.GetType(name)
}
