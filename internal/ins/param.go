package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type paramDiary struct {
	groupsSet bool
	rangeSet  bool
	widthSet  bool
}

func insParam(typeChain []prs.Element) (*elem.Param, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	param := elem.Param{}
	param.Elem = e

	diary := paramDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyParamType(&param, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillParamProps(&param, diary)

	return &param, nil
}

func applyParamType(param *elem.Param, typ prs.Element, diary *paramDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "param"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(prop); err != nil {
			return fmt.Errorf("%s: line %d: %v", typ.File().Path, prop.Line, err)
		}

		v, err := prop.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}

		switch prop.Name {
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, "groups")
			}
			param.Groups = makeGroupList(v)
			diary.groupsSet = true
		case "range":
			if diary.rangeSet {
				return fmt.Errorf(propAlreadySetMsg, "range")
			}
			var rang fbdlVal.Range
			switch r := v.(type) {
			case val.Int:
				rang = []int64{0, int64(r)}
			case val.List:
				for _, bound := range r {
					rang = append(rang, int64(bound.(val.Int)))
				}
			}
			param.Range = rang
			diary.rangeSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			param.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillParamProps(param *elem.Param, diary paramDiary) {
	if !diary.widthSet {
		if !diary.rangeSet {
			param.Width = busWidth
		} else {
			param.Width = param.Range.Width()
		}
	}
}
