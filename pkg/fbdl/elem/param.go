package elem

import (
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
