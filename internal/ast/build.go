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
		err error
		f   File
		c   ctx
		doc Doc
		con Const
		imp Import
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
			con, err = buildConst(toks, &c)
			if con, ok := con.(SingleConst); ok {
				if doc.endLine() == con.Name.Line()+1 {
					con.Doc = doc
				}
			}
			f.Consts = append(f.Consts, con)
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
