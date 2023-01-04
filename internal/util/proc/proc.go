package proc

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func HasElement(p *elem.Proc, name string) bool {
	for i := range p.Params {
		if p.Params[i].Name == name {
			return true
		}
	}
	for i := range p.Returns {
		if p.Returns[i].Name == name {
			return true
		}
	}
	return false
}
