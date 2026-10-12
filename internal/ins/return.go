package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parser"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

type returnDiary struct {
	widthSet bool
}

func insReturn(typeChain []parser.Functionality) (*fn.Return, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, err
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
			return nil, err
		}
	}

	fillReturnProps(&ret, diary)

	return &ret, nil
}

func applyReturnType(ret *fn.Return, typ parser.Functionality, diary *returnDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "return"); err != nil {
			return err
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
			ret.Width = int64(v.(types.Int))
			diary.widthSet = true
		default:
			panic("should never happen")
		}
	}

	for _, sym := range typ.Symbols() {
		pi, ok := sym.(*parser.Inst)
		if !ok {
			continue
		}

		f := insFunctionality(sym.(parser.Functionality))
		err := checkInnerFuncType(f, ret)
		if err != nil {
			return pi.NewError(
				token.Error{
					Msg:  fmt.Sprintf("%v", err),
					Toks: []token.Token{pi.Ast().Name},
				},
			)
		}
	}

	return nil
}

func fillReturnProps(ret *fn.Return, diary returnDiary) {
	if !diary.widthSet {
		ret.Width = busWidth
	}
}
