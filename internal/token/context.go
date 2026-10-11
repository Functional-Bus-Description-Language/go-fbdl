package token

// Parsing context
type context struct {
	lastTok Token
	indent  int // Current indent level
	idx     int // Current buffer index
	src     []byte
}

func (ctx context) end() bool {
	return ctx.idx >= len(ctx.src)
}

// Creates span from the current context state.
func (ctx context) span() span {
	return span{ctx.idx, ctx.idx}
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
func (ctx *context) getWord() []byte {
	startIdx := ctx.idx

	endIdx := len(ctx.src)
	for idx := startIdx; idx < len(ctx.src); idx++ {
		b := ctx.peek(0)
		if !isLetter(b) && !isDigit(b) && b != '_' && b != '-' && b != '.' {
			endIdx = idx
			break
		}
		ctx.next()
	}
	return ctx.src[startIdx:endIdx]
}
