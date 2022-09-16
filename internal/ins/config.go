package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type configDiary struct {
	atomicSet bool
	dfltSet   bool
	dflt      val.Value
	groupsSet bool
	//rangeSet  bool
	onceSet  bool
	widthSet bool
}

func insConfig(typeChain []prs.Element) (*elem.Config, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, err
	}
	cfg := elem.Config{}
	cfg.SetElem(e)

	diary := configDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyConfigType(&cfg, typ, &diary)
		if err != nil {
			return nil, err
		}
	}

	fillConfigProps(&cfg, diary)

	if diary.dfltSet {
		dflt, err := processDefault(cfg.Width(), diary.dflt)
		if err != nil {
			return &cfg, err
		}
		cfg.SetDefault(fbdlVal.MakeBitStr(dflt))
	}

	return &cfg, nil
}

func applyConfigType(cfg *elem.Config, typ prs.Element, diary *configDiary) error {
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
			if diary.atomicSet {
				return fmt.Errorf(propAlreadySetMsg, "atomic")
			}
			cfg.SetAtomic(bool(v.(val.Bool)))
			diary.atomicSet = true
		case "default":
			if diary.dfltSet {
				return fmt.Errorf(propAlreadySetMsg, "default")
			}
			diary.dflt = v
			diary.dfltSet = true
		case "range":
			panic("not yet implemented")
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, "groups")
			}
			cfg.SetGroups(makeGroupList(v))
			diary.groupsSet = true
		case "once":
			if diary.onceSet {
				return fmt.Errorf(propAlreadySetMsg, "once")
			}
			cfg.SetOnce(bool(v.(val.Bool)))
			diary.onceSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			cfg.SetWidth(int64(v.(val.Int)))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillConfigProps(cfg *elem.Config, diary configDiary) {
	if !diary.atomicSet {
		cfg.SetAtomic(true)
	}
	if !diary.widthSet {
		cfg.SetWidth(busWidth)
	}
}
