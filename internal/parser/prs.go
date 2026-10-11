package parser

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ast"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

func ParsePackages(packages Packages) error {
	var wg sync.WaitGroup
	var count int
	for _, pkgs := range packages {
		count += len(pkgs)
	}
	failures := make(chan error, count)
	for _, pkgs := range packages {
		for _, pkg := range pkgs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := parsePackage(pkg); err != nil {
					failures <- err
				}
			}()
		}
	}
	wg.Wait()
	close(failures)
	for err := range failures {
		return err
	}
	return bindImports(packages)
}

func parsePackage(pkg *Package) error {
	if pkg.Name == "main" {
		if err := parseFile(pkg.Path, pkg); err != nil {
			return err
		}
		return checkInstantiations(pkg)
	}

	pkgDirContent, err := os.ReadDir(pkg.Path)
	if err != nil {
		return Error{token.Error{Msg: err.Error()}, nil}
	}

	cwd, err := os.Getwd()
	if err != nil {
		return Error{token.Error{Msg: err.Error()}, nil}
	}
	cwd += string(os.PathSeparator)

	for _, file := range pkgDirContent {
		if file.IsDir() {
			continue
		}
		if file.Name()[len(file.Name())-4:] != ".fbd" {
			continue
		}

		filePath := path.Join(pkg.Path, file.Name())
		filePath = strings.TrimPrefix(filePath, cwd)

		if err := parseFile(filePath, pkg); err != nil {
			return err
		}
	}

	return checkInstantiations(pkg)
}

func checkInstantiations(pkg *Package) error {
	for _, f := range pkg.Files {
		for _, ins := range f.Insts {
			if ins.typ != "bus" && util.IsBaseType(ins.typ) {
				return Error{token.Error{
					Msg:  fmt.Sprintf("functionality '%s' of type %s cannot be instantiated at package level", ins.name, ins.typ),
					Toks: []token.Token{ins.Tok()},
				}, f}
			} else if ins.typ == "bus" {
				if pkg.Name != "main" {
					return Error{token.Error{Msg: "bus instantiation must be placed within 'main' package", Toks: []token.Token{ins.Tok()}}, f}
				}
			}
		}
	}
	return nil
}

func parseFile(path string, pkg *Package) (err error) {
	file := File{Path: path, Pkg: pkg, Imports: make(map[string]Import)}
	defer func() {
		switch e := err.(type) {
		case token.Error:
			err = Error{e, &file}
		case *token.Error:
			err = Error{*e, &file}
		}
	}()

	src, err := os.ReadFile(path)
	if err != nil {
		return Error{token.Error{Msg: fmt.Sprintf("cannot read %s: %v", path, err)}, &file}
	}

	file.Src = src

	toks, err := token.Parse(src)
	if err != nil {
		return err
	}

	astFile, err := ast.Build(toks)
	if err != nil {
		return err
	}

	// Handle file imports
	imports := buildImports(astFile.Imports, src)
	for _, i := range imports {
		if _, exist := file.Imports[i.Name]; exist {
			return token.Error{Msg: fmt.Sprintf("%s: line %d: at least two packages imported as '%s'", path, i.Line, i.Name)}
		}
		file.Imports[i.Name] = i
	}

	// Handle file and package constants
	consts, err := buildConsts(astFile.Consts, &file, &file)
	if err != nil {
		return err
	}
	file.Consts = consts
	for _, c := range consts {
		c.setScope(&file)
		err = pkg.addConst(c)
		if err != nil {
			return err
		}
	}

	// Handle type definitions
	types, err := buildTypes(astFile.Types, &file)
	if err != nil {
		return err
	}
	file.Types = types
	for _, t := range types {
		t.setScope(&file)
		err = pkg.addType(t)
		if err != nil {
			return err
		}
	}

	// Handle instantiations
	insts, err := buildInsts(astFile.Insts, &file)
	if err != nil {
		return err
	}
	file.Insts = insts
	for _, i := range insts {
		i.setScope(&file)
		err = pkg.addInst(i)
		if err != nil {
			return err
		}
	}

	pkg.AddFile(&file)
	return nil
}
