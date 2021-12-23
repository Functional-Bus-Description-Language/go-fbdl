// Package prs implements parser based on the tree-sitter parser.
package prs

import (
	"bufio"
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ts"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"unicode"
)

func ParsePackages(packages Packages) {
	var wg sync.WaitGroup

	for name, _ := range packages {
		for i, _ := range packages[name] {
			wg.Add(1)
			go parsePackage(packages[name][i], &wg)
		}
	}

	wg.Wait()

	bindImports(packages)
}

func parsePackage(pkg *Package, wg *sync.WaitGroup) {
	defer wg.Done()

	var files_wg sync.WaitGroup
	defer files_wg.Wait()

	if pkg.Name == "main" {
		files_wg.Add(1)
		parseFile(pkg.Path, pkg, &files_wg)
		checkInstantiations(pkg)
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
		parseFile(path.Join(pkg.Path, file.Name()), pkg, &files_wg)
	}

	checkInstantiations(pkg)
}

func checkInstantiations(pkg *Package) {
	for _, f := range pkg.Files {
		for _, symbol := range f.Symbols {
			if e, ok := symbol.(*ElementDefinition); ok {
				if e.type_ != "bus" && util.IsBaseType(e.type_) {
					log.Fatalf(
						"%s: line %d: element of type '%s' cannot be instantiated at package level",
						f.Path, e.LineNumber(), e.Type(),
					)
				} else if e.type_ == "bus" {
					if e.Name() != "main" || pkg.Name != "main" {
						log.Fatalf(
							"%s: line %d: bus instantiation must be named 'main' and must be placed in 'main' package",
							f.Path, e.LineNumber(),
						)
					}
				}
			}
		}
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

func checkIndentAndTrailingSemicolon(code []string) error {
	var currentIndent uint32 = 0

	for i, line := range code {
		if line == "\n" {
			continue
		}

		if len(line) > 0 && line[len(line)-1] == ';' {
			return fmt.Errorf("line %d: extra ';' at end of line", i+1)
		}

		indent, err := getIndent(line)
		if err != nil {
			return fmt.Errorf("line %d: %v", i, err)
		} else if indent > currentIndent+1 {
			return fmt.Errorf("line %d: multi indent detected", i+1)
		}
		currentIndent = indent
	}

	return nil
}

func readFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	io_scanner := bufio.NewScanner(f)
	var code []string
	for io_scanner.Scan() {
		code = append(code, strings.TrimRight(io_scanner.Text(), " \t"))
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	return code
}

func parseFile(path string, pkg *Package, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	codeLines := readFile(path)

	err = checkIndentAndTrailingSemicolon(codeLines)
	if err != nil {
		log.Fatalf("%s: %v", path, err)
	}

	var codeBytes []byte
	for _, line := range codeLines {
		codeBytes = append(codeBytes, []byte(line)...)
		codeBytes = append(codeBytes, byte('\n'))
	}

	node := ts.MakeRootNode(codeBytes)

	file := File{Path: path, Pkg: pkg, Symbols: SymbolContainer{}, Imports: make(map[string]Import)}

	var symbols []Symbol
	for {
		switch node.Type() {
		case "element_anonymous_multi_line_instantiation":
			symbols, err = parseElementAnonymousMultiLineInstantiation(node, &file)
		case "element_anonymous_single_line_instantiation":
			symbols, err = parseElementAnonymousSingleLineInstantiation(node, &file)
		case "element_definitive_instantiation":
			symbols, err = parseElementDefinitiveInstantiation(node, &file)
		case "element_type_definition":
			symbols, err = parseElementTypeDefinition(node, &file)
		case "multi_constant_definition":
			symbols, err = parseMultiConstantDefinition(node)
		case "single_constant_definition":
			symbols, err = parseSingleConstantDefinition(node)
		case "single_import_statement":
			i := parseSingleImportStatement(node)
			if _, exist := file.Imports[i.ImportName]; exist {
				log.Fatalf(
					"%s: line %d: at least two packages imported as '%s'",
					path, node.LineNumber(), i.ImportName,
				)
			}
			file.Imports[i.ImportName] = i
			goto nextNode
		case "ERROR":
			log.Fatalf("%s: line %d: invalid syntax, tree-sitter ERROR", path, node.LineNumber())
		default:
			panic(fmt.Sprintf("parsing %q not yet supported", node.Type()))
		}

		if err != nil {
			log.Fatalf("%s: %v", path, err)
		}

		for i := 0; i < len(symbols); i++ {
			err = file.AddSymbol(symbols[i])
			if err != nil {
				log.Fatalf("%s: %v", path, err)
			}

			err = pkg.AddSymbol(symbols[i])
			if err != nil {
				log.Fatalf("%s: %v", path, err)
			}
		}

	nextNode:
		if node.HasNextSibling() == false {
			break
		}
		node = node.NextSibling()
	}

	pkg.AddFile(&file)
}

func parseArgumentList(n ts.Node, parent Searchable) ([]Argument, error) {
	args := []Argument{}

	names := []string{}
	var err error
	var hasName = false
	name := ""
	var val Expression
	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "(" || t == "," || t == "=" || t == ")" {
			continue
		}

		if t == "identifier" {
			name = nc.Content()
			hasName = true
		} else {
			val, err = MakeExpression(nc, parent)
			if err != nil {
				return args, fmt.Errorf("argument list: %v", err)
			}
		}

		next_node_type := n.Child(i + 1).Type()
		if next_node_type == "," || next_node_type == ")" {
			if name != "" {
				for i, _ := range names {
					if name == names[i] {
						return args, fmt.Errorf("argument '%s' assigned at least twice in argument list", name)
					}
				}
			}

			names = append(names, name)
			args = append(args, Argument{HasName: hasName, Name: name, Value: val})
			hasName = false
			name = ""
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

func parseElementAnonymousMultiLineInstantiation(n ts.Node, parent Searchable) ([]Symbol, error) {
	var err error

	isArray := false
	var count Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		if n.Child(2).Child(0).IsMissing() {
			return nil, fmt.Errorf(
				"line %d: '%s' element, missing array size expression",
				n.Child(2).LineNumber(), n.Child(0).Content(),
			)
		}
		expr, err := MakeExpression(n.Child(2), parent)
		if err != nil {
			return nil, fmt.Errorf(": %v", err)
		}
		count = expr
	}

	var type_ string
	if n.Child(1).Type() == "element_type" {
		type_ = n.Child(1).Content()
	} else {
		type_ = n.Child(4).Content()
	}

	if util.IsBaseType(type_) == false {
		return nil,
			fmt.Errorf(
				"line %d: invalid type '%s', only base types can be used in anonymous instantiation",
				n.LineNumber(), type_,
			)
	}

	elem := ElementDefinition{
		base: base{
			lineNumber: n.LineNumber(),
			name:       n.Child(0).Content(),
		},
		IsArray:           isArray,
		Count:             count,
		type_:             type_,
		InstantiationType: Anonymous,
	}

	var props map[string]Property
	var symbols SymbolContainer
	lastNode := n.LastChild()
	if lastNode.Type() == "element_body" {
		props, symbols, err = parseElementBody(lastNode, &elem)
		if err != nil {
			return nil, fmt.Errorf(
				"line %d: '%s' element anonymous instantiation: %v", n.LineNumber(), elem.name, err,
			)
		}

		for prop, v := range props {
			if err = util.IsValidProperty(prop, type_); err != nil {
				return nil,
					fmt.Errorf(
						"line %d: element anonymous instantiation: "+
							"line %d: %v",
						n.LineNumber(), v.LineNumber, err,
					)
			}
		}
	}

	elem.properties = props
	elem.symbols = symbols

	if len(elem.symbols) > 0 {
		for name, _ := range elem.symbols {
			elem.symbols[name].SetParent(&elem)
		}
	}

	return []Symbol{&elem}, nil
}

func parseElementAnonymousSingleLineInstantiation(n ts.Node, parent Searchable) ([]Symbol, error) {
	var err error

	isArray := false
	var count Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		if n.Child(2).Child(0).IsMissing() {
			return nil, fmt.Errorf(
				"line %d: '%s' element, missing array size expression",
				n.Child(2).LineNumber(), n.Child(0).Content(),
			)
		}
		expr, err := MakeExpression(n.Child(2), parent)
		if err != nil {
			return nil, fmt.Errorf(": %v", err)
		}
		count = expr
	}

	var type_ string
	if n.Child(1).Type() == "element_type" {
		type_ = n.Child(1).Content()
	} else {
		type_ = n.Child(4).Content()
	}

	if util.IsBaseType(type_) == false {
		return nil,
			fmt.Errorf(
				"line %d: invalid type '%s', only base types can be used in anonymous instantiation",
				n.LineNumber(), type_,
			)
	}

	elem := ElementDefinition{
		base: base{
			lineNumber: n.LineNumber(),
			name:       n.Child(0).Content(),
		},
		IsArray:           isArray,
		Count:             count,
		type_:             type_,
		InstantiationType: Anonymous,
	}

	var props map[string]Property

	lastNode := n.LastChild()
	if lastNode.Type() == "multi_property_assignment" {
		props, err = parseMultiPropertyAssignment(lastNode, &elem)
		if err != nil {
			return nil, fmt.Errorf(
				"line %d: '%s' element anonymous instantiation: %v", n.LineNumber(), elem.name, err,
			)
		}

		for prop, v := range props {
			if err = util.IsValidProperty(prop, type_); err != nil {
				return nil,
					fmt.Errorf(
						"line %d: '%s' element anonymous instantiation: "+
							"line %d: %v",
						n.LineNumber(), elem.name, v.LineNumber, err,
					)
			}
		}
	}

	elem.properties = props

	return []Symbol{&elem}, nil
}

func parseElementDefinitiveInstantiation(n ts.Node, parent Searchable) ([]Symbol, error) {
	var err error

	isArray := false
	var count Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		if n.Child(2).Child(0).IsMissing() {
			return nil, fmt.Errorf(
				"line %d: '%s' element, missing array size expression",
				n.Child(2).LineNumber(), n.Child(0).Content(),
			)
		}
		count, err = MakeExpression(n.Child(2), parent)
		if err != nil {
			return nil, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
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
			return nil,
				fmt.Errorf(
					"line %d: symbol '%s' imported from package '%s' starts with lower case letter",
					n.LineNumber(), id, pkg,
				)
		}
	}

	args := []Argument{}
	if n.Child(int(n.ChildCount()-2)).Type() == "argument_list" {
		args, err = parseArgumentList(n.Child(int(n.ChildCount()-2)), parent)
		if err != nil {
			return nil, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
		}
	}

	last_child := n.LastChild()
	if last_child.Type() == "argument_list" {
		args, err = parseArgumentList(last_child, parent)
		if err != nil {
			return nil, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
		}
	}

	elem := ElementDefinition{
		base: base{
			lineNumber: n.LineNumber(),
			name:       n.Child(0).Content(),
		},
		IsArray:           isArray,
		Count:             count,
		type_:             type_,
		InstantiationType: Definitive,
		args:              args,
	}

	props := make(map[string]Property)
	symbols := SymbolContainer{}
	if last_child.Type() == "element_body" {
		props, symbols, err = parseElementBody(last_child, &elem)
		if err != nil {
			return nil, fmt.Errorf("line %d: element definitve instantiation: %v", n.LineNumber(), err)
		}
	}

	elem.properties = props
	elem.symbols = symbols

	if len(elem.symbols) > 0 {
		for name, _ := range elem.symbols {
			elem.symbols[name].SetParent(&elem)
		}
	}

	return []Symbol{&elem}, nil
}

func parseElementBody(n ts.Node, element Searchable) (map[string]Property, SymbolContainer, error) {
	var err error
	props := make(map[string]Property)
	symbols := SymbolContainer{}

	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()
		switch t {
		case "ERROR":
			if nc.PrevSibling().Type() == "element_anonymous_single_line_instantiation" &&
				nc.NextSibling().Type() == "single_property_assignment" {
				return props, symbols, fmt.Errorf(
					"line %d: column %d: missing ';' or newline", nc.LineNumber(), nc.Column()-1,
				)
			} else {
				panic("implement me")
			}
		case "single_property_assignment":
			name := nc.Child(0).Content()
			if _, ok := props[name]; ok {
				return props,
					symbols,
					fmt.Errorf("line %d: property '%s' assigned at least twice in the same element body", nc.LineNumber(), name)
			}
			expr, err := MakeExpression(nc.Child(2), element)
			if err != nil {
				return props,
					symbols,
					fmt.Errorf("line %d: single property assignment: %v", nc.LineNumber(), err)
			}
			props[name] = Property{LineNumber: nc.LineNumber(), Value: expr}
		default:
			var ss []Symbol
			switch t {
			case "element_type_definition":
				ss, err = parseElementTypeDefinition(nc, element)
			case "element_anonymous_multi_line_instantiation":
				ss, err = parseElementAnonymousMultiLineInstantiation(nc, element)
			case "element_anonymous_single_line_instantiation":
				ss, err = parseElementAnonymousSingleLineInstantiation(nc, element)
			case "element_definitive_instantiation":
				ss, err = parseElementDefinitiveInstantiation(nc, element)
			case "single_constant_definition":
				ss, err = parseSingleConstantDefinition(nc)
			case "multi_constant_definition":
				panic("not yet implemented")
			case "ERROR":
				return props, symbols, fmt.Errorf("line %d: invalid syntax, tree-sitter ERROR", n.LineNumber())
			default:
				panic("this should never happen %s")
			}

			if err != nil {
				return props,
					symbols,
					fmt.Errorf("element body: %v", err)
			}

			for i := 0; i < len(ss); i++ {
				s, exists := symbols.Get(ss[i].Name())
				if exists {
					return props,
						symbols,
						fmt.Errorf(
							"line %d: symbol '%s' defined at least twice in the same element body, "+
								"first occurrence line %d",
							nc.LineNumber(), ss[i].Name(), s.LineNumber(),
						)
				}
				_ = symbols.Add(ss[i])
			}
		}
	}

	return props, symbols, nil
}

