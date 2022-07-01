package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

const propAlreadySetMsg string = "cannot set property '%s', property is already set in one of ancestor types"
const invalidInnerTypeMsg string = "element '%s' of base type '%s' cannot be instantiated in element of base type '%s'"
const elemWithNameAlreadyInstMsg string = "cannot instantiate element '%s', element with such name is already instantiated in one of ancestor types"

func makeElem(typeChain []prs.Element) (elem.Elem, error) {
	// Instantiation is always the last one in the type chain.
	inst := typeChain[len(typeChain)-1].(*prs.Inst)

	name := inst.Name()
	doc := inst.Doc()
	isArray := false
	count := int64(1)

	if inst.IsArray {
		isArray = true
		v, err := inst.Count.Eval()

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

	e := elem.Elem{}
	e.SetName(name)
	e.SetDoc(doc)
	e.SetIsArray(isArray)
	e.SetCount(count)

	return e, nil
}

func makeGroupList(v val.Value) []string {
	vGrps := v.(val.List)
	grps := make([]string, 0, len(vGrps))
	for _, g := range vGrps {
		grps = append(grps, string(g.(val.Str)))
	}
	return grps
}
