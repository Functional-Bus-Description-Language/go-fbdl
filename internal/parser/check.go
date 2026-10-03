package parser
import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func checkPropConflict(typ string, prop Prop, props PropContainer) error {
	msg := `cannot set '%s' property, because '%s' property is already set in line %d column %d`

	if w, ok := props.Get("width"); ok {
		if prop.Name == "range" {
			return token.Error{
				Msg:  fmt.Sprintf(msg, "range", "width", w.Line(), w.Col()),
				Toks: []token.Token{prop.NameTok, w.NameTok},
			}
		}
	}

	if r, ok := props.Get("range"); ok {
		if prop.Name == "width" {
			return fmt.Errorf(msg, "width", "range", r.Line(), r.Col())
		}
	}

	return nil
}
