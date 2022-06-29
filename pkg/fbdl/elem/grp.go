package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type Group interface {
	Name() string
	Statuses() []*Status
}

type GroupStatusArraySameSizesSingle struct {
	name     string
	statuses []*Status
}

func (g *GroupStatusArraySameSizesSingle) Name() string        { return g.name }
func (g *GroupStatusArraySameSizesSingle) Statuses() []*Status { return g.statuses }

func makeGroupStatusArraySameSizesSingle(insGrp *ins.Group, addr int64) (Group, int64) {
	grp := GroupStatusArraySameSizesSingle{
		name:     insGrp.Name,
		statuses: []*Status{},
	}

	startBit := int64(0)

	for _, e := range insGrp.Elems {
		st, _ := regStatusArraySingle(e, addr, startBit)
		startBit += st.Width
		grp.statuses = append(grp.statuses, st)
	}

	return &grp, addr
}

func regGroupStatusArray(blk *Block, insGrp *ins.Group, addr int64) (Group, int64) {
	sameSizes := true
	for _, e := range insGrp.Elems {
		if e.Count != insGrp.Elems[0].Count {
			sameSizes = false
			break
		}
	}

	var grp Group
	if sameSizes {
		grp, addr = regGroupStatusArraySameSizes(blk, insGrp, addr)
	} else {
		panic("not yet implemented")
	}

	return grp, addr
}

func regGroupStatusArraySameSizes(blk *Block, insGrp *ins.Group, addr int64) (Group, int64) {
	widths := []int64{}
	singleIndexWidth := int64(0)

	for _, e := range insGrp.Elems {
		w := int64(e.Props["width"].(val.Int))
		widths = append(widths, w)
		singleIndexWidth += w
	}

	var grp Group
	if busWidth/2 < singleIndexWidth && singleIndexWidth <= busWidth {
		grp, addr = makeGroupStatusArraySameSizesSingle(insGrp, addr)
	}

	return grp, addr
}
