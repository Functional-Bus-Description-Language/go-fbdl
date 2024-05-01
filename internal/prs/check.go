package prs

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

func checkPropConflict(typ string, prop Prop, props PropContainer) error {
	msg := `cannot set '%s' property, because '%s' property is already set in line %d column %d`

	if w, ok := props.Get("width"); ok {
		if prop.Name == "range" {
			return tok.Error{
				Msg:  fmt.Sprintf(msg, "range", "width", w.Line(), w.Col()),
				Toks: []tok.Token{prop.NameTok, w.NameTok},
			}
		}
	}

	if r, ok := props.Get("range"); ok {
		if prop.Name == "width" {
			return fmt.Errorf(msg, "width", "range", r.Line, r.Col)
		}
	}

	return nil
}
