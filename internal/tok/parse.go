package tok

import (
	"bytes"
	"fmt"
	"unicode"
)

// Parsing context
type ctx struct {
	line       int
	indent     int // Current indent level
	i          int // Current buffer index
	newlineIdx int // Last newline index
}

// Col returns column number for given index.
func (c ctx) col(idx int) int {
	return idx - c.newlineIdx
}

// nextByte returns byte with index equal idx + 1.
// If (idx + 1) >= len(src), then 0 is returned.
func nextByte(src []byte, idx int) byte {
	if idx+1 >= len(src) {
		return 0
	}
	return src[idx+1]
}

// lastToken returns last token from the Token list.
// If list is empty, the second return is false.
func lastToken(toks []Token) (Token, bool) {
	if len(toks) == 0 {
		return nil, false
	}
	return toks[len(toks)-1], true
}

// getWord returns word from the source starting from index idx.
// The function assumes byte under idx is not a whitespace character.
// The second return is true if word contains hyphen '-' character.
// The third return is true if word contains dot '.' character.
func getWord(src []byte, idx int) ([]byte, bool, bool) {
	hasHyphen := false
	hasDot := false
	end_idx := idx

	for {
		if end_idx >= len(src) {
			return src[idx:end_idx], hasHyphen, hasDot
		}

		b := src[end_idx]
		if isLetter(b) || isDigit(b) || b == '_' || b == '-' || b == '.' {
			if b == '-' {
				hasHyphen = true
			} else if b == '.' {
				hasDot = true
			}
			end_idx++
			continue
		} else {
			return src[idx:end_idx], hasHyphen, hasDot
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
func Parse(src []byte) ([]Token, error) {
	var (
		c   ctx
		err error
		s   []Token
		t   Token
	)
	c.line = 1
	c.newlineIdx = -1

	for {
		if c.i == len(src) {
			break
		}

		t = None{}
		err = nil
		b := src[c.i]            // Current byte
		nb := nextByte(src, c.i) // Next byte

		if b == ' ' {
			err = parseSpace(&c, src, s)
		} else if b == '\t' {
			err = parseTab(&c, src, &s)
		} else if b == '\n' {
			err = parseNewline(&c, src, &s)
		} else if b == '#' {
			t = parseComment(&c, src, s)
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
			t, err = parseWord(&c, src, &s)
		} else {
			panic(fmt.Sprintf("unhandled byte '%c'", b))
		}

		if err != nil {
			return s, err
		}

		if _, ok := t.(None); !ok {
			s = append(s, t)
		}
	}

	s = append(s, Eof{start: c.i, end: c.i, line: c.line, column: c.col(c.i)})

	return s, nil
}

func parseSpace(c *ctx, src []byte, s []Token) error {
	if t, ok := lastToken(s); ok {
		if _, ok := t.(Newline); ok {
			return Error{
				Indent{c.i, c.i, c.line, c.col(c.i)},
				"space character ' ' not allowed for indent",
			}
		}
	}

	// Eat all spaces
	c.i++
	for {
		if src[c.i] == ' ' {
			c.i++
		} else {
			break
		}
	}

	return nil
}

func parseTab(c *ctx, src []byte, s *[]Token) error {
	start := c.i

	tab := Indent{start, start, c.line, c.col(c.i)}
	errMsg := "tab character '\\t' not allowed for alignment"
	if t, ok := lastToken(*s); ok {
		if _, ok := t.(Newline); !ok {
			return Error{tab, errMsg}
		}
	} else {
		return Error{tab, errMsg}
	}

	indent := 1
	for {
		c.i++
		if c.i >= len(src) {
			break
		}

		b := src[c.i]
		if b == '\t' {
			indent++
		} else if b == ' ' {
			return Error{
				Indent{c.i, c.i, c.line, c.col(c.i)},
				"space character ' ' right after tab character '\\t'",
			}
		} else {
			break
		}
	}

	if indent == c.indent+1 {
		t := Indent{start, c.i - 1, c.line, c.col(start)}
		*s = append(*s, t)
	} else if indent > c.indent+1 {
		return Error{
			Indent{start, start, c.line, c.col(start)}, "multi indent increase",
		}
	} else if indent < c.indent {
		// Insert proper number of INDENT_DEC tokens.
		t := Dedent{start, start, c.line, c.col(start)}
		for i := 0; indent+i < c.indent; i++ {
			*s = append(*s, t)
		}
	}

	c.indent = indent

	return nil
}

func parseNewline(c *ctx, src []byte, s *[]Token) error {
	if t, ok := lastToken(*s); ok {
		if _, ok := t.(Semicolon); ok {
			return Error{t, "extra ';' at line end"}
		}
	}

	nl := Newline{c.i, c.i, c.line, c.col(c.i)}

	// Eat all newlines
	for {
		c.newlineIdx = c.i
		c.line++
		c.i++
		if c.i == len(src) || src[c.i] != '\n' {
			break
		}
		nl.end++
	}

	*s = append(*s, nl)

	if c.i < len(src) && src[c.i] != '\t' && c.indent != 0 {
		// Insert proper number of Dedent tokens.
		t := Dedent{c.i, c.i, c.line, c.col(c.i)}
		for i := 0; i < c.indent; i++ {
			*s = append(*s, t)
		}
		c.indent = 0
	}

	return nil
}

func parseComment(c *ctx, src []byte, s []Token) Token {
	t := Comment{start: c.i, line: c.line, column: c.col(c.i)}

	for {
		c.i++
		if c.i >= len(src) || src[c.i] == '\n' {
			t.end = c.i - 1
			break
		}
	}

	// Add comment to the token stream only if it is a potential documentation comment.
	if prevTok, ok := lastToken(s); ok {
		switch prevTok.(type) {
		case Newline, Indent, Dedent:
			return t
		}
	} else {
		return t
	}

	return None{}
}

func parseComma(c *ctx, s []Token) (Token, error) {
	if t, ok := lastToken(s); ok {
		if _, ok := t.(Comma); ok {
			return nil, Error{
				Comma{c.i, c.i, c.line, c.col(c.i)}, "redundant ','",
			}
		}
	}

	t := Comma{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return t, nil
}

func parseSemicolon(c *ctx, s []Token) (Token, error) {
	if t, ok := lastToken(s); ok {
		if _, ok := t.(Semicolon); ok {
			return nil, Error{
				Semicolon{c.i, c.i, c.line, c.col(c.i)}, "redundant ';'",
			}
		}
	}

	t := Semicolon{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return t, nil
}

func parseNonequalityOperator(c *ctx) Neq {
	n := Neq{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return n
}

func parseNegationOperator(c *ctx) Neg {
	n := Neg{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return n
}

func parseEqualityOperator(c *ctx) Eq {
	e := Eq{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return e
}

func parseAssignmentOperator(c *ctx) Ass {
	a := Ass{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return a
}

func parseAdditionOperator(c *ctx) Add {
	a := Add{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return a
}

func parseSubtractionOperator(c *ctx) Sub {
	s := Sub{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return s
}

func parseRemainderOperator(c *ctx) Rem {
	r := Rem{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return r
}

func parseExponentiationOperator(c *ctx) Exp {
	e := Exp{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return e
}

func parseMultiplicationOperator(c *ctx) Mul {
	m := Mul{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return m
}

func parseDivisionOperator(c *ctx) Div {
	d := Div{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return d
}

func parseLessThanEqualOperator(c *ctx) LessEq {
	le := LessEq{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return le
}

func parseLeftShiftOperator(c *ctx) LeftShift {
	ls := LeftShift{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return ls
}

func parseLessThanOperator(c *ctx) Less {
	l := Less{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return l
}

func parseGreaterThanEqualOperator(c *ctx) GreaterEq {
	ge := GreaterEq{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return ge
}

func parseRightShiftOperator(c *ctx) RightShift {
	rs := RightShift{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return rs
}

func parseGreaterThanOperator(c *ctx) Greater {
	g := Greater{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return g
}

func parseLeftParenthesis(c *ctx) LeftParen {
	lp := LeftParen{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return lp
}

func parseRightParenthesis(c *ctx) RightParen {
	rp := RightParen{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return rp
}

func parseLeftBracket(c *ctx) LeftBracket {
	lb := LeftBracket{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return lb
}

func parseRightBracket(c *ctx) RightBracket {
	rb := RightBracket{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return rb
}

func parseLogicalAnd(c *ctx) And {
	a := And{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return a
}

func parseBitAnd(c *ctx) BitAnd {
	ba := BitAnd{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return ba
}

func parseLogicalOr(c *ctx) Or {
	o := Or{c.i, c.i + 1, c.line, c.col(c.i)}
	c.i += 2
	return o
}

func parseBitOr(c *ctx) BitOr {
	bo := BitOr{c.i, c.i, c.line, c.col(c.i)}
	c.i++
	return bo
}

func parseString(c *ctx, src []byte) (String, error) {
	t := String{c.i, c.i, c.line, c.col(c.i)}

	for {
		c.i++
		if c.i >= len(src) {
			return t, Error{t, "unterminated string, probably missing '\"'"}
		}
		b := src[c.i]
		if b != '\n' {
			t.end++
		}
		if b == '"' {
			break
		}
	}
	c.i++
	return t, nil
}

func parseBinaryBitString(c *ctx, src []byte) (Token, error) {
	t := BitString{c.i, c.i + 1, c.line, c.col(c.i)}

	// Skip b"
	c.i += 2
	for {
		if c.i >= len(src) {
			return t, Error{t, "unterminated binary bit string, probably missing '\"'"}
		}

		switch b := src[c.i]; b {
		case '"':
			t.end++
			c.i++
			return t, nil
		case '0', '1',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			c.i++
		default:
			switch b {
			case ' ', '\n', ';', ',':
				return t, Error{t, "unterminated binary bit string, probably missing '\"'"}
			default:
				return t, Error{
					BitString{c.i, c.i, c.line, c.col(c.i)},
					fmt.Sprintf("invalid character '%c' in binary bit string", b),
				}
			}
		}
	}
}

func parseOctalBitString(c *ctx, src []byte) (Token, error) {
	t := BitString{c.i, c.i + 1, c.line, c.col(c.i)}

	// Skip o"
	c.i += 2
	for {
		if c.i >= len(src) {
			return t, Error{t, "unterminated octal bit string, probably missing '\"'"}
		}

		switch b := src[c.i]; b {
		case '"':
			t.end++
			c.i++
			return t, nil
		case '0', '1', '2', '3', '4', '5', '6', '7',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			c.i++
		default:
			switch b {
			case ' ', '\n', ';', ',':
				return t, Error{t, "unterminated octal bit string, probably missing '\"'"}
			default:
				return t, Error{
					BitString{c.i, c.i, c.line, c.col(c.i)},
					fmt.Sprintf("invalid character '%c' in octal bit string", b),
				}
			}
		}
	}
}

func parseHexBitString(c *ctx, src []byte) (Token, error) {
	t := BitString{c.i, c.i + 1, c.line, c.col(c.i)}

	// Skip x"
	c.i += 2
	for {
		if c.i >= len(src) {
			return t, Error{t, "unterminated hex bit string, probably missing '\"'"}
		}

		switch b := src[c.i]; b {
		case '"':
			t.end++
			c.i++
			return t, nil
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			c.i++
		case ' ', '\n', ';', ',':
			return t, Error{t, "unterminated hex bit string, probably missing '\"'"}
		default:
			return t, Error{
				BitString{c.i, c.i, c.line, c.col(c.i)},
				fmt.Sprintf("invalid character '%c' in hex bit string", b),
			}
		}
	}
}

func parseNumber(c *ctx, src []byte) (Number, error) {
	b := src[c.i]
	nb := nextByte(src, c.i)

	if b == '0' && (nb == 'b' || nb == 'B') {
		return parseBinaryInt(c, src)
	} else if b == '0' && (nb == 'o' || nb == 'O') {
		return parseOctalInt(c, src)
	} else if b == '0' && (nb == 'x' || nb == 'X') {
		return parseHexInt(c, src)
	}

	i := Int{start: c.i, line: c.line, column: c.col(c.i)}
	hasPoint := false
	hasE := false

	for {
		c.i++
		if c.i >= len(src) {
			break
		}

		b := src[c.i]
		if isDigit(b) {
			continue
		}

		if b == '.' {
			if hasPoint {
				return nil, Error{
					Real{c.i, c.i, c.line, c.col(c.i)},
					"second point character '.' in number",
				}
			} else {
				if hasE {
					return nil, Error{
						Real{c.i, c.i, c.line, c.col(c.i)},
						"point character '.' after exponent in number",
					}
				}
				hasPoint = true
			}
		} else if b == 'e' || b == 'E' {
			if hasE {
				return nil, Error{
					Real{c.i, c.i, c.line, c.col(c.i)},
					"second exponent in number",
				}
			} else {
				hasE = true
			}
		} else if isValidAfterNumber(b) {
			break
		} else {
			return nil, Error{
				Int{c.i, c.i, c.line, c.col(c.i)},
				fmt.Sprintf("invalid character '%c' in number", b),
			}
		}
	}

	i.end = c.i - 1
	var n Number = i
	if hasPoint || hasE {
		n = Real(i)
	}

	return n, nil
}

func parseBinaryInt(c *ctx, src []byte) (Int, error) {
	t := Int{start: c.i, line: c.line, column: c.col(c.i)}

	// Skip 0b
	c.i += 2
	for {
		if c.i >= len(src) {
			break
		}
		b := src[c.i]
		if isBinDigit(b) {
			c.i++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, Error{
				Int{c.i, c.i, c.line, c.col(c.i)},
				fmt.Sprintf("invalid character '%c' in binary", b),
			}
		}
	}
	t.end = c.i - 1
	return t, nil
}

func parseOctalInt(c *ctx, src []byte) (Int, error) {
	t := Int{start: c.i, line: c.line, column: c.col(c.i)}

	// Skip 0o
	c.i += 2
	for {
		if c.i >= len(src) {
			break
		}
		b := src[c.i]
		if isOctalDigit(b) {
			c.i++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, Error{
				Int{c.i, c.i, c.line, c.col(c.i)},
				fmt.Sprintf("invalid character '%c' in octal", b),
			}
		}
	}
	t.end = c.i - 1
	return t, nil
}

func parseHexInt(c *ctx, src []byte) (Int, error) {
	t := Int{start: c.i, line: c.line, column: c.col(c.i)}

	// Skip 0x
	c.i += 2
	for {
		if c.i >= len(src) {
			break
		}
		b := src[c.i]
		if isHexDigit(b) {
			c.i++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, Error{
				Int{c.i, c.i, c.line, c.col(c.i)},
				fmt.Sprintf("invalid character '%c' in hex", b),
			}
		}
	}
	t.end = c.i - 1
	return t, nil
}

// isValidQualifiedIdentifier returns false if symbol name in
// qualified identifier does not start with upper case letter.
func isValidQualifiedIdentifier(qi []byte) bool {
	aux := bytes.Split(qi, []byte("."))
	sym := aux[1]
	return unicode.IsUpper([]rune(string(sym))[0])
}

// TODO: Refactor, too complex, split into 2 (or more) functions.
func parseWord(c *ctx, src []byte, s *[]Token) (Token, error) {
	var t Token
	defer func() { c.i = t.End() + 1 }()
	word, hasHyphen, hasDot := getWord(src, c.i)

	qualIdentErrMsg := "symbol name in qualified identifier must start with upper case letter"
	if hasHyphen && hasDot {
		// This is for sure part of an expression
		chunks := bytes.Split(word, []byte{'-'})
		for i, chunk := range chunks {
			if bytes.Contains(chunk, []byte{'.'}) {
				t = QualIdent{
					start: c.i, end: c.i + len(chunk) - 1, line: c.line, column: c.col(c.i),
				}
				if !isValidQualifiedIdentifier(chunk) {
					return t, Error{t, qualIdentErrMsg}
				}
			} else {
				t = Ident{
					start: c.i, end: c.i + len(chunk) - 1, line: c.line, column: c.col(c.i),
				}
			}
			if i == len(chunks)-1 {
				return t, nil
			}
			*s = append(*s, t)
			c.i += len(chunks[i])
			t = Sub{start: c.i, end: c.i, line: c.line, column: c.col(c.i)}
			*s = append(*s, t)
			c.i++
		}
	} else if hasDot {
		// It is qualified identifier
		t = QualIdent{start: c.i, end: c.i + len(word) - 1, line: c.line, column: c.col(c.i)}

		if !isValidQualifiedIdentifier(word) {
			return t, Error{t, qualIdentErrMsg}
		}

		return t, nil
	}

	splitHyphenatedWord := func() (Ident, Sub, Ident) {
		i1 := Ident{start: c.i, line: c.line, column: c.col(c.i)}
		s := Sub{line: c.line}
		i2 := Ident{end: c.i + len(word) - 1, line: c.line}

		for i := 0; i < len(word); i++ {
			if word[i] == '-' {
				i1.end = c.i + i - 1

				s.start = c.i + i
				s.end = c.i + i
				s.column = c.col(c.i + i)

				i2.start = c.i + i + 1
				i2.column = c.col(c.i + i + 1)
			}
		}
		return i1, s, i2
	}

	if !hasHyphen {
		// Firstly assume word is a keyword
		t = parseKeyword(word, c)
		// If it is not a keyword, then it might be a property or identifier.
		if _, ok := t.(None); ok {
			t = parseProperty(word, c)
			// If it is not property, then it must be an identifier.
			if _, ok := t.(None); ok {
				t = Ident{t.Start(), t.End(), t.Line(), t.Column()}
			} else {
				// However, properties are properties only if they are in valid place,
				// otherwise, these are regular identifiers.
				if prevTok, ok := lastToken(*s); ok {
					switch prevTok.(type) {
					case Newline, Semicolon, Indent:
						// Do nothing, this is property
					default:
						t = Ident{t.Start(), t.End(), t.Line(), t.Column()}
					}
				}
			}
		}

		// Allow functionality keywords to be instantiation names
		if _, ok := t.(Functionality); ok {
			if prevTok, ok := lastToken(*s); ok {
				switch prevTok.(type) {
				case Newline, Indent, Dedent:
					t = Ident{t.Start(), t.End(), t.Line(), t.Column()}
				}
			}
		}
	} else {
		// Firstly assume word is a property
		t = parseProperty(word, c)
		// If it is not property, then it is part of an expression.
		if _, ok := t.(None); ok {
			// t is last, as deferred function has to calculate new context index.
			i1, sub, i2 := splitHyphenatedWord()
			t = i2
			*s = append(*s, []Token{i1, sub, i2}...)
			// Assing to t for updating context index in deferred function
			t = i2
			return None{}, nil
		} else {
			// It might be property, or part of an expression.
			prevTok, ok := lastToken(*s)
			if !ok {
				// Safe to return, time literal units do not contain hyphen '-'.
				return t, nil
			}
			// It is part of an expression.
			switch prevTok.(type) {
			case Newline, Indent, Semicolon:
				// It is property
			default:
				i1, sub, i2 := splitHyphenatedWord()
				*s = append(*s, []Token{i1, sub, i2}...)
				// Assing to t for updating context index in deferred function
				t = i2
				return None{}, nil
			}
		}
	}

	// The word might be the unit of time literal
	if _, ok := t.(Ident); ok {
		if prevTok, ok := lastToken(*s); ok {
			if _, ok := prevTok.(Int); ok {
				switch string(word) {
				case "ns", "us", "ms", "s":
					t = Time{
						start:  prevTok.Start(),
						end:    t.End(),
						line:   prevTok.Line(),
						column: prevTok.Column(),
					}
					// Remove previous Int from the list
					// New Time token will be inserted
					*s = (*s)[:len(*s)-1]
				}
			}
		}
	}

	return t, nil
}

func parseKeyword(word []byte, c *ctx) Token {
	s := c.i
	e := s + len(word) - 1
	l := c.line
	col := c.col(c.i)

	switch string(word) {
	case "false", "true":
		return Bool{s, e, l, col}
	case "block":
		return Block{s, e, l, col}
	case "bus":
		return Bus{s, e, l, col}
	case "config":
		return Config{s, e, l, col}
	case "const":
		return Const{s, e, l, col}
	case "import":
		return Import{s, e, l, col}
	case "irq":
		return Irq{s, e, l, col}
	case "mask":
		return Mask{s, e, l, col}
	case "memory":
		return Memory{s, e, l, col}
	case "param":
		return Param{s, e, l, col}
	case "proc":
		return Proc{s, e, l, col}
	case "return":
		return Return{s, e, l, col}
	case "static":
		return Static{s, e, l, col}
	case "status":
		return Status{s, e, l, col}
	case "stream":
		return Stream{s, e, l, col}
	case "type":
		return Type{s, e, l, col}
	}

	return None{s, e, l, col}
}

func parseProperty(word []byte, c *ctx) Token {
	s := c.i
	e := s + len(word) - 1
	l := c.line
	col := c.col(c.i)

	switch string(word) {
	case "access":
		return Access{s, e, l, col}
	case "add-enable":
		return AddEnable{s, e, l, col}
	case "atomic":
		return Atomic{s, e, l, col}
	case "byte-write-enable":
		return ByteWriteEnable{s, e, l, col}
	case "clear":
		return Clear{s, e, l, col}
	case "delay":
		return Delay{s, e, l, col}
	case "enable-init-value":
		return EnableInitValue{s, e, l, col}
	case "enable-reset-value":
		return EnableResetValue{s, e, l, col}
	case "groups":
		return Groups{s, e, l, col}
	case "init-value":
		return InitValue{s, e, l, col}
	case "in-trigger":
		return InTrigger{s, e, l, col}
	case "masters":
		return Masters{s, e, l, col}
	case "out-trigger":
		return OutTrigger{s, e, l, col}
	case "range":
		return Range{s, e, l, col}
	case "read-latency":
		return ReadLatency{s, e, l, col}
	case "read-value":
		return ReadValue{s, e, l, col}
	case "reset":
		return Reset{s, e, l, col}
	case "reset-value":
		return ResetValue{s, e, l, col}
	case "size":
		return Size{s, e, l, col}
	case "width":
		return Width{s, e, l, col}
	}

	return None{s, e, l, col}
}
