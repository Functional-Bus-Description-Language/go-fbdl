package ins

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/adler32"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"

	"sort"
)

// If you add field to the Element struct, remember to include this field
// in the hash() method.
type Element struct {
	Name    string
	Doc     string
	Type    string // ins.Element type can only be base type
	IsArray bool
	Count   int64
	Props   map[string]val.Value
	Consts  map[string]val.Value
	Elems   ElemContainer
	Grps    []*Group

	hashCacheValid bool
	hashCache      uint32
}

func (elem *Element) applyType(typ prs.Element, resolvedArgs map[string]prs.Expr) error {
	if elem.Type == "" {
		if !util.IsBaseType(typ.Type()) {
			return fmt.Errorf("cannot start element instantiation from non base type '%s'", typ.Type())
		}

		elem.Type = typ.Type()
	}

	if i, ok := typ.(*prs.Inst); ok {
		elem.Name = i.Name()
	}

	if resolvedArgs != nil {
		typ.SetResolvedArgs(resolvedArgs)
	}

	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, elem.Type); err != nil {
			return fmt.Errorf(": %v", err)
		}
		err := checkProp(prop.Name, prop)
		if err != nil {
			return fmt.Errorf("\n  %s: line %d: %v", typ.File().Path, prop.LineNum, err)
		}
		if _, exist := elem.Props[prop.Name]; exist {
			return fmt.Errorf(
				"cannot set property '%s', property is already set in one of ancestor types",
				prop.Name,
			)
		}
		v, err := prop.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}
		err = checkPropConflict(elem, prop.Name)
		if err != nil {
			return fmt.Errorf("line %d: %v", prop.LineNum, err)
		}
		elem.Props[prop.Name] = v
	}

	for _, s := range typ.Symbols() {
		if c, ok := s.(*prs.Const); ok {
			if _, has := elem.Consts[c.Name()]; has {
				return fmt.Errorf(
					"const '%s' is already defined in one of ancestor types", c.Name(),
				)
			}

			val, err := c.Value.Eval()
			if err != nil {
				return fmt.Errorf(
					"cannot evaluate expression for const '%s': %v", c.Name(), err,
				)
			}
			elem.Consts[c.Name()] = val
		}

		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := instantiateElement(pe)

		if util.IsValidInnerType(e.Type, elem.Type) == false {
			return fmt.Errorf(
				"element '%s' of base type '%s' cannot be instantiated in element of base type '%s'",
				e.Name, e.Type, elem.Type,
			)
		}

		err := checkElemConflict(elem, e)
		if err != nil {
			return fmt.Errorf("line %d: cannot instantiate element '%s': %v", pe.LineNum(), e.Name, err)
		}

		if !elem.Elems.Add(e) {
			return fmt.Errorf(
				"cannot instantiate element '%s', element with such name is already instantiated in one of ancestor types",
				e.Name,
			)
		}
	}

	if inst, ok := typ.(*prs.Inst); ok {
		if elem.IsArray {
			panic("should never happen")
		}
		if inst.IsArray {
			elem.IsArray = true
			count, err := inst.Count.Eval()

			if count.Type() != "integer" {
				return fmt.Errorf("size of array must be of 'integer' type, current type '%s'", count.Type())
			}

			if err != nil {
				return fmt.Errorf("applying type '%s': %v", typ.Name(), err)
			}
			elem.Count = int64(count.(val.Int))
		} else {
			elem.Count = int64(1)
		}
	}

	return nil
}

