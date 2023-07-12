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
		} else if b == '[' {
			t = parseLeftBracket(&c)
		} else if b == ']' {
			t = parseRightBracket(&c)
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
	// TODO: Eat all spaces.
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
	return t, nil
}

func parseNegationOperator(c *context) Token {
	return Token{Kind: NEG, Pos: Position{Start: c.idx, End: c.idx}}
}

func parseEqualityOperator(c *context) Token {
	return Token{Kind: EQ, Pos: Position{Start: c.idx, End: c.idx + 1}}
}

func parseAssignmentOperator(c *context) Token {
	return Token{Kind: ASS, Pos: Position{Start: c.idx, End: c.idx}}
}

func parseLeftParenthesis(c *context) Token {
	return Token{Kind: LPAREN, Pos: Position{Start: c.idx, End: c.idx}}
}

func parseRightParenthesis(c *context) Token {
	return Token{Kind: RPAREN, Pos: Position{Start: c.idx, End: c.idx}}
}

func parseLeftBracket(c *context) Token {
	return Token{Kind: LBRACK, Pos: Position{Start: c.idx, End: c.idx}}
}

func parseRightBracket(c *context) Token {
	return Token{Kind: RBRACK, Pos: Position{Start: c.idx, End: c.idx}}
}

func parseBinaryBitStringLiteral(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Pos: Position{Start: c.idx}}

	// Skip b"
	i := c.idx + 2
	for {
		if i >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in binary bit string literal",
				c.line, c.col(t.Pos.Start),
			)
		}

		b := src[i]

		if b == '"' {
			t.Pos.End = i
			return t, nil
		}

		switch b {
		case '0', '1',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			i++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in binary bit string literal",
				c.line, c.col(i), b,
			)
		}
	}
}

func parseOctalBitStringLiteral(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Pos: Position{Start: c.idx}}

	// Skip o"
	i := c.idx + 2
	for {
		if i >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in octal bit string literal",
				c.line, c.col(t.Pos.Start),
			)
		}

		b := src[i]

		if b == '"' {
			t.Pos.End = i
			return t, nil
		}

		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			i++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in octal bit string literal",
				c.line, c.col(i), b,
			)
		}
	}
}

func parseHexBitStringLiteral(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Pos: Position{Start: c.idx}}

	// Skip x"
	i := c.idx + 2
	for {
		if i >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in hex bit string literal",
				c.line, c.col(t.Pos.Start),
			)
		}

		b := src[i]

		if b == '"' {
			t.Pos.End = i
			return t, nil
		}

		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			i++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in hex bit string literal",
				c.line, c.col(i), b,
			)
		}
	}
}

func parseNumberLiteral(c *context, src []byte) (Token, error) {
	b := src[c.idx]
	nb := nextByte(src, c.idx)

	if b == '0' && (nb == 'b' || nb == 'B') {
		return parseBinaryInt(c, src)
	} else if b == '0' && (nb == 'o' || nb == 'O') {
		return parseOctalInt(c, src)
	} else if b == '0' && (nb == 'x' || nb == 'X') {
		return parseHexInt(c, src)
	}

	t := Token{Kind: INT, Pos: Position{Start: c.idx}}
	hasPoint := false
	hasE := false
	end_idx := c.idx

byteLoop:
	for {
		if end_idx >= len(src) {
			break
		}

		end_idx++
		b := src[end_idx]
		if isDigit(b) {
			continue
		}
		switch b {
		case '.':
			if hasPoint {
				return t, fmt.Errorf(
					"%d:%d: second point character '.' in number literal",
					c.line, c.col(end_idx),
				)
			} else {
				if hasE {
					return t, fmt.Errorf(
						"%d:%d: point character '.' after exponent in number literal",
						c.line, c.col(end_idx),
					)
				}
				hasPoint = true
			}
		case 'e', 'E':
			if hasE {
				return t, fmt.Errorf(
					"%d:%d: second exponent in number literal",
					c.line, c.col(end_idx),
				)
			} else {
				hasE = true
			}
		case ' ', '\t', '(', ')', ']', '-', '+', '*', '/', '%', '=', '<', '>', ';', ':', ',':
			break byteLoop
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in number literal",
				c.line, c.col(end_idx), b,
			)
		}
	}

	t.Pos.End = end_idx - 1
	if hasPoint || hasE {
		t.Kind = REAL
	}

	return t, nil
}

