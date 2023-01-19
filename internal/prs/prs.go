// Package prs implements parser based on the tree-sitter parser.
package prs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ts"
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
		if line == "" {
			continue
		}

		if len(line) > 0 && line[len(line)-1] == ';' {
			if strings.TrimSpace(line)[0] != '#' {
				return fmt.Errorf("line %d: extra ';' at end of line", i+1)
			}
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
		log.Fatalf("reading file '%s': %v", path, err)
	}

	ioScanner := bufio.NewScanner(f)
	var code []string
	for ioScanner.Scan() {
		code = append(code, strings.TrimRight(ioScanner.Text(), " \t"))
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
	var cmnt comment
	for {
		switch node.Type() {
		case "multi_line_instantiation":
			symbols, err = parseMultiLineInstantiation(node, &file)
		case "single_line_instantiation":
			symbols, err = parseSingleLineInstantiation(node, &file)
		case "multi_line_type_definition":
			symbols, err = parseMultiLineTypeDefinition(node, &file)
		case "single_line_type_definition":
			symbols, err = parseSingleLineTypeDefinition(node, &file)
		case "multi_constant_definition":
			symbols, err = parseMultiConstantDefinition(node)
		case "single_constant_definition":
			symbols, err = parseSingleConstantDefinition(node)
		case "single_import_statement":
			i := parseSingleImportStatement(node)
			if _, exist := file.Imports[i.ImportName]; exist {
				log.Fatalf(
					"%s: line %d: at least two packages imported as '%s'",
					path, node.LineNum(), i.ImportName,
				)
			}
			file.Imports[i.ImportName] = i
			goto nextNode
		case "comment":
			if cmnt.isEmpty() {
				cmnt = makeComment(node.Content(), node.LineNum())
			} else if node.LineNum() == cmnt.endLineNum+1 {
				cmnt.append(node.Content())
			} else {
				cmnt = makeComment(node.Content(), node.LineNum())
			}
			symbols = []Symbol{}
		case "ERROR":
			log.Fatalf("%s: line %d: invalid syntax, tree-sitter ERROR", path, node.LineNum())
		default:
			panic(fmt.Sprintf("parsing %q not yet supported", node.Type()))
		}

		if err != nil {
			log.Fatalf("%s: %v", path, err)
		}

		// Attach comment to symbol as its documentation.
		if len(symbols) == 1 {
			if symbols[0].LineNum() == cmnt.endLineNum+1 {
				symbols[0].SetDoc(cmnt)
			}
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
		if !node.HasNextSibling() {
			break
		}
		node = node.NextSibling()
	}

	pkg.AddFile(&file)
}

func parseArgumentList(n ts.Node, parent Searchable) ([]Arg, error) {
	args := []Arg{}

	names := []string{}
	var err error
	var hasName = false
	name := ""
	var val Expr
	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "(" || t == "," || t == "=" || t == ")" {
			continue
		}

		if t == "declared_identifier" {
			name = nc.Content()
			hasName = true
		} else {
			val, err = MakeExpr(nc, parent)
			if err != nil {
				return args, fmt.Errorf("argument list: %v", err)
			}
		}

		nextNodeType := n.Child(i + 1).Type()
		if nextNodeType == "," || nextNodeType == ")" {
			if name != "" {
				for i := range names {
					if name == names[i] {
						return args, fmt.Errorf("argument '%s' assigned at least twice in argument list", name)
					}
				}
			}

			names = append(names, name)
			args = append(args, Arg{HasName: hasName, Name: name, Value: val})
			hasName = false
			name = ""
		}
	}

	// Check if arguments without name precede arguments with name.
	withName := false
	for _, a := range args {
		if withName && !a.HasName {
			return args, fmt.Errorf("arguments without name must precede the ones with name")
		}

		if a.HasName {
			withName = true
		}
	}

	return args, nil
}

func parseArrayMarker(n ts.Node, parent Searchable) (Expr, error) {
	var count Expr

	if n.Child(0).Type() == "[" {
		if n.Child(1).Child(0).IsMissing() || n.Child(1).Content() == "" {
			return nil, fmt.Errorf("missing array size expression")
		}
		expr, err := MakeExpr(n.Child(1), parent)
		if err != nil {
			return nil, fmt.Errorf(": %v", err)
		}
		count = expr
	}

	return count, nil
}

