package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Status represents status element.
type Status struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Atomic  bool
	Default BitStr
	Groups  []string
	Once    bool
	Width   int64
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

func registerifyStatus(insSt *ins.Element, addr int64, gp *gapPool) (*Status, int64) {
	st := makeStatus(insSt)

	if insSt.IsArray {
		return registerifyStatusArray(st, addr, gp)
	} else {
		return registerifyStatusSingle(st, addr, gp)
	}
}

func registerifyStatusSingle(st *Status, addr int64, gp *gapPool) (*Status, int64) {
	if g, ok := gp.getSingle(st.Width, false); ok {
		st.Access = makeAccessSingleSingle(g.endAddr, g.StartBit(), st.Width)
	} else {
		st.Access = makeAccessSingle(addr, 0, st.Width)
		addr += st.Access.RegCount()
	}
	if st.Access.EndBit() < busWidth-1 {
		gp.Add(gap{
			isArray:   false,
			startAddr: st.Access.EndAddr(),
			endAddr:   st.Access.EndAddr(),
			mask:      AccessMask{Upper: busWidth - 1, Lower: st.Access.EndBit() + 1},
			writeSafe: true,
		})
	}

	return st, addr
}

func registerifyStatusArray(st *Status, addr int64, gp *gapPool) (*Status, int64) {
	if busWidth/2 < st.Width && st.Width <= busWidth {
		st.Access = makeAccessArraySingle(st.Count, addr, 0, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width == 0 || st.Count <= busWidth/st.Width || st.Width < busWidth/2 {
		st.Access = makeAccessArrayMultiplePacked(st.Count, addr, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("not yet implemented")
	}
	addr += st.Access.RegCount()

	return st, addr
}

func registerifyStatusArraySingle(insSt *ins.Element, addr, startBit int64) (*Status, int64) {
	st := makeStatus(insSt)

	st.Access = makeAccessArraySingle(insSt.Count, addr, startBit, st.Width)

	addr += st.Access.RegCount()

	return st, addr
}
