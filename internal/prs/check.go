package prs

import (
	"fmt"
)

func checkPropConflict(typ string, prop Prop, props PropContainer) error {
	msg := `line %d: cannot set '%s' property, because '%s' property is already set in line %d`

	if w, ok := props.Get("width"); ok {
		if prop.Name == "range" {
			return fmt.Errorf(msg, prop.LineNum, "range", "width", w.LineNum)
		}
	}

	if r, ok := props.Get("range"); ok {
		if prop.Name == "width" {
			return fmt.Errorf(msg, prop.LineNum, "width", "range", r.LineNum)
		}
	}

	return nil
}
