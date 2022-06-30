package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type statusAlreadySet struct {
	atomic bool
	groups bool
	once   bool
	width  bool
}

func insStatus(typeChain []prs.Element) (*elem.Status, error) {
	inst := typeChain[len(typeChain)-1].(*prs.Inst)

	st := elem.Status{}
	st.SetName(inst.Name())
	st.SetDoc(inst.Doc())
	st.SetIsArray(false)
	st.SetCount(1)

	if inst.IsArray {
		st.SetIsArray(true)
		v, err := inst.Count.Eval()

		if v.Type() != "integer" {
			return nil, fmt.Errorf("size of array must be of 'integer' type, current type '%s'", v.Type())
		}

		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		st.SetCount(int64(v.(val.Int)))
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
			st.SetAtomic(bool(v.(val.Bool)))
			alreadySet.atomic = true
		case "groups":
			vGrps := v.(val.List)
			grps := make([]string, 0, len(vGrps))
			for _, g := range vGrps {
				grps = append(grps, string(g.(val.Str)))
			}
			st.SetGroups(grps)
		case "once":
			if alreadySet.once {
				return fmt.Errorf(propAlreadySetMsg, "once")
			}
			st.SetOnce(bool(v.(val.Bool)))
			alreadySet.once = true
		case "width":
			if alreadySet.width {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			st.SetWidth(int64(v.(val.Int)))
			alreadySet.width = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillStatusProps(st *elem.Status, alreadySet statusAlreadySet) {
	if !alreadySet.atomic {
		st.SetAtomic(true)
	}
	if !alreadySet.width {
		st.SetWidth(busWidth)
	}
}
