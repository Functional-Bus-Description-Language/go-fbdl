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
	"strings"
	"sync"
	"unicode"
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

func (n Node) ChildCount() uint32 {
	return n.n.ChildCount()
}

func (n Node) LastChild() Node {
	return n.Child(int(n.ChildCount() - 1))
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
	//fmt.Println(root)
	tsnode := root.Child(0)
	if tsnode == nil {
		panic("TODO")
	}
	node := Node{n: tsnode, code: code_bytes}

	file := File{Path: path, Pkg: pkg, Symbols: make(map[string]Symbol)}

	var symbol Symbol
	for {
		switch node.Type() {
		case "element_anonymous_instantiation":
			symbol, err = parseElementAnonymousInstantiation(node)
		case "element_definitive_instantiation":
			symbol, err = parseElementDefinitiveInstantiation(node)
		case "element_type_definition":
			symbol, err = parseElementTypeDefinition(node)
		case "single_constant_definition":
			symbol, err = parseSingleConstantDefinition(node)
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

func parseArgumentList(n Node) ([]Argument, error) {
	args := []Argument{}

	names := []string{}
	var err error
	var hasName = true
	var name string
	var val Expression
	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "(" || t == "," || t == "=" || t == ")" {
			continue
		}

		if t == "identifier" {
			name = nc.Content()
		} else {
			val, err = MakeExpression(nc)
			if err != nil {
				return args, fmt.Errorf("argument list: %v", err)
			}
		}

		next_node_type := n.Child(i + 1).Type()
		if next_node_type == "," || next_node_type == ")" {
			for i, _ := range names {
				if name == names[i] {
					return args, fmt.Errorf("argument '%s' assigned at least twice in argument list", name)
				}
			}

			args = append(args, Argument{HasName: hasName, Name: name, Value: val})
			hasName = false
		}
	}

	// Check if arguments without name precede arguments with name.
	with_name := false
	for _, a := range args {
		if with_name && a.HasName == false {
			return args, fmt.Errorf("arguments without name must precede the ones with name")
		}

		if a.HasName {
			with_name = true
		}
	}

	return args, nil
}

func parseElementAnonymousInstantiation(n Node) (*Element, error) {
	var err error

	isArray := false
	var count Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		expr, err := MakeExpression(n.Child(2))
		if err != nil {
			return &Element{}, fmt.Errorf(": %v", err)
		}
		count = expr
	}

	var type_ string
	if n.Child(1).Type() == "element_type" {
		type_ = n.Child(1).Content()
	} else {
		type_ = n.Child(4).Content()
	}

	if IsBaseType(type_) == false {
		return &Element{},
			fmt.Errorf(
				"line %d: invalid type '%s', only base types can be used in anonymous instantiation",
				n.LineNumber(), type_,
			)
	}

	var props map[string]Property
	var symbols map[string]Symbol
	last_node := n.LastChild()
	if last_node.Type() == "element_body" {
		props, symbols, err = parseElementBody(last_node)
		if err != nil {
			return &Element{}, fmt.Errorf("line %d: element anonymous instantiation: %v", n.LineNumber(), err)
		}

		for prop, v := range props {
			if IsValidProperty(type_, prop) == false {
				return &Element{},
					fmt.Errorf(
						"line %d: element anonymous instantiation: "+
							"line %d: invalid property '%s' for element of type '%v'",
						n.LineNumber(), v.LineNumber, prop, type_,
					)
			}
		}
	}

	elem := Element{
		common: common{
			Id:         generateId(),
			lineNumber: n.LineNumber(),
			name:       n.Child(0).Content(),
		},
		IsArray:           isArray,
		Count:             count,
		Type:              type_,
		InstantiationType: Anonymous,
		Properties:        props,
		Symbols:           symbols,
	}

	if len(elem.Symbols) > 0 {
		for name, _ := range elem.Symbols {
			elem.Symbols[name].SetParent(&elem)
		}
	}

	return &elem, nil
}

func parseElementDefinitiveInstantiation(n Node) (*Element, error) {
	var err error

	isArray := false
	var count Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		count, err = MakeExpression(n.Child(2))
		if err != nil {
			return &Element{}, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
		}
	}

	var type_ string
	if n.Child(1).Type() == "identifier" || n.Child(1).Type() == "qualified_identifier" {
		type_ = n.Child(1).Content()
	} else {
		type_ = n.Child(4).Content()
	}

	if strings.Contains(type_, ".") {
		aux := strings.Split(type_, ".")
		pkg := aux[0]
		id := aux[1]
		if unicode.IsUpper([]rune(id)[0]) == false {
			return &Element{},
				fmt.Errorf(
					"line %d: symbol '%s' imported from package '%s' starts with lower case letter",
					n.LineNumber(), id, pkg,
				)
		}
	}

	args := []Argument{}
	if n.Child(int(n.ChildCount()-2)).Type() == "argument_list" {
		args, err = parseArgumentList(n.Child(int(n.ChildCount() - 2)))
		if err != nil {
			return &Element{}, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
		}
	}

	last_child := n.LastChild()
	if last_child.Type() == "argument_list" {
		args, err = parseArgumentList(last_child)
		if err != nil {
			return &Element{}, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
		}
	}

	props := make(map[string]Property)
	symbols := make(map[string]Symbol)
	if last_child.Type() == "element_body" {
		props, symbols, err = parseElementBody(last_child)
		if err != nil {
			return &Element{}, fmt.Errorf("line %d: element definitve instantiation: %v", n.LineNumber(), err)
		}
	}

	name := n.Child(0).Content()
	if IsBaseType(name) {
		return &Element{},
			fmt.Errorf("line %d: invalid instance name '%s', element instance can not have the same name as base type",
				n.LineNumber(), name,
			)
	}

	elem := Element{
		common: common{
			Id:         generateId(),
			lineNumber: n.LineNumber(),
			name:       name,
		},
		IsArray:           isArray,
		Count:             count,
		Type:              type_,
		InstantiationType: Definitive,
		Properties:        props,
		Symbols:           symbols,
		Arguments:         args,
	}

	if len(elem.Symbols) > 0 {
		for name, _ := range elem.Symbols {
			elem.Symbols[name].SetParent(&elem)
		}
	}

	return &elem, nil
}

