package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type paramAlreadySet struct {
	atomic bool
	groups bool
	once   bool
	width  bool
}

func insParam(typeChain []prs.Element) (*elem.Param, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	param := elem.Param{}
	param.SetElem(e)

	alreadySet := paramAlreadySet{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyParamType(&param, typ, &alreadySet)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	return &param, nil
}

func applyParamType(param *elem.Param, typ prs.Element, alreadySet *paramAlreadySet) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "param"); err != nil {
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
			param.SetGroups(grps)
		case "range":
			panic("not yet implemented")
		case "width":
			if alreadySet.width {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			param.SetWidth(int64(v.(val.Int)))
			alreadySet.width = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillParamProps(param *elem.Param, alreadySet statusAlreadySet) {
	if !alreadySet.width {
		param.SetWidth(busWidth)
	}
}
