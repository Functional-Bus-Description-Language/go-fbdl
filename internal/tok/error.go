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

	return fmt.Sprintf(
		"%serror%s: %s\n%s +%d:%d\n%s",
		colorPrefix, colorSuffix, err.Msg, err.Toks[0].Path(), err.Toks[0].Line(), err.Toks[0].Column(), err.code(),
	)
}

// Returns error token code.
func (err Error) code() string {
	src := err.Toks[0].Src()
	b := strings.Builder{}

	lineNum := strconv.FormatInt(int64(err.Toks[0].Line()), 10)
	lineNumWidth := len(lineNum)
	for i := 0; i < lineNumWidth+2; i++ {
		b.WriteRune(' ')
	}
	b.WriteString("|\n")

	lineStartIdx := err.Toks[0].Start()
	for {
		if lineStartIdx == 0 || src[lineStartIdx-1] == '\n' {
			break
		}
		lineStartIdx--
	}

	lineEndIdx := err.Toks[0].End()
	if _, ok := err.Toks[0].(Newline); !ok {
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
	if err.Toks[0].Column() > 1 {
		for i := 0; i < indent; i++ {
			b.WriteRune('\t')
			col++
		}
	}

	for {
		if col == err.Toks[0].Column() {
			break
		}
		b.WriteRune(' ')
		col++
	}

	colorPrefix, colorSuffix := err.getColor()

	b.WriteString(colorPrefix)
	for {
		if col == err.Toks[0].Column()+(err.Toks[0].End()-err.Toks[0].Start()+1) {
			break
		}
		b.WriteRune('^')
		col++
	}
	b.WriteString(colorSuffix)

	b.WriteRune('\n')

	return b.String()
}
