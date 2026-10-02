package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parser"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func insBlock(typeChain []parser.Functionality) (*fn.Block, error) {
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

func applyBlockType(blk *fn.Block, typ parser.Functionality) error {
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
		case "align":
			if blk.Align != 0 {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "align")
			}
			blk.Align = int64(v.(types.Int))
		case "masters":
			if blk.Masters != 0 {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "masters")
			}
			blk.Masters = int64(v.(types.Int))
		case "reset":
			if blk.Reset != "" {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "reset")
			}
			blk.Reset = string(v.(types.Str))
		case "width":
			if blk.Width != 0 {
				return fmt.Errorf(propAlreadySetMsg, p.Loc(), "width")
			}
			width := int64(v.(types.Int))
			blk.Width = width
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	for _, s := range typ.Symbols() {
		if c, ok := s.(*parser.Const); ok {
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

		_, ok := s.(*parser.Inst)
		if !ok {
			continue
		}

		f := insFunctionality(s.(parser.Functionality))

		if !util.IsValidInnerType(f.Type(), "block") {
			return fmt.Errorf(
				invalidInnerTypeMsg, f.GetName(), f.Type(), "block",
			)
		}

		if blk.HasFunctionality(f.GetName()) {
			return fmt.Errorf(funcWithNameAlreadyInstMsg, f.GetName())
		}
		addBlockInnerFunc(blk, f)
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

func addBlockInnerFunc(blk *fn.Block, f fn.Functionality) {
	switch f := f.(type) {
	case (*fn.Blackbox):
		blk.Blackboxes = append(blk.Blackboxes, f)
	case (*fn.Config):
		blk.Configs = append(blk.Configs, f)
	case (*fn.Group):
		blk.Groups = append(blk.Groups, f)
	case (*fn.Irq):
		blk.Irqs = append(blk.Irqs, f)
	case (*fn.Mask):
		blk.Masks = append(blk.Masks, f)
	case (*fn.Proc):
		blk.Procs = append(blk.Procs, f)
	case (*fn.Static):
		blk.Statics = append(blk.Statics, f)
	case (*fn.Status):
		blk.Statuses = append(blk.Statuses, f)
	case (*fn.Stream):
		blk.Streams = append(blk.Streams, f)
	case (*fn.Block):
		blk.Subblocks = append(blk.Subblocks, f)
	default:
		panic("should never happen")
	}
}
