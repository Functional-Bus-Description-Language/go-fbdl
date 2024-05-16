package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type statusDiary struct {
	atomicSet  bool
	groupsSet  bool
	readValSet bool
	readVal    val.Value
	widthSet   bool
}

func insStatus(typeChain []prs.Functionality) (*fn.Status, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	st := fn.Status{}
	st.Func = f

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

	if diary.readValSet {
		val, err := processValue(diary.readVal, st.Width)
		if err != nil {
			return nil, fmt.Errorf("'read-value': %v", err)
		}
		st.ReadValue = fbdlVal.MakeBitStr(val)
	}

	return &st, nil
}

func applyStatusType(st *fn.Status, typ prs.Functionality, diary *statusDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "status"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(p); err != nil {
			return err
		}

		v, err := p.Value.Eval()
		if err != nil {
			return err
		}

		switch p.Name {
		case "atomic":
			if diary.atomicSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "atomic")
			}
			st.Atomic = bool(v.(val.Bool))
			diary.atomicSet = true
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "groups")
			}
			st.Groups = makeGroupList(v)
			diary.groupsSet = true
		case "read-value":
			if diary.readValSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "read-value")
			}
			diary.readVal = v
			diary.readValSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "width")
			}
			st.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillStatusProps(st *fn.Status, diary statusDiary) {
	if !diary.atomicSet {
		st.Atomic = true
	}
	if !diary.widthSet {
		st.Width = busWidth
	}
}