func (elem *Element) makeGrps() error {
	elemsWithGrps := []*Element{}

	for _, e := range elem.Elems {
		if _, ok := e.Props["groups"]; ok {
			elemsWithGrps = append(elemsWithGrps, e)
		}
	}

	if len(elemsWithGrps) == 0 {
		return nil
	}

	groups := make(map[string][]*Element)

	for _, e := range elemsWithGrps {
		grps := e.Props["groups"].(val.List)
		for _, g := range grps {
			g := string(g.(val.Str))
			if _, ok := groups[g]; !ok {
				groups[g] = []*Element{}
			}
			groups[g] = append(groups[g], e)
		}
	}

	// Check for element and group names conflict.
	for grpName, _ := range groups {
		if _, ok := elem.Elems.Get(grpName); ok {
			return fmt.Errorf("invalid group name %q, there is inner element with the same name", grpName)
		}
	}

	// Check for groups with single element.
	for name, g := range groups {
		if len(g) == 1 {
			return fmt.Errorf("group %q has only one element '%s'", name, g[0].Name)
		}
	}

	// Check groups order.
	for i, e1 := range elemsWithGrps[:len(elemsWithGrps)-1] {
		grps1 := e1.Props["groups"].(val.List)
		for _, e2 := range elemsWithGrps[i+1:] {
			grps2 := e2.Props["groups"].(val.List)
			indexes := []int{}
			for _, g1 := range grps1 {
				for j2, g2 := range grps2 {
					if string(g1.(val.Str)) == string(g2.(val.Str)) {
						indexes = append(indexes, j2)
					}
				}
			}

			prevId := -1
			for _, id := range indexes {
				if id <= prevId {
					return fmt.Errorf(
						"conflicting order of groups, "+
							"group %q is after group %q in element '%s', "+
							"but before group %q in element '%s'",
						string(grps2[id].(val.Str)),
						string(grps2[id+1].(val.Str)),
						e1.Name,
						string(grps2[id+1].(val.Str)),
						e2.Name,
					)
				}
				prevId = id
			}
		}
	}

	var grpsOrder []string

	if _, ok := elem.Props["groupsOrder"]; ok {
		panic("not yet implemented")
	} else {
		graph := newGrpGraph()

		for _, e := range elemsWithGrps {
			grps := e.Props["groups"].(val.List)
			var prevGrp string = ""
			for _, g := range grps {
				g := string(g.(val.Str))
				graph.addEdge(prevGrp, g)
				prevGrp = g
			}
		}

		grpsOrder = graph.sort()
	}

	log.Printf("debug: groups order for element '%s': %v", elem.Name, grpsOrder)

	elem.Grps = []*Group{}
	for _, grpName := range grpsOrder {
		elem.Grps = append(
			elem.Grps,
			&Group{
				Name:  grpName,
				Elems: groups[grpName],
			},
		)
	}

	return nil
}

// processDflt processes the 'default' property.
// If element has no 'default' property it immediately returns.
// Otherwise it checks the type of 'default' value.
// If the value is BitStr, it checks whether its width is not greater than value of 'width' property.
// If the value is Int, it tries to convert it to BitStr with width of 'width' property value.
func (e *Element) processDflt() error {
	dflt, ok := e.Props["default"]

	if !ok {
		return nil
	}

	width := int64(e.Props["width"].(val.Int))

	if bs, ok := dflt.(val.BitStr); ok {
		if bs.BitWidth() > width {
			return fmt.Errorf(
				"width of 'default' bit string (%d) is greater than value of 'width' property (%d)",
				bs.BitWidth(), width,
			)
		}
	}
	if i, ok := dflt.(val.Int); ok {
		bs, err := val.BitStrFromInt(i, width)
		if err != nil {
			return fmt.Errorf("processing 'default' property: %v", err)
		}
		e.Props["default"] = bs
	}

	return nil
}

func (e *Element) hash() uint32 {
	if e.hashCacheValid {
		return e.hashCache
	}

	b := bytes.Buffer{}

	b.Write([]byte(e.Name))
	b.Write([]byte(e.Doc))
	b.Write([]byte(e.Type))
	if e.IsArray {
		b.WriteByte(1)
	} else {
		b.WriteByte(0)
	}
	aux := make([]byte, 8)
	binary.LittleEndian.PutUint64(aux, uint64(e.Count))
	b.Write(aux)

	// Properties
	props := make([]string, 0, len(e.Props))
	for name := range e.Props {
		props = append(props, name)
	}
	sort.Strings(props)
	for _, name := range props {
		b.Write([]byte(name))
		b.Write(e.Props[name].Bytes())
	}

	// Constants
	consts := make([]string, 0, len(e.Consts))
	for name := range e.Consts {
		consts = append(consts, name)
	}
	sort.Strings(consts)
	for _, name := range consts {
		b.Write([]byte(name))
		b.Write(e.Consts[name].Bytes())
	}

	// Inner Elements
	aux = make([]byte, 4)
	for _, ee := range e.Elems {
		binary.LittleEndian.PutUint32(aux, ee.hash())
		b.Write(aux)
	}

	// Groups
	for _, g := range e.Grps {
		binary.LittleEndian.PutUint32(aux, g.hash())
		b.Write(aux)
	}

	e.hashCache = adler32.Checksum(b.Bytes())
	e.hashCacheValid = true

	return e.hashCache
}
