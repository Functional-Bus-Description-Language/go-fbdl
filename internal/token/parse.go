package token

import (
	"bytes"
	"fmt"
)

// Parses src byte array containing the source code and returns token Stream.
func Parse(src []byte, path string) ([]Token, error) {
	var (
		ctx    context
		err    error
		toks   []Token
		tokens []Token // Token stream
	)
	ctx.src = src
	ctx.path = path
	ctx.lastTok = None{}

	for !ctx.end() {
		b := ctx.peek(0)
		nb := ctx.peek(1)

		switch b {
		case ' ':
			toks, err = parseSpace(&ctx)
		case '\t':
			toks, err = parseTab(&ctx)
		case '\n':
			toks, err = parseNewline(&ctx)
		case '#':
			toks, err = parseComment(&ctx)
		case ',':
			toks, err = parseComma(&ctx)
		case ';':
			toks, err = parseSemicolon(&ctx)
		case ':':
			toks, err = parseColon(&ctx)
		case '!':
			if nb == '=' {
				toks, err = parseNonequalityOperator(&ctx)
			} else {
				toks, err = parseNegationOperator(&ctx)
			}
		case '=':
			if nb == '=' {
				toks, err = parseEqualityOperator(&ctx)
			} else {
				toks, err = parseAssignmentOperator(&ctx)
			}
		case '+':
			toks, err = parseAdditionOperator(&ctx)
		case '-':
			toks, err = parseSubtractionOperator(&ctx)
		case '%':
			toks, err = parseRemainderOperator(&ctx)
		case '*':
			if nb == '*' {
				toks, err = parseExponentiationOperator(&ctx)
			} else {
				toks, err = parseMultiplicationOperator(&ctx)
			}
		case '/':
			toks, err = parseDivisionOperator(&ctx)
		case '<':
			switch nb {
			case '=':
				toks, err = parseLessThanEqualOperator(&ctx)
			case '<':
				toks, err = parseLeftShiftOperator(&ctx)
			default:
				toks, err = parseLessThanOperator(&ctx)
			}
		case '>':
			switch nb {
			case '=':
				toks, err = parseGreaterThanEqualOperator(&ctx)
			case '>':
				toks, err = parseRightShiftOperator(&ctx)
			default:
				toks, err = parseGreaterThanOperator(&ctx)
			}
		case '(':
			toks, err = parseLeftParenthesis(&ctx)
		case ')':
			toks, err = parseRightParenthesis(&ctx)
		case '[':
			toks, err = parseLeftBracket(&ctx)
		case ']':
			toks, err = parseRightBracket(&ctx)
		case '&':
			if nb == '&' {
				toks, err = parseLogicalAnd(&ctx)
			} else {
				toks, err = parseBitAnd(&ctx)
			}
		case '|':
			if nb == '|' {
				toks, err = parseLogicalOr(&ctx)
			} else {
				toks, err = parseBitOr(&ctx)
			}
		case '"':
			toks, err = parseString(&ctx)
		case 'b', 'B':
			if nb == '"' {
				toks, err = parseBinBitString(&ctx)
			} else {
				toks, err = parseWord(&ctx)
			}
		case 'o', 'O':
			if nb == '"' {
				toks, err = parseOctalBitString(&ctx)
			} else {
				toks, err = parseWord(&ctx)
			}
		case 'x', 'X':
			if nb == '"' {
				toks, err = parseHexBitString(&ctx)
			} else {
				toks, err = parseWord(&ctx)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			toks, err = parseNumber(&ctx)
		default:
			if isLetter(b) || b == '_' {
				toks, err = parseWord(&ctx)
			} else {
				return nil, Error{
					Msg:  fmt.Sprintf("invalid byte 0x%x ('%c')", b, b),
					Toks: []Token{None{position: ctx.pos()}},
				}
			}
		}

		if err != nil {
			return tokens, err
		}

		tokens = append(tokens, toks...)
		if len(toks) > 0 {
			ctx.lastTok = toks[len(toks)-1]
		}
	}

	tokens = append(tokens, Eof{ctx.pos()})

	return tokens, nil
}

func parseSpace(ctx *context) ([]Token, error) {
	if _, ok := ctx.lastTok.(Newline); ok {
		return parseIndent(ctx)
	}

	// Eat all spaces
	startIdx := ctx.idx
	ctx.next()
	for ctx.peek(0) == ' ' {
		ctx.next()
	}

	spaceCount := ctx.idx - startIdx

	if ctx.peek(0) == '\n' {
		tok := None{ctx.pos()}
		tok.start = startIdx
		tok.end--
		if spaceCount > 1 {
			return nil, Error{
				fmt.Sprintf("extra %d spaces at line end", spaceCount),
				[]Token{tok},
			}
		} else {
			return nil, Error{
				"extra space at line end",
				[]Token{tok},
			}
		}
	}

	return nil, nil
}

func parseIndent(ctx *context) ([]Token, error) {
	indent := Indent{ctx.pos()}

	spaceCount := 0
	// Eat all spaces
	for ctx.peek(0) == ' ' {
		ctx.next()
		spaceCount++
	}

	if ctx.peek(0) == '\n' {
		if spaceCount > 1 {
			indent.end += spaceCount - 1
			return nil, Error{
				fmt.Sprintf("extra %d spaces at line end", spaceCount),
				[]Token{indent},
			}
		} else {
			return nil, Error{
				"extra space at line end",
				[]Token{indent},
			}
		}
	}

	indent.end += spaceCount - 1

	if spaceCount%2 != 0 {
		return []Token{indent}, Error{
			fmt.Sprintf("odd number (%d) of spaces in indent, expected even number", spaceCount),
			[]Token{indent},
		}
	}

	indentLvl := spaceCount / 2
	if indentLvl == ctx.indent+1 {
		ctx.indent = indentLvl
		return []Token{indent}, nil
	} else if indentLvl > ctx.indent+1 {
		return []Token{indent}, Error{
			fmt.Sprintf("multi indent increase, previous indent %d , current indent %d", ctx.indent, indentLvl),
			[]Token{indent},
		}
	} else if indentLvl < ctx.indent {
		// Insert proper number of INDENT_DEC tokens.
		var parsed []Token
		t := Dedent(indent)
		for i := 0; indentLvl+i < ctx.indent; i++ {
			parsed = append(parsed, t)
		}
		ctx.indent = indentLvl
		return parsed, nil
	}

	return nil, nil
}

func parseTab(ctx *context) ([]Token, error) {
	tab := Indent{ctx.pos()}
	return nil, Error{
		"tab character '\\t' allowed only in comments, use spaces",
		[]Token{tab},
	}
}

func parseNewline(ctx *context) ([]Token, error) {
	if _, ok := ctx.lastTok.(Semicolon); ok {
		return nil, Error{"extra ';' at line end", []Token{ctx.lastTok}}
	}

	nl := Newline{ctx.pos()}
	ctx.next()
	parsed := []Token{nl}

	if !ctx.end() && ctx.peek(0) != ' ' && ctx.peek(0) != '\n' && ctx.indent != 0 {
		// Insert proper number of Dedent tokens.
		t := Dedent{ctx.pos()}
		for range ctx.indent {
			parsed = append(parsed, t)
		}
		ctx.indent = 0
	}

	return parsed, nil
}

func parseComment(ctx *context) ([]Token, error) {
	t := Comment{ctx.pos()}

	for {
		ctx.next()
		if ctx.end() || ctx.peek(0) == '\n' {
			t.end = ctx.idx - 1
			break
		}
	}

	// Add comment to the token stream only if it is a potential documentation comment.
	switch ctx.lastTok.(type) {
	case None, Newline, Indent, Dedent:
		return []Token{t}, nil
	}

	return nil, nil
}

func parseComma(ctx *context) ([]Token, error) {
	if _, ok := ctx.lastTok.(Comma); ok {
		return nil, Error{"redundant ','", []Token{Comma{ctx.pos()}}}
	}

	t := Comma{ctx.pos()}
	ctx.next()
	return []Token{t}, nil
}

func parseSemicolon(ctx *context) ([]Token, error) {
	if _, ok := ctx.lastTok.(Semicolon); ok {
		return nil, Error{"redundant ';'", []Token{Semicolon{ctx.pos()}}}
	}

	t := Semicolon{ctx.pos()}
	ctx.next()
	return []Token{t}, nil
}

func parseColon(ctx *context) ([]Token, error) {
	c := Colon{ctx.pos()}
	ctx.next()
	return []Token{c}, nil
}

func parseNonequalityOperator(ctx *context) ([]Token, error) {
	n := Neq{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{n}, nil
}

func parseNegationOperator(ctx *context) ([]Token, error) {
	n := Neg{ctx.pos()}
	ctx.next()
	return []Token{n}, nil
}

func parseEqualityOperator(ctx *context) ([]Token, error) {
	e := Eq{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{e}, nil
}

func parseAssignmentOperator(ctx *context) ([]Token, error) {
	a := Ass{ctx.pos()}
	ctx.next()
	return []Token{a}, nil
}

func parseAdditionOperator(ctx *context) ([]Token, error) {
	a := Add{ctx.pos()}
	ctx.next()
	return []Token{a}, nil
}

func parseSubtractionOperator(ctx *context) ([]Token, error) {
	toks := Sub{ctx.pos()}
	ctx.next()
	return []Token{toks}, nil
}

func parseRemainderOperator(ctx *context) ([]Token, error) {
	r := Rem{ctx.pos()}
	ctx.next()
	return []Token{r}, nil
}

func parseExponentiationOperator(ctx *context) ([]Token, error) {
	e := Exp{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{e}, nil
}

func parseMultiplicationOperator(ctx *context) ([]Token, error) {
	m := Mul{ctx.pos()}
	ctx.next()
	return []Token{m}, nil
}

func parseDivisionOperator(ctx *context) ([]Token, error) {
	d := Div{ctx.pos()}
	ctx.next()
	return []Token{d}, nil
}

func parseLessThanEqualOperator(ctx *context) ([]Token, error) {
	le := LessEq{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{le}, nil
}

func parseLeftShiftOperator(ctx *context) ([]Token, error) {
	ls := LShift{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{ls}, nil
}

func parseLessThanOperator(ctx *context) ([]Token, error) {
	l := Less{ctx.pos()}
	ctx.next()
	return []Token{l}, nil
}

func parseGreaterThanEqualOperator(ctx *context) ([]Token, error) {
	ge := GreaterEq{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{ge}, nil
}

func parseRightShiftOperator(ctx *context) ([]Token, error) {
	rs := RShift{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{rs}, nil
}

func parseGreaterThanOperator(ctx *context) ([]Token, error) {
	g := Greater{ctx.pos()}
	ctx.next()
	return []Token{g}, nil
}

func parseLeftParenthesis(ctx *context) ([]Token, error) {
	lp := LParen{ctx.pos()}
	ctx.next()
	return []Token{lp}, nil
}

func parseRightParenthesis(ctx *context) ([]Token, error) {
	rp := RParen{ctx.pos()}
	ctx.next()
	return []Token{rp}, nil
}

func parseLeftBracket(ctx *context) ([]Token, error) {
	lb := LBracket{ctx.pos()}
	ctx.next()
	return []Token{lb}, nil
}

func parseRightBracket(ctx *context) ([]Token, error) {
	rb := RBracket{ctx.pos()}
	ctx.next()
	return []Token{rb}, nil
}

func parseLogicalAnd(ctx *context) ([]Token, error) {
	a := And{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{a}, nil
}

func parseBitAnd(ctx *context) ([]Token, error) {
	ba := BitAnd{ctx.pos()}
	ctx.next()
	return []Token{ba}, nil
}

func parseLogicalOr(ctx *context) ([]Token, error) {
	o := Or{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}
	ctx.next()
	ctx.next()
	return []Token{o}, nil
}

func parseBitOr(ctx *context) ([]Token, error) {
	bo := BitOr{ctx.pos()}
	ctx.next()
	return []Token{bo}, nil
}

func parseString(ctx *context) ([]Token, error) {
	t := String{ctx.pos()}

	for {
		ctx.next()
		if ctx.end() {
			return []Token{t}, Error{"unterminated string, probably missing '\"'", []Token{t}}
		}
		b := ctx.peek(0)
		if b != '\n' {
			t.end++
		}
		if b == '"' {
			break
		}
	}
	ctx.next()
	return []Token{t}, nil
}

func parseBinBitString(ctx *context) ([]Token, error) {
	t := BitString{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}

	// Skip b"
	ctx.next()
	ctx.next()
	for {
		if ctx.end() {
			return nil, Error{
				"unterminated binary bit string, probably missing '\"'",
				[]Token{t},
			}
		}

		switch b := ctx.peek(0); b {
		case '"':
			t.end++
			ctx.next()
			return []Token{t}, nil
		case '0', '1',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			ctx.next()
		default:
			switch b {
			case ' ', '\n', ';', ',':
				return nil, Error{
					"unterminated binary bit string, probably missing '\"'",
					[]Token{t},
				}
			default:
				return nil, Error{
					fmt.Sprintf("invalid character '%c' in binary bit string", b),
					[]Token{BitString{ctx.pos()}},
				}
			}
		}
	}
}

func parseOctalBitString(ctx *context) ([]Token, error) {
	t := BitString{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}

	// Skip o"
	ctx.next()
	ctx.next()
	for {
		if ctx.end() {
			return nil, Error{
				"unterminated octal bit string, probably missing '\"'",
				[]Token{t},
			}
		}

		switch b := ctx.peek(0); b {
		case '"':
			t.end++
			ctx.next()
			return []Token{t}, nil
		case '0', '1', '2', '3', '4', '5', '6', '7',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			ctx.next()
		default:
			switch b {
			case ' ', '\n', ';', ',':
				return nil, Error{
					"unterminated octal bit string, probably missing '\"'",
					[]Token{t},
				}
			default:
				return nil, Error{
					fmt.Sprintf("invalid character '%c' in octal bit string", b),
					[]Token{BitString{ctx.pos()}},
				}
			}
		}
	}
}

func parseHexBitString(ctx *context) ([]Token, error) {
	t := BitString{position{ctx.idx, ctx.idx + 1, ctx.src, ctx.path}}

	// Skip x"
	ctx.next()
	ctx.next()
	for {
		if ctx.end() {
			return nil, Error{
				"unterminated hex bit string, probably missing '\"'",
				[]Token{t},
			}
		}

		switch b := ctx.peek(0); b {
		case '"':
			t.end++
			ctx.next()
			return []Token{t}, nil
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
			'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F',
			'-', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
			t.end++
			ctx.next()
		case ' ', '\n', ';', ',':
			return nil, Error{
				"unterminated hex bit string, probably missing '\"'",
				[]Token{t},
			}
		default:
			return nil, Error{
				fmt.Sprintf("invalid character '%c' in hex bit string", b),
				[]Token{BitString{ctx.pos()}},
			}
		}
	}
}

func parseNumber(ctx *context) ([]Token, error) {
	b := ctx.peek(0)
	nb := ctx.peek(1)

	if b == '0' && (nb == 'b' || nb == 'B') {
		return parseBinInt(ctx)
	} else if b == '0' && (nb == 'o' || nb == 'O') {
		return parseOctalInt(ctx)
	} else if b == '0' && (nb == 'x' || nb == 'X') {
		return parseHexInt(ctx)
	}

	i := Int{ctx.pos()}
	hasPoint := false
	hasE := false

	for {
		ctx.next()
		if ctx.end() {
			break
		}

		b := ctx.peek(0)
		if isDigit(b) {
			continue
		}

		if b == '.' {
			if hasPoint {
				return nil, Error{
					"second point character '.' in number",
					[]Token{Float{ctx.pos()}},
				}
			} else {
				if hasE {
					return nil, Error{
						"point character '.' after exponent in number",
						[]Token{Float{ctx.pos()}},
					}
				}
				hasPoint = true
			}
		} else if b == 'e' || b == 'E' {
			if hasE {
				return nil, Error{
					"second exponent in number",
					[]Token{Float{ctx.pos()}},
				}
			} else {
				hasE = true
			}
		} else if b == '\'' && !hasPoint && !hasE {
			return parseTime(ctx, i.start)
		} else if isValidAfterNumber(b) {
			break
		} else {
			return nil, Error{
				fmt.Sprintf("invalid character '%c' in number", b),
				[]Token{Int{ctx.pos()}},
			}
		}
	}

	i.end = ctx.idx - 1
	var n Number = i
	if hasPoint || hasE {
		n = Float(i)
	}

	return []Token{n}, nil
}

func parseTime(ctx *context, start int) ([]Token, error) {
	ctx.next() // Skip the apostrophe.
	unitStart := ctx.idx
	for !ctx.end() && !isValidAfterNumber(ctx.peek(0)) {
		ctx.next()
	}

	t := Time{position{start, ctx.idx - 1, ctx.src, ctx.path}}
	switch string(ctx.src[unitStart:ctx.idx]) {
	case "ns", "us", "ms", "s":
		return []Token{t}, nil
	default:
		return nil, Error{
			"invalid unit in time literal, valid units for time literal are: 'ns', 'us', 'ms', 's'",
			[]Token{t},
		}
	}
}

func parseBinInt(ctx *context) ([]Token, error) {
	t := Int{ctx.pos()}

	// Skip 0b
	ctx.next()
	ctx.next()
	for !ctx.end() {
		b := ctx.peek(0)
		if b == '0' || b == '1' {
			ctx.next()
		} else if isValidAfterNumber(b) {
			break
		} else {
			return []Token{t}, Error{
				fmt.Sprintf("invalid character '%c' in binary", b),
				[]Token{Int{ctx.pos()}},
			}
		}
	}
	t.end = ctx.idx - 1
	return []Token{t}, nil
}

func parseOctalInt(ctx *context) ([]Token, error) {
	t := Int{ctx.pos()}

	// Skip 0o
	ctx.next()
	ctx.next()
	for !ctx.end() {
		b := ctx.peek(0)
		if '0' <= b && b <= '7' {
			ctx.next()
		} else if isValidAfterNumber(b) {
			break
		} else {
			return []Token{t}, Error{
				fmt.Sprintf("invalid character '%c' in octal", b),
				[]Token{Int{ctx.pos()}},
			}
		}
	}
	t.end = ctx.idx - 1
	return []Token{t}, nil
}

func parseHexInt(ctx *context) ([]Token, error) {
	t := Int{ctx.pos()}

	// Skip 0x
	ctx.next()
	ctx.next()
	for !ctx.end() {
		b := ctx.peek(0)
		if isHexDigit(b) {
			ctx.next()
		} else if isValidAfterNumber(b) {
			break
		} else {
			return []Token{t}, Error{
				fmt.Sprintf("invalid character '%c' in hex", b),
				[]Token{Int{ctx.pos()}},
			}
		}
	}
	t.end = ctx.idx - 1
	return []Token{t}, nil
}

// Checks that the symbol following the dot starts with a letter.
func isValidQualIdent(word []byte) bool {
	dot := bytes.IndexByte(word, '.')
	return dot >= 0 && dot+1 < len(word) && isLetter(word[dot+1])
}

func parseWord(ctx *context) ([]Token, error) {
	word := ctx.getWord()
	dotCount := bytes.Count(word, []byte{'.'})
	hyphenCount := bytes.Count(word, []byte{'-'})

	if dotCount > 1 && hyphenCount == 0 {
		return nil, Error{
			Msg: "multiple dots in identifier",
			Toks: []Token{
				Ident{
					position{
						start: ctx.idx - len(word),
						end:   ctx.idx - 1,
						src:   ctx.src,
						path:  ctx.path,
					},
				},
			},
		}
	}

	if dotCount == 0 && hyphenCount == 0 {
		return parseNoHyphenKeywordOrIdent(ctx, word)
	} else if dotCount == 1 && hyphenCount == 0 {
		tok := QualIdent{ctx.pos()}
		tok.start = tok.end - len(word)
		tok.end--
		if !isValidQualIdent(word) {
			return nil, Error{
				"symbol name in qualified identifier must start with letter",
				[]Token{tok},
			}
		}
		return []Token{tok}, nil
	} else if dotCount == 0 && hyphenCount == 1 {
		return parseHyphenPropOrOneSubExpr(ctx, word)
	} else if dotCount == 0 && hyphenCount > 1 {
		return parseMultiSubExpr(ctx, word)
	}

	return parseMultiSubExprWithQualIdent(ctx, word)
}

func parseNoHyphenKeywordOrIdent(ctx *context, word []byte) ([]Token, error) {
	toks, err := parseKeyword(ctx, word)
	if err != nil || len(toks) != 0 {
		return toks, err
	}

	switch ctx.lastTok.(type) {
	case None, Newline, Semicolon, Indent, Dedent:
		toks, err := parseNoHyphenProperty(ctx, word)
		if err != nil || len(toks) != 0 {
			return toks, err
		}
	}

	ident := Ident{
		position{
			start: ctx.idx - len(word),
			end:   ctx.idx - 1,
			src:   ctx.src,
			path:  ctx.path,
		},
	}

	switch ctx.lastTok.(type) {
	case Newline, Indent, Dedent:
		return []Token{ident}, nil
	}

	toks, err = parseFunc(ctx, word)
	if err != nil || len(toks) != 0 {
		return toks, err
	}

	return []Token{ident}, nil
}

func parseHyphenPropOrOneSubExpr(ctx *context, word []byte) ([]Token, error) {
	switch ctx.lastTok.(type) {
	case Newline, Semicolon, Indent, Dedent:
		toks, err := parseHyphenProperty(ctx, word)
		if err != nil || len(toks) != 0 {
			return toks, err
		}
	}

	hyphen := bytes.IndexByte(word, '-')
	start := ctx.idx - len(word)
	pos := position{start, start + len(word) - 1, ctx.src, ctx.path}

	left := Ident{pos}
	left.end = start + hyphen - 1

	sub := Sub{pos}
	sub.start = start + hyphen
	sub.end = sub.start

	right := Ident{pos}
	right.start = sub.end + 1

	return []Token{left, sub, right}, nil
}

func parseMultiSubExpr(ctx *context, word []byte) ([]Token, error) {
	chunks := bytes.Split(word, []byte{'-'})
	var toks []Token
	start := ctx.idx - len(word)

	for i, chunk := range chunks {
		if len(chunk) != 0 {
			toks = append(toks, Ident{position{start, start + len(chunk) - 1, ctx.src, ctx.path}})
		}

		start += len(chunk)
		if i < len(chunks)-1 {
			toks = append(toks, Sub{position{start, start, ctx.src, ctx.path}})
			start++
		}
	}

	return toks, nil
}

// Parses expression with multiple subtraction and potentially qualified identifiers.
func parseMultiSubExprWithQualIdent(ctx *context, word []byte) ([]Token, error) {
	chunks := bytes.Split(word, []byte{'-'})
	var toks []Token
	start := ctx.idx - len(word)

	for i, chunk := range chunks {
		pos := position{start, start + len(chunk) - 1, ctx.src, ctx.path}
		if len(chunk) != 0 {
			switch bytes.Count(chunk, []byte{'.'}) {
			case 0:
				toks = append(toks, Ident{pos})
			case 1:
				tok := QualIdent{pos}
				if !isValidQualIdent(chunk) {
					return nil, Error{
						"symbol name in qualified identifier must start with letter",
						[]Token{tok},
					}
				}
				toks = append(toks, tok)
			default:
				return nil, Error{"multiple dots in identifier", []Token{Ident{pos}}}
			}
		}

		start += len(chunk)
		if i < len(chunks)-1 {
			toks = append(toks, Sub{position{start, start, ctx.src, ctx.path}})
			start++
		}
	}

	return toks, nil
}

// Hard keywords that must not be used anywhere.
func parseKeyword(ctx *context, word []byte) ([]Token, error) {
	s := ctx.idx - len(word)
	e := ctx.idx - 1

	pos := position{s, e, ctx.src, ctx.path}

	switch string(word) {
	case "false", "true":
		return []Token{Bool{pos}}, nil
	case "const":
		return []Token{Const{pos}}, nil
	case "import":
		return []Token{Import{pos}}, nil
	case "type":
		return []Token{Type{pos}}, nil
	}

	return nil, nil
}

func parseFunc(ctx *context, word []byte) ([]Token, error) {
	s := ctx.idx - len(word)
	e := ctx.idx - 1

	pos := position{s, e, ctx.src, ctx.path}

	switch string(word) {
	case "block":
		return []Token{Block{pos}}, nil
	case "bus":
		return []Token{Bus{pos}}, nil
	case "config":
		return []Token{Config{pos}}, nil
	case "group":
		return []Token{Group{pos}}, nil
	case "irq":
		return []Token{Irq{pos}}, nil
	case "mask":
		return []Token{Mask{pos}}, nil
	case "memory":
		return []Token{Memory{pos}}, nil
	case "param":
		return []Token{Param{pos}}, nil
	case "proc":
		return []Token{Proc{pos}}, nil
	case "return":
		return []Token{Return{pos}}, nil
	case "static":
		return []Token{Static{pos}}, nil
	case "status":
		return []Token{Status{pos}}, nil
	case "stream":
		return []Token{Stream{pos}}, nil
	}

	return nil, nil
}

func parseNoHyphenProperty(ctx *context, word []byte) ([]Token, error) {
	s := ctx.idx - len(word)
	e := ctx.idx - 1

	pos := position{s, e, ctx.src, ctx.path}

	switch string(word) {
	case "access":
		return []Token{Access{pos}}, nil
	case "align":
		return []Token{Align{pos}}, nil
	case "atomic":
		return []Token{Atomic{pos}}, nil
	case "clear":
		return []Token{Clear{pos}}, nil
	case "delay":
		return []Token{Delay{pos}}, nil
	case "masters":
		return []Token{Masters{pos}}, nil
	case "range":
		return []Token{Range{pos}}, nil
	case "reset":
		return []Token{Reset{pos}}, nil
	case "size":
		return []Token{Size{pos}}, nil
	case "virtual":
		return []Token{Virtual{pos}}, nil
	case "width":
		return []Token{Width{pos}}, nil
	}

	return nil, nil
}

func parseHyphenProperty(ctx *context, word []byte) ([]Token, error) {
	s := ctx.idx - len(word)
	e := ctx.idx - 1

	pos := position{s, e, ctx.src, ctx.path}

	switch string(word) {
	case "add-enable":
		return []Token{AddEnable{pos}}, nil
	case "byte-write-enable":
		return []Token{ByteWriteEnable{pos}}, nil
	case "enable-init-value":
		return []Token{EnableInitValue{pos}}, nil
	case "enable-reset-value":
		return []Token{EnableResetValue{pos}}, nil
	case "init-value":
		return []Token{InitValue{pos}}, nil
	case "in-trigger":
		return []Token{InTrigger{pos}}, nil
	case "out-trigger":
		return []Token{OutTrigger{pos}}, nil
	case "read-latency":
		return []Token{ReadLatency{pos}}, nil
	case "read-value":
		return []Token{ReadValue{pos}}, nil
	case "reset-value":
		return []Token{ResetValue{pos}}, nil
	}

	return nil, nil
}
