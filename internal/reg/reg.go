// Package reg implements code responsible for registerificaiton.
// This includes packing functionalities into registers and assigning addresses.
package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
	"log"
)

var busWidth uint

func Registerify(insBus *ins.Element) *Element {
	if insBus == nil {
		log.Println("registerification: there is no main bus; returning nil")
		return nil
	}

	busWidth = uint(insBus.Properties["width"].(val.Int).V)

	regBus := Element{InsElem: insBus, Elements: make(map[string]*Element)}

	// addr is current block internal access address, not global address.
	// 0 and 1 are reserved for x_uuid_x and x_timestamp_x.
	addr := uint(2)

	addr = registerifyFunctionalities(&regBus, addr)

	regBus.Sizes.Compact = addr
	regBus.Sizes.Own = addr

	return &regBus
}

func registerifyFunctionalities(elem *Element, addr uint) uint {
	if len(elem.InsElem.Elements) == 0 {
		return addr
	}

	addr = registerifyStatuses(elem, addr)

	return addr
}

// Current approach is trivial. Even groups are not respected.
func registerifyStatuses(elem *Element, addr uint) uint {
	var statuses = []*ins.Element{}
	for _, ie := range elem.InsElem.Elements {
		if ie.BaseType == "status" {
			statuses = append(statuses, ie)
		}
	}

	for _, st := range statuses {
		e := Element{InsElem: st}

		width := uint(st.Properties["width"].(val.Int).V)

		if st.IsArray {
			e.Access = MakeAccessSingle(addr, width)
		} else {
			e.Access = MakeAccessArray(st.Count, addr, width)
		}
		addr += e.Access.Count()

		elem.Elements[st.Name] = &e
	}

	return addr
}
