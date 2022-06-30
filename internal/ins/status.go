package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

type statusAlreadySet struct {
	atomic bool
	groups bool
	once   bool
	width  bool
}

func insStatus(typeChain []prs.Element) (*elem.Status, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	st := elem.Status{
		Elem: e,
	}

	alreadySet := statusAlreadySet{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyStatusType(&st, typ, &alreadySet)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillStatusProps(&st, alreadySet)

	return &st, nil
}

func applyStatusType(st *elem.Status, typ prs.Element, alreadySet *statusAlreadySet) error {
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
			if alreadySet.atomic {
				return fmt.Errorf(propAlreadySetMsg, "atomic")
			}
			st.Atomic = bool(v.(val.Bool))
			alreadySet.atomic = true
		case "groups":
			grps := v.(val.List)
			st.Groups = make([]string, 0, len(grps))
			for _, g := range v.(val.List) {
				st.Groups = append(st.Groups, string(g.(val.Str)))
			}
		case "once":
			if alreadySet.once {
				return fmt.Errorf(propAlreadySetMsg, "once")
			}
			st.Atomic = bool(v.(val.Bool))
			alreadySet.once = true
		case "width":
			if alreadySet.width {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			st.Width = int64(v.(val.Int))
			alreadySet.width = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillStatusProps(st *elem.Status, alreadySet statusAlreadySet) {
	if !alreadySet.atomic {
		st.Atomic = true
	}
	if !alreadySet.width {
		st.Width = busWidth
	}
}
