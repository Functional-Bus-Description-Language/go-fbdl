package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

const propAlreadySetMsg string = "cannot set property '%s', property is already set in one of ancestor types"
const invalidInnerTypeMsg string = "element '%s' of base type '%s' cannot be instantiated in element of base type '%s'"
const elemWithNameAlreadyInstMsg string = "cannot instantiate element '%s', element with such name is already instantiated in one of ancestor types"

func makeElem(typeChain []prs.Element) (elem.Elem, error) {
	// Instantiation is always the last one in the type chain.
	inst := typeChain[len(typeChain)-1]

	name := inst.Name()
	doc := inst.Doc()
	isArray := false
	count := int64(1)

	if inst.IsArray() {
		isArray = true
		v, err := inst.Count().Eval()

		if v.Type() != "integer" {
			return elem.Elem{}, fmt.Errorf("size of array must be of 'integer' type, current type '%s'", v.Type())
		}

		if err != nil {
			return elem.Elem{}, fmt.Errorf("%v", err)
		}
		count = int64(v.(val.Int))
		if count < 0 {
			return elem.Elem{}, fmt.Errorf("negative array size (%d)", count)
		}
	}

	e := elem.Elem{
		Name:    name,
		Doc:     doc,
		IsArray: isArray,
		Count:   count,
	}

	return e, nil
}

func makeGroupList(propVal val.Value) []string {
	var grps []string
	switch v := propVal.(type) {
	case val.Str:
		grps = []string{string(v)}
	case val.List:
		grps = make([]string, 0, len(v))
		for _, g := range v {
			grps = append(grps, string(g.(val.Str)))
		}
	default:
		panic("should never happen")
	}

	return grps
}

// processDefault processes the 'default' property.
// If the value is BitStr, it checks whether its width is not greater than the width.
// If the value is Int, it tries to convert it to the BitStr with width of width argument.
func processDefault(width int64, v val.Value) (val.BitStr, error) {
	dflt := val.BitStr("")

	if bs, ok := v.(val.BitStr); ok {
		if bs.BitWidth() > width {
			return dflt, fmt.Errorf(
				"width of 'default' bit string (%d) is greater than value of 'width' property (%d)",
				bs.BitWidth(), width,
			)
		}
		dflt = bs
	}
	if i, ok := v.(val.Int); ok {
		bs, err := val.BitStrFromInt(i, width)
		if err != nil {
			return dflt, fmt.Errorf("processing 'default' property: %v", err)
		}
		dflt = bs
	}

	return dflt, nil
}
