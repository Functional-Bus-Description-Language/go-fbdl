package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Param represents param element.
type Return struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  access.Access

	// Properties
	Groups []string
	Width  int64
}

func makeReturn(insRet *ins.Element) *Return {
	r := Return{
		Name:    insRet.Name,
		Doc:     insRet.Doc,
		IsArray: insRet.IsArray,
		Count:   insRet.Count,
		Groups:  []string{},
		Width:   int64(insRet.Props["width"].(val.Int)),
	}

	if groups, ok := insRet.Props["groups"].(val.List); ok {
		for _, g := range groups {
			r.Groups = append(r.Groups, string(g.(val.Str)))
		}
	}

	return &r
}
