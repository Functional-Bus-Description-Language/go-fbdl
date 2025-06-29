package tok

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"strconv"
	"strings"
)

type Error struct {
	Msg  string
	Toks []Token
}

func (err Error) getColor() (string, string) {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return "\033[1;31m", "\033[0m"
	}
	return "", ""
}

func (err Error) Error() string {
	colorPrefix, colorSuffix := err.getColor()

	msg := fmt.Sprintf(
		"%serror%s: %s\n",
		colorPrefix, colorSuffix, err.Msg,
	)

	for _, tok := range err.Toks {
		msg += err.code(tok)
	}

	return msg
}

// Returns error token code.
func (err Error) code(tok Token) string {
	b := strings.Builder{}

	b.WriteString(
		fmt.Sprintf(
			"%s +%d:%d\n",
			tok.Path(), tok.Line(), tok.Column(),
		),
	)

	lineNum := strconv.FormatInt(int64(tok.Line()), 10)
	lineNumWidth := len(lineNum)
	for range lineNumWidth + 2 {
		b.WriteRune(' ')
	}
	b.WriteString("|\n")

	src := tok.Src()
	lineStartIdx := tok.Start()
	for lineStartIdx > 0 {
		if src[lineStartIdx-1] == '\n' {
			break
		}
		lineStartIdx--
	}

	lineEndIdx := tok.End()
	if _, ok := tok.(Newline); !ok {
		for lineEndIdx < len(src)-1 {
			if src[lineEndIdx+1] == '\n' {
				break
			}
			lineEndIdx++
		}
	} else {
		lineEndIdx--
	}

	line := src[lineStartIdx : lineEndIdx+1]

	b.WriteRune(' ')
	b.WriteString(lineNum)
	b.WriteRune(' ')
	b.WriteRune('|')
	b.WriteRune(' ')
	b.Write(line)
	b.WriteRune('\n')

	for range lineNumWidth + 2 {
		b.WriteRune(' ')
	}
	b.WriteRune('|')
	b.WriteRune(' ')

	col := 1
	for col < tok.Column() {
		b.WriteRune(' ')
		col++
	}

	colorPrefix, colorSuffix := err.getColor()
	b.WriteString(colorPrefix)
	for col < tok.Column()+(tok.End()-tok.Start()+1) {
		b.WriteRune('^')
		col++
	}
	b.WriteString(colorSuffix)

	b.WriteRune('\n')

	return b.String()
}
