package token

type Position struct {
	Start  uint // First byte index of the token
	End    uint // Last byte index of the token
	Line   uint // Line number, starting at 1
	Column uint // Column number, starting at 1 (byte count)
}
