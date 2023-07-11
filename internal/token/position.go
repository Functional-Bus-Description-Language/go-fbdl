package token

type Position struct {
	Start  int // First byte index of the token
	End    int // Last byte index of the token
	Line   int // Line number, starting at 1
	Column int // Column number, starting at 1 (byte count)
}
