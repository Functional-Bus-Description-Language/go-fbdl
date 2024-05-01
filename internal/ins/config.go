package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type configDiary struct {
	atomicSet   bool
	initValSet  bool
	initVal     val.Value
	groupsSet   bool
	rangeSet    bool
	readValSet  bool
	readVal     val.Value
	resetValSet bool
	resetVal    val.Value
	widthSet    bool
}

func insConfig(typeChain []prs.Functionality) (*fn.Config, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, err
	}
	cfg := fn.Config{}
	cfg.Func = f

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
	err = fillConfigValues(&cfg, diary)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func applyConfigType(cfg *fn.Config, typ prs.Functionality, diary *configDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "config"); err != nil {
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
		case "atomic":
			if diary.atomicSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "atomic")
			}
			cfg.Atomic = (bool(v.(val.Bool)))
			diary.atomicSet = true
		case "init-value":
			if diary.initValSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "init-value")
			}
			diary.initVal = v
			diary.initValSet = true
		case "range":
			if diary.rangeSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "range")
			}
			if diary.widthSet {
				return fmt.Errorf(propConflictMsg, p.Loc(), "range", "width")
			}
			var rang fbdlVal.Range
			switch r := v.(type) {
			case val.Int:
				rang = []int64{0, int64(r)}
			case val.List:
				for _, bound := range r {
					rang = append(rang, int64(bound.(val.Int)))
				}
			}
			cfg.Range = rang
			diary.rangeSet = true
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "groups")
			}
			cfg.Groups = makeGroupList(v)
			diary.groupsSet = true
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
			if diary.rangeSet {
				return fmt.Errorf(propConflictMsg, p.Loc(), "width", "range")
			}
			cfg.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	return nil
}

func fillConfigProps(cfg *fn.Config, diary configDiary) {
	if !diary.atomicSet {
		cfg.Atomic = true
	}
	if !diary.widthSet {
		if !diary.rangeSet {
			cfg.Width = busWidth
		} else {
			cfg.Width = cfg.Range.Width()
		}
	}
}

func fillConfigValues(cfg *fn.Config, diary configDiary) error {
	if diary.initValSet {
		val, err := processValue(diary.initVal, cfg.Width)
		if err != nil {
			return fmt.Errorf("'init-value': %v", err)
		}
		cfg.InitValue = fbdlVal.MakeBitStr(val)
	}

	if diary.resetValSet {
		val, err := processValue(diary.resetVal, cfg.Width)
		if err != nil {
			return fmt.Errorf("'reset-value': %v", err)
		}
		cfg.ResetValue = fbdlVal.MakeBitStr(val)
	}

	if diary.readValSet {
		val, err := processValue(diary.readVal, cfg.Width)
		if err != nil {
			return fmt.Errorf("'read-value': %v", err)
		}
		cfg.ReadValue = fbdlVal.MakeBitStr(val)
	}

	return nil
}
