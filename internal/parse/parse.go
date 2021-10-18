package parse

import (
	"bufio"
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/expr"
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
	defer wg.Wait()

	for name, _ := range packages {
		for i, _ := range packages[name] {
			wg.Add(1)
			go parsePackage(packages[name][i], &wg)
		}
	}
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

func checkIndent(code []string) error {
	var currentIndent uint32 = 0

	for i, line := range code {
		if line == "\n" {
			continue
		}

		indent, err := getIndent(line)
		if err != nil {
			return fmt.Errorf("line %d: %v", i, err)
		} else if indent > currentIndent+1 {
			return fmt.Errorf("line %d: multi indent detected", i)
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
	defer f.Close()

	io_scanner := bufio.NewScanner(f)
	var code []string
	for io_scanner.Scan() {
		code = append(code, io_scanner.Text())
	}

	return code
}

func parseFile(path string, pkg *Package, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	codeLines := readFile(path)

	err = checkIndent(codeLines)
	if err != nil {
		log.Fatalf("%s: %v", path, err)
	}

	var codeBytes []byte
	for _, line := range codeLines {
		codeBytes = append(codeBytes, []byte(line)...)
		codeBytes = append(codeBytes, byte('\n'))
	}

	node := ts.MakeRootNode(codeBytes)

	file := File{Path: path, Pkg: pkg, Symbols: make(map[string]Symbol), Imports: make(map[string]Import)}

	var symbols []Symbol
	for {
		switch node.Type() {
		case "element_anonymous_instantiation":
			symbols, err = parseElementAnonymousInstantiation(node)
		case "element_definitive_instantiation":
			symbols, err = parseElementDefinitiveInstantiation(node)
		case "element_type_definition":
			symbols, err = parseElementTypeDefinition(node)
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
			log.Fatalf("parsing %s not yet supported", node.Type())
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

	pkg.AddFile(file)
}

func parseArgumentList(n ts.Node) ([]Argument, error) {
	args := []Argument{}

	names := []string{}
	var err error
	var hasName = true
	name := ""
	var val expr.Expression
	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "(" || t == "," || t == "=" || t == ")" {
			continue
		}

		if t == "identifier" {
			name = nc.Content()
		} else {
			val, err = expr.Make(nc)
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

func parseElementAnonymousInstantiation(n ts.Node) ([]Symbol, error) {
	var err error

	isArray := false
	var count expr.Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		expr, err := expr.Make(n.Child(2))
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

	var props map[string]Property
	var symbols map[string]Symbol
	last_node := n.LastChild()
	if last_node.Type() == "element_body" {
		props, symbols, err = parseElementBody(last_node)
		if err != nil {
			return nil, fmt.Errorf("line %d: element anonymous instantiation: %v", n.LineNumber(), err)
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

	elem := ElementDefinition{
		base: base{
			lineNumber: n.LineNumber(),
			name:       n.Child(0).Content(),
		},
		IsArray:           isArray,
		Count:             count,
		type_:             type_,
		InstantiationType: Anonymous,
		properties:        props,
		symbols:           symbols,
	}

	if len(elem.symbols) > 0 {
		for name, _ := range elem.symbols {
			elem.symbols[name].SetParent(&elem)
		}
	}

	return []Symbol{&elem}, nil
}

func parseElementDefinitiveInstantiation(n ts.Node) ([]Symbol, error) {
	var err error

	isArray := false
	var count expr.Expression
	if n.Child(1).Type() == "[" {
		isArray = true
		count, err = expr.Make(n.Child(2))
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
		args, err = parseArgumentList(n.Child(int(n.ChildCount() - 2)))
		if err != nil {
			return nil, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
		}
	}

	last_child := n.LastChild()
	if last_child.Type() == "argument_list" {
		args, err = parseArgumentList(last_child)
		if err != nil {
			return nil, fmt.Errorf("line %d: element definitive instantiation: %v", n.LineNumber(), err)
		}
	}

	props := make(map[string]Property)
	symbols := make(map[string]Symbol)
	if last_child.Type() == "element_body" {
		props, symbols, err = parseElementBody(last_child)
		if err != nil {
			return nil, fmt.Errorf("line %d: element definitve instantiation: %v", n.LineNumber(), err)
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
		properties:        props,
		symbols:           symbols,
		args:              args,
	}

	if len(elem.symbols) > 0 {
		for name, _ := range elem.symbols {
			elem.symbols[name].SetParent(&elem)
		}
	}

	return []Symbol{&elem}, nil
}

func parseElementBody(n ts.Node) (map[string]Property, map[string]Symbol, error) {
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
			expr, err := expr.Make(nc.Child(2))
			if err != nil {
				return props,
					symbols,
					fmt.Errorf("line %d: property assignment: %v", nc.LineNumber(), err)
			}
			props[name] = Property{LineNumber: nc.LineNumber(), Value: expr}
		default:
			var ss []Symbol
			switch t {
			case "element_type_definition":
				ss, err = parseElementTypeDefinition(nc)
			case "element_anonymous_instantiation":
				ss, err = parseElementAnonymousInstantiation(nc)
			case "element_definitive_instantiation":
				ss, err = parseElementDefinitiveInstantiation(nc)
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
				if _, exist := symbols[ss[i].Name()]; exist {
					return props,
						symbols,
						fmt.Errorf(
							"line %d: symbol '%s' defined at least twice in the same element body", nc.LineNumber(), ss[i].Name(),
						)
				}
				symbols[ss[i].Name()] = ss[i]
			}
		}
	}

	return props, symbols, nil
}

func parseElementTypeDefinition(n ts.Node) ([]Symbol, error) {
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
				return nil,
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
			return nil, fmt.Errorf("line %d: element type definition: %v", n.LineNumber(), err)
		}
	}

	if len(args) > 0 {
		if util.IsBaseType(type_) {
			return nil,
				fmt.Errorf("line %d: base type '%s' does not accept argument list", n.LineNumber(), type_)
		}
	}

	name := n.Child(1).Content()
	if util.IsBaseType(name) {
		return nil, fmt.Errorf("line %d: invalid type name '%s', type name cannot be the same as base type", n.LineNumber(), name)
	}

	type__ := TypeDefinition{
		base: base{
			lineNumber: n.LineNumber(),
			name:       name,
		},
		type_:      type_,
		properties: props,
		symbols:    symbols,
		params:     params,
		args:       args,
	}

	if len(type__.symbols) > 0 {
		for name, _ := range type__.symbols {
			type__.symbols[name].SetParent(&type__)
		}
	}

	return []Symbol{&type__}, nil
}

func parseMultiConstantDefinition(n ts.Node) ([]Symbol, error) {
	var symbols []Symbol

	for i := 0; i < (int(n.ChildCount())-1)/3; i++ {
		expr, err := expr.Make(n.Child(i*3 + 3))
		if err != nil {
			return nil, fmt.Errorf("line %d: %v", n.Child(i*3+1).LineNumber(), err)
		}

		symbols = append(symbols,
			&Constant{
				base: base{
					lineNumber: n.Child(i*3 + 1).LineNumber(),
					name:       n.Child(i*3 + 1).Content(),
				},
				value: expr,
			},
		)
	}

	return symbols, nil
}

func parseParameterList(n ts.Node) ([]Parameter, error) {
	params := []Parameter{}

	var err error
	var name string
	var hasDefaultValue bool
	var defaultValue expr.Expression

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
			defaultValue, err = expr.Make(nc)
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
	v, err := expr.Make(n.Child(3))
	if err != nil {
		return nil, fmt.Errorf("line %d: single constant definition: %v", n.LineNumber(), err)
	}

	return []Symbol{&Constant{
		base: base{
			lineNumber: n.LineNumber(),
			name:       n.Child(1).Content(),
		},
		value: v,
	}}, nil
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
