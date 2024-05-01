package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
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
		last := typeChain[len(typeChain)-1]
		return nil, tok.Error{
			Msg:  fmt.Sprintf("%v", err),
			Toks: []tok.Token{last.Tok()},
		}
	}

	return &st, nil
}

func applyStaticType(st *fn.Static, typ prs.Functionality, diary *staticDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "static"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(p); err != nil {
			return err
		}

		v, err := p.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}

		switch p.Name {
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "groups")
			}
			st.Groups = makeGroupList(v)
			diary.groupsSet = true
		case "init-value":
			if diary.initValSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "init-value")
			}
			diary.initVal = v
			diary.initValSet = true
		case "read-value":
			if diary.readValSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "read-value")
			}
			diary.readVal = v
			diary.readValSet = true
		case "reset-value":
			if diary.resetValSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "reset-value")
			}
			diary.resetVal = v
			diary.resetValSet = true
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
		return fmt.Errorf("'static' functionality must have 'init-value' property set")
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
