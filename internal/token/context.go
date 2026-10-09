package token

// Parsing context
type context struct {
	indent int // Current indent level
	idx    int // Current buffer index
	src    []byte
	path   string
}

func (ctx context) end() bool {
	return ctx.idx >= len(ctx.src)
}

// Creates position from the current context state.
func (ctx context) pos() position {
	return position{ctx.idx, ctx.idx, ctx.src, ctx.path}
}

// Returns the byte at the given offset from the current index without advancing.
// If the resulting index is outside the source, then 0 is returned.
func (ctx context) peek(offset int) byte {
	idx := ctx.idx + offset
	if idx < 0 || idx >= len(ctx.src) {
		return 0
	}
	return ctx.src[idx]
}

// Returns the current byte and advances the index.
// At the end of the source, returns 0 without advancing.
func (ctx *context) next() byte {
	if ctx.end() {
		return 0
	}
	b := ctx.peek(0)
	ctx.idx++
	return b
}

// Returns word from the source starting at the current context index.
//
// The method assumes the current byte is not a whitespace character.
// The second return is true if word contains hyphen '-' character.
// The third return is true if word contains dot '.' character.
func (ctx context) getWord() ([]byte, bool, bool) {
	hasHyphen := false
	hasDot := false
	startIdx := ctx.idx

	for !ctx.end() {
		b := ctx.peek(0)
		if !isLetter(b) && !isDigit(b) && b != '_' && b != '-' && b != '.' {
			break
		}
		switch ctx.next() {
		case '-':
			hasHyphen = true
		case '.':
			hasDot = true
		}
	}
	return ctx.src[startIdx:ctx.idx], hasHyphen, hasDot
}
