package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func insBlock(typeChain []prs.Element) (*elem.Block, error) {
	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	blk := elem.Block{
		Elem: e,
	}

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

func applyBlockType(blk *elem.Block, typ prs.Element) error {
	for _, prop := range typ.Props() {
		if err := util.IsValidProperty(prop.Name, "bus"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(prop); err != nil {
			return fmt.Errorf("%s: line %d: %v", typ.File().Path, prop.LineNum, err)
		}

		v, err := prop.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}

		switch prop.Name {
		case "masters":
			if blk.Masters != 0 {
				return fmt.Errorf(propAlreadySetMsg, "masters")
			}
			blk.Masters = int64(v.(val.Int))
		case "width":
			if blk.Width != 0 {
				return fmt.Errorf(propAlreadySetMsg, "width")
			}
			blk.Width = int64(v.(val.Int))
		default:
			panic("should never happen")
		}
	}

	for _, s := range typ.Symbols() {
		if c, ok := s.(*prs.Const); ok {
			if blk.HasConst(c.Name()) {
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
			blk.AddConst(c.Name(), val)
		}

		pe, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(pe)

		if util.IsValidInnerType(pe.Type(), "block") == false {
			return fmt.Errorf(
				"element '%s' of base type '%s' cannot be instantiated in element of base type '%s'",
				e.Name(), e.Type(), blk.Type(),
			)
		}

		/*
			err := checkElemConflict(elem, e)
			if err != nil {
				return fmt.Errorf("line %d: cannot instantiate element '%s': %v", pe.LineNum(), e.Name, err)
			}
		*/

		if blk.HasElement(e.Name()) {
			return fmt.Errorf(
				"cannot instantiate element '%s', element with such name is already instantiated in one of ancestor types",
				e.Name(),
			)
		}
		addBlockInnerElement(blk, e)
	}

	return nil
}

func fillBlockProps(blk *elem.Block) {
	if blk.Masters == 0 {
		blk.Masters = 1
	}
	if blk.Width == 0 {
		blk.Width = 32
	}
}

func addBlockInnerElement(blk *elem.Block, e elem.Element) {
	switch e.(type) {
	case (*elem.Config):
		blk.Configs = append(blk.Configs, e.(*elem.Config))
	default:
		panic("should never happen")
	}
}
