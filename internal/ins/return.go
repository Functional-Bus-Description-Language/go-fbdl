package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type returnAlreadySet struct {
	groups bool
	width  bool
}

func insReturn(typeChain []prs.Element) (*elem.Return, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	ret := elem.Return{}
	ret.SetElem(e)

	alreadySet := returnAlreadySet{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyReturnType(&ret, typ, &alreadySet)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	return &ret, nil
}

func applyReturnType(ret *elem.Return, typ prs.Element, alreadySet *returnAlreadySet) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "return"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(prop); err != nil {
			return fmt.Errorf("%s: line %d: %v", typ.File().Path, prop.LineNum, err)
		}

		v, err := prop.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}

		switch prop.Name {
		case "groups":
			vGrps := v.(val.List)
			grps := make([]string, 0, len(vGrps))
			for _, g := range vGrps {
				grps = append(grps, string(g.(val.Str)))
			}
			ret.SetGroups(grps)
		case "width":
			if alreadySet.width {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			ret.SetWidth(int64(v.(val.Int)))
			alreadySet.width = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillReturnProps(ret *elem.Return, alreadySet returnAlreadySet) {
	if !alreadySet.width {
		ret.SetWidth(busWidth)
	}
}
