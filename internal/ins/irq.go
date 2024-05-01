package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type irqDiary struct {
	addEnableSet      bool
	clearSet          bool
	enableInitValSet  bool
	enableInitVal     val.Value
	groupsSet         bool
	enableResetValSet bool
	enableResetVal    val.Value
	inTriggerSet      bool
	outTriggerSet     bool
}

func insIrq(typeChain []prs.Functionality) (*fn.Irq, error) {
	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	irq := fn.Irq{}
	irq.Func = f

	diary := irqDiary{}

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyIrqType(&irq, typ, &diary)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillIrqProps(&irq, diary)
	err = fillIrqValues(&irq, diary)
	if err != nil {
		return nil, err
	}

	return &irq, nil
}

func applyIrqType(irq *fn.Irq, typ prs.Functionality, diary *irqDiary) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "irq"); err != nil {
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
		case "add-enable":
			if diary.addEnableSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "add-enable")
			}
			irq.AddEnable = (bool(v.(val.Bool)))
			diary.addEnableSet = true
		case "clear":
			if diary.clearSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "clear")
			}
			irq.Clear = (string(v.(val.Str)))
			diary.clearSet = true
		case "enable-init-value":
			if diary.enableInitValSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "enable-init-value")
			}
			diary.enableInitVal = v
			diary.enableInitValSet = true
		case "enable-reset-value":
			if diary.enableResetValSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "enable-reset-value")
			}
			diary.enableResetVal = v
			diary.enableResetValSet = true
		case "groups":
			if diary.groupsSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "groups")
			}
			irq.Groups = makeGroupList(v)
			diary.groupsSet = true
		case "in-trigger":
			if diary.inTriggerSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "in-trigger")
			}
			irq.InTrigger = (string(v.(val.Str)))
			diary.inTriggerSet = true
		case "out-trigger":
			if diary.outTriggerSet {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "out-trigger")
			}
			irq.OutTrigger = (string(v.(val.Str)))
			diary.outTriggerSet = true
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	return nil
}

func fillIrqProps(irq *fn.Irq, diary irqDiary) {
	if !diary.clearSet {
		irq.Clear = "Explicit"
	}

	if !diary.inTriggerSet {
		irq.InTrigger = "Level"
	}

	if !diary.outTriggerSet {
		irq.OutTrigger = "Level"
	}
}

func fillIrqValues(irq *fn.Irq, diary irqDiary) error {
	if diary.enableInitValSet {
		if !irq.AddEnable {
			return fmt.Errorf("'enable-init-value' set but 'add-enable' is false")
		}

		val, err := processValue(diary.enableInitVal, 1)
		if err != nil {
			return fmt.Errorf("'enable-init-value': %v", err)
		}
		irq.EnableInitValue = fbdlVal.MakeBitStr(val)
	}

	if diary.enableResetValSet {
		if !irq.AddEnable {
			return fmt.Errorf("'enable-reset-value' set but 'add-enable' is false")
		}

		val, err := processValue(diary.enableResetVal, 1)
		if err != nil {
			return fmt.Errorf("'enable-reset-value': %v", err)
		}
		irq.EnableResetValue = fbdlVal.MakeBitStr(val)
	}

	return nil
}
