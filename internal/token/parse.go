package token

import (
	"bytes"
	"fmt"
)

// Parses src byte array containing the source code and returns token Stream.
func Parse(src []byte) ([]Token, error) {
	var (
		ctx    context
		err    error
		toks   []Token
		tokens []Token // Token stream
	)
	ctx.src = src
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
		case '^':
			toks, err = parseXorOperator(&ctx)
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
					Toks: []Token{None{span: ctx.span()}},
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

	tokens = append(tokens, Eof{ctx.span()})

	return tokens, nil
}

func parseSpace(ctx *context) ([]Token, error) {
	switch ctx.lastTok.(type) {
	case None, Newline:
		return parseIndent(ctx)
	}

	// Eat all spaces
	startIdx := ctx.idx
	ctx.next()
	for ctx.peek(0) == ' ' {
		ctx.next()
	}

	spaceCount := ctx.idx - startIdx

	if ctx.end() || ctx.peek(0) == '\n' {
		tok := None{span{startIdx, ctx.idx}}
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
	indent := Indent{ctx.span()}

	spaceCount := 0
	// Eat all spaces
	for ctx.peek(0) == ' ' {
		ctx.next()
		spaceCount++
	}

	if ctx.end() || ctx.peek(0) == '\n' {
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
	tab := Indent{ctx.span()}
	return nil, Error{
		"tab character '\\t' allowed only in comments, use spaces",
		[]Token{tab},
	}
}

func parseNewline(ctx *context) ([]Token, error) {
	if _, ok := ctx.lastTok.(Semicolon); ok {
		return nil, Error{"extra ';' at line end", []Token{ctx.lastTok}}
	}

	nl := Newline{ctx.span()}
	ctx.next()
	parsed := []Token{nl}

	if !ctx.end() && ctx.peek(0) != ' ' && ctx.peek(0) != '\n' && ctx.indent != 0 {
		// Insert proper number of Dedent tokens.
		t := Dedent{ctx.span()}
		for range ctx.indent {
			parsed = append(parsed, t)
		}
		ctx.indent = 0
	}

	return parsed, nil
}

func parseComment(ctx *context) ([]Token, error) {
	t := Comment{ctx.span()}

	for {
		ctx.next()
		if ctx.end() || ctx.peek(0) == '\n' {
			t.end = ctx.idx
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
		return nil, Error{"redundant ','", []Token{Comma{ctx.span()}}}
	}

	t := Comma{ctx.span()}
	ctx.next()
	return []Token{t}, nil
}

func parseSemicolon(ctx *context) ([]Token, error) {
	if _, ok := ctx.lastTok.(Semicolon); ok {
		return nil, Error{"redundant ';'", []Token{Semicolon{ctx.span()}}}
	}

	t := Semicolon{ctx.span()}
	ctx.next()
	return []Token{t}, nil
}

func parseColon(ctx *context) ([]Token, error) {
	c := Colon{ctx.span()}
	ctx.next()
	return []Token{c}, nil
}

func parseNonequalityOperator(ctx *context) ([]Token, error) {
	n := Neq{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{n}, nil
}

func parseNegationOperator(ctx *context) ([]Token, error) {
	n := Neg{ctx.span()}
	ctx.next()
	return []Token{n}, nil
}

func parseEqualityOperator(ctx *context) ([]Token, error) {
	e := Eq{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{e}, nil
}

func parseAssignmentOperator(ctx *context) ([]Token, error) {
	a := Ass{ctx.span()}
	ctx.next()
	return []Token{a}, nil
}

func parseAdditionOperator(ctx *context) ([]Token, error) {
	a := Add{ctx.span()}
	ctx.next()
	return []Token{a}, nil
}

func parseSubtractionOperator(ctx *context) ([]Token, error) {
	toks := Sub{ctx.span()}
	ctx.next()
	return []Token{toks}, nil
}

func parseRemainderOperator(ctx *context) ([]Token, error) {
	r := Rem{ctx.span()}
	ctx.next()
	return []Token{r}, nil
}

func parseExponentiationOperator(ctx *context) ([]Token, error) {
	e := Exp{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{e}, nil
}

func parseMultiplicationOperator(ctx *context) ([]Token, error) {
	m := Mul{ctx.span()}
	ctx.next()
	return []Token{m}, nil
}

func parseDivisionOperator(ctx *context) ([]Token, error) {
	d := Div{ctx.span()}
	ctx.next()
	return []Token{d}, nil
}

func parseXorOperator(ctx *context) ([]Token, error) {
	x := Xor{ctx.span()}
	ctx.next()
	return []Token{x}, nil
}

func parseLessThanEqualOperator(ctx *context) ([]Token, error) {
	le := LessEq{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{le}, nil
}

func parseLeftShiftOperator(ctx *context) ([]Token, error) {
	ls := LShift{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{ls}, nil
}

func parseLessThanOperator(ctx *context) ([]Token, error) {
	l := Less{ctx.span()}
	ctx.next()
	return []Token{l}, nil
}

func parseGreaterThanEqualOperator(ctx *context) ([]Token, error) {
	ge := GreaterEq{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{ge}, nil
}

func parseRightShiftOperator(ctx *context) ([]Token, error) {
	rs := RShift{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{rs}, nil
}

func parseGreaterThanOperator(ctx *context) ([]Token, error) {
	g := Greater{ctx.span()}
	ctx.next()
	return []Token{g}, nil
}

func parseLeftParenthesis(ctx *context) ([]Token, error) {
	lp := LParen{ctx.span()}
	ctx.next()
	return []Token{lp}, nil
}

func parseRightParenthesis(ctx *context) ([]Token, error) {
	rp := RParen{ctx.span()}
	ctx.next()
	return []Token{rp}, nil
}

func parseLeftBracket(ctx *context) ([]Token, error) {
	lb := LBracket{ctx.span()}
	ctx.next()
	return []Token{lb}, nil
}

func parseRightBracket(ctx *context) ([]Token, error) {
	rb := RBracket{ctx.span()}
	ctx.next()
	return []Token{rb}, nil
}

func parseLogicalAnd(ctx *context) ([]Token, error) {
	a := And{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{a}, nil
}

func parseBitAnd(ctx *context) ([]Token, error) {
	ba := BitAnd{ctx.span()}
	ctx.next()
	return []Token{ba}, nil
}

func parseLogicalOr(ctx *context) ([]Token, error) {
	o := Or{span{ctx.idx, ctx.idx + 2}}
	ctx.next()
	ctx.next()
	return []Token{o}, nil
}

func parseBitOr(ctx *context) ([]Token, error) {
	bo := BitOr{ctx.span()}
	ctx.next()
	return []Token{bo}, nil
}

func parseString(ctx *context) ([]Token, error) {
	t := String{ctx.span()}

	for {
		ctx.next()
		if ctx.end() || ctx.peek(0) == '\n' {
			return []Token{t}, Error{"unterminated string literal", []Token{t}}
		}
		b := ctx.peek(0)
		t.end++
		if b == '"' {
			break
		}
	}
	ctx.next()
	return []Token{t}, nil
}

func parseBinBitString(ctx *context) ([]Token, error) {
	t := BitString{span{ctx.idx, ctx.idx + 2}}

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
			'-', 'h', 'H', 'l', 'L', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
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
					[]Token{BitString{ctx.span()}},
				}
			}
		}
	}
}

func parseOctalBitString(ctx *context) ([]Token, error) {
	t := BitString{span{ctx.idx, ctx.idx + 2}}

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
			'-', 'h', 'H', 'l', 'L', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
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
					[]Token{BitString{ctx.span()}},
				}
			}
		}
	}
}

func parseHexBitString(ctx *context) ([]Token, error) {
	t := BitString{span{ctx.idx, ctx.idx + 2}}

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
			'-', 'h', 'H', 'l', 'L', 'u', 'U', 'w', 'W', 'x', 'X', 'z', 'Z':
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
				[]Token{BitString{ctx.span()}},
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

	i := Int{ctx.span()}
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
					[]Token{Float{ctx.span()}},
				}
			} else {
				if hasE {
					return nil, Error{
						"point character '.' after exponent in number",
						[]Token{Float{ctx.span()}},
					}
				}
				hasPoint = true
			}
		} else if b == 'e' || b == 'E' {
			if hasE {
				return nil, Error{
					"second exponent in number",
					[]Token{Float{ctx.span()}},
				}
			} else {
				hasE = true
			}
			ctx.next() // Skip the exponent marker.
			if ctx.peek(0) == '+' || ctx.peek(0) == '-' {
				ctx.next()
			}
			if !isDigit(ctx.peek(0)) {
				return nil, Error{
					"missing digits in number exponent",
					[]Token{Float{span{i.start, ctx.idx}}},
				}
			}
		} else if b == '\'' && !hasPoint && !hasE {
			return parseTime(ctx, i.start)
		} else if isValidAfterNumber(b) {
			break
		} else {
			return nil, Error{
				fmt.Sprintf("invalid character '%c' in number", b),
				[]Token{Int{ctx.span()}},
			}
		}
	}

	i.end = ctx.idx
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

	t := Time{span{start, ctx.idx}}
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
	t := Int{ctx.span()}

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
				[]Token{Int{ctx.span()}},
			}
		}
	}
	t.end = ctx.idx
	if t.end == t.start+2 {
		return nil, Error{"missing digit after bin prefix", []Token{t}}
	}
	return []Token{t}, nil
}

