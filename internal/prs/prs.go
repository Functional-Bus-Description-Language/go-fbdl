// Package prs implements parser based on the tree-sitter parser.
package prs

import (
	_ "fmt"
	"log"
	"os"
	"path"
	_ "strings"
	"sync"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

func ParsePackages(packages Packages) {
	var wg sync.WaitGroup

	for name := range packages {
		for i := range packages[name] {
			wg.Add(1)
			go parsePackage(packages[name][i], &wg)
		}
	}

	wg.Wait()

	bindImports(packages)
}

func parsePackage(pkg *Package, wg *sync.WaitGroup) {
	defer wg.Done()

	var filesWG sync.WaitGroup
	defer filesWG.Wait()

	if pkg.Name == "main" {
		filesWG.Add(1)
		parseFile(pkg.Path, pkg, &filesWG)
		checkInstantiations(pkg)
		return
	}

	pkgDirContent, err := os.ReadDir(pkg.Path)
	if err != nil {
		panic(err)
	}

	for _, file := range pkgDirContent {
		if file.IsDir() {
			continue
		}
		if file.Name()[len(file.Name())-4:] != ".fbd" {
			continue
		}

		filesWG.Add(1)
		parseFile(path.Join(pkg.Path, file.Name()), pkg, &filesWG)
	}

	checkInstantiations(pkg)
}

func checkInstantiations(pkg *Package) {
	for _, f := range pkg.Files {
		for _, symbol := range f.Symbols {
			if e, ok := symbol.(*Inst); ok {
				if e.typ != "bus" && util.IsBaseType(e.typ) {
					log.Fatalf(
						"%s: line %d: element of type '%s' cannot be instantiated at package level",
						f.Path, e.LineNum(), e.Type(),
					)
				} else if e.typ == "bus" {
					if pkg.Name != "main" {
						log.Fatalf(
							"%s: line %d: bus instantiation must be placed within 'main' package",
							f.Path, e.LineNum(),
						)
					}
				}
			}
		}
	}
}

func parseFile(path string, pkg *Package, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	src, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot read %s: %v", path, err)
	}

	astFile, err := ast.Build(src)
	if err != nil {
		log.Fatalf("%s:%v", err)
	}

	file := File{
		Path:    path,
		Pkg:     pkg,
		Symbols: SymbolContainer{},
		Imports: make(map[string]Import),
	}

	consts, err := buildConsts(astFile.Consts, src)
	if err != nil {
		log.Fatalf("%s:%v", path, err)
	}

	for _, c := range consts {
		err = file.AddSymbol(c)
		if err != nil {
			log.Fatalf("%s:%v", path, err)
		}
		err = pkg.AddSymbol(c)
		if err != nil {
			log.Fatalf("%s:%v", path, err)
		}
	}

	pkg.AddFile(&file)
}
