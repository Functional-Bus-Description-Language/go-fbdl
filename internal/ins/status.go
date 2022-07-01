package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type statusDiary struct {
	atomicSet bool
	groupsSet bool
	onceSet   bool
	widthSet  bool
}

func insStatus(typeChain []prs.Element) (*elem.Status, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	st := elem.Status{}
	st.SetElem(e)

	diary := statusDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyStatusType(&st, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillStatusProps(&st, diary)

	return &st, nil
}

func applyStatusType(st *elem.Status, typ prs.Element, diary *statusDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "status"); err != nil {
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
			st.SetAtomic(bool(v.(val.Bool)))
			diary.atomicSet = true
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, "groups")
			}
			st.SetGroups(makeGroupList(v))
			diary.groupsSet = true
		case "once":
			if diary.onceSet {
				return fmt.Errorf(propAlreadySetMsg, "once")
			}
			st.SetOnce(bool(v.(val.Bool)))
			diary.onceSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			st.SetWidth(int64(v.(val.Int)))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillStatusProps(st *elem.Status, diary statusDiary) {
	if !diary.atomicSet {
		st.SetAtomic(true)
	}
	if !diary.widthSet {
		st.SetWidth(busWidth)
	}
}
