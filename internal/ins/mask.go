package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type maskDiary struct {
	atomicSet bool
	dfltSet   bool
	dflt      val.Value
	groupsSet bool
	//rangeSet   bool
	onceSet  bool
	widthSet bool
}

func insMask(typeChain []prs.Element) (*elem.Mask, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	mask := elem.Mask{}
	mask.SetElem(e)

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

	if diary.dfltSet {
		dflt, err := processDefault(mask.Width(), diary.dflt)
		if err != nil {
			return &mask, err
		}
		mask.SetDefault(fbdlVal.MakeBitStr(dflt))
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
			mask.SetAtomic(bool(v.(val.Bool)))
			diary.atomicSet = true
		case "default":
			panic("not yet implemented")
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, "groups")
			}
			mask.SetGroups(makeGroupList(v))
			diary.groupsSet = true
		case "once":
			if diary.onceSet {
				return fmt.Errorf(propAlreadySetMsg, "once")
			}
			mask.SetOnce(bool(v.(val.Bool)))
			diary.onceSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			mask.SetWidth(int64(v.(val.Int)))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillMaskProps(mask *elem.Mask, diary maskDiary) {
	if !diary.atomicSet {
		mask.SetAtomic(true)
	}
	if !diary.widthSet {
		mask.SetWidth(busWidth)
	}
}
