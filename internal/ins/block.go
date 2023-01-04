package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/block"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	"golang.org/x/exp/maps"
	"log"
	"sort"
)

func insBlock(typeChain []prs.Element) (*elem.Block, error) {
	typeChainStr := fmt.Sprintf("debug: instantiating block, type chain: %s", typeChain[0].Name())
	for i := 1; i < len(typeChain); i++ {
		typeChainStr = fmt.Sprintf("%s -> %s", typeChainStr, typeChain[i].Name())
	}
	log.Print(typeChainStr)

	e, err := makeElem(typeChain)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	blk := elem.Block{}
	blk.Elem = e

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
		return nil, fmt.Errorf("%v", err)
	}

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
			if constContainer.HasConst(blk.ConstContainer, c.Name()) {
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
			constContainer.AddConst(&blk.ConstContainer, c.Name(), val)
		}

		_, ok := s.(*prs.Inst)
		if !ok {
			continue
		}

		e := insElement(s.(prs.Element))

		if !util.IsValidInnerType(elem.Type(e), "block") {
			return fmt.Errorf(
				invalidInnerTypeMsg, elem.Name(e), elem.Type(e), "block",
			)
		}

		if block.HasElement(blk, elem.Name(e)) {
			return fmt.Errorf(elemWithNameAlreadyInstMsg, elem.Name(e))
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

func addBlockInnerElement(blk *elem.Block, e any) {
	switch e := e.(type) {
	case (*elem.Config):
		block.AddConfig(blk, e)
	case (*elem.Mask):
		block.AddMask(blk, e)
	case (*elem.Proc):
		block.AddProc(blk, e)
	case (*elem.Static):
		block.AddStatic(blk, e)
	case (*elem.Status):
		block.AddStatus(blk, e)
	case (*elem.Stream):
		block.AddStream(blk, e)
	case (*elem.Block):
		block.AddSubblock(blk, e)
	default:
		panic("should never happen")
	}
}

func checkBlockGroups(blk *elem.Block) error {
	elemsWithGrps := blk.GroupedElems()

	if len(elemsWithGrps) == 0 {
		return nil
	}

	groups := make(map[string][]elem.Groupable)

	for _, e := range elemsWithGrps {
		grps := e.GroupNames()
		for _, g := range grps {
			if _, ok := groups[g]; !ok {
				groups[g] = []elem.Groupable{}
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
			return fmt.Errorf("group %q has only one element '%s'", name, elem.Name(g[0]))
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
						grps2[id], grps2[id+1], elem.Name(e1), grps2[id+1], elem.Name(e2),
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
			elemNames1 = append(elemNames1, elem.Name(e))
		}
		for _, grpName2 := range grpNames {
			g2 := groups[grpName2]
			if grpName1 == grpName2 {
				continue
			}
			elemNames2 := make([]string, 0, len(g2))
			for _, e := range g2 {
				elemNames2 = append(elemNames2, elem.Name(e))
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
				return fmt.Errorf("groups %q and %q are identical", grpName1, grpName2)
			}
		}
	}

	return nil
}