func parseOctalInt(ctx *context) ([]Token, error) {
	t := Int{ctx.span()}

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
				[]Token{Int{ctx.span()}},
			}
		}
	}
	t.end = ctx.idx
	if t.end == t.start+2 {
		return nil, Error{"missing digit after octal prefix", []Token{t}}
	}
	return []Token{t}, nil
}

func parseHexInt(ctx *context) ([]Token, error) {
	t := Int{ctx.span()}

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
				[]Token{Int{ctx.span()}},
			}
		}
	}
	t.end = ctx.idx
	if t.end == t.start+2 {
		return nil, Error{"missing digit after hex prefix", []Token{t}}
	}
	return []Token{t}, nil
}

// Checks that the package is present and the symbol starts with a letter.
func validateQualIdent(word []byte, tok QualIdent) error {
	dot := bytes.IndexByte(word, '.')
	if dot == 0 {
		return Error{"missing package in qualified identifier", []Token{tok}}
	}
	if dot < 0 || dot+1 >= len(word) || !isLetter(word[dot+1]) {
		return Error{"symbol name in qualified identifier must start with letter", []Token{tok}}
	}
	return nil
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
					span{
						start: ctx.idx - len(word),
						end:   ctx.idx,
					},
				},
			},
		}
	}

	if dotCount == 0 && hyphenCount == 0 {
		return parseNoHyphenWord(ctx, word)
	} else if dotCount == 1 && hyphenCount == 0 {
		tok := QualIdent{span{ctx.idx - len(word), ctx.idx}}
		if err := validateQualIdent(word, tok); err != nil {
			return nil, err
		}
		return []Token{tok}, nil
	} else if dotCount == 0 && hyphenCount >= 1 {
		return parseMultiHyphenWord(ctx, word)
	}

	return parseMultiSubExprWithQualIdent(ctx, word)
}