func parseElementTypeDefinition(n ts.Node, parent Searchable) ([]Symbol, error) {
	args := []Argument{}
	params := []Parameter{}
	props := make(map[string]Property)
	symbols := SymbolContainer{}

	name := n.Child(1).Content()
	if util.IsBaseType(name) {
		return nil, fmt.Errorf("line %d: invalid type name '%s', type name cannot be the same as base type", n.LineNumber(), name)
	}

	type__ := TypeDefinition{
		base: base{
			lineNumber: n.LineNumber(),
			name:       name,
		},
	}

	var err error
	var type_ string

	for i := 2; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		switch t {
		case "parameter_list":
			params, err = parseParameterList(nc, parent)
		case "identifier":
			type_ = nc.Content()
		case "qualified_identifier":
			type_ = nc.Content()
			aux := strings.Split(type_, ".")
			pkg := aux[0]
			id := aux[1]
			if unicode.IsUpper([]rune(id)[0]) == false {
				return nil,
					fmt.Errorf(
						"line %d: symbol '%s' imported from package '%s' starts with lower case letter",
						nc.LineNumber(), id, pkg,
					)
			}
		case "argument_list":
			args, err = parseArgumentList(nc, parent)
		case "element_body":
			props, symbols, err = parseElementBody(nc, &type__)
		case "ERROR":
			return nil, fmt.Errorf("line %d: invalid syntax, tree-sitter ERROR", nc.LineNumber())
		default:
			panic("should never happen")
		}

		if err != nil {
			return nil, fmt.Errorf(
				"line %d: '%s' element type definition: %v", n.LineNumber(), type__.name, err,
			)
		}
	}

	if len(args) > 0 {
		if util.IsBaseType(type_) {
			return nil,
				fmt.Errorf("line %d: base type '%s' does not accept argument list", n.LineNumber(), type_)
		}
	}

	type__.type_ = type_
	type__.properties = props
	type__.symbols = symbols
	type__.params = params
	type__.args = args

	if len(type__.symbols) > 0 {
		for name, _ := range type__.symbols {
			type__.symbols[name].SetParent(&type__)
		}
	}

	return []Symbol{&type__}, nil
}

