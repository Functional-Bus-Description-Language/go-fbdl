package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Import struct {
	Line int
	Name string
	Path string
	Pkg  *Package
}

// buildImports builds list of file Imports based on the list of ast.Import.
func buildImports(astImports []ast.Import, src []byte) []Import {
	imports := []Import{}
	name := ""
	path := ""

	for _, ai := range astImports {
		if ai.Name == nil {
			path = token.Text(ai.Path, src)
			path = path[1 : len(path)-1]
			// TODO: Should it be [0] or the last element?
			name = strings.Split(path, string(os.PathSeparator))[0]
			if len(name) > 4 && name[0:3] == "fbd-" {
				name = name[4:]
			}
		} else {
			name = token.Text(ai.Name, src)
			path = token.Text(ai.Path, src)
			path = path[1 : len(path)-1]
		}
		imports = append(
			imports, Import{lineNumber(ai.Path, src), name, path, nil},
		)
	}

	return imports
}

func bindImports(packages Packages) error {
	for _, pkgs := range packages {
		for _, pkg := range pkgs {
			err := bindPkgImports(pkg, packages)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func bindPkgImports(pkg *Package, packages Packages) error {
	for _, file := range pkg.Files {
		err := bindFileImports(file, packages)
		if err != nil {
			return Error{err.(token.Error), file}
		}
	}

	return nil
}

func bindFileImports(file *File, packages Packages) error {
	for importName := range file.Imports {
		import_ := file.Imports[importName]
		matches := packages.GetMatching(import_.Path)
		if len(matches) == 0 {
			return token.Error{Msg: fmt.Sprintf("cannot find package %q", import_.Path)}
		} else if len(matches) == 1 {
			import_.Pkg = matches[0]
			file.Imports[importName] = import_
		} else {
			return token.Error{Msg: fmt.Sprintf("%d packages match path %q", len(matches), import_.Path)}
		}
	}

	return nil
}