func parseMultiLineInstantiation(n ts.Node, parent Searchable) ([]Symbol, error) {
	i := Inst{base: base{lineNum: n.LineNum()}}

	var err error

	for j := 0; uint32(j) < n.ChildCount(); j++ {
		nc := n.Child(j)
		switch nc.Type() {
		case "ERROR":
			return nil, fmt.Errorf("line %d: invalid syntax, tree-sitter ERROR", nc.LineNum())
		case "identifier":
			i.name = nc.Content()
		case "array_marker":
			i.isArray = true
			i.count, err = parseArrayMarker(nc, parent)
		case "declared_identifier":
			i.typ = nc.Content()
		case "qualified_identifier":
			i.typ = nc.Content()
			err = util.IsValidQualifiedIdentifier(i.typ)
		case "argument_list":
			i.args, err = parseArgumentList(nc, parent)
		case "element_body":
			i.props, i.symbols, err = parseElementBody(nc, &i)
		default:
			panic("should never happen")
		}

		if err != nil {
			return nil, fmt.Errorf(
				"line %d: '%s' instantiation: %v", n.LineNum(), i.name, err,
			)
		}
	}

	err = i.validate()
	if err != nil {
		return nil, fmt.Errorf("line %d: '%s' instantiation: %v", n.LineNum(), i.name, err)
	}

	if len(i.symbols) > 0 {
		for name := range i.symbols {
			i.symbols[name].SetParent(&i)
		}
	}

	return []Symbol{&i}, nil
}

func parseSingleLineInstantiation(n ts.Node, parent Searchable) ([]Symbol, error) {
	i := Inst{base: base{lineNum: n.LineNum()}}

	var err error

	for j := 0; uint32(j) < n.ChildCount(); j++ {
		nc := n.Child(j)
		switch nc.Type() {
		case "ERROR":
			return nil, fmt.Errorf("line %d: invalid syntax, tree-sitter ERROR", nc.LineNum())
		case "identifier":
			i.name = nc.Content()
		case "array_marker":
			i.isArray = true
			i.count, err = parseArrayMarker(nc, parent)
		case "declared_identifier":
			i.typ = nc.Content()
		case "qualified_identifier":
			i.typ = nc.Content()
			err = util.IsValidQualifiedIdentifier(i.typ)
		case "argument_list":
			i.args, err = parseArgumentList(nc, parent)
		case ";", "comment":
			continue
		case "property_assignments":
			i.props, err = parsePropertyAssignments(nc, &i)
		default:
			panic("should never happen")
		}

		if err != nil {
			return nil, fmt.Errorf(
				"line %d: '%s' instantiation: %v", n.LineNum(), i.name, err,
			)
		}
	}

	err = i.validate()
	if err != nil {
		return nil, fmt.Errorf("line %d: '%s' instantiation: %v", n.LineNum(), i.name, err)
	}

	return []Symbol{&i}, nil
}

func parseElementBody(n ts.Node, element Searchable) (PropContainer, SymbolContainer, error) {
	var err error
	props := PropContainer{}
	symbols := SymbolContainer{}

	var cmnt comment

	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()
		switch t {
		case "ERROR":
			if nc.PrevSibling().Type() == "element_anonymous_single_line_instantiation" &&
				nc.NextSibling().Type() == "property_assignments" {
				return props, symbols, fmt.Errorf(
					"line %d: column %d: missing ';' or newline", nc.LineNum(), nc.Column()-1,
				)
			} else {
				return props, symbols, fmt.Errorf("line %d: invalid syntax, tree-sitter ERROR", n.LineNum())
			}
		case "property_assignments":
			ps, err := parsePropertyAssignments(nc, element)
			if err != nil {
				return props,
					symbols,
					fmt.Errorf("line %d: property assignments: %v", nc.LineNum(), err)
			}
			for _, p := range ps {
				if _, ok := props.Get(p.Name); ok {
					return props,
						symbols,
						fmt.Errorf(
							"line %d: property '%s' assigned at least twice in the same element body", nc.LineNum(), p.Name,
						)
				}
				props.Add(p)
			}
		default:
			var ss []Symbol
			switch t {
			case "multi_line_type_definition":
				ss, err = parseMultiLineTypeDefinition(nc, element)
			case "single_line_type_definition":
				ss, err = parseSingleLineTypeDefinition(nc, element)
			case "multi_line_instantiation":
				ss, err = parseMultiLineInstantiation(nc, element)
			case "single_line_instantiation":
				ss, err = parseSingleLineInstantiation(nc, element)
			case "single_constant_definition":
				ss, err = parseSingleConstantDefinition(nc)
			case "multi_constant_definition":
				panic("not yet implemented")
			case "comment":
				if cmnt.isEmpty() {
					cmnt = makeComment(nc.Content(), nc.LineNum())
				} else if nc.LineNum() == cmnt.endLineNum+1 {
					cmnt.append(nc.Content())
				} else {
					cmnt = makeComment(nc.Content(), nc.LineNum())
				}
			default:
				panic("should never happen")
			}

			if err != nil {
				return props,
					symbols,
					fmt.Errorf("element body: %v", err)
			}

			// Attach comment to symbol as its documentation.
			if len(ss) == 1 {
				if ss[0].LineNum() == cmnt.endLineNum+1 {
					ss[0].SetDoc(cmnt)
				}
			}

			for i := 0; i < len(ss); i++ {
				s, exists := symbols.GetByName(ss[i].Name())
				if exists {
					return props,
						symbols,
						fmt.Errorf(
							"line %d: symbol '%s' defined at least twice in the same element body, "+
								"first occurrence line %d",
							nc.LineNum(), ss[i].Name(), s.LineNum(),
						)
				}
				_ = symbols.Add(ss[i])
			}
		}
	}

	return props, symbols, nil
}

