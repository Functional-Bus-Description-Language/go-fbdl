package prs

import (
	"fmt"
	"log"
)

type Import struct {
	Path       string
	ImportName string
	Pkg        *Package
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
