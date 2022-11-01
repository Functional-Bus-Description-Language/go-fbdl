package stream

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func HasElement(s *elem.Stream, name string) bool {
	for i := range s.Params {
		if s.Params[i].Name == name {
			return true
		}
	}
	for i := range s.Returns {
		if s.Returns[i].Name == name {
			return true
		}
	}
	return false
}