func parseNoHyphenWord(ctx *context, word []byte) ([]Token, error) {
	toks, err := parseKeyword(ctx, word)
	if err != nil || len(toks) != 0 {
		return toks, err
	}

	toks, err = parseFunc(ctx, word)
	if err != nil || len(toks) != 0 {
		return toks, err
	}

	toks, err = parseWordProp(ctx, word)
	if err != nil || len(toks) != 0 {
		return toks, err
	}

	ident := Ident{
		span{
			start: ctx.idx - len(word),
			end:   ctx.idx,
		},
	}

	return []Token{ident}, nil
}

func parseMultiHyphenWord(ctx *context, word []byte) ([]Token, error) {
	toks, err := parseHyphenProp(ctx, word)
	if err != nil || len(toks) != 0 {
		return toks, err
	}

	return parseMultiSubExprWithQualIdent(ctx, word)
}

// Parses expression with multiple subtraction and potentially qualified identifiers.
func parseMultiSubExprWithQualIdent(ctx *context, word []byte) ([]Token, error) {
	chunks := bytes.Split(word, []byte{'-'})
	var toks []Token
	start := ctx.idx - len(word)

	for i, chunk := range chunks {
		spn := span{start, start + len(chunk)}
		if len(chunk) != 0 {
			switch bytes.Count(chunk, []byte{'.'}) {
			case 0:
				chunkCtx := *ctx
				chunkCtx.idx = spn.end
				parsed, err := parseNoHyphenWord(&chunkCtx, chunk)
				if err != nil {
					return nil, err
				}
				toks = append(toks, parsed...)
			case 1:
				tok := QualIdent{spn}
				if err := validateQualIdent(chunk, tok); err != nil {
					return nil, err
				}
				toks = append(toks, tok)
			default:
				return nil, Error{"multiple dots in identifier", []Token{Ident{spn}}}
			}
		}

		start += len(chunk)
		if i < len(chunks)-1 {
			toks = append(toks, Sub{span{start, start + 1}})
			start++
		}
	}

	return toks, nil
}

