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

// Node is a wrapper for tree-sitter node.
type Node struct {
	n    *ts.Node
	code []byte
}

func (n Node) Content() string {
	return n.n.Content(n.code)
}

func (n Node) Type() string {
	return n.n.Type()
}

func (n Node) Child(idx int) Node {
	tsn := n.n.Child(idx)
	if tsn == nil {
		panic("can't get child")
	}

	return Node{n: tsn, code: n.code}
}

func (n Node) LineNumber() uint32 {
	return n.n.StartPoint().Row + 1
}

func (n Node) HasNextSibling() bool {
	tsn := n.n.NextSibling()

	if tsn == nil {
		return false
	}

	return true
}

func (n Node) NextSibling() Node {
	return Node{n: n.n.NextSibling(), code: n.code}
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
	tsnode := root.Child(0)
	if tsnode == nil {
		panic("TODO")
	}
	node := Node{n: tsnode, code: code_bytes}

	file := File{Path: path, Pkg: pkg, Symbols: make(map[string]Symbol)}

	var symbol Symbol
	for {
		switch node.Type() {
		case "single_constant_definition":
			symbol, err = parseSingleConstantDefinition(node)
		case "element_anonymous_instantiation":
			symbol, err = parseElementAnonymousInstantiation(node)
		default:
			log.Fatalf("parsing %s not yet supported", node.Type())
		}

		if err != nil {
			log.Fatalf("%s: %v", path, err)
		}

		err = file.AddSymbol(symbol)
		if err != nil {
			log.Fatalf("%s: %v", path, err)
		}

		if node.HasNextSibling() == false {
			break
		}
		node = node.NextSibling()
	}

	pkg.AddFile(file)
}

func parseElementAnonymousInstantiation(n Node) (Element, error) {
	var err error

	isArray := false
	var count Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		expr, err := MakeExpression(n.Child(2))
		if err != nil {
			return Element{}, fmt.Errorf(": %v", err)
		}
		count = expr
	}

	var type_ ElementType
	if n.Child(1).Type() == "element_type" {
		type_, err = ToElementType(n.Child(1).Content())
	} else {
		type_, err = ToElementType(n.Child(4).Content())
	}
	if err != nil {
		return Element{}, fmt.Errorf("%d: element anonymous instantiation: %v", n.LineNumber(), err)
	}

	//if parser.node.children[-1].type == 'element_body':
	//    properties, symbols = parse_element_body(
	//        ParserFromNode(parser, parser.node.children[-1]), symbol
	//    )
	//    if properties:
	//        symbol['Properties'] = properties
	//    if symbols:
	//        for _, sym in symbols.items():
	//            sym['Parent'] = RefDict(symbol)
	//        symbol['Symbols'] = symbols

	element := Element{
		common: common{
			Id:         generateId(),
			lineNumber: n.LineNumber(),
			name:       n.Child(0).Content(),
		},
		IsArray:           isArray,
		Count:             count,
		Type:              type_,
		InstantiationType: Anonymous,
	}

	return element, nil
}

func parseSingleConstantDefinition(n Node) (Constant, error) {
	v, err := MakeExpression(n.Child(3))
	if err != nil {
		return Constant{}, fmt.Errorf("%d: single constant definition: %v", n.LineNumber(), err)
	}

	constant := Constant{
		common: common{
			Id:         generateId(),
			lineNumber: n.LineNumber(),
			name:       n.Child(1).Content(),
		},
		value: v,
	}

	return constant, nil
}
