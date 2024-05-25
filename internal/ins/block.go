package ins

import (
	"fmt"
	"log"
	"sort"

	"golang.org/x/exp/maps"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
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
			if !util.IsValidBusWidth(width) {
				return tok.Error{
					Msg: fmt.Sprintf(
						"invalid bus width %d, valid bus width must be greater than 7 and must be a power of 2",
						width,
					),
					Toks: []tok.Token{p.ValueTok},
				}
			}
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

func checkBlockGroups(blk *fn.Block) error {
	instsWithGrps := blk.GroupedInsts()

	if len(instsWithGrps) == 0 {
		return nil
	}

	groups := make(map[string][]fn.Groupable)

	for _, i := range instsWithGrps {
		grps := i.GroupNames()
		for _, g := range grps {
			if _, ok := groups[g]; !ok {
				groups[g] = []fn.Groupable{}
			}
			groups[g] = append(groups[g], i)
		}
	}

	// Check for functionality and group names conflict.
	for grpName := range groups {
		if block.HasFunctionality(blk, grpName) {
			return fmt.Errorf("invalid group name %q, there is inner functionality with the same name", grpName)
		}
	}

	// Check for groups with single functionality.
	for name, g := range groups {
		if len(g) == 1 {
			return fmt.Errorf("group %q has only one functionality '%s'", name, g[0].GetName())
		}
	}

	// Check groups order.
	for i, e1 := range instsWithGrps[:len(instsWithGrps)-1] {
		grps1 := e1.GroupNames()
		for _, e2 := range instsWithGrps[i+1:] {
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
							"group %q is after group %q in functionality '%s', "+
							"but before group %q in functionality '%s'",
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
		instNames1 := make([]string, 0, len(g1))
		for _, i := range g1 {
			instNames1 = append(instNames1, i.GetName())
		}
		for _, grpName2 := range grpNames {
			g2 := groups[grpName2]
			if grpName1 == grpName2 {
				continue
			}
			instNames2 := make([]string, 0, len(g2))
			for _, i := range g2 {
				instNames2 = append(instNames2, i.GetName())
			}
			if len(instNames1) != len(instNames2) {
				continue
			}
			identical := true
			for _, name1 := range instNames1 {
				found := false
				for _, name2 := range instNames2 {
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
