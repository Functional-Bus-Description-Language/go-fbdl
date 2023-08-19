package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type staticDiary struct {
	initValSet  bool
	initVal     val.Value
	groupsSet   bool
	readValSet  bool
	readVal     val.Value
	resetValSet bool
	resetVal    val.Value
	widthSet    bool
}

func insStatic(typeChain []prs.Functionality) (*fn.Static, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	st := fn.Static{}
	st.Func = f

	diary := staticDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyStaticType(&st, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillStaticProps(&st, diary)
	err = fillStaticValues(&st, diary)
	if err != nil {
		return nil, err
	}

	return &st, nil
}

func applyStaticType(st *fn.Static, typ prs.Functionality, diary *staticDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "static"); err != nil {
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
			st.Groups = makeGroupList(v)
			diary.groupsSet = true
		case "init-value":
			if diary.initValSet {
				return fmt.Errorf(propAlreadySetMsg, "init-value")
			}
			diary.initVal = v
			diary.initValSet = true
		case "read-value":
			if diary.readValSet {
				return fmt.Errorf(propAlreadySetMsg, "read-value")
			}
			diary.readVal = v
			diary.readValSet = true
		case "reset-value":
			if diary.resetValSet {
				return fmt.Errorf(propAlreadySetMsg, "reset-value")
			}
			diary.resetVal = v
			diary.resetValSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			st.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillStaticProps(st *fn.Static, diary staticDiary) {
	if !diary.widthSet {
		st.Width = busWidth
	}
}

func fillStaticValues(st *fn.Static, diary staticDiary) error {
	if diary.initValSet {
		val, err := processValue(diary.initVal, st.Width)
		if err != nil {
			return fmt.Errorf("'init-value': %v", err)
		}
		st.InitValue = fbdlVal.MakeBitStr(val)
	} else {
		return fmt.Errorf("'static' element must have 'init-value' property set")
	}

	if diary.resetValSet {
		val, err := processValue(diary.resetVal, st.Width)
		if err != nil {
			return fmt.Errorf("'reset-value': %v", err)
		}
		st.ResetValue = fbdlVal.MakeBitStr(val)
	}

	if diary.readValSet {
		val, err := processValue(diary.readVal, st.Width)
		if err != nil {
			return fmt.Errorf("'read-value': %v", err)
		}
		st.ReadValue = fbdlVal.MakeBitStr(val)
	}

	return nil
}
