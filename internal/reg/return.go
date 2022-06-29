package reg

import (
	//"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func makeReturn(r *elem.Return) *elem.Return {
	/*
		r := elem.Return{
			Name:    insRet.Name,
			Doc:     insRet.Doc,
			IsArray: insRet.IsArray,
			Count:   insRet.Count,
			Groups:  []string{},
			Width:   int64(insRet.Props["width"].(val.Int)),
		}
	*/

	/*
		if groups, ok := insRet.Props["groups"].(val.List); ok {
			for _, g := range groups {
				r.Groups = append(r.Groups, string(g.(val.Str)))
			}
		}
	*/

	return r
}