func parseMultiLineTypeDefinition(n ts.Node, parent Searchable) ([]Symbol, error) {
	name := n.Child(1).Content()
	if util.IsBaseType(name) {
		return nil, fmt.Errorf("line %d: invalid type name '%s', type name cannot be the same as base type", n.LineNum(), name)
	}

	t := Type{
		base: base{
			lineNum: n.LineNum(),
			name:    name,
		},
	}

	var err error

	for i := 2; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)

		switch nc.Type() {
		case "parameter_list":
			t.params, err = parseParameterList(nc, parent)
		case "identifier":
			t.typ = nc.Content()
		case "declared_identifier":
			t.typ = nc.Content()
		case "qualified_identifier":
			t.typ = nc.Content()
			err = util.IsValidQualifiedIdentifier(t.typ)
		case "argument_list":
			t.args, err = parseArgumentList(nc, parent)
		case "element_body":
			t.props, t.symbols, err = parseElementBody(nc, &t)
		case "ERROR":
			return nil, fmt.Errorf("line %d: invalid syntax, tree-sitter ERROR", nc.LineNum())
		default:
			panic("should never happen")
		}

		if err != nil {
			return nil, fmt.Errorf(
				"line %d: '%s' type definition: %v", n.LineNum(), t.name, err,
			)
		}
	}

	if len(t.args) > 0 && util.IsBaseType(t.typ) {
		return nil,
			fmt.Errorf("line %d: base type '%s' does not accept argument list", n.LineNum(), t.typ)
	}

	if util.IsBaseType(t.typ) {
		for _, p := range t.props {
			if err = util.IsValidProperty(p.Name, t.typ); err != nil {
				return nil,
					fmt.Errorf(
						"line %d: type definition: line %d: %v",
						n.LineNum(), p.LineNum, err,
					)
			}
		}
	}

	if len(t.symbols) > 0 {
		for name := range t.symbols {
			t.symbols[name].SetParent(&t)
		}
	}

	return []Symbol{&t}, nil
}

func parseSingleLineTypeDefinition(n ts.Node, parent Searchable) ([]Symbol, error) {
	name := n.Child(1).Content()
	if util.IsBaseType(name) {
		return nil, fmt.Errorf("line %d: invalid type name '%s', type name cannot be the same as base type", n.LineNum(), name)
	}

	t := Type{
		base: base{
			lineNum: n.LineNum(),
			name:    name,
		},
	}

	var err error

	for i := 2; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)

		switch nc.Type() {
		case "parameter_list":
			t.params, err = parseParameterList(nc, parent)
		case "declared_identifier":
			t.typ = nc.Content()
		case "qualified_identifier":
			t.typ = nc.Content()
			err = util.IsValidQualifiedIdentifier(t.typ)
		case "argument_list":
			t.args, err = parseArgumentList(nc, parent)
		case ";":
			continue
		case "property_assignments":
			t.props, err = parsePropertyAssignments(nc, &t)
		case "ERROR":
			return nil, fmt.Errorf("line %d: invalid syntax, tree-sitter ERROR", nc.LineNum())
		default:
			panic("should never happen")
		}

		if err != nil {
			return nil, fmt.Errorf(
				"line %d: '%s' type definition: %v", n.LineNum(), t.name, err,
			)
		}
	}

	if len(t.args) > 0 && util.IsBaseType(t.typ) {
		return nil,
			fmt.Errorf("line %d: base type '%s' does not accept argument list", n.LineNum(), t.typ)
	}

	if util.IsBaseType(t.typ) {
		for _, p := range t.props {
			if err = util.IsValidProperty(p.Name, t.typ); err != nil {
				return nil,
					fmt.Errorf(
						"line %d: '%s' type definition: line %d: %v",
						n.LineNum(), t.name, p.LineNum, err,
					)
			}
		}
	}

	return []Symbol{&t}, nil
}