// Hard keywords that must not be used anywhere.
func parseKeyword(ctx *context, word []byte) ([]Token, error) {
	s := span{
		ctx.idx - len(word),
		ctx.idx,
	}

	switch string(word) {
	case "false", "true":
		return []Token{Bool{s}}, nil
	case "const":
		return []Token{Const{s}}, nil
	case "import":
		return []Token{Import{s}}, nil
	case "type":
		return []Token{Type{s}}, nil
	}

	return nil, nil
}

func parseFunc(ctx *context, word []byte) ([]Token, error) {
	s := span{
		ctx.idx - len(word),
		ctx.idx,
	}

	switch string(word) {
	case "blackbox":
		return []Token{Blackbox{s}}, nil
	case "block":
		return []Token{Block{s}}, nil
	case "bus":
		return []Token{Bus{s}}, nil
	case "config":
		return []Token{Config{s}}, nil
	case "group":
		return []Token{Group{s}}, nil
	case "irq":
		return []Token{Irq{s}}, nil
	case "mask":
		return []Token{Mask{s}}, nil
	case "param":
		return []Token{Param{s}}, nil
	case "proc":
		return []Token{Proc{s}}, nil
	case "return":
		return []Token{Return{s}}, nil
	case "static":
		return []Token{Static{s}}, nil
	case "status":
		return []Token{Status{s}}, nil
	case "stream":
		return []Token{Stream{s}}, nil
	}

	return nil, nil
}

func parseWordProp(ctx *context, word []byte) ([]Token, error) {
	s := span{
		ctx.idx - len(word),
		ctx.idx,
	}

	switch string(word) {
	case "access":
		return []Token{Access{s}}, nil
	case "align":
		return []Token{Align{s}}, nil
	case "atomic":
		return []Token{Atomic{s}}, nil
	case "clear":
		return []Token{Clear{s}}, nil
	case "delay":
		return []Token{Delay{s}}, nil
	case "masters":
		return []Token{Masters{s}}, nil
	case "range":
		return []Token{Range{s}}, nil
	case "reset":
		return []Token{Reset{s}}, nil
	case "size":
		return []Token{Size{s}}, nil
	case "virtual":
		return []Token{Virtual{s}}, nil
	case "width":
		return []Token{Width{s}}, nil
	}

	return nil, nil
}

func parseHyphenProp(ctx *context, word []byte) ([]Token, error) {
	s := span{
		ctx.idx - len(word),
		ctx.idx,
	}

	switch string(word) {
	case "add-enable":
		return []Token{AddEnable{s}}, nil
	case "byte-write-enable":
		return []Token{ByteWriteEnable{s}}, nil
	case "enable-init-value":
		return []Token{EnableInitValue{s}}, nil
	case "enable-reset-value":
		return []Token{EnableResetValue{s}}, nil
	case "init-value":
		return []Token{InitValue{s}}, nil
	case "in-trigger":
		return []Token{InTrigger{s}}, nil
	case "out-trigger":
		return []Token{OutTrigger{s}}, nil
	case "read-latency":
		return []Token{ReadLatency{s}}, nil
	case "read-value":
		return []Token{ReadValue{s}}, nil
	case "reset-value":
		return []Token{ResetValue{s}}, nil
	}

	return nil, nil
}
