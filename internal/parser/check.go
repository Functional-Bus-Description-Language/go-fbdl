package parser

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func checkPropConflict(prop Prop, props PropContainer) *token.Error {
	msg := "cannot set %s property, because %s property is already set"

	if prop.Name == "range" {
		if first, ok := props.Get("width"); ok {
			return &token.Error{
				Msg:  fmt.Sprintf(msg, prop.NameTok.Name(), first.NameTok.Name()),
				Toks: []token.Token{prop.NameTok, first.NameTok},
			}
		}
	}

	if prop.Name == "width" {
		if first, ok := props.Get("range"); ok {
			return &token.Error{
				Msg:  fmt.Sprintf(msg, prop.NameTok.Name(), first.NameTok.Name()),
				Toks: []token.Token{prop.NameTok, first.NameTok},
			}
		}
	}

	return nil
}
