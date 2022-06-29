package elem

import (
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

func (r *Return) Hash() int64 {
	return 0
}
