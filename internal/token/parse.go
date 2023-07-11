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

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func isLetter(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z')
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
		} else if b == ';' {
			t, err = parseSemicolon(&c, s)
		} else if (b == 'b' || b == 'B') && nextByte(src, c.idx) == '"' {
			t, err = parseBinaryBitStringLiteral(&c, src)
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
	if prev_tok, ok := s.LastToken(); ok {
		if prev_tok.Kind == SEMICOLON {
			return Token{}, fmt.Errorf(
				"%d:%d: extra ';' at the end of line", prev_tok.Pos.Line, prev_tok.Pos.Column,
			)
		}
	}

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

func parseSemicolon(c *context, s Stream) (Token, error) {
	if prev_tok, ok := s.LastToken(); ok {
		if prev_tok.Kind == SEMICOLON {
			return Token{}, fmt.Errorf(
				"%d:%d: redundant ';'", prev_tok.Pos.Line, c.col(c.idx),
			)
		}
	}

	t := Token{
		Kind: SEMICOLON,
		Pos: Position{
			Start: c.idx,
			End:   c.idx,
		},
	}
	c.idx++
	return t, nil
}

func parseBinaryBitStringLiteral(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Pos: Position{Start: c.idx}}

	// Skip b"
	c.idx += 2
	for {
		if c.idx >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in binary bit string literal",
				c.line, c.col(t.Pos.Start),
			)
		}

		b := src[c.idx]

		if b == '"' {
			t.Pos.End = c.idx
			return t, nil
		}

		switch b {
		case '0', '1', '-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			c.idx++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in binary bit string literal",
				c.line, c.col(c.idx), b,
			)
		}
	}
}

func parseNumberLiteral(c *context, src []byte) (Token, error) {
	panic("unimplemented")
	t := Token{Pos: Position{Start: c.idx}}

	return t, nil
}
