package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

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
	}

	return elem.MakeElem(name, doc, isArray, count), nil
}
