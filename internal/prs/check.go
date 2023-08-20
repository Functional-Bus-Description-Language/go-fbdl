package prs

import (
	"fmt"
)

func checkPropConflict(typ string, prop Prop, props PropContainer) error {
	msg := `cannot set '%s' property, because '%s' property is already set in line %d`

	if w, ok := props.Get("width"); ok {
		if prop.Name == "range" {
			return fmt.Errorf(msg, "range", "width", w.Line)
		}
	}

	if r, ok := props.Get("range"); ok {
		if prop.Name == "width" {
			return fmt.Errorf(msg, "width", "range", r.Line)
		}
	}

	return nil
}