func parseMultiConstantDefinition(n ts.Node) ([]Symbol, error) {
	var symbols []Symbol

	var c *Const

	var doc comment

	for i := 0; i < int(n.ChildCount()); i++ {
		child := n.Child(i)

		switch child.Type() {
		case "const":
			continue
		case "identifier":
			if c != nil {
				symbols = append(symbols, c)
			}

			c = &Const{
				base: base{
					lineNum: child.LineNum(),
					name:    child.Content(),
				},
			}
			if c.lineNum == doc.endLineNum+1 {
				c.doc = doc.msg
			}
			doc = emptyComment()
		case "comment":
			if c == nil || child.LineNum() != c.lineNum {
				if doc.isEmpty() {
					doc = makeComment(child.Content(), child.LineNum())
				} else {
					doc.append(child.Content())
				}
			}
		case "primary_expression", "expression_list":
			expr, err := MakeExpr(child, c)
			if err != nil {
				return nil, fmt.Errorf("line %d: constant %s: %v", c.LineNum(), c.name, err)
			}

			c.Value = expr
		}
	}

	symbols = append(symbols, c)

	return symbols, nil
}

func parsePropertyAssignments(n ts.Node, element Searchable) (PropContainer, error) {
	props := PropContainer{}

	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		switch nc.Type() {
		case "identifier":
			name := nc.Content()
			if _, ok := props.Get(name); ok {
				return props,
					fmt.Errorf("line %d: property '%s' assigned at least twice in the same element body", nc.LineNum(), name)
			}
			expr, err := MakeExpr(n.Child(i+2), element)
			if err != nil {
				return props,
					fmt.Errorf("line %d: '%s' property assignment: %v", nc.LineNum(), name, err)
			}
			props.Add(Prop{LineNum: nc.LineNum(), Name: name, Value: expr})
		default:
			continue
		}
	}

	return props, nil
}

func parseParameterList(n ts.Node, parent Searchable) ([]Param, error) {
	params := []Param{}

	var err error
	var name string
	var hasDfltValue bool
	var dfltValue Expr

	for i := 0; uint32(i) < n.ChildCount(); i++ {
		nc := n.Child(i)
		t := nc.Type()

		if t == "(" || t == "=" || t == "," || t == ")" {
			continue
		}

		hasDfltValue = false

		if t == "identifier" {
			name = nc.Content()
		} else {
			dfltValue, err = MakeExpr(nc, parent)
			if err != nil {
				return nil, fmt.Errorf("parameter list: %v", err)
			}

			hasDfltValue = true
		}

		nextNodeType := n.Child(i + 1).Type()
		if nextNodeType == "," || nextNodeType == ")" {
			for i := range params {
				if name == params[i].Name {
					return nil, fmt.Errorf("parameter '%s' defined at least twice", name)
				}
			}
			params = append(
				params,
				Param{Name: name, HasDfltValue: hasDfltValue, DfltValue: dfltValue},
			)
		}
	}

	// Check if parameters without default value precede parameters with default value.
	withDflt := false
	for i, p := range params {
		if withDflt && !p.HasDfltValue {
			return nil, fmt.Errorf("parameters without default value must precede the ones with default value")
		}

		if params[i].HasDfltValue {
			withDflt = true
		}
	}

	return params, nil
}

func parseSingleConstantDefinition(n ts.Node) ([]Symbol, error) {
	c := &Const{
		base: base{
			lineNum: n.LineNum(),
			name:    n.Child(1).Content(),
		},
	}

	v, err := MakeExpr(n.Child(3), c)
	if err != nil {
		return nil, fmt.Errorf("line %d: single constant definition: %v", n.LineNum(), err)
	}

	c.Value = v

	return []Symbol{c}, nil
}

func parseSingleImportStatement(n ts.Node) Import {
	var path string
	var name string

	if n.ChildCount() == 2 {
		path = n.Child(1).Content()
		path = path[1 : len(path)-1]
		name = strings.Split(path, "/")[0]
		if len(name) > 4 && name[0:3] == "fbd-" {
			name = name[4:]
		}
	} else {
		path = n.Child(2).Content()
		path = path[1 : len(path)-1]
		name = n.Child(1).Content()
	}

	return Import{Path: path, ImportName: name}
}