func parseElementBody(n Node) (map[string]Property, map[string]Symbol, error) {
	var err error
	props := make(map[string]Property)
	symbols := make(map[string]Symbol)

	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()
		switch t {
		case "property_assignment":
			name := nc.Child(0).Content()
			if _, ok := props[name]; ok {
				return props,
					symbols,
					fmt.Errorf("line %d: property '%s' assigned at least twice in the same element body", nc.LineNumber(), name)
			}
			expr, err := MakeExpression(nc.Child(2))
			if err != nil {
				return props,
					symbols,
					fmt.Errorf("line %d: property assignment: %v", nc.LineNumber(), err)
			}
			props[name] = Property{LineNumber: nc.LineNumber(), Value: expr}
		default:
			var s Symbol
			switch t {
			case "element_type_definition":
				s, err = parseElementTypeDefinition(nc)
			case "element_anonymous_instantiation":
				s, err = parseElementAnonymousInstantiation(nc)
			case "element_definitive_instantiation":
				s, err = parseElementDefinitiveInstantiation(nc)
			case "single_constant_definition":
				s, err = parseSingleConstantDefinition(nc)
			case "multi_constant_definition":
				panic("not yet implemented")
			default:
				panic("this should never happen %s")
			}

			if err != nil {
				return props,
					symbols,
					fmt.Errorf("element body: %v", err)
			}

			if _, exist := symbols[s.Name()]; exist {
				return props,
					symbols,
					fmt.Errorf("line %d: symbol '%s' defined at least twice in the same element body", nc.LineNumber(), s.Name())
			}
			symbols[s.Name()] = s
		}
	}

	return props, symbols, nil
}

func parseElementTypeDefinition(n Node) (*Type, error) {
	args := []Argument{}
	params := []Parameter{}
	props := make(map[string]Property)
	symbols := make(map[string]Symbol)

	var err error
	var type_ string

	for i := 2; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		switch t {
		case "parameter_list":
			params, err = parseParameterList(nc)
		case "identifier":
			type_ = nc.Content()
		case "qualified_identifier":
			type_ = nc.Content()
			aux := strings.Split(type_, ".")
			pkg := aux[0]
			id := aux[1]
			if unicode.IsUpper([]rune(id)[0]) == false {
				return &Type{},
					fmt.Errorf(
						"line %d: symbol '%s' imported from package '%s' starts with lower case letter",
						nc.LineNumber(), id, pkg,
					)
			}
		case "argument_list":
			args, err = parseArgumentList(nc)
		case "element_body":
			props, symbols, err = parseElementBody(nc)
		default:
			panic("should never happen")
		}

		if err != nil {
			return &Type{}, fmt.Errorf("line %d: element type definition: %v", n.LineNumber(), err)
		}
	}

	if len(args) > 0 {
		if IsBaseType(type_) {
			return &Type{},
				fmt.Errorf("line %d: base type '%s' does not accept argument list", n.LineNumber(), type_)
		}
	}

	type__ := Type{
		common: common{
			Id:         generateId(),
			lineNumber: n.LineNumber(),
			name:       n.Child(1).Content(),
		},
		Parameters: params,
		Arguments:  args,
		Type:       type_,
		Properties: props,
		Symbols:    symbols,
	}

	if len(type__.Symbols) > 0 {
		for name, _ := range type__.Symbols {
			type__.Symbols[name].SetParent(&type__)
		}
	}

	return &type__, nil
}

func parseParameterList(n Node) ([]Parameter, error) {
	params := []Parameter{}

	var err error
	var name string
	var hasDefaultValue bool
	var defaultValue Expression

	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		// TODO: check if switch case works as expected here.
		if t == "(" || t == "=" || t == "," || t == ")" {
			continue
		}

		hasDefaultValue = false

		if t == "identifier" {
			name = nc.Content()
		} else {
			defaultValue, err = MakeExpression(nc)
			if err != nil {
				return params, fmt.Errorf("parameter list: %v", err)
			}

			hasDefaultValue = true
		}

		next_node_type := n.Child(i + 1).Type()
		if next_node_type == "," || next_node_type == ")" {
			for i, _ := range params {
				if name == params[i].Name {
					return params, fmt.Errorf("parameter '%s' defined at least twice", name)
				}
			}
			params = append(
				params,
				Parameter{Name: name, HasDefaultValue: hasDefaultValue, DefaultValue: defaultValue},
			)
		}
	}

	return params, nil
}

func parseSingleConstantDefinition(n Node) (*Constant, error) {
	v, err := MakeExpression(n.Child(3))
	if err != nil {
		return &Constant{}, fmt.Errorf("line %d: single constant definition: %v", n.LineNumber(), err)
	}

	return &Constant{
		common: common{
			Id:         generateId(),
			lineNumber: n.LineNumber(),
			name:       n.Child(1).Content(),
		},
		value: v,
	}, nil
}
