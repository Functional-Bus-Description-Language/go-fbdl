package fun

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func HasElement(f *elem.Func, name string) bool {
	for i := range f.Params {
		if f.Params[i].Name == name {
			return true
		}
	}
	for i := range f.Returns {
		if f.Returns[i].Name == name {
			return true
		}
	}
	return false
}
