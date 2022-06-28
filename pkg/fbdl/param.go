package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Param represents param element.
type Param struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  access.Access

	// Properties
	// TODO: Should Default be supported for param?
	// Default BitStr
	//Range Range
	Groups []string
	Width  int64
}

func makeParam(insParam *ins.Element) *Param {
	p := Param{
		Name:    insParam.Name,
		Doc:     insParam.Doc,
		IsArray: insParam.IsArray,
		Count:   insParam.Count,
		Groups:  []string{},
		Width:   int64(insParam.Props["width"].(val.Int)),
	}

	if groups, ok := insParam.Props["groups"].(val.List); ok {
		for _, g := range groups {
			p.Groups = append(p.Groups, string(g.(val.Str)))
		}
	}

	return &p
}
