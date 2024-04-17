package tok

// Parsing context
type context struct {
	line   int // Current line number
	indent int // Current indent level
	idx    int // Current buffer index
	nlIdx  int // Last newline index
	src    []byte
}

func (ctx context) end() bool {
	return ctx.idx >= len(ctx.src)
}

// Returns column number for given index.
func (ctx context) col(idx int) int {
	return idx - ctx.nlIdx
}

// Creates position from the current context state.
func (ctx context) pos() position {
	return position{ctx.idx, ctx.idx, ctx.line, ctx.col(ctx.idx)}
}

// Returns byte with index equal idx.
// If idx >= len(src), then 0 is returned.
func (ctx context) byte() byte {
	if ctx.idx >= len(ctx.src) {
		return 0
	}
	return ctx.src[ctx.idx]
}

// Returns byte with index equal idx + 1.
// If (idx + 1) >= len(src), then 0 is returned.
func (ctx context) nextByte() byte {
	if ctx.idx+1 >= len(ctx.src) {
		return 0
	}
	return ctx.src[ctx.idx+1]
}
