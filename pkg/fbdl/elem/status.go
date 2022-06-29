package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Status represents status element.
type Status struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  access.Access

	// Properties
	Atomic  bool
	Default BitStr
	Groups  []string
	Once    bool
	Width   int64
}

// HasDecreasingAccessOrder returns true if status must be accessed
// from the end register to the start register order.
// It is useful only in case of some atomic statuses.
// If the end register is wider, then starting reading from the end register
// saves some flip flops, becase the atomic shadow regsiter can be narrower.
func (s *Status) HasDecreasingAccessOrder() bool {
	if !s.Atomic {
		return false
	}

	if asc, ok := s.Access.(access.SingleContinuous); ok {
		if asc.IsEndMaskWider() {
			return true
		}
	}

	return false
}

func makeStatus(insSt *ins.Element) *Status {
	st := Status{
		Name:    insSt.Name,
		Doc:     insSt.Doc,
		IsArray: insSt.IsArray,
		Count:   insSt.Count,
		Atomic:  bool(insSt.Props["atomic"].(val.Bool)),
		Groups:  []string{},
		Width:   int64(insSt.Props["width"].(val.Int)),
	}

	if groups, ok := insSt.Props["groups"].(val.List); ok {
		for _, g := range groups {
			st.Groups = append(st.Groups, string(g.(val.Str)))
		}
	}

	return &st
}

// regStatus registerifies a Status element.
func regStatus(insSt *ins.Element, addr int64, gp *gap.Pool) (*Status, int64) {
	st := makeStatus(insSt)

	if insSt.IsArray {
		return regStatusArray(st, addr, gp)
	} else {
		return regStatusSingle(st, addr, gp)
	}
}

func regStatusSingle(st *Status, addr int64, gp *gap.Pool) (*Status, int64) {
	if g, ok := gp.GetSingle(st.Width, false); ok {
		st.Access = access.MakeSingleSingle(g.EndAddr, g.StartBit(), st.Width)
	} else {
		st.Access = access.MakeSingle(addr, 0, st.Width)
		addr += st.Access.RegCount()
	}
	if st.Access.EndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: st.Access.EndAddr(),
			EndAddr:   st.Access.EndAddr(),
			Mask:      access.Mask{Upper: busWidth - 1, Lower: st.Access.EndBit() + 1},
			WriteSafe: true,
		})
	}

	return st, addr
}

func regStatusArray(st *Status, addr int64, gp *gap.Pool) (*Status, int64) {
	if busWidth/2 < st.Width && st.Width <= busWidth {
		st.Access = access.MakeArraySingle(st.Count, addr, 0, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width == 0 || st.Count <= busWidth/st.Width || st.Width < busWidth/2 {
		st.Access = access.MakeArrayMultiplePacked(st.Count, addr, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("not yet implemented")
	}
	addr += st.Access.RegCount()

	return st, addr
}

func regStatusArraySingle(insSt *ins.Element, addr, startBit int64) (*Status, int64) {
	st := makeStatus(insSt)

	st.Access = access.MakeArraySingle(insSt.Count, addr, startBit, st.Width)

	addr += st.Access.RegCount()

	return st, addr
}
