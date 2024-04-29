// Package prs implements parser based on the tree-sitter parser.
package prs

import (
	"log"
	"os"
	"path"
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
		for _, ins := range f.Insts {
			if ins.typ != "bus" && util.IsBaseType(ins.typ) {
				log.Fatalf(
					"%s:%d:%d: functionality '%s' of type %s cannot be instantiated at package level",
					f.Path, ins.Line(), ins.Col(), ins.name, ins.typ,
				)
			} else if ins.typ == "bus" {
				if pkg.Name != "main" {
					log.Fatalf(
						"%s:%d:%d: bus instantiation must be placed within 'main' package",
						f.Path, ins.Line(), ins.Col(),
					)
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

	astFile, err := ast.Build(src, path)
	if err != nil {
		log.Fatalf("%v", err)
	}

	file := File{
		Path:    path,
		Pkg:     pkg,
		Imports: make(map[string]Import),
	}

	// Handle file imports
	imports := buildImports(astFile.Imports, src)
	for _, i := range imports {
		if _, exist := file.Imports[i.Name]; exist {
			log.Fatalf(
				"%s: line %d: at least two packages imported as '%s'",
				path, i.Line, i.Name,
			)
		}
		file.Imports[i.Name] = i
	}

	// Handle file and package constants
	consts, err := buildConsts(astFile.Consts, src, &file)
	if err != nil {
		log.Fatalf("%s", err)
	}
	file.Consts = consts
	for _, c := range consts {
		c.setFile(&file)
		c.setScope(&file)
		err = pkg.addConst(c)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	// Handle type definitions
	types, err := buildTypes(astFile.Types, src)
	if err != nil {
		log.Fatalf("%s", err)
	}
	file.Types = types
	for _, t := range types {
		t.setFile(&file)
		t.setScope(&file)
		err = pkg.addType(t)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	// Handle instantiations
	insts, err := buildInsts(astFile.Insts, src)
	if err != nil {
		log.Fatalf("%s", err)
	}
	file.Insts = insts
	for _, i := range insts {
		i.setFile(&file)
		i.setScope(&file)
		err = pkg.addInst(i)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	pkg.AddFile(&file)
}
