package tok

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func isHexDigit(b byte) bool {
	return ('0' <= b && b <= '9') ||
		('a' <= b && b <= 'f') ||
		('A' <= b && b <= 'F')
}

// Returns true if character is a valid character after number literal.
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
