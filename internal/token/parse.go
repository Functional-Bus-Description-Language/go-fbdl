package token

import (
	"fmt"
)

type context struct {
	line        int
	indent      int // Current indent level
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

func isBinDigit(b byte) bool {
	return '0' <= b && b <= '1'
}

func isOctalDigit(b byte) bool {
	return '0' <= b && b <= '7'
}

func isHexDigit(b byte) bool {
	return ('0' <= b && b <= '9') ||
		('a' <= b && b <= 'f') ||
		('A' <= b && b <= 'F')
}

// isValidAfterNumber returns true if character is a valid character
// after number literal.
func isValidAfterNumber(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '(', ')', ']', '-', '+', '*', '/', '%', '=', '<', '>', ';', ':', ',', '|', '&':
		return true
	}
	return false
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
			err = parseSpace(&c, src, s)
		} else if b == '\t' {
			t, err = parseTab(&c, src, s)
		} else if b == '\n' {
			t, err = parseNewline(&c, src, s)
		} else if b == '#' {
			t = parseComment(&c, src)
		} else if b == ',' {
			t, err = parseComma(&c, s)
		} else if b == ';' {
			t, err = parseSemicolon(&c, s)
		} else if b == '!' && nb == '=' {
			t = parseNonequalityOperator(&c)
		} else if b == '!' {
			t = parseNegationOperator(&c)
		} else if b == '=' && nb == '=' {
			t = parseEqualityOperator(&c)
		} else if b == '=' {
			t = parseAssignmentOperator(&c)
		} else if b == '+' {
			t = parseAdditionOperator(&c)
		} else if b == '-' {
			t = parseSubtractionOperator(&c)
		} else if b == '%' {
			t = parseRemainderOperator(&c)
		} else if b == '*' && nb == '*' {
			t = parseExponentiationOperator(&c)
		} else if b == '*' {
			t = parseMultiplicationOperator(&c)
		} else if b == '/' {
			t = parseDivisionOperator(&c)
		} else if b == '<' && nb == '=' {
			t = parseLessThanEqualOperator(&c)
		} else if b == '<' && nb == '<' {
			t = parseLeftShiftOperator(&c)
		} else if b == '<' {
			t = parseLessThanOperator(&c)
		} else if b == '>' && nb == '=' {
			t = parseGreaterThanEqualOperator(&c)
		} else if b == '>' && nb == '>' {
			t = parseRightShiftOperator(&c)
		} else if b == '>' {
			t = parseGreaterThanOperator(&c)
		} else if b == '(' {
			t = parseLeftParenthesis(&c)
		} else if b == ')' {
			t = parseRightParenthesis(&c)
		} else if b == '[' {
			t = parseLeftBracket(&c)
		} else if b == ']' {
			t = parseRightBracket(&c)
		} else if b == '&' && nb == '&' {
			t = parseLogicalAnd(&c)
		} else if b == '&' {
			t = parseBitAnd(&c)
		} else if b == '|' && nb == '|' {
			t = parseLogicalOr(&c)
		} else if b == '|' {
			t = parseBitOr(&c)
		} else if b == '"' {
			t, err = parseString(&c, src)
		} else if (b == 'b' || b == 'B') && nb == '"' {
			t, err = parseBinaryBitString(&c, src)
		} else if (b == 'o' || b == 'O') && nb == '"' {
			t, err = parseOctalBitString(&c, src)
		} else if (b == 'x' || b == 'X') && nb == '"' {
			t, err = parseHexBitString(&c, src)
		} else if isDigit(b) {
			t, err = parseNumber(&c, src)
		} else if isLetter(b) {
			t, err = parseWord(&c, src, s)
		} else {
			panic(fmt.Sprintf("unhandled byte '%c'", b))
		}

		if err != nil {
			return s, err
		}

		if t.Kind != INVALID {
			t.Line = c.line
			if t.Kind == NEWLINE {
				c.line++
			} else {
				t.Column = c.col(t.Start)
			}

			s = append(s, t)
			c.idx = t.End + 1
		}
	}

	s = append(s, Token{Kind: EOF, Start: c.idx, End: c.idx, Line: c.line, Column: c.col(c.idx)})

	return s, nil
}

