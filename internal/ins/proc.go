package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/proc"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type procDiary struct {
	delaySet bool
}

func insProc(typeChain []prs.Functionality) (*fn.Proc, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
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
			return nil, fmt.Errorf("%v", err)
		}
	}

	return &proc, nil
}

func applyProcType(p *fn.Proc, typ prs.Functionality, diary *procDiary) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "proc"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(prop); err != nil {
			return fmt.Errorf("%s: %v", prop.Loc(), err)
		}

		v, err := prop.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}

		switch prop.Name {
		case "delay":
			if diary.delaySet {
				return fmt.Errorf(propAlreadySetMsg, prop.Loc(), "delay")
			}
			t := v.(val.Time)
			delay := fbdlVal.Time{S: t.S, Ns: t.Ns}

			p.Delay = &delay
			diary.delaySet = true
		default:
			panic("should never happen")
		}

	}

	for _, s := range typ.Symbols().Symbols() {
		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(pe)

		if !util.IsValidInnerType(e.Type(), "proc") {
			return fmt.Errorf(invalidInnerTypeMsg, e.GetName(), e.Type(), "proc")
		}

		if proc.HasElement(p, e.GetName()) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, e.GetName())
		}
		addProcInnerElement(p, e)
	}

	return nil
}

func addProcInnerElement(p *fn.Proc, e fn.Functionality) {
	switch e := e.(type) {
	case (*fn.Param):
		p.Params = append(p.Params, e)
	case (*fn.Return):
		p.Returns = append(p.Returns, e)
	default:
		panic("should never happen")
	}
}
