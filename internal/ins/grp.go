package ins

type Group struct {
	Name     string
	Elements []*Element
}

// IsStatus returns true if group contains only status elements.
func (g *Group) IsStatus() bool {
	for _, e := range g.Elements {
		if e.Type != "status" {
			return false
		}
	}
	return true
}

// IsConfig returns true if group contains only config elements.
func (g *Group) IsConfig() bool {
	for _, e := range g.Elements {
		if e.Type != "config" {
			return false
		}
	}
	return true
}

// IsArray returns true if group contains only array elements.
func (g *Group) IsArray() bool {
	for _, e := range g.Elements {
		if !e.IsArray {
			return false
		}
	}
	return true
}
