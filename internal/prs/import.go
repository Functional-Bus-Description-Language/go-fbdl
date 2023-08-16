package prs

import (
	"fmt"
	"log"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

type Import struct {
	LineNum int
	Name    string
	Path    string
	Pkg     *Package
}

// buildImports builds list of file Imports based on the list of ast.Import.
func buildImports(astImports []ast.Import, src []byte) ([]Import, error) {
	imports := []Import{}
	name := ""
	path := ""

	for _, ai := range astImports {
		if ai.Name == nil {
			path = tok.Text(ai.Name, src)
			path = path[1 : len(path)-1]
			// TODO: Should it be [0] or the last element?
			name = strings.Split(path, "/")[0]
			if len(name) > 4 && name[0:3] == "fbd-" {
				name = name[4:]
			}
		} else {
			name = tok.Text(ai.Name, src)
			path = tok.Text(ai.Path, src)
			path = path[1 : len(path)-1]
		}
		imports = append(
			imports, Import{ai.Path.Line(), name, path, nil},
		)
	}

	return imports, nil
}

func bindImports(packages Packages) {
	for pkgName, pkgs := range packages {
		for _, pkg := range pkgs {
			err := bindPkgImports(pkg, packages)
			if err != nil {
				log.Fatalf("package %s: %v", pkgName, err)
			}
		}
	}
}

func bindPkgImports(pkg *Package, packages Packages) error {
	for _, file := range pkg.Files {
		err := bindFileImports(file, packages)
		if err != nil {
			return fmt.Errorf("file %s: %v", file.Path, err)
		}
	}

	return nil
}

func bindFileImports(file *File, packages Packages) error {
	for importName := range file.Imports {
		import_ := file.Imports[importName]
		matches := packages.GetMatching(import_.Path)
		if len(matches) == 0 {
			return fmt.Errorf("cannot find package %q", import_.Path)
		} else if len(matches) == 1 {
			import_.Pkg = matches[0]
			file.Imports[importName] = import_
		} else {
			return fmt.Errorf("%d packages match path %q", len(matches), import_.Path)
		}
	}

	return nil
}
