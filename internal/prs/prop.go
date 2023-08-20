package prs

import "fmt"

// Prop struct represents functionality property.
type Prop struct {
	Line  int
	Col   int
	Name  string
	Value Expr
}

func (p Prop) Loc() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Col)
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
