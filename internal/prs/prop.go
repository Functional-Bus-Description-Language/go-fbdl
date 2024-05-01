package prs

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Prop struct represents functionality property.
type Prop struct {
	NameTok tok.Token
	Name    string

	Value    Expr
	ValueTok tok.Token
}

func (p Prop) Line() int { return p.NameTok.Line() }
func (p Prop) Col() int  { return p.NameTok.Column() }

func (p Prop) Loc() string {
	return fmt.Sprintf("%d:%d", p.Line(), p.Col())
}

// PropContainer represents a list of properties.
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