func parseMultiConstantDefinition(n ts.Node) ([]Symbol, error) {
	var symbols []Symbol

	var c *Constant

	for i := 0; i < int(n.ChildCount()); i++ {
		child := n.Child(i)

		switch child.Type() {
		case "const", "comment":
			continue
		case "identifier":
			c = &Constant{
				base: base{
					lineNumber: child.LineNumber(),
					name:       child.Content(),
				},
			}
		case "primary_expression", "expression_list":
			expr, err := MakeExpression(child, c)
			if err != nil {
				return nil, fmt.Errorf("line %d: constant %s: %v", c.LineNumber(), c.name, err)
			}

			c.Value = expr

			symbols = append(symbols, c)
		}
	}

	return symbols, nil
}

func parseMultiPropertyAssignment(n ts.Node, element Searchable) (map[string]Property, error) {
	props := make(map[string]Property)

	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		switch nc.Type() {
		case "identifier":
			name := nc.Content()
			if _, ok := props[name]; ok {
				return props,
					fmt.Errorf("line %d: property '%s' assigned at least twice in the same element body", nc.LineNumber(), name)
			}
			expr, err := MakeExpression(n.Child(i+2), element)
			if err != nil {
				return props,
					fmt.Errorf("line %d: '%s' property assignment: %v", nc.LineNumber(), name, err)
			}
			props[name] = Property{LineNumber: nc.LineNumber(), Value: expr}
		default:
			continue
		}
	}

	return props, nil
}

