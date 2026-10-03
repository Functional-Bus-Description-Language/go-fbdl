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
