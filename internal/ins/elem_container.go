package ins

type ElementContainer []*Element

// Add adds element to ElementContainer.
// If element with given name already exists it returns false.
// If the operation is successful it returns true.
func (ec *ElementContainer) Add(elem *Element) bool {
	for _, e := range *ec {
		if e.Name == elem.Name {
			return false
		}
	}

	*ec = append(*ec, elem)

	return true
}

func (ec *ElementContainer) Get(name string) (*Element, bool) {
	for _, e := range *ec {
		if e.Name == name {
			return e, true
		}
	}

	return nil, false
}

func (ec *ElementContainer) GetAllByType(typ string) []*Element {
	ret := []*Element{}

	for _, e := range *ec {
		if e.Type == typ {
			ret = append(ret, e)
		}
	}

	return ret
}
