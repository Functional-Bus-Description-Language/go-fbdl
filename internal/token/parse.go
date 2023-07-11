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

// getWord returns word from the source starting from index idx.
// The function assumes byte under idx is not a whitespace character.
// The second return is true if word contains hyphen '-' character.
func getWord(src []byte, idx int) ([]byte, bool) {
	hasHyphen := false
	end_idx := idx

	for {
		if end_idx >= len(src) {
			return src[idx:end_idx], hasHyphen
		}

		b := src[end_idx]
		if isLetter(b) || isDigit(b) || b == '_' || b == '-' {
			if b == '-' {
				hasHyphen = true
			}
			end_idx++
			continue
		} else {
			return src[idx:end_idx], hasHyphen
		}
	}
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
		b := src[c.idx]            // Current byte
		nb := nextByte(src, c.idx) // Next byte

		if b == ' ' {
			err = parseSpace(&c, s)
		} else if b == '\t' {
			err = parseTab(&c, s)
		} else if b == '\n' {
			t, err = parseNewline(&c, s)
		} else if b == '#' {
			t = parseComment(&c, src)
		} else if b == ',' {
			t, err = parseComma(&c, s)
		} else if b == ';' {
			t, err = parseSemicolon(&c, s)
		} else if b == '!' {
			t = parseNegationOperator(&c)
		} else if b == '=' && nb == '=' {
			t = parseEqualityOperator(&c)
		} else if b == '=' {
			t = parseAssignmentOperator(&c)
		} else if b == '(' {
			t = parseLeftParenthesis(&c)
		} else if b == ')' {
			t = parseRightParenthesis(&c)
		} else if (b == 'b' || b == 'B') && nb == '"' {
			t, err = parseBinaryBitStringLiteral(&c, src)
		} else if (b == 'o' || b == 'O') && nb == '"' {
			t, err = parseOctalBitStringLiteral(&c, src)
		} else if (b == 'x' || b == 'X') && nb == '"' {
			t, err = parseHexBitStringLiteral(&c, src)
		} else if isDigit(b) {
			t, err = parseNumberLiteral(&c, src)
		} else if isLetter(b) {
			t, err = parseWord(&c, src, s)
		} else {
			panic(fmt.Sprintf("unhandled byte '%c'", b))
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

func parseComma(c *context, s Stream) (Token, error) {
	if prev_tok, ok := s.LastToken(); ok {
		if prev_tok.Kind == COMMA {
			return Token{}, fmt.Errorf(
				"%d:%d: redundant ','", prev_tok.Pos.Line, c.col(c.idx),
			)
		}
	}

	t := Token{
		Kind: COMMA,
		Pos: Position{
			Start: c.idx,
			End:   c.idx,
		},
	}
	c.idx++
	return t, nil
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

func parseNegationOperator(c *context) Token {
	t := Token{Kind: NEG, Pos: Position{Start: c.idx, End: c.idx}}
	c.idx++
	return t
}

func parseEqualityOperator(c *context) Token {
	t := Token{Kind: EQ, Pos: Position{Start: c.idx, End: c.idx + 1}}
	c.idx += 2
	return t
}

func parseAssignmentOperator(c *context) Token {
	t := Token{Kind: ASS, Pos: Position{Start: c.idx, End: c.idx}}
	c.idx++
	return t
}

func parseLeftParenthesis(c *context) Token {
	t := Token{Kind: LPAREN, Pos: Position{Start: c.idx, End: c.idx}}
	c.idx++
	return t
}

func parseRightParenthesis(c *context) Token {
	t := Token{Kind: RPAREN, Pos: Position{Start: c.idx, End: c.idx}}
	c.idx++
	return t
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
		case '0', '1',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			c.idx++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in binary bit string literal",
				c.line, c.col(c.idx), b,
			)
		}
	}
}

func parseOctalBitStringLiteral(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Pos: Position{Start: c.idx}}

	// Skip o"
	c.idx += 2
	for {
		if c.idx >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in octal bit string literal",
				c.line, c.col(t.Pos.Start),
			)
		}

		b := src[c.idx]

		if b == '"' {
			t.Pos.End = c.idx
			return t, nil
		}

		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			c.idx++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in octal bit string literal",
				c.line, c.col(c.idx), b,
			)
		}
	}
}

func parseHexBitStringLiteral(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Pos: Position{Start: c.idx}}

	// Skip x"
	c.idx += 2
	for {
		if c.idx >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in hex bit string literal",
				c.line, c.col(t.Pos.Start),
			)
		}

		b := src[c.idx]

		if b == '"' {
			t.Pos.End = c.idx
			return t, nil
		}

		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			c.idx++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in hex bit string literal",
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

func parseWord(c *context, src []byte, s Stream) (Token, error) {
	var t Token
	word, hasHyphen := getWord(src, c.idx)

	if !hasHyphen {
		// Firstly assume word is a keyword
		t = parseKeyword(word, c)
		// Start and End are already set by the parseKeyword function
		if t.Kind == INVALID {
			// If it is not keyword, then it must be identifier
			t.Kind = IDENT
		} else {
			// In other case this might be a keyword, but not necessarily,
			// as for example "const block = true" is valid semantically.
			if prev_tok, ok := s.LastToken(); ok {
				k := prev_tok.Kind
				if k == CONST ||
					(isOperator(k) && t.Kind != BOOL) ||
					(k == NEWLINE && isFunctionality(t.Kind)) {
					t.Kind = IDENT
				}
			} else {
				if isFunctionality(t.Kind) {
					t.Kind = IDENT
				}
			}
		}
	}

	return t, nil
}

func parseKeyword(word []byte, c *context) Token {
	t := Token{Kind: INVALID, Pos: Position{Start: c.idx, End: c.idx + len(word) - 1}}

	switch string(word) {
	case "false", "true":
		t.Kind = BOOL
	case "block":
		t.Kind = BLOCK
	case "bus":
		t.Kind = BUS
	case "const":
		t.Kind = CONST
	case "import":
		t.Kind = IMPORT
	case "IRQ":
		t.Kind = IRQ
	case "mask":
		t.Kind = MASK
	case "memory":
		t.Kind = MEMORY
	case "param":
		t.Kind = PARAM
	case "proc":
		t.Kind = PROC
	case "return":
		t.Kind = RETURN
	case "static":
		t.Kind = STATIC
	case "stream":
		t.Kind = STREAM
	case "type":
		t.Kind = TYPE
	}

	return t
}