func parseSpace(c *context, src []byte, s Stream) error {
	if t, ok := s.LastToken(); ok {
		if t.Kind == NEWLINE {
			return fmt.Errorf(
				"%d:%d: space character ' ' not allowed for indent", c.line, c.col(c.idx),
			)
		}
	}

	// Eat all spaces
	c.idx++
	for {
		if src[c.idx] == ' ' {
			c.idx++
		} else {
			break
		}
	}

	return nil
}

func parseTab(c *context, src []byte, s Stream) (Token, error) {
	t := Token{Kind: INVALID, Start: c.idx}

	errMsg := fmt.Sprintf(
		"%d:%d: tab character '\t' not allowed for alignment", c.line, c.col(c.idx),
	)
	if t, ok := s.LastToken(); ok {
		if t.Kind != NEWLINE {
			return t, fmt.Errorf(errMsg)
		}
	} else {
		return t, fmt.Errorf(errMsg)
	}

	indent := 1
	for {
		c.idx++
		if c.idx >= len(src) {
			break
		}

		b := src[c.idx]
		if b == '\t' {
			indent++
		} else if b == ' ' {
			return t, fmt.Errorf(
				"%d:%d: space character ' ' right after tab character '\t'",
				c.line, c.col(c.idx),
			)
		} else {
			break
		}
	}

	if indent == c.indent+1 {
		t.Kind = INDENT_INC
		t.End = c.idx - 1
	} else if indent > c.indent+1 {
		return t, fmt.Errorf(
			"%d:%d: multi indent increase",
			c.line, c.col(t.Start),
		)
	} else if indent < c.indent {
		// Insert proper number of INDENT_DEC tokens.
		t.Kind = INDENT_DEC
		t.End = t.Start
		for i := 0; indent+i < c.indent; i++ {
			s = append(s, t)
		}
		t.Kind = INVALID
	}

	c.indent = indent

	return t, nil
}

func parseNewline(c *context, src []byte, s Stream) (Token, error) {
	if prev_tok, ok := s.LastToken(); ok {
		if prev_tok.Kind == SEMICOLON {
			return Token{}, fmt.Errorf(
				"%d:%d: extra ';' at the end of line", prev_tok.Line, prev_tok.Column,
			)
		}
	}

	if nextByte(src, c.idx) != '\t' && c.indent != 0 {
		// Insert proper number of INDENT_DEC tokens.
		t := Token{Kind: INDENT_DEC, Start: c.idx, End: c.idx}
		t.End = t.Start
		for i := 0; i < c.indent; i++ {
			s = append(s, t)
		}
		c.indent = 0
	}

	t := Token{Kind: NEWLINE, Start: c.idx, End: c.idx, Column: c.col(c.idx)}
	c.newline_idx = c.idx
	return t, nil
}

func parseComment(c *context, src []byte) Token {
	t := Token{Kind: COMMENT, Start: c.idx}

	for {
		c.idx++
		if c.idx >= len(src) || src[c.idx] == '\n' {
			t.End = c.idx - 1
			return t
		}
	}
}

func parseComma(c *context, s Stream) (Token, error) {
	if prev_tok, ok := s.LastToken(); ok {
		if prev_tok.Kind == COMMA {
			return Token{}, fmt.Errorf(
				"%d:%d: redundant ','", prev_tok.Line, c.col(c.idx),
			)
		}
	}

	t := Token{Kind: COMMA, Start: c.idx, End: c.idx}
	return t, nil
}

func parseSemicolon(c *context, s Stream) (Token, error) {
	if prev_tok, ok := s.LastToken(); ok {
		if prev_tok.Kind == SEMICOLON {
			return Token{}, fmt.Errorf(
				"%d:%d: redundant ';'", prev_tok.Line, c.col(c.idx),
			)
		}
	}

	t := Token{Kind: SEMICOLON, Start: c.idx, End: c.idx}
	return t, nil
}

func parseNonequalityOperator(c *context) Token {
	return Token{Kind: NEQ, Start: c.idx, End: c.idx + 1}
}

