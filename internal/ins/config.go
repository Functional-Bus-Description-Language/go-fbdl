package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type configAlreadySet struct {
	atomic bool
	dflt   bool
	groups bool
	rang   bool
	once   bool
	width  bool
}

func insConfig(typeChain []prs.Element) (*elem.Config, error) {
	inst := typeChain[len(typeChain)-1].(*prs.Inst)

	cfg := elem.Config{}
	cfg.SetName(inst.Name())
	cfg.SetDoc(inst.Doc())
	cfg.SetIsArray(false)
	cfg.SetCount(1)

	if inst.IsArray {
		cfg.SetIsArray(true)
		v, err := inst.Count.Eval()

		if v.Type() != "integer" {
			return nil, fmt.Errorf("size of array must be of 'integer' type, current type '%s'", v.Type())
		}

		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		cfg.SetCount(int64(v.(val.Int)))
	}

	alreadySet := configAlreadySet{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyConfigType(&cfg, typ, &alreadySet)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillConfigProps(&cfg, alreadySet)

	return &cfg, nil
}

func applyConfigType(cfg *elem.Config, typ prs.Element, alreadySet *configAlreadySet) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "config"); err != nil {
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
			cfg.SetAtomic(bool(v.(val.Bool)))
			alreadySet.atomic = true
		case "default", "range":
			panic("not yet implemented")
		case "groups":
			vGrps := v.(val.List)
			grps := make([]string, 0, len(vGrps))
			for _, g := range vGrps {
				grps = append(grps, string(g.(val.Str)))
			}
			cfg.SetGroups(grps)
		case "once":
			if alreadySet.once {
				return fmt.Errorf(propAlreadySetMsg, "once")
			}
			cfg.SetOnce(bool(v.(val.Bool)))
			alreadySet.once = true
		case "width":
			if alreadySet.width {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			cfg.SetWidth(int64(v.(val.Int)))
			alreadySet.width = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillConfigProps(cfg *elem.Config, alreadySet configAlreadySet) {
	if !alreadySet.atomic {
		cfg.SetAtomic(true)
	}
	if !alreadySet.width {
		cfg.SetWidth(busWidth)
	}
}
