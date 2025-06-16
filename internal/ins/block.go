package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/block"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func insBlock(typeChain []prs.Functionality) (*fn.Block, error) {
	typeChainStr := fmt.Sprintf("debug: instantiating block, type chain: %s", typeChain[0].Name())
	for i := 1; i < len(typeChain); i++ {
		typeChainStr = fmt.Sprintf("%s -> %s", typeChainStr, typeChain[i].Name())
	}
	log.Print(typeChainStr)

	f, err := makeFunctionality(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	blk := fn.Block{}
	blk.Func = f

	tci := typeChainIter(typeChain)
	for {
		typ, ok := tci()
		if !ok {
			break
		}
		err := applyBlockType(&blk, typ)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillBlockProps(&blk)

	return &blk, nil
}

func applyBlockType(blk *fn.Block, typ prs.Functionality) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "bus"); err != nil {
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
		case "masters":
			if blk.Masters != 0 {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "masters")
			}
			blk.Masters = int64(v.(val.Int))
		case "reset":
			if blk.Reset != "" {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "reset")
			}
			blk.Reset = string(v.(val.Str))
		case "width":
			if blk.Width != 0 {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "width")
			}
			width := int64(v.(val.Int))
			blk.Width = width
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	for _, s := range typ.Symbols() {
		if c, ok := s.(*prs.Const); ok {
			if constContainer.HasConst(blk.Consts, c.Name()) {
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
			constContainer.AddConst(&blk.Consts, c.Name(), val)
		}

		_, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		f := insFunctionality(s.(prs.Functionality))

		if !util.IsValidInnerType(f.Type(), "block") {
			return fmt.Errorf(
				invalidInnerTypeMsg, f.GetName(), f.Type(), "block",
			)
		}

		if block.HasFunctionality(blk, f.GetName()) {
			return fmt.Errorf(funcWithNameAlreadyInstMsg, f.GetName())
		}
		addBlockInnerElement(blk, f)
	}

	return nil
}

func fillBlockProps(blk *fn.Block) {
	if blk.Masters == 0 {
		blk.Masters = 1
	}
	if blk.Width == 0 {
		blk.Width = 32
	}
}

func addBlockInnerElement(blk *fn.Block, f any) {
	switch f := f.(type) {
	case (*fn.Blackbox):
		block.AddBlackbox(blk, f)
	case (*fn.Config):
		block.AddConfig(blk, f)
	case (*fn.Irq):
		block.AddIrq(blk, f)
	case (*fn.Mask):
		block.AddMask(blk, f)
	case (*fn.Memory):
		block.AddMemory(blk, f)
	case (*fn.Proc):
		block.AddProc(blk, f)
	case (*fn.Static):
		block.AddStatic(blk, f)
	case (*fn.Status):
		block.AddStatus(blk, f)
	case (*fn.Stream):
		block.AddStream(blk, f)
	case (*fn.Block):
		block.AddSubblock(blk, f)
	default:
		panic("should never happen")
	}
}
