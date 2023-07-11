package token

import (
	"fmt"
)

type context struct {
	line uint
	//indent uint       // Current indent level
	idx         uint // Start index
	newline_idx uint // Last newline index
}

// Col returns column number.
func (c context) col() uint {
	return c.idx - c.newline_idx
}

// Parse parses src byte array containing the source code and returns token Stream.
func Parse(src []byte) (Stream, error) {
	var (
		c   context
		err error
		s   Stream
		t   Token
	)
	c.line = 1

	for {
		if int(c.idx) == len(src) {
			break
		}

		t.Kind = INVALID
		err = nil
		b := src[c.idx]

		if b == ' ' {
			err = parseSpace(&c, s)
		} else if b == '\t' {
			err = parseTab(&c, s)
		} else if b == '\n' {
			t, err = parseNewline(&c, s)
		} else if isDigit(b) {
			t, err = parseNumberLiteral(&c, src)
		}

		if err != nil {
			return s, err
		}

		if t.Kind != INVALID {
			t.Pos.Line = c.line
			if t.Kind == NEWLINE {
				c.line++
			}

			t.Pos.Column = c.col()
			s = append(s, t)
			c.idx = t.Pos.End + 1
		}
	}

	return s, nil
}

func parseSpace(c *context, s Stream) error {
	if t, ok := s.LastToken(); ok {
		if t.Kind == NEWLINE {
			return fmt.Errorf(
				"%d:%d: space character ' ' not allowed for indent", c.line, c.col(),
			)
		}
	}
	c.idx++
	return nil
}

func parseTab(c *context, s Stream) error {
	c.idx++
	return nil
}

func parseNewline(c *context, s Stream) (Token, error) {
	t := Token{
		Kind: NEWLINE,
		Pos:  Position{Start: c.idx, End: c.idx},
	}
	c.newline_idx = c.idx
	c.idx++
	return t, nil
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func parseNumberLiteral(c *context, src []byte) (Token, error) {
	panic("unimplemented")
	t := Token{Pos: Position{Start: c.idx}}

	return t, nil
}
