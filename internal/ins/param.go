package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/value"
)

type paramDiary struct {
	rangeSet bool
	widthSet bool
}

func insParam(typeChain []prs.Functionality) (*fn.Param, error) {
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

func applyParamType(param *fn.Param, typ prs.Functionality, diary *paramDiary) error {
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
			case val.Int:
				param.Range = fbdlVal.SingleRange{Left: 0, Right: int64(rng)}
			case val.Range:
				param.Range = fbdlVal.SingleRange{Left: rng.L, Right: rng.R}
			case val.List:
				mr := fbdlVal.MultiRange{}
				for _, r := range rng {
					mr = append(
						mr,
						fbdlVal.SingleRange{
							Left:  r.(val.Range).L,
							Right: r.(val.Range).R,
						},
					)
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
			param.Width = int64(v.(val.Int))
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
			param.Width = param.Range.Width()
		}
	}
}
