package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Building context
type ctx struct {
	i int // Current token index
}

// Build builds ast from provided source.
func Build(src []byte) (File, error) {
	var (
		err    error
		f      File
		c      ctx
		doc    Doc
		consts []Const
		imps   []Import
		ins    Inst
		typ    Type
	)

	toks, err := tok.Parse([]byte(src))
	if err != nil {
		return File{}, err
	}

	for {
		if _, ok := toks[c.i].(tok.Eof); ok {
			break
		}

		switch t := toks[c.i].(type) {
		case tok.Newline:
			c.i++
		case tok.Comment:
			doc = buildDoc(toks, &c)
		case tok.Const:
			consts, err = buildConst(toks, &c)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()+1 {
					consts[0].Doc = doc
				}
				f.Consts = append(f.Consts, consts...)
			}
		case tok.Ident:
			ins, err = buildInst(toks, &c)
			f.Insts = append(f.Insts, ins)
		case tok.Import:
			imps, err = buildImport(toks, &c)
			if len(imps) > 0 {
				f.Imports = append(f.Imports, imps...)
			}
		case tok.Type:
			typ, err = buildType(toks, &c)
			f.Types = append(f.Types, typ)
		default:
			panic(fmt.Sprintf("%s: unhandled token %s", tok.Loc(t), t.Kind()))
		}

		if err != nil {
			return f, err
		}
	}

	return f, nil
}
