package token

type Stream []Token

// LastToken returns last token from the stream.
// If stream is empty, the second return is false.
func (s Stream) LastToken() (Token, bool) {
	if len(s) == 0 {
		return Token{}, false
	}
	return s[len(s)-1], true
}
