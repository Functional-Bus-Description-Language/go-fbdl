package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parser"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type paramDiary struct {
	rangeSet bool
	widthSet bool
}

func insParam(typeChain []parser.Functionality) (*fn.Param, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	param := fn.Param{}
	param.Func = f

	diary := paramDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyParamType(&param, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillParamProps(&param, diary)

	return &param, nil
}

func applyParamType(param *fn.Param, typ parser.Functionality, diary *paramDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "param"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(p); err != nil {
			return fmt.Errorf("%s: %v", p.Loc(), err)
		}

		v, err := p.Value.Eval()
		if err != nil {
			return err
		}

		switch p.Name {
		case "range":
			if diary.rangeSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "range")
			}
			if diary.widthSet {
				return fmt.Errorf(propConflictMsg, p.Loc(), "range", "width")
			}

			switch rng := v.(type) {
			case types.Int:
				param.Range = types.SingleRange{Start: 0, End: int64(rng)}
			case types.SingleRange:
				param.Range = rng
			case types.List:
				mr := types.ArrayRange{}
				for _, r := range rng {
					mr = append(mr, r.(types.SingleRange))
				}
				param.Range = mr
			}
			diary.rangeSet = true
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "width")
			}
			if diary.rangeSet {
				return fmt.Errorf(propConflictMsg, p.Loc(), "width", "range")
			}
			param.Width = int64(v.(types.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillParamProps(param *fn.Param, diary paramDiary) {
	if !diary.widthSet {
		if !diary.rangeSet {
			param.Width = busWidth
		} else {
			param.Width = param.Range.BitWidth()
		}
	}
}
