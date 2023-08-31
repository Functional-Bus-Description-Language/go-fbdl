package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/block"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"golang.org/x/exp/maps"
	"log"
	"sort"
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

	err = checkBlockGroups(&blk)
	if err != nil {
		last := typeChain[len(typeChain)-1]
		return nil, fmt.Errorf("%d:%d: %v", last.Line(), last.Col(), err)
	}

	return &blk, nil
}

func applyBlockType(blk *fn.Block, typ prs.Functionality) error {
	for _, p := range typ.Props() {
		if err := util.IsValidProperty(p.Name, "bus"); err != nil {
			return fmt.Errorf(": %v", err)
		}
		if err := checkProp(p); err != nil {
			return fmt.Errorf("%s: %v", p.Loc(), err)
		}

		v, err := p.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
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
			blk.Width = int64(v.(val.Int))
		default:
			panic(fmt.Sprintf("unhandled '%s' property", p.Name))
		}
	}

	for _, s := range typ.Symbols().Symbols() {
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

		e := insElement(s.(prs.Functionality))

		if !util.IsValidInnerType(e.Type(), "block") {
			return fmt.Errorf(
				invalidInnerTypeMsg, e.GetName(), e.Type(), "block",
			)
		}

		if block.HasElement(blk, e.GetName()) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, e.GetName())
		}
		addBlockInnerElement(blk, e)
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

func addBlockInnerElement(blk *fn.Block, e any) {
	switch e := e.(type) {
	case (*fn.Config):
		block.AddConfig(blk, e)
	case (*fn.Irq):
		block.AddIrq(blk, e)
	case (*fn.Mask):
		block.AddMask(blk, e)
	case (*fn.Memory):
		block.AddMemory(blk, e)
	case (*fn.Proc):
		block.AddProc(blk, e)
	case (*fn.Static):
		block.AddStatic(blk, e)
	case (*fn.Status):
		block.AddStatus(blk, e)
	case (*fn.Stream):
		block.AddStream(blk, e)
	case (*fn.Block):
		block.AddSubblock(blk, e)
	default:
		panic("should never happen")
	}
}

func checkBlockGroups(blk *fn.Block) error {
	elemsWithGrps := blk.GroupedElems()

	if len(elemsWithGrps) == 0 {
		return nil
	}

	groups := make(map[string][]fn.Groupable)

	for _, e := range elemsWithGrps {
		grps := e.GroupNames()
		for _, g := range grps {
			if _, ok := groups[g]; !ok {
				groups[g] = []fn.Groupable{}
			}
			groups[g] = append(groups[g], e)
		}
	}

	// Check for element and group names conflict.
	for grpName := range groups {
		if block.HasElement(blk, grpName) {
			return fmt.Errorf("invalid group name %q, there is inner element with the same name", grpName)
		}
	}

	// Check for groups with single element.
	for name, g := range groups {
		if len(g) == 1 {
			return fmt.Errorf("group %q has only one element '%s'", name, g[0].GetName())
		}
	}

	// Check groups order.
	for i, e1 := range elemsWithGrps[:len(elemsWithGrps)-1] {
		grps1 := e1.GroupNames()
		for _, e2 := range elemsWithGrps[i+1:] {
			grps2 := e2.GroupNames()
			indexes := []int{}
			for _, g1 := range grps1 {
				for j2, g2 := range grps2 {
					if g1 == g2 {
						indexes = append(indexes, j2)
					}
				}
			}

			prevId := -1
			for _, id := range indexes {
				if id <= prevId {
					return fmt.Errorf(
						"conflicting order of groups, "+
							"group %q is after group %q in element '%s', "+
							"but before group %q in element '%s'",
						grps2[id], grps2[id+1], e1.GetName(), grps2[id+1], e2.GetName(),
					)
				}
				prevId = id
			}
		}
	}

	// Check for identical groups.
	grpNames := maps.Keys(groups)
	sort.Strings(grpNames)
	for _, grpName1 := range grpNames {
		g1 := groups[grpName1]
		elemNames1 := make([]string, 0, len(g1))
		for _, e := range g1 {
			elemNames1 = append(elemNames1, e.GetName())
		}
		for _, grpName2 := range grpNames {
			g2 := groups[grpName2]
			if grpName1 == grpName2 {
				continue
			}
			elemNames2 := make([]string, 0, len(g2))
			for _, e := range g2 {
				elemNames2 = append(elemNames2, e.GetName())
			}
			if len(elemNames1) != len(elemNames2) {
				continue
			}
			identical := true
			for _, name1 := range elemNames1 {
				found := false
				for _, name2 := range elemNames2 {
					if name1 == name2 {
						found = true
						break
					}
				}
				if !found {
					identical = false
					break
				}
			}
			if identical {
				return fmt.Errorf(
					"groups %q and %q of '%s' functionality are identical",
					grpName1, grpName2, blk.Name,
				)
			}
		}
	}

	return nil
}
