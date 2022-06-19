package ins

type ElemContainer []*Element

// Add adds element to ElemContainer.
// If element with given name already exists it returns false.
// If the operation is successful it returns true.
func (ec *ElemContainer) Add(elem *Element) bool {
	for _, e := range *ec {
		if e.Name == elem.Name {
			return false
		}
	}

	*ec = append(*ec, elem)

	return true
}

func (ec *ElemContainer) Get(name string) (*Element, bool) {
	for _, e := range *ec {
		if e.Name == name {
			return e, true
		}
	}

	return nil, false
}

func (ec *ElemContainer) GetAllByType(typ string) []*Element {
	ret := []*Element{}

	for _, e := range *ec {
		if e.Type == typ {
			ret = append(ret, e)
		}
	}

	return ret
}

// HasType returns true if element container has at least one element of given type.
func (ec *ElemContainer) HasType(typ string) bool {
	for _, e := range *ec {
		if e.Type == typ {
			return true
		}
	}
	return false
}
