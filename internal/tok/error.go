package tok

import (
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	Msg string
	Tok Token
}

func (e Error) Error() string {
	return fmt.Sprintf(
		"%d:%d: %s", e.Tok.Line(), e.Tok.Column(), e.Msg,
	)
}

// ErrorLoc returns error token location.
func ErrorLoc(err error, src []byte) string {
	e := err.(Error)

	b := strings.Builder{}

	lineNum := strconv.FormatInt(int64(e.Tok.Line()), 10)
	lineNumWidth := len(lineNum)
	for i := 0; i < lineNumWidth+2; i++ {
		b.WriteRune(' ')
	}
	b.WriteString("|\n")

	lineStartIdx := e.Tok.Start()
	for {
		if lineStartIdx == 0 || src[lineStartIdx-1] == '\n' {
			break
		}
		lineStartIdx--
	}

	lineEndIdx := e.Tok.End()
	if _, ok := e.Tok.(Newline); !ok {
		for {
			if lineEndIdx == len(src)-1 || src[lineEndIdx+1] == '\n' {
				break
			}
			lineEndIdx++
		}
	} else {
		lineEndIdx--
	}

	line := src[lineStartIdx : lineEndIdx+1]
	indent := 0
	for i := 0; i < len(line); i++ {
		if line[i] == '\t' {
			indent++
		} else {
			break
		}
	}

	b.WriteRune(' ')
	b.WriteString(lineNum)
	b.WriteRune(' ')
	b.WriteRune('|')
	b.WriteRune(' ')
	b.Write(line)
	b.WriteRune('\n')

	for i := 0; i < lineNumWidth+2; i++ {
		b.WriteRune(' ')
	}
	b.WriteRune('|')
	b.WriteRune(' ')

	col := 1
	if e.Tok.Column() > 1 {
		for i := 0; i < indent; i++ {
			b.WriteRune('\t')
			col++
		}
	}

	for {
		if col == e.Tok.Column() {
			break
		}
		b.WriteRune(' ')
		col++
	}

	for {
		if col == e.Tok.Column()+(e.Tok.End()-e.Tok.Start()+1) {
			break
		}
		b.WriteRune('^')
		col++
	}

	b.WriteRune('\n')

	return b.String()
}
