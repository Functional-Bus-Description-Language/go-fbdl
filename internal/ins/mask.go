package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type maskDiary struct {
	atomicSet   bool
	initValSet  bool
	initVal     val.Value
	readValSet  bool
	readVal     val.Value
	resetValSet bool
	resetVal    val.Value
	widthSet    bool
}

func insMask(typeChain []prs.Functionality) (*fn.Mask, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	mask := fn.Mask{}
	mask.Func = f

	diary := maskDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyMaskType(&mask, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillMaskProps(&mask, diary)
	err = fillMaskValues(&mask, diary)
	if err != nil {
		return nil, err
	}

	return &mask, nil
}

func applyMaskType(mask *fn.Mask, typ prs.Functionality, diary *maskDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "mask"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(p); err != nil {
			return err
		}

		v, err := p.Value.Eval()
		if err != nil {
			return err
		}

		switch p.Name {
		case "atomic":
			if diary.atomicSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "atomic")
			}
			mask.Atomic = bool(v.(val.Bool))
			diary.atomicSet = true
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
			mask.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillMaskProps(mask *fn.Mask, diary maskDiary) {
	if !diary.atomicSet {
		mask.Atomic = true
	}
	if !diary.widthSet {
		mask.Width = busWidth
	}
}

func fillMaskValues(mask *fn.Mask, diary maskDiary) error {
	if diary.initValSet {
		val, err := processValue(diary.initVal, mask.Width)
		if err != nil {
			return fmt.Errorf("'init-value': %v", err)
		}
		mask.InitValue = fbdlVal.MakeBitStr(val)
	}

	if diary.resetValSet {
		val, err := processValue(diary.resetVal, mask.Width)
		if err != nil {
			return fmt.Errorf("'reset-value': %v", err)
		}
		mask.ResetValue = fbdlVal.MakeBitStr(val)
	}

	if diary.readValSet {
		val, err := processValue(diary.readVal, mask.Width)
		if err != nil {
			return fmt.Errorf("'read-value': %v", err)
		}
		mask.ReadValue = fbdlVal.MakeBitStr(val)
	}

	return nil
}
