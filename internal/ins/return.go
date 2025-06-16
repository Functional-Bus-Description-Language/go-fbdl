package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

type returnDiary struct {
	widthSet bool
}

func insReturn(typeChain []prs.Functionality) (*fn.Return, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	ret := fn.Return{}
	ret.Func = f

	diary := returnDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyReturnType(&ret, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillReturnProps(&ret, diary)

	return &ret, nil
}

func applyReturnType(ret *fn.Return, typ prs.Functionality, diary *returnDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "return"); err != nil {
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
		case "width":
			if diary.widthSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "width")
			}
			ret.Width = int64(v.(val.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	return nil
}

func fillReturnProps(ret *fn.Return, diary returnDiary) {
	if !diary.widthSet {
		ret.Width = busWidth
	}
}
