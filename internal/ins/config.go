package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parser"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

type configDiary struct {
	atomicSet bool
	rangeSet  bool
	widthSet  bool
	initVal   types.Value
	readVal   types.Value
	resetVal  types.Value
}

func insConfig(typeChain []parser.Functionality) (*fn.Config, error) {
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

func applyConfigType(cfg *fn.Config, typ parser.Functionality, diary *configDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "config"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(prop); err != nil {
			return err
		}

		val, err := prop.Value.Eval()
		if err != nil {
			return err
		}

		switch prop.Name {
		case "atomic":
			if diary.atomicSet {
				return fmt.Errorf(
					propAlreadySetMsg, prop.Loc(), "atomic",
				)
			}
			cfg.Atomic = (bool(val.(types.Bool)))
			diary.atomicSet = true
		case "init-value":
			if diary.initVal != nil {
				return fmt.Errorf(
					propAlreadySetMsg, prop.Loc(), "init-value",
				)
			}
			diary.initVal = val
		case "range":
			if diary.rangeSet {
				return fmt.Errorf(
					propAlreadySetMsg, prop.Loc(), "range",
				)
			}
			if diary.widthSet {
				return fmt.Errorf(
					propConflictMsg, prop.Loc(), "range", "width",
				)
			}

			switch rng := val.(type) {
			case types.Int:
				cfg.Range = types.SingleRange{Start: 0, End: int64(rng)}
			case types.SingleRange:
				cfg.Range = rng
			case types.List:
				mr := types.ArrayRange{}
				for _, r := range rng {
					mr = append(mr, r.(types.SingleRange))
				}
				cfg.Range = mr
			}
			diary.rangeSet = true
		case "read-value":
			if diary.readVal != nil {
				return fmt.Errorf(
					propAlreadySetMsg, prop.Loc(), "read-value",
				)
			}
			diary.readVal = val
		case "reset-value":
			if diary.resetVal != nil {
				return fmt.Errorf(
					propAlreadySetMsg, prop.Loc(), "reset-value",
				)
			}
			diary.resetVal = val
		case "width":
			if diary.widthSet {
				return fmt.Errorf(
					propAlreadySetMsg, prop.Loc(), "width",
				)
			}
			if diary.rangeSet {
				return fmt.Errorf(
					propConflictMsg, prop.Loc(), "width", "range",
				)
			}
			cfg.Width = int64(val.(types.Int))
			diary.widthSet = true
		default:
			panic(fmt.Sprintf("unhandled '%s' property", prop.Name))
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
			cfg.Width = cfg.Range.BitWidth()
		}
	}
}

func fillConfigValues(cfg *fn.Config, diary configDiary) error {
	if diary.initVal != nil {
		val, err := processValue(diary.initVal, cfg.Width)
		if err != nil {
			return fmt.Errorf("'init-value': %v", err)
		}
		cfg.InitValue = val
	}

	if diary.resetVal != nil {
		val, err := processValue(diary.resetVal, cfg.Width)
		if err != nil {
			return fmt.Errorf("'reset-value': %v", err)
		}
		cfg.ResetValue = val
	}

	if diary.readVal != nil {
		val, err := processValue(diary.readVal, cfg.Width)
		if err != nil {
			return fmt.Errorf("'read-value': %v", err)
		}
		cfg.ReadValue = val
	}

	return nil
}
