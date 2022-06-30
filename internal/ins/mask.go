package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

type maskAlreadySet struct {
	atomic bool
	dflt   bool
	groups bool
	rang   bool
	once   bool
	width  bool
}

func insMask(typeChain []prs.Element) (*elem.Mask, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	mask := elem.Mask{
		Elem: e,
	}

	alreadySet := maskAlreadySet{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyMaskType(&mask, typ, &alreadySet)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillMaskProps(&mask, alreadySet)

	return &mask, nil
}

func applyMaskType(mask *elem.Mask, typ prs.Element, alreadySet *maskAlreadySet) error {
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
			if alreadySet.atomic {
				return fmt.Errorf(propAlreadySetMsg, "atomic")
			}
			mask.Atomic = bool(v.(val.Bool))
			alreadySet.atomic = true
		case "default":
			panic("not yet implemented")
		case "groups":
			grps := v.(val.List)
			mask.Groups = make([]string, 0, len(grps))
			for _, g := range v.(val.List) {
				mask.Groups = append(mask.Groups, string(g.(val.Str)))
			}
		case "once":
			if alreadySet.once {
				return fmt.Errorf(propAlreadySetMsg, "once")
			}
			mask.Atomic = bool(v.(val.Bool))
			alreadySet.once = true
		case "width":
			if alreadySet.width {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			mask.Width = int64(v.(val.Int))
			alreadySet.width = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillMaskProps(mask *elem.Mask, alreadySet maskAlreadySet) {
	if !alreadySet.atomic {
		mask.Atomic = true
	}
	if !alreadySet.width {
		mask.Width = busWidth
	}
}
