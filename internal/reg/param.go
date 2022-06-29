package reg

import (
	_ "github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func makeParam(p *elem.Param) *elem.Param {
	/*
		p := elem.Param{
			Name:    insParam.Name,
			Doc:     insParam.Doc,
			IsArray: insParam.IsArray,
			Count:   insParam.Count,
			Groups:  []string{},
			Width:   int64(insParam.Props["width"].(val.Int)),
		}
	*/

	/*
		if groups, ok := insParam.Props["groups"].(val.List); ok {
			for _, g := range groups {
				p.Groups = append(p.Groups, string(g.(val.Str)))
			}
		}
	*/

	return p
}
