package token

type TokenStream []Token

// LastToken returns last token from the stream.
// If stream is empty, the second return is false.
func (ts TokenStream) LastToken() (Token, bool) {
	if len(ts) == 0 {
		return nil, false
	}
	return ts[len(ts)-1], true
}
