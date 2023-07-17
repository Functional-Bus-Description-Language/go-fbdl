package token

import "fmt"

type Token struct {
	Kind Kind

	// Token position
	Start  int // First byte index of the token
	End    int // Last byte index of the token
	Line   int // Line number, starting at 1
	Column int // Column number, starting at 1 (byte count)
}

// Loc return location of the token within the file in "line:column" format.
func (t Token) Loc() string {
	return fmt.Sprintf("%d:%d", t.Line, t.Column)
}
