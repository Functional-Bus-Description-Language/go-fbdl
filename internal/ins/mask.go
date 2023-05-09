package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type maskDiary struct {
	atomicSet  bool
	initValSet bool
	initVal    val.Value
	groupsSet  bool
	//rangeSet   bool
	widthSet bool
}

func insMask(typeChain []prs.Element) (*elem.Mask, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	mask := elem.Mask{}
	mask.Elem = e

	diary := maskDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyMaskType(&mask, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillMaskProps(&mask, diary)

	if diary.initValSet {
		initVal, err := processValue(diary.initVal, mask.Width)
		if err != nil {
			return &mask, err
		}
		mask.InitValue = fbdlVal.MakeBitStr(initVal)
	}

	return &mask, nil
}

func applyMaskType(mask *elem.Mask, typ prs.Element, diary *maskDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "mask"); err != nil {
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
		case "atomic":
			if diary.atomicSet {
				return fmt.Errorf(propAlreadySetMsg, "atomic")
			}
			mask.Atomic = bool(v.(val.Bool))
			diary.atomicSet = true
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, "groups")
			}
			mask.Groups = makeGroupList(v)
			diary.groupsSet = true
		case "init-value":
			panic("not yet implemented")
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			mask.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillMaskProps(mask *elem.Mask, diary maskDiary) {
	if !diary.atomicSet {
		mask.Atomic = true
	}
	if !diary.widthSet {
		mask.Width = busWidth
	}
}
