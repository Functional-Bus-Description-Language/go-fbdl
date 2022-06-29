package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func insBlock(typeChain []prs.Element) (*elem.Block, error) {
	// Instantiation is always the last one in the type chain.
	inst := typeChain[len(typeChain)-1].(*prs.Inst)

	blk := elem.Block{
		Name:    inst.Name(),
		Doc:     inst.Doc(),
		IsArray: false,
		Count:   1,
	}

	if inst.IsArray {
		blk.IsArray = true
		count, err := inst.Count.Eval()

		if count.Type() != "integer" {
			return nil, fmt.Errorf("size of array must be of 'integer' type, current type '%s'", count.Type())
		}

		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		blk.Count = int64(count.(val.Int))
	}

	for i, typ := range typeChain {
		resolvedArgs := make(map[string]prs.Expr)
		if (i+1) < len(typeChain) && typeChain[i+1].ResolvedArgs() != nil {
			resolvedArgs = typeChain[i+1].ResolvedArgs()
		}
		err := applyBlockType(&blk, typ, resolvedArgs)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	fillBlockProps(&blk)

	return &blk, nil
}

// TODO: Can resolvedArgs be set in insBlock?
func applyBlockType(blk *elem.Block, typ prs.Element, resolvedArgs map[string]prs.Expr) error {
	if resolvedArgs != nil {
		typ.SetResolvedArgs(resolvedArgs)
	}

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
