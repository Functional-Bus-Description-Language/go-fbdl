package elem

import (
	"fmt"
)

type Groupable interface {
	Element
	isGroupable() bool
}

func (c *Config) isGroupable() bool { return true }
func (m *Mask) isGroupable() bool   { return true }
func (s *Status) isGroupable() bool { return true }
func (p *Param) isGroupable() bool  { return true }
func (r *Return) isGroupable() bool { return true }

// Grouped returns inner elements with groups.
func Grouped(elem any) []Groupable {
	switch e := elem.(type) {
	case *Block:
		return blockGroups(e)
	default:
		panic(
			fmt.Sprintf("%T is not a group container\n", elem),
		)
	}
}

func blockGroups(blk *Block) []Groupable {
	elemsWithGrps := []Groupable{}

	for _, c := range blk.Configs {
		if len(c.Groups) > 0 {
			elemsWithGrps = append(elemsWithGrps, c)
		}
	}
	for _, m := range blk.Masks {
		if len(m.Groups) > 0 {
			elemsWithGrps = append(elemsWithGrps, m)
		}
	}
	for _, s := range blk.Statuses {
		if len(s.Groups) > 0 {
			elemsWithGrps = append(elemsWithGrps, s)
		}
	}

	return elemsWithGrps
}

// Groups returns element groups.
func Groups(elem any) []string {
	switch e := elem.(type) {
	case *Config:
		return e.Groups
	case *Mask:
		return e.Groups
	case *Param:
		return e.Groups
	case *Return:
		return e.Groups
	case *Status:
		return e.Groups
	default:
		panic(
			fmt.Sprintf("%T doesn't have 'groups' property\n", elem),
		)
	}
}