func parseParameterList(n ts.Node, parent Searchable) ([]Parameter, error) {
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
			defaultValue, err = MakeExpression(nc, parent)
			if err != nil {
				return nil, fmt.Errorf("parameter list: %v", err)
			}

			hasDefaultValue = true
		}

		next_node_type := n.Child(i + 1).Type()
		if next_node_type == "," || next_node_type == ")" {
			for i, _ := range params {
				if name == params[i].Name {
					return nil, fmt.Errorf("parameter '%s' defined at least twice", name)
				}
			}
			params = append(
				params,
				Parameter{Name: name, HasDefaultValue: hasDefaultValue, DefaultValue: defaultValue},
			)
		}
	}

	// Check if parameters without default value precede parameters with default value.
	with_default := false
	for i, p := range params {
		if with_default && p.HasDefaultValue == false {
			return nil, fmt.Errorf("parameters without default value must precede the ones with default value")
		}

		if params[i].HasDefaultValue {
			with_default = true
		}
	}

	return params, nil
}

func parseSingleConstantDefinition(n ts.Node) ([]Symbol, error) {
	c := &Constant{
		base: base{
			lineNumber: n.LineNumber(),
			name:       n.Child(1).Content(),
		},
	}

	v, err := MakeExpression(n.Child(3), c)
	if err != nil {
		return nil, fmt.Errorf("line %d: single constant definition: %v", n.LineNumber(), err)
	}

	c.Value = v

	return []Symbol{c}, nil
}

func parseSingleImportStatement(n ts.Node) Import {
	var path string
	var import_name string

	if n.ChildCount() == 2 {
		path = n.Child(1).Content()
		path = path[1 : len(path)-1]
		import_name = strings.Split(path, "/")[0]
		if len(import_name) > 4 && import_name[0:3] == "fbd-" {
			import_name = import_name[4:]
		}
	} else {
		path = n.Child(2).Content()
		path = path[1 : len(path)-2]
		import_name = n.Child(1).Content()
	}

	return Import{Path: path, ImportName: import_name}
}
