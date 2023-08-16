package prs

// Prop struct represents functionality property.
type Prop struct {
	LineNum uint32
	Name    string
	Value   Expr
}

type PropContainer []Prop

func (pc *PropContainer) Add(prop Prop) bool {
	for _, p := range *pc {
		if p.Name == prop.Name {
			return false
		}
	}

	*pc = append(*pc, prop)

	return true
}

func (pc PropContainer) Get(name string) (Prop, bool) {
	for _, p := range pc {
		if p.Name == name {
			return p, true
		}
	}

	return Prop{}, false
}
