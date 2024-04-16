package tok

import (
	"bytes"
	"fmt"
	"unicode"
)

// Parsing context
type context struct {
	line   int // Current line number
	indent int // Current indent level
	idx    int // Current buffer index
	nlIdx  int // Last newline index
}

// Col returns column number for given index.
func (ctx context) col(idx int) int {
	return idx - ctx.nlIdx
}

// Creates position from the current context state.
func (ctx context) pos() position {
	return position{ctx.idx, ctx.idx, ctx.line, ctx.col(ctx.idx)}
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
	endIdx := idx

	for {
		if endIdx >= len(src) {
			return src[idx:endIdx], hasHyphen, hasDot
		}

		b := src[endIdx]
		if isLetter(b) || isDigit(b) || b == '_' || b == '-' || b == '.' {
			if b == '-' {
				hasHyphen = true
			} else if b == '.' {
				hasDot = true
			}
			endIdx++
			continue
		} else {
			return src[idx:endIdx], hasHyphen, hasDot
		}
	}
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
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
		ctx  context
		tok  Token
		err  error
		toks []Token // Token stream
	)
	ctx.line = 1
	ctx.nlIdx = -1

	for {
		if ctx.idx == len(src) {
			break
		}

		tok = None{}
		err = nil
		b := src[ctx.idx]            // Current byte
		nb := nextByte(src, ctx.idx) // Next byte

		if b == ' ' {
			err = parseSpace(&ctx, src, toks)
		} else if b == '\t' {
			err = parseTab(&ctx, src, &toks)
		} else if b == '\n' {
			err = parseNewline(&ctx, src, &toks)
		} else if b == '#' {
			tok = parseComment(&ctx, src, toks)
		} else if b == ',' {
			tok, err = parseComma(&ctx, toks)
		} else if b == ';' {
			tok, err = parseSemicolon(&ctx, toks)
		} else if b == '!' && nb == '=' {
			tok = parseNonequalityOperator(&ctx)
		} else if b == '!' {
			tok = parseNegationOperator(&ctx)
		} else if b == '=' && nb == '=' {
			tok = parseEqualityOperator(&ctx)
		} else if b == '=' {
			tok = parseAssignmentOperator(&ctx)
		} else if b == '+' {
			tok = parseAdditionOperator(&ctx)
		} else if b == '-' {
			tok = parseSubtractionOperator(&ctx)
		} else if b == '%' {
			tok = parseRemainderOperator(&ctx)
		} else if b == '*' && nb == '*' {
			tok = parseExponentiationOperator(&ctx)
		} else if b == '*' {
			tok = parseMultiplicationOperator(&ctx)
		} else if b == '/' {
			tok = parseDivisionOperator(&ctx)
		} else if b == '<' && nb == '=' {
			tok = parseLessThanEqualOperator(&ctx)
		} else if b == '<' && nb == '<' {
			tok = parseLeftShiftOperator(&ctx)
		} else if b == '<' {
			tok = parseLessThanOperator(&ctx)
		} else if b == '>' && nb == '=' {
			tok = parseGreaterThanEqualOperator(&ctx)
		} else if b == '>' && nb == '>' {
			tok = parseRightShiftOperator(&ctx)
		} else if b == '>' {
			tok = parseGreaterThanOperator(&ctx)
		} else if b == '(' {
			tok = parseLeftParenthesis(&ctx)
		} else if b == ')' {
			tok = parseRightParenthesis(&ctx)
		} else if b == '[' {
			tok = parseLeftBracket(&ctx)
		} else if b == ']' {
			tok = parseRightBracket(&ctx)
		} else if b == '&' && nb == '&' {
			tok = parseLogicalAnd(&ctx)
		} else if b == '&' {
			tok = parseBitAnd(&ctx)
		} else if b == '|' && nb == '|' {
			tok = parseLogicalOr(&ctx)
		} else if b == '|' {
			tok = parseBitOr(&ctx)
		} else if b == '"' {
			tok, err = parseString(&ctx, src)
		} else if (b == 'b' || b == 'B') && nb == '"' {
			tok, err = parseBinBitString(&ctx, src)
		} else if (b == 'o' || b == 'O') && nb == '"' {
			tok, err = parseOctalBitString(&ctx, src)
		} else if (b == 'x' || b == 'X') && nb == '"' {
			tok, err = parseHexBitString(&ctx, src)
		} else if isDigit(b) {
			tok, err = parseNumber(&ctx, src)
		} else if isLetter(b) {
			tok, err = parseWord(&ctx, src, &toks)
		} else {
			panic(fmt.Sprintf("unhandled byte '%c'", b))
		}

		if err != nil {
			return toks, err
		}

		if _, ok := tok.(None); !ok {
			toks = append(toks, tok)
		}
	}

	toks = append(toks, Eof{ctx.pos()})

	return toks, nil
}

