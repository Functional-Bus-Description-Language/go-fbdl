package fbdl

import (
	fbdl "./ts"
	"bufio"
	"fmt"
	ts "github.com/smacker/go-tree-sitter"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

var parser *ts.Parser

func init() {
	parser = ts.NewParser()
	parser.SetLanguage(fbdl.GetLanguage())
}

func ParsePackages(packages Packages) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for name, _ := range packages {
		for i, _ := range packages[name] {
			wg.Add(1)
			go ParsePackage(packages[name][i], &wg)
		}
	}
}

func ParsePackage(pkg *Package, wg *sync.WaitGroup) {
	defer wg.Done()

	var files_wg sync.WaitGroup
	defer files_wg.Wait()

	if pkg.Name == "main" {
		files_wg.Add(1)
		ParseFile(pkg.Path, pkg, &files_wg)
		return
	}

	pkg_dir_content, err := ioutil.ReadDir(pkg.Path)
	if err != nil {
		panic(err)
	}

	for _, file := range pkg_dir_content {
		if file.IsDir() {
			continue
		}
		if file.Name()[len(file.Name())-4:] != ".fbd" {
			continue
		}

		files_wg.Add(1)
		ParseFile(path.Join(pkg.Path, file.Name()), pkg, &files_wg)
	}
}

func getIndent(line string) (uint32, error) {
	var indent uint32 = 0

	for _, char := range line {
		if char == ' ' {
			return 0, fmt.Errorf("space character ' ' is not allowed in indent")
		} else if char == '\t' {
			indent += 1
		} else {
			break
		}
	}

	return indent, nil
}

func checkIndent(code []string) error {
	var current_indent uint32 = 0

	for i, line := range code {
		if line == "\n" {
			continue
		}

		indent, err := getIndent(line)
		if err != nil {
			return fmt.Errorf("line %d: %v", i, err)
		} else if indent > current_indent+1 {
			return fmt.Errorf("line %d: multi indent detected", i)
		}
		current_indent = indent
	}

	return nil
}

func readFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io_scanner := bufio.NewScanner(f)
	var code []string
	for io_scanner.Scan() {
		code = append(code, io_scanner.Text())
	}

	return code
}

func ParseFile(path string, pkg *Package, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	code_lines := readFile(path)

	err = checkIndent(code_lines)
	if err != nil {
		log.Fatalf("%s: %v", path, err)
	}

	var code_bytes []byte
	for _, line := range code_lines {
		code_bytes = append(code_bytes, []byte(line)...)
		code_bytes = append(code_bytes, byte('\n'))
	}

	tree := parser.Parse(nil, []byte(code_bytes))
	root := tree.RootNode()
	node := root.Child(0)
	if node == nil {
		panic("TODO")
	}

	file := File{Path: path, Pkg: pkg, Symbols: make(map[string]Symbol)}

	var symbol Symbol
	for {
		if node.Type() == "single_constant_definition" {
			symbol, err = parseSingleConstantDefinition(node, code_bytes)
		}

		if err != nil {
			log.Fatalf("%s: %v", path, err)
		}

		err = file.AddSymbol(symbol)
		if err != nil {
			log.Fatalf("%s: %v", path, err)
		}

		node = node.NextSibling()
		if node == nil {
			break
		}
	}

	pkg.AddFile(file)
}

func getLineNumber(node *ts.Node) uint32 {
	return node.StartPoint().Row + 1
}

func parseSingleConstantDefinition(node *ts.Node, code_bytes []byte) (Constant, error) {
	constant := Constant{
		common: common{
			Id:         generateId(),
			lineNumber: getLineNumber(node),
			name:       node.Child(1).Content(code_bytes),
		},
	}

	return constant, nil
}
