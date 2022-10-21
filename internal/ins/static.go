package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type staticDiary struct {
	dfltSet   bool
	dflt      val.Value
	groupsSet bool
	onceSet   bool
	widthSet  bool
}

func insStatic(typeChain []prs.Element) (*elem.Static, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	st := elem.Static{}
	st.SetElem(e)

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

	if diary.dfltSet {
		dflt, err := processDefault(st.Width(), diary.dflt)
		if err != nil {
			return &st, err
		}
		st.SetDefault(fbdlVal.MakeBitStr(dflt))
	}

	return &st, nil
}

func applyStaticType(st *elem.Static, typ prs.Element, diary *staticDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "static"); err != nil {
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
		case "default":
			if diary.dfltSet {
				return fmt.Errorf(propAlreadySetMsg, "default")
			}
			diary.dflt = v
			diary.dfltSet = true
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

func fillStaticProps(st *elem.Static, diary staticDiary) {
	if !diary.widthSet {
		st.SetWidth(busWidth)
	}
}