func parseSpace(ctx *context, src []byte, toks []Token) error {
	if t, ok := lastToken(toks); ok {
		if _, ok := t.(Newline); ok {
			return Error{
				Indent{ctx.pos()},
				"space character ' ' not allowed for indent",
			}
		}
	}

	// Eat all spaces
	ctx.idx++
	for {
		if src[ctx.idx] == ' ' {
			ctx.idx++
		} else {
			break
		}
	}

	return nil
}

func parseTab(ctx *context, src []byte, toks *[]Token) error {
	tab := Indent{ctx.pos()}
	start := ctx.idx
	errMsg := "tab character '\\t' not allowed for alignment"

	if t, ok := lastToken(*toks); ok {
		if _, ok := t.(Newline); !ok {
			return Error{tab, errMsg}
		}
	} else {
		return Error{tab, errMsg}
	}

	indent := 1
	for {
		ctx.idx++
		if ctx.idx >= len(src) {
			break
		}

		b := src[ctx.idx]
		if b == '\t' {
			indent++
		} else if b == ' ' {
			return Error{
				Indent{ctx.pos()},
				"space character ' ' right after tab character '\\t'",
			}
		} else {
			break
		}
	}

	if indent == ctx.indent+1 {
		t := Indent{position{start, ctx.idx - 1, ctx.line, ctx.col(start)}}
		*toks = append(*toks, t)
	} else if indent > ctx.indent+1 {
		return Error{
			Indent{position{start, start, ctx.line, ctx.col(start)}}, "multi indent increase",
		}
	} else if indent < ctx.indent {
		// Insert proper number of INDENT_DEC tokens.
		t := Dedent{position{start, start, ctx.line, ctx.col(start)}}
		for i := 0; indent+i < ctx.indent; i++ {
			*toks = append(*toks, t)
		}
	}

	ctx.indent = indent

	return nil
}

func parseNewline(ctx *context, src []byte, toks *[]Token) error {
	if t, ok := lastToken(*toks); ok {
		if _, ok := t.(Semicolon); ok {
			return Error{t, "extra ';' at line end"}
		}
	}

	nl := Newline{ctx.pos()}

	// Eat all newlines
	for {
		ctx.nlIdx = ctx.idx
		ctx.line++
		ctx.idx++
		if ctx.idx == len(src) || src[ctx.idx] != '\n' {
			break
		}
		nl.end++
	}

	*toks = append(*toks, nl)

	if ctx.idx < len(src) && src[ctx.idx] != '\t' && ctx.indent != 0 {
		// Insert proper number of Dedent tokens.
		t := Dedent{ctx.pos()}
		for i := 0; i < ctx.indent; i++ {
			*toks = append(*toks, t)
		}
		ctx.indent = 0
	}

	return nil
}

func parseComment(ctx *context, src []byte, toks []Token) Token {
	t := Comment{ctx.pos()}

	for {
		ctx.idx++
		if ctx.idx >= len(src) || src[ctx.idx] == '\n' {
			t.end = ctx.idx - 1
			break
		}
	}

	// Add comment to the token stream only if it is a potential documentation comment.
	if prevTok, ok := lastToken(toks); ok {
		switch prevTok.(type) {
		case Newline, Indent, Dedent:
			return t
		}
	} else {
		return t
	}

	return None{}
}

