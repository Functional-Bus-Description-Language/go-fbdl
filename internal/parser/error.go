package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
	"github.com/mattn/go-isatty"
)

// Error describes a token error in a source file.
type Error struct {
	err  token.Error
	file *File
}

func NewError(tokErr token.Error, file *File) Error {
	return Error{tokErr, file}
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
		colorPrefix, colorSuffix, err.err.Msg,
	)

	for _, tok := range err.err.Toks {
		msg += err.code(tok)
	}

	return msg
}

// Returns error token code.
func (err Error) code(tok token.Token) string {
	file := err.file
	b := strings.Builder{}

	fmt.Fprintf(
		&b,
		"%s +%d:%d\n",
		file.Path, lineNumber(tok, file.Src), columnNumber(tok, file.Src),
	)

	lineNum := strconv.FormatInt(int64(lineNumber(tok, file.Src)), 10)
	lineNumWidth := len(lineNum)
	for range lineNumWidth + 2 {
		b.WriteRune(' ')
	}
	b.WriteString("|\n")

	src := file.Src
	lineStartIdx := tok.Start()
	for lineStartIdx > 0 {
		if src[lineStartIdx-1] == '\n' {
			break
		}
		lineStartIdx--
	}

	lineEndIdx := min(tok.End()-1, len(src)-1)
	if _, ok := tok.(token.Newline); !ok {
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
	for col < columnNumber(tok, file.Src) {
		b.WriteRune(' ')
		col++
	}

	colorPrefix, colorSuffix := err.getColor()
	b.WriteString(colorPrefix)
	for col < columnNumber(tok, file.Src)+max(1, tok.End()-tok.Start()) {
		b.WriteRune('^')
		col++
	}
	b.WriteString(colorSuffix)

	b.WriteRune('\n')

	return b.String()
}

func (err Error) Unwrap() error { return err.err }

func lineNumber(tok token.Token, src []byte) int {
	line := 1
	for _, b := range src[:tok.Start()] {
		if b == '\n' {
			line++
		}
	}
	return line
}

func columnNumber(tok token.Token, src []byte) int {
	col := 1
	for i := tok.Start() - 1; i >= 0 && src[i] != '\n'; i-- {
		col++
	}
	return col
}
