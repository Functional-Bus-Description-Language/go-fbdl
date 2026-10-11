package parser

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
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
			return nil, Error{
				token.Error{Msg: fmt.Sprintf("package '%s' is not imported", pkgName)},
				file,
			}
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
			return nil, Error{
				token.Error{Msg: fmt.Sprintf("package '%s' is not imported", pkgName)},
				file,
			}
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
			return nil, Error{
				token.Error{Msg: fmt.Sprintf("package '%s' is not imported", pkgName)},
				file,
			}
		}

		return pkg.Pkg.GetType(symName)
	}

	return file.Pkg.GetType(name)
}
