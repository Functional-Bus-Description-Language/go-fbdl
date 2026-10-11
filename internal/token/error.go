package token

// Error describes an error at one or more token spans.
type Error struct {
	Msg  string
	Toks []Token
}

func (err Error) Error() string { return err.Msg }
