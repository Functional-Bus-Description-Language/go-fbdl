package fbdl

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

func makeGroupStatusArraySameSizesSingle(group *ins.Group, addr int64) (Group, int64) {
	grp := GroupStatusArraySameSizesSingle{
		name:     group.Name,
		statuses: []*Status{},
	}

	startBit := int64(0)

	for _, e := range group.Elements {
		st, _ := registerifyStatusArraySingle(e, addr, startBit)
		startBit += st.Width
		grp.statuses = append(grp.statuses, st)
	}

	return &grp, addr
}

func registerifyGroupStatusArray(blk *Block, group *ins.Group, addr int64) (Group, int64) {
	sameSizes := true
	for _, e := range group.Elements {
		if e.Count != group.Elements[0].Count {
			sameSizes = false
			break
		}
	}

	var grp Group
	if sameSizes {
		grp, addr = registerifyGroupStatusArraySameSizes(blk, group, addr)
	} else {
		panic("not yet implemented")
	}

	return grp, addr
}

func registerifyGroupStatusArraySameSizes(blk *Block, group *ins.Group, addr int64) (Group, int64) {
	widths := []int64{}
	singleIndexWidth := int64(0)

	for _, e := range group.Elements {
		w := int64(e.Props["width"].(val.Int))
		widths = append(widths, w)
		singleIndexWidth += w
	}

	var grp Group
	if busWidth/2 < singleIndexWidth && singleIndexWidth <= busWidth {
		grp, addr = makeGroupStatusArraySameSizesSingle(group, addr)
	}

	return grp, addr
}
