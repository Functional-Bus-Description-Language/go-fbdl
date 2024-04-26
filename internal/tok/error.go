package tok

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
	"strconv"
	"strings"
)

type Error struct {
	Msg string
	Tok Token
}

func (err Error) getColor() (string, string) {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return "\033[1;31m", "\033[0m"
	}
	return "", ""
}

func (err Error) Error() string {
	colorPrefix, colorSuffix := err.getColor()

	return fmt.Sprintf(
		"%serror%s: %s\n%s +%d:%d\n%s",
		colorPrefix, colorSuffix, err.Msg, err.Tok.Path(), err.Tok.Line(), err.Tok.Column(), err.code(),
	)
}

// Returns error token code.
func (err Error) code() string {
	src := err.Tok.Src()
	b := strings.Builder{}

	lineNum := strconv.FormatInt(int64(err.Tok.Line()), 10)
	lineNumWidth := len(lineNum)
	for i := 0; i < lineNumWidth+2; i++ {
		b.WriteRune(' ')
	}
	b.WriteString("|\n")

	lineStartIdx := err.Tok.Start()
	for {
		if lineStartIdx == 0 || src[lineStartIdx-1] == '\n' {
			break
		}
		lineStartIdx--
	}

	lineEndIdx := err.Tok.End()
	if _, ok := err.Tok.(Newline); !ok {
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
	if err.Tok.Column() > 1 {
		for i := 0; i < indent; i++ {
			b.WriteRune('\t')
			col++
		}
	}

	for {
		if col == err.Tok.Column() {
			break
		}
		b.WriteRune(' ')
		col++
	}

	colorPrefix, colorSuffix := err.getColor()

	b.WriteString(colorPrefix)
	for {
		if col == err.Tok.Column()+(err.Tok.End()-err.Tok.Start()+1) {
			break
		}
		b.WriteRune('^')
		col++
	}
	b.WriteString(colorSuffix)

	b.WriteRune('\n')

	return b.String()
}
