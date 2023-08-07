package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Building context
type ctx struct {
	i int // Current token index
}

// Build builds ast based on token stream.
func Build(toks []token.Token) (File, error) {
	var (
		err    error
		f      File
		c      ctx
		doc    Doc
		consts []Const
		imp    Import
	)

	for {
		if _, ok := toks[c.i].(token.Eof); ok {
			break
		}

		switch t := toks[c.i].(type) {
		case token.Newline:
			c.i++
		case token.Comment:
			doc = buildDoc(toks, &c)
		case token.Const:
			consts, err = buildConst(toks, &c)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()+1 {
					consts[0].Doc = doc
				}
				f.Consts = append(f.Consts, consts...)
			}
		case token.Import:
			imp, err = buildImport(toks, &c)
			f.Imports = append(f.Imports, imp)
		default:
			panic(fmt.Sprintf("%s: unhandled token %s", token.Loc(t), t.Kind()))
		}

		if err != nil {
			return f, err
		}
	}

	return f, nil
}
