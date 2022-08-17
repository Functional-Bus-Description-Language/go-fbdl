package prs

import (
	"fmt"
)

type comment struct {
	msg        string
	endLineNum uint32
}

func emptyComment() comment {
	return comment{msg: "", endLineNum: 0}
}

func (c comment) isEmpty() bool {
	if c.endLineNum == 0 {
		return true
	}

	return false
}
func makeComment(s string, lineNum uint32) comment {
	firstByte := 1
	if s[1] == ' ' {
		firstByte = 2
	}

	return comment{msg: s[firstByte:], endLineNum: lineNum}
}

func (c *comment) append(s string) {
	if len(s) == 1 {
		c.msg += "\n"
		c.endLineNum += 1
		return
	}

	firstByte := 1
	if s[1] == ' ' {
		firstByte = 2
	}
	c.msg += fmt.Sprintf("\n%s", s[firstByte:])
	c.endLineNum += 1
}