func parseComma(ctx *context, toks []Token) (Token, error) {
	if t, ok := lastToken(toks); ok {
		if _, ok := t.(Comma); ok {
			return nil, Error{
				Comma{ctx.pos()}, "redundant ','",
			}
		}
	}

	t := Comma{ctx.pos()}
	ctx.idx++
	return t, nil
}

func parseSemicolon(ctx *context, toks []Token) (Token, error) {
	if t, ok := lastToken(toks); ok {
		if _, ok := t.(Semicolon); ok {
			return nil, Error{
				Semicolon{ctx.pos()}, "redundant ';'",
			}
		}
	}

	t := Semicolon{ctx.pos()}
	ctx.idx++
	return t, nil
}

func parseNonequalityOperator(ctx *context) Neq {
	n := Neq{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return n
}

func parseNegationOperator(ctx *context) Neg {
	n := Neg{ctx.pos()}
	ctx.idx++
	return n
}

func parseEqualityOperator(ctx *context) Eq {
	e := Eq{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return e
}

func parseAssignmentOperator(ctx *context) Ass {
	a := Ass{ctx.pos()}
	ctx.idx++
	return a
}

func parseAdditionOperator(ctx *context) Add {
	a := Add{ctx.pos()}
	ctx.idx++
	return a
}

func parseSubtractionOperator(ctx *context) Sub {
	toks := Sub{ctx.pos()}
	ctx.idx++
	return toks
}

func parseRemainderOperator(ctx *context) Rem {
	r := Rem{ctx.pos()}
	ctx.idx++
	return r
}

func parseExponentiationOperator(ctx *context) Exp {
	e := Exp{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return e
}

func parseMultiplicationOperator(ctx *context) Mul {
	m := Mul{ctx.pos()}
	ctx.idx++
	return m
}

func parseDivisionOperator(ctx *context) Div {
	d := Div{ctx.pos()}
	ctx.idx++
	return d
}

func parseLessThanEqualOperator(ctx *context) LessEq {
	le := LessEq{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return le
}

func parseLeftShiftOperator(ctx *context) LeftShift {
	ls := LeftShift{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return ls
}

func parseLessThanOperator(ctx *context) Less {
	l := Less{ctx.pos()}
	ctx.idx++
	return l
}

func parseGreaterThanEqualOperator(ctx *context) GreaterEq {
	ge := GreaterEq{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return ge
}

func parseRightShiftOperator(ctx *context) RightShift {
	rs := RightShift{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return rs
}

func parseGreaterThanOperator(ctx *context) Greater {
	g := Greater{ctx.pos()}
	ctx.idx++
	return g
}

func parseLeftParenthesis(ctx *context) LeftParen {
	lp := LeftParen{ctx.pos()}
	ctx.idx++
	return lp
}

func parseRightParenthesis(ctx *context) RightParen {
	rp := RightParen{ctx.pos()}
	ctx.idx++
	return rp
}

func parseLeftBracket(ctx *context) LeftBracket {
	lb := LeftBracket{ctx.pos()}
	ctx.idx++
	return lb
}

func parseRightBracket(ctx *context) RightBracket {
	rb := RightBracket{ctx.pos()}
	ctx.idx++
	return rb
}

func parseLogicalAnd(ctx *context) And {
	a := And{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return a
}

func parseBitAnd(ctx *context) BitAnd {
	ba := BitAnd{ctx.pos()}
	ctx.idx++
	return ba
}

func parseLogicalOr(ctx *context) Or {
	o := Or{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}
	ctx.idx += 2
	return o
}

func parseBitOr(ctx *context) BitOr {
	bo := BitOr{ctx.pos()}
	ctx.idx++
	return bo
}

func parseString(ctx *context, src []byte) (String, error) {
	t := String{ctx.pos()}

	for {
		ctx.idx++
		if ctx.idx >= len(src) {
			return t, Error{t, "unterminated string, probably missing '\"'"}
		}
		b := src[ctx.idx]
		if b != '\n' {
			t.end++
		}
		if b == '"' {
			break
		}
	}
	ctx.idx++
	return t, nil
}

func parseBinBitString(ctx *context, src []byte) (Token, error) {
	t := BitString{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}

	// Skip b"
	ctx.idx += 2
	for {
		if ctx.idx >= len(src) {
			return t, Error{t, "unterminated binary bit string, probably missing '\"'"}
		}

		switch b := src[ctx.idx]; b {
		case '"':
			t.end++
			ctx.idx++
			return t, nil
		case '0', '1',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			ctx.idx++
		default:
			switch b {
			case ' ', '\n', ';', ',':
				return t, Error{t, "unterminated binary bit string, probably missing '\"'"}
			default:
				return t, Error{
					BitString{ctx.pos()},
					fmt.Sprintf("invalid character '%c' in binary bit string", b),
				}
			}
		}
	}
}

func parseOctalBitString(ctx *context, src []byte) (Token, error) {
	t := BitString{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}

	// Skip o"
	ctx.idx += 2
	for {
		if ctx.idx >= len(src) {
			return t, Error{t, "unterminated octal bit string, probably missing '\"'"}
		}

		switch b := src[ctx.idx]; b {
		case '"':
			t.end++
			ctx.idx++
			return t, nil
		case '0', '1', '2', '3', '4', '5', '6', '7',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			ctx.idx++
		default:
			switch b {
			case ' ', '\n', ';', ',':
				return t, Error{t, "unterminated octal bit string, probably missing '\"'"}
			default:
				return t, Error{
					BitString{ctx.pos()},
					fmt.Sprintf("invalid character '%c' in octal bit string", b),
				}
			}
		}
	}
}

func parseHexBitString(ctx *context, src []byte) (Token, error) {
	t := BitString{position{ctx.idx, ctx.idx + 1, ctx.line, ctx.col(ctx.idx)}}

	// Skip x"
	ctx.idx += 2
	for {
		if ctx.idx >= len(src) {
			return t, Error{t, "unterminated hex bit string, probably missing '\"'"}
		}

		switch b := src[ctx.idx]; b {
		case '"':
			t.end++
			ctx.idx++
			return t, nil
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			ctx.idx++
		case ' ', '\n', ';', ',':
			return t, Error{t, "unterminated hex bit string, probably missing '\"'"}
		default:
			return t, Error{
				BitString{ctx.pos()},
				fmt.Sprintf("invalid character '%c' in hex bit string", b),
			}
		}
	}
}

func parseNumber(ctx *context, src []byte) (Number, error) {
	b := src[ctx.idx]
	nb := nextByte(src, ctx.idx)

	if b == '0' && (nb == 'b' || nb == 'B') {
		return parseBinInt(ctx, src)
	} else if b == '0' && (nb == 'o' || nb == 'O') {
		return parseOctalInt(ctx, src)
	} else if b == '0' && (nb == 'x' || nb == 'X') {
		return parseHexInt(ctx, src)
	}

	i := Int{ctx.pos()}
	hasPoint := false
	hasE := false

	for {
		ctx.idx++
		if ctx.idx >= len(src) {
			break
		}

		b := src[ctx.idx]
		if isDigit(b) {
			continue
		}

		if b == '.' {
			if hasPoint {
				return nil, Error{
					Real{ctx.pos()},
					"second point character '.' in number",
				}
			} else {
				if hasE {
					return nil, Error{
						Real{ctx.pos()},
						"point character '.' after exponent in number",
					}
				}
				hasPoint = true
			}
		} else if b == 'e' || b == 'E' {
			if hasE {
				return nil, Error{
					Real{ctx.pos()},
					"second exponent in number",
				}
			} else {
				hasE = true
			}
		} else if isValidAfterNumber(b) {
			break
		} else {
			return nil, Error{
				Int{ctx.pos()},
				fmt.Sprintf("invalid character '%c' in number", b),
			}
		}
	}

	i.end = ctx.idx - 1
	var n Number = i
	if hasPoint || hasE {
		n = Real(i)
	}

	return n, nil
}

func parseBinInt(ctx *context, src []byte) (Int, error) {
	t := Int{ctx.pos()}

	// Skip 0b
	ctx.idx += 2
	for {
		if ctx.idx >= len(src) {
			break
		}
		b := src[ctx.idx]
		if b == '0' || b == '1' {
			ctx.idx++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, Error{
				Int{ctx.pos()},
				fmt.Sprintf("invalid character '%c' in binary", b),
			}
		}
	}
	t.end = ctx.idx - 1
	return t, nil
}

func parseOctalInt(ctx *context, src []byte) (Int, error) {
	t := Int{ctx.pos()}

	// Skip 0o
	ctx.idx += 2
	for {
		if ctx.idx >= len(src) {
			break
		}
		b := src[ctx.idx]
		if '0' <= b && b <= '7' {
			ctx.idx++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, Error{
				Int{ctx.pos()},
				fmt.Sprintf("invalid character '%c' in octal", b),
			}
		}
	}
	t.end = ctx.idx - 1
	return t, nil
}

func parseHexInt(ctx *context, src []byte) (Int, error) {
	t := Int{ctx.pos()}

	// Skip 0x
	ctx.idx += 2
	for {
		if ctx.idx >= len(src) {
			break
		}
		b := src[ctx.idx]
		if isHexDigit(b) {
			ctx.idx++
		} else if isValidAfterNumber(b) {
			break
		} else {
			return t, Error{
				Int{ctx.pos()},
				fmt.Sprintf("invalid character '%c' in hex", b),
			}
		}
	}
	t.end = ctx.idx - 1
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
func parseWord(ctx *context, src []byte, toks *[]Token) (Token, error) {
	var t Token
	defer func() { ctx.idx = t.End() + 1 }()
	word, hasHyphen, hasDot := getWord(src, ctx.idx)

	qualIdentErrMsg := "symbol name in qualified identifier must start with upper case letter"
	if hasHyphen && hasDot {
		// This is for sure part of an expression
		chunks := bytes.Split(word, []byte{'-'})
		for i, chunk := range chunks {
			if bytes.Contains(chunk, []byte{'.'}) {
				t = QualIdent{
					position{
						start: ctx.idx, end: ctx.idx + len(chunk) - 1, line: ctx.line, column: ctx.col(ctx.idx),
					},
				}
				if !isValidQualifiedIdentifier(chunk) {
					return t, Error{t, qualIdentErrMsg}
				}
			} else {
				t = Ident{
					position{
						start: ctx.idx, end: ctx.idx + len(chunk) - 1, line: ctx.line, column: ctx.col(ctx.idx),
					},
				}
			}
			if i == len(chunks)-1 {
				return t, nil
			}
			*toks = append(*toks, t)
			ctx.idx += len(chunks[i])
			t = Sub{ctx.pos()}
			*toks = append(*toks, t)
			ctx.idx++
		}
	} else if hasDot {
		// It is qualified identifier
		t = QualIdent{position{start: ctx.idx, end: ctx.idx + len(word) - 1, line: ctx.line, column: ctx.col(ctx.idx)}}

		if !isValidQualifiedIdentifier(word) {
			return t, Error{t, qualIdentErrMsg}
		}

		return t, nil
	}

	splitHyphenatedWord := func() (Ident, Sub, Ident) {
		i1 := Ident{ctx.pos()}
		s := Sub{position{line: ctx.line}}
		i2 := Ident{position{end: ctx.idx + len(word) - 1, line: ctx.line}}

		for i := 0; i < len(word); i++ {
			if word[i] == '-' {
				i1.end = ctx.idx + i - 1

				s.start = ctx.idx + i
				s.end = ctx.idx + i
				s.column = ctx.col(ctx.idx + i)

				i2.start = ctx.idx + i + 1
				i2.column = ctx.col(ctx.idx + i + 1)
			}
		}
		return i1, s, i2
	}

	if !hasHyphen {
		// Firstly assume word is a keyword
		t = parseKeyword(word, ctx)
		// If it is not a keyword, then it might be a property or identifier.
		if _, ok := t.(None); ok {
			t = parseProperty(word, ctx)
			// If it is not property, then it must be an identifier.
			if _, ok := t.(None); ok {
				t = Ident{position{t.Start(), t.End(), t.Line(), t.Column()}}
			} else {
				// However, properties are properties only if they are in valid place,
				// otherwise, these are regular identifiers.
				if prevTok, ok := lastToken(*toks); ok {
					switch prevTok.(type) {
					case Newline, Semicolon, Indent:
						// Do nothing, this is property
					default:
						t = Ident{position{t.Start(), t.End(), t.Line(), t.Column()}}
					}
				}
			}
		}

		// Allow functionality keywords to be instantiation names
		if _, ok := t.(Functionality); ok {
			if prevTok, ok := lastToken(*toks); ok {
				switch prevTok.(type) {
				case Newline, Indent, Dedent:
					t = Ident{position{t.Start(), t.End(), t.Line(), t.Column()}}
				}
			}
			if len(*toks) == 0 {
				t = Ident{position{t.Start(), t.End(), t.Line(), t.Column()}}
			}
		}
	} else {
		// Firstly assume word is a property
		t = parseProperty(word, ctx)
		// If it is not property, then it is part of an expression.
		if _, ok := t.(None); ok {
			// t is last, as deferred function has to calculate new context index.
			i1, sub, i2 := splitHyphenatedWord()
			t = i2
			*toks = append(*toks, []Token{i1, sub, i2}...)
			// Assing to t for updating context index in deferred function
			t = i2
			return None{}, nil
		} else {
			// It might be property, or part of an expression.
			prevTok, ok := lastToken(*toks)
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
				*toks = append(*toks, []Token{i1, sub, i2}...)
				// Assing to t for updating context index in deferred function
				t = i2
				return None{}, nil
			}
		}
	}

	// The word might be the unit of time literal
	if _, ok := t.(Ident); ok {
		if prevTok, ok := lastToken(*toks); ok {
			if _, ok := prevTok.(Int); ok {
				switch string(word) {
				case "ns", "us", "ms", "s":
					t = Time{
						position{
							start:  prevTok.Start(),
							end:    t.End(),
							line:   prevTok.Line(),
							column: prevTok.Column(),
						},
					}
					// Remove previous Int from the list
					// New Time token will be inserted
					*toks = (*toks)[:len(*toks)-1]
				}
			}
		}
	}

	return t, nil
}

func parseKeyword(word []byte, ctx *context) Token {
	s := ctx.idx
	e := s + len(word) - 1
	l := ctx.line
	col := ctx.col(ctx.idx)

	pos := position{s, e, l, col}

	switch string(word) {
	case "false", "true":
		return Bool{pos}
	case "block":
		return Block{pos}
	case "bus":
		return Bus{pos}
	case "config":
		return Config{pos}
	case "const":
		return Const{pos}
	case "import":
		return Import{pos}
	case "irq":
		return Irq{pos}
	case "mask":
		return Mask{pos}
	case "memory":
		return Memory{pos}
	case "param":
		return Param{pos}
	case "proc":
		return Proc{pos}
	case "return":
		return Return{pos}
	case "static":
		return Static{pos}
	case "status":
		return Status{pos}
	case "stream":
		return Stream{pos}
	case "type":
		return Type{pos}
	}

	return None{pos}
}

func parseProperty(word []byte, ctx *context) Token {
	s := ctx.idx
	e := s + len(word) - 1
	l := ctx.line
	col := ctx.col(ctx.idx)

	pos := position{s, e, l, col}

	switch string(word) {
	case "access":
		return Access{pos}
	case "add-enable":
		return AddEnable{pos}
	case "atomic":
		return Atomic{pos}
	case "byte-write-enable":
		return ByteWriteEnable{pos}
	case "clear":
		return Clear{pos}
	case "delay":
		return Delay{pos}
	case "enable-init-value":
		return EnableInitValue{pos}
	case "enable-reset-value":
		return EnableResetValue{pos}
	case "groups":
		return Groups{pos}
	case "init-value":
		return InitValue{pos}
	case "in-trigger":
		return InTrigger{pos}
	case "masters":
		return Masters{pos}
	case "out-trigger":
		return OutTrigger{pos}
	case "range":
		return Range{pos}
	case "read-latency":
		return ReadLatency{pos}
	case "read-value":
		return ReadValue{pos}
	case "reset":
		return Reset{pos}
	case "reset-value":
		return ResetValue{pos}
	case "size":
		return Size{pos}
	case "width":
		return Width{pos}
	}

	return None{pos}
}
