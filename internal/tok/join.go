package tok

// Join joins two tokens and returns None token with position spanning both of them.
// It is useful for reporting errors spanning multiple tokens.
//
// It panics if:
//   - tokens are from different source,
//   - tokens are not in the same line,
//   - tok1 is after tok2.
func Join(tok1, tok2 Token) Token {
	if tok1.Line() != tok2.Line() {
		panic("cannot join tokens from different files")
	}
	if tok1.Line() != tok2.Line() {
		panic("cannot join tokens placed in different lines")
	}
	if tok1.Column() > tok2.Column() {
		panic("cannot join tokens, tok1 starts after tok2")
	}

	return None{
		position{
			start:  tok1.Start(),
			end:    tok2.End(),
			line:   tok1.Line(),
			column: tok1.Column(),
			src:    tok1.Src(),
			path:   tok1.Path(),
		},
	}
}
