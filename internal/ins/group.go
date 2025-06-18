package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/group"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

type groupDiary struct {
	virtualSet bool
}

func insGroup(typeChain []prs.Functionality) (*fn.Group, error) {
	typeChainStr := fmt.Sprintf("debug: instantiating group, type chain: %s", typeChain[0].Name())
	for i := 1; i < len(typeChain); i++ {
		typeChainStr = fmt.Sprintf("%s -> %s", typeChainStr, typeChain[i].Name())
	}
	log.Print(typeChainStr)

	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	grp := fn.Group{}
	grp.Func = f

	diary := groupDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyGroupType(&grp, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	if group.IsEmpty(grp) {
		return &grp, tok.Error{
			Msg:  fmt.Sprintf("group '%s' is empty", grp.Name),
			Toks: []tok.Token{typeChain[len(typeChain)-1].Tok()},
		}
	}

	err = checkGroup(grp)
	if err != nil {
		return &grp, tok.Error{
			Msg:  fmt.Sprintf("%v", err),
			Toks: []tok.Token{typeChain[len(typeChain)-1].Tok()},
		}
	}

	return &grp, nil
}

func applyGroupType(grp *fn.Group, typ prs.Functionality, diary *groupDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "group"); err != nil {
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
		case "virtual":
			if diary.virtualSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "virtual")
			}

			grp.Virtual = bool(v.(val.Bool))
			diary.virtualSet = true
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	for _, sym := range typ.Symbols() {
		if c, ok := sym.(*prs.Const); ok {
			if constContainer.HasConst(grp.Consts, c.Name()) {
				return fmt.Errorf(
					"const '%s' is already defined in one of ancestor types", c.Name(),
				)
			}

			val, err := c.Value.Eval()
			if err != nil {
				return fmt.Errorf(
					"cannot evaluate expression for const '%s': %v", c.Name(), err,
				)
			}
			constContainer.AddConst(&grp.Consts, c.Name(), val)
		}

		_, ok := sym.(*prs.Inst)
		if !ok {
			continue
		}

		f := insFunctionality(sym.(prs.Functionality))

		if !util.IsValidInnerType(f.Type(), "group") {
			return fmt.Errorf(
				invalidInnerTypeMsg, f.GetName(), f.Type(), "group",
			)
		}

		if group.HasFunctionality(grp, f.GetName()) {
			return fmt.Errorf(funcWithNameAlreadyInstMsg, f.GetName())
		}

		err := addGroupInnerElement(grp, f)
		if err != nil {
			return tok.Error{
				Msg:  fmt.Sprintf("%v", err),
				Toks: []tok.Token{sym.Tok()},
			}
		}
	}

	return nil
}

func addGroupInnerElement(grp *fn.Group, f any) error {
	var err error

	switch f := f.(type) {
	case (*fn.Config):
		err = group.AddConfig(grp, f)
	case (*fn.Irq):
		err = group.AddIrq(grp, f)
	case (*fn.Mask):
		err = group.AddMask(grp, f)
	case (*fn.Param):
		err = group.AddParam(grp, f)
	case (*fn.Return):
		err = group.AddReturn(grp, f)
	case (*fn.Static):
		err = group.AddStatic(grp, f)
	case (*fn.Status):
		err = group.AddStatus(grp, f)
	default:
		panic("should never happen")
	}

	return err
}

func checkGroup(grp fn.Group) error {
	if len(grp.Irqs) > 0 {
		return checkIrqGroup(grp)
	}
	return nil
}

func checkIrqGroup(grp fn.Group) error {
	// Make sure all irqs within the group have the same out-trigger.
	irq0 := grp.Irqs[0]
	for _, irq := range grp.Irqs[1:] {
		if irq0.OutTrigger != irq.OutTrigger {
			return fmt.Errorf(
				"mismatched output trigger within irq group '%s'\n"+
					"all irqs within irq group must have the same 'out-trigger' property value\n"+
					"irq '%s' 'out-trigger' = %q\n"+
					"irq '%s' 'out-trigger' = %q",
				grp.Name, irq0.Name, irq0.OutTrigger, irq.Name, irq.OutTrigger,
			)
		}
	}

	return nil
}