func parseNegationOperator(c *context) Token {
	return Token{Kind: NEG, Start: c.idx, End: c.idx}
}

func parseEqualityOperator(c *context) Token {
	return Token{Kind: EQ, Start: c.idx, End: c.idx + 1}
}

func parseAssignmentOperator(c *context) Token {
	return Token{Kind: ASS, Start: c.idx, End: c.idx}
}

func parseAdditionOperator(c *context) Token {
	return Token{Kind: ADD, Start: c.idx, End: c.idx}
}

func parseSubtractionOperator(c *context) Token {
	return Token{Kind: SUB, Start: c.idx, End: c.idx}
}

func parseRemainderOperator(c *context) Token {
	return Token{Kind: REM, Start: c.idx, End: c.idx}
}

func parseExponentiationOperator(c *context) Token {
	return Token{Kind: EXP, Start: c.idx, End: c.idx + 1}
}

func parseMultiplicationOperator(c *context) Token {
	return Token{Kind: MUL, Start: c.idx, End: c.idx}
}

func parseDivisionOperator(c *context) Token {
	return Token{Kind: DIV, Start: c.idx, End: c.idx}
}

func parseLessThanEqualOperator(c *context) Token {
	return Token{Kind: LEQ, Start: c.idx, End: c.idx + 1}
}

func parseLeftShiftOperator(c *context) Token {
	return Token{Kind: SHL, Start: c.idx, End: c.idx + 1}
}

func parseLessThanOperator(c *context) Token {
	return Token{Kind: LSS, Start: c.idx, End: c.idx}
}

func parseGreaterThanEqualOperator(c *context) Token {
	return Token{Kind: GEQ, Start: c.idx, End: c.idx + 1}
}

func parseRightShiftOperator(c *context) Token {
	return Token{Kind: SHR, Start: c.idx, End: c.idx + 1}
}

func parseGreaterThanOperator(c *context) Token {
	return Token{Kind: GTR, Start: c.idx, End: c.idx}
}

func parseLeftParenthesis(c *context) Token {
	return Token{Kind: LPAREN, Start: c.idx, End: c.idx}
}

func parseRightParenthesis(c *context) Token {
	return Token{Kind: RPAREN, Start: c.idx, End: c.idx}
}

func parseLeftBracket(c *context) Token {
	return Token{Kind: LBRACK, Start: c.idx, End: c.idx}
}

func parseRightBracket(c *context) Token {
	return Token{Kind: RBRACK, Start: c.idx, End: c.idx}
}

func parseLogicalAnd(c *context) Token {
	return Token{Kind: LAND, Start: c.idx, End: c.idx + 1}
}

func parseBitAnd(c *context) Token {
	return Token{Kind: AND, Start: c.idx, End: c.idx}
}

func parseLogicalOr(c *context) Token {
	return Token{Kind: LOR, Start: c.idx, End: c.idx + 1}
}

func parseBitOr(c *context) Token {
	return Token{Kind: OR, Start: c.idx, End: c.idx}
}

func parseString(c *context, src []byte) (Token, error) {
	t := Token{Kind: STRING, Start: c.idx}

	idx := c.idx
	for {
		idx++
		if idx >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: unterminated string literal", c.line, c.col(c.idx),
			)
		}
		b := src[idx]
		if b == '"' || b == '\n' {
			break
		}
	}
	t.End = idx
	return t, nil
}

func parseBinaryBitString(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Start: c.idx}

	// Skip b"
	idx := c.idx + 2
	for {
		if idx >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in binary bit string literal",
				c.line, c.col(t.Start),
			)
		}

		b := src[idx]

		if b == '"' {
			t.End = idx
			return t, nil
		}

		switch b {
		case '0', '1',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			idx++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in binary bit string literal",
				c.line, c.col(idx), b,
			)
		}
	}
}

func parseOctalBitString(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Start: c.idx}

	// Skip o"
	idx := c.idx + 2
	for {
		if idx >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in octal bit string literal",
				c.line, c.col(t.Start),
			)
		}

		b := src[idx]

		if b == '"' {
			t.End = idx
			return t, nil
		}

		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			idx++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in octal bit string literal",
				c.line, c.col(idx), b,
			)
		}
	}
}

