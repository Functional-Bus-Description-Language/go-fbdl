package token

import (
	"fmt"
)

type context struct {
	line int
	//indent int       // Current indent level
	idx         int // Start index
	newline_idx int // Last newline index
}

// Col returns column number for given index.
func (c context) col(idx int) int {
	return idx - c.newline_idx
}

// nextByte returns byte with index equal idx + 1.
// If (idx + 1) >= len(src), then 0 is returned.
func nextByte(src []byte, idx int) byte {
	if idx+1 >= len(src) {
		return 0
	}
	return src[idx+1]
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
	c.newline_idx = -1

	for {
		if c.idx == len(src) {
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
		} else if b == '#' {
			t = parseComment(&c, src)
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
			} else {
				t.Pos.Column = c.col(t.Pos.Start)
			}

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
				"%d:%d: space character ' ' not allowed for indent", c.line, c.col(c.idx),
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
		Pos:  Position{Start: c.idx, End: c.idx, Column: c.col(c.idx)},
	}
	c.newline_idx = c.idx
	c.idx++
	return t, nil
}

func parseComment(c *context, src []byte) Token {
	t := Token{Kind: COMMENT, Pos: Position{Start: c.idx}}

	for {
		c.idx++
		if c.idx >= len(src) || src[c.idx] == '\n' {
			t.Pos.End = c.idx - 1
			return t
		}
	}
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func parseNumberLiteral(c *context, src []byte) (Token, error) {
	panic("unimplemented")
	t := Token{Pos: Position{Start: c.idx}}

	return t, nil
}