func parseBinaryInt(c *context, src []byte) (Token, error) {
	panic("unimplemented")
}

func parseOctalInt(c *context, src []byte) (Token, error) {
	panic("unimplemented")
}

func parseHexInt(c *context, src []byte) (Token, error) {
	panic("unimplemented")
}

func parseWord(c *context, src []byte, s Stream) (Token, error) {
	var t Token
	word, hasHyphen := getWord(src, c.idx)

	if !hasHyphen {
		// Firstly assume word is a keyword
		t = parseKeyword(word, c)
		// If it is not a keyword, then it might be a property or identifier.
		if t.Kind == INVALID {
			t = parseProperty(word, c)
			// If it is not property, then it must be an identifier.
			if t.Kind == INVALID {
				t.Kind = IDENT
			} else {
				// However, properties are properties only if they are in valid place,
				// otherwise, these are regular identifiers.
				if prev_tok, ok := s.LastToken(); ok {
					if prev_tok.Kind != NEWLINE && prev_tok.Kind != SEMICOLON {
						t.Kind = IDENT
					}
				}
			}
		}
	} else {
		// Firstly assume word is a property
		t = parseProperty(word, c)
		// If it is not property, then it is part of an expression.
		if t.Kind == INVALID {
			panic("unimplemented")
		} else {
			// It might be property, or part of an expression.
			prev_tok, ok := s.LastToken()
			if !ok {
				// Safe to return, time literal units do not contain hyphen '-'.
				return t, nil
			}
			// It is part of an expression.
			if prev_tok.Kind != NEWLINE && prev_tok.Kind != SEMICOLON {
				panic("unimplemented")
			}
		}
	}

	// The word might be the unit of time literal
	if t.Kind == IDENT {
		if prev_tok, ok := s.LastToken(); ok {
			if prev_tok.Kind == INT {
				switch string(word) {
				case "ns", "us", "ms", "s":
					idx := len(s) - 1
					s[idx].Kind = TIME
					s[idx].Pos.End = t.Pos.End
					t.Kind = INVALID
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
	case "config":
		t.Kind = CONFIG
	case "const":
		t.Kind = CONST
	case "import":
		t.Kind = IMPORT
	case "irq":
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

func parseProperty(word []byte, c *context) Token {
	t := Token{Kind: INVALID, Pos: Position{Start: c.idx, End: c.idx + len(word) - 1}}

	switch string(word) {
	case "access":
		t.Kind = ACCESS
	case "add-enable":
		t.Kind = ADD_ENABLE
	case "atomic":
		t.Kind = ATOMIC
	case "byte-write-enable":
		t.Kind = BYTE_WRITE_ENABLE
	case "clear":
		t.Kind = CLEAR
	case "delay":
		t.Kind = DELAY
	case "enable-init-value":
		t.Kind = ENABLE_INIT_VALUE
	case "enable-reset-value":
		t.Kind = ENABLE_RESET_VALUE
	case "groups":
		t.Kind = GROUPS
	case "init-value":
		t.Kind = INIT_VALUE
	case "in-trigger":
		t.Kind = IN_TRIGGER
	case "masters":
		t.Kind = MASTERS
	case "out-trigger":
		t.Kind = OUT_TRIGGER
	case "range":
		t.Kind = RANGE
	case "read-latency":
		t.Kind = READ_LATENCY
	case "read-value":
		t.Kind = READ_VALUE
	case "reset":
		t.Kind = RESET
	case "reset-value":
		t.Kind = RESET_VALUE
	case "size":
		t.Kind = SIZE
	case "width":
		t.Kind = WIDTH
	}

	return t
}