func parseHexBitString(c *context, src []byte) (Token, error) {
	t := Token{Kind: BIT_STRING, Start: c.idx}

	// Skip x"
	idx := c.idx + 2
	for {
		if idx >= len(src) {
			return t, fmt.Errorf(
				"%d:%d: missing terminating '\"' in hex bit string literal",
				c.line, c.col(t.Start),
			)
		}

		b := src[idx]

		if b == '"' {
			t.End = idx
			return t, nil
		}

		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			idx++
		default:
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in hex bit string literal",
				c.line, c.col(idx), b,
			)
		}
	}
}

func parseNumber(c *context, src []byte) (Token, error) {
	b := src[c.idx]
	nb := nextByte(src, c.idx)

	if b == '0' && (nb == 'b' || nb == 'B') {
		return parseBinaryInt(c, src)
	} else if b == '0' && (nb == 'o' || nb == 'O') {
		return parseOctalInt(c, src)
	} else if b == '0' && (nb == 'x' || nb == 'X') {
		return parseHexInt(c, src)
	}

	t := Token{Kind: INT, Start: c.idx}
	hasPoint := false
	hasE := false
	idx := c.idx

	for {
		idx++
		if idx >= len(src) {
			break
		}

		b := src[idx]
		if isDigit(b) {
			continue
		}

		if b == '.' {
			if hasPoint {
				return t, fmt.Errorf(
					"%d:%d: second point character '.' in number literal",
					c.line, c.col(idx),
				)
			} else {
				if hasE {
					return t, fmt.Errorf(
						"%d:%d: point character '.' after exponent in number literal",
						c.line, c.col(idx),
					)
				}
				hasPoint = true
			}
		} else if b == 'e' || b == 'E' {
			if hasE {
				return t, fmt.Errorf(
					"%d:%d: second exponent in number literal",
					c.line, c.col(idx),
				)
			} else {
				hasE = true
			}
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in number literal",
				c.line, c.col(idx), b,
			)
		}
	}

	t.End = idx - 1
	if hasPoint || hasE {
		t.Kind = REAL
	}

	return t, nil
}

func parseBinaryInt(c *context, src []byte) (Token, error) {
	t := Token{Kind: INT, Start: c.idx}

	// Skip 0b
	idx := c.idx + 2
	for {
		if idx >= len(src) {
			break
		}
		b := src[idx]
		if isBinDigit(b) {
			idx++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in binary literal",
				c.line, c.col(c.idx), b,
			)
		}
	}
	t.End = idx - 1
	return t, nil
}

func parseOctalInt(c *context, src []byte) (Token, error) {
	t := Token{Kind: INT, Start: c.idx}

	// Skip 0o
	idx := c.idx + 2
	for {
		if idx >= len(src) {
			break
		}
		b := src[idx]
		if isOctalDigit(b) {
			idx++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in octal literal",
				c.line, c.col(c.idx), b,
			)
		}
	}
	t.End = idx - 1
	return t, nil
}

func parseHexInt(c *context, src []byte) (Token, error) {
	t := Token{Kind: INT, Start: c.idx}

	// Skip 0x
	idx := c.idx + 2
	for {
		if idx >= len(src) {
			break
		}
		b := src[idx]
		if isHexDigit(b) {
			idx++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, fmt.Errorf(
				"%d:%d: invalid character '%c' in hex literal",
				c.line, c.col(c.idx), b,
			)
		}
	}
	t.End = idx - 1
	return t, nil
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
					if prev_tok.Kind != NEWLINE &&
						prev_tok.Kind != SEMICOLON &&
						prev_tok.Kind != INDENT_INC {
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
			if prev_tok.Kind != INDENT_INC && prev_tok.Kind != SEMICOLON {
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
					s[idx].End = t.End
					t.Kind = INVALID
					c.idx = t.End + 1
				}
			}
		}
	}

	return t, nil
}

func parseKeyword(word []byte, c *context) Token {
	t := Token{Kind: INVALID, Start: c.idx, End: c.idx + len(word) - 1}

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
	t := Token{Kind: INVALID, Start: c.idx, End: c.idx + len(word) - 1}

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
