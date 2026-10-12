package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parser"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
)

type procDiary struct {
	delaySet bool
}

func insProc(typeChain []parser.Functionality) (*fn.Proc, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, err
	}
	proc := fn.Proc{}
	proc.Func = f

	diary := procDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyProcType(&proc, typ, &diary)
		if err != nil {
			return nil, err
		}
	}

	return &proc, nil
}

func applyProcType(p *fn.Proc, typ parser.Functionality, diary *procDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "proc"); err != nil {
			return err
		}
		if err := checkProp(prop); err != nil {
			return err
		}

		v, err := prop.Value.Eval()
		if err != nil {
			return err
		}

		switch prop.Name {
		case "delay":
			if diary.delaySet {
				return fmt.Errorf(propAlreadySetMsg, prop.Loc(), "delay")
			}
			t := v.(types.Time)
			delay := types.Time{S: t.S, Ns: t.Ns}

			p.Delay = &delay
			diary.delaySet = true
		default:
			panic("should never happen")
		}

	}

	for _, sym := range typ.Symbols() {
		pi, ok := sym.(*parser.Inst)
		if !ok {
			continue
		}

		f := insFunctionality(pi)

		if !util.IsValidInnerType(f.Type(), "proc") {
			return pi.NewError(
				token.Error{
					Msg: fmt.Sprintf(
						invalidInnerTypeMsg,
						f.GetName(), f.Type(), p.Name, "proc",
					),
					Toks: []token.Token{pi.Ast().Name},
				},
			)
		}

		if p.HasFunctionality(f.GetName()) {
			return fmt.Errorf(funcWithNameAlreadyInstMsg, f.GetName())
		}
		addProcInnerFunc(p, f)
	}

	return nil
}

func addProcInnerFunc(p *fn.Proc, f fn.Functionality) {
	switch f := f.(type) {
	case (*fn.Param):
		p.Params = append(p.Params, f)
	case (*fn.Return):
		p.Returns = append(p.Returns, f)
	default:
		panic("should never happen")
	}
}
