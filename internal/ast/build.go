package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Build builds an AST based on Token vector.
func Build(toks []token.Token) (File, error) {
	var (
		err    error
		f      File
		tokens = tokenStream{toks: toks}
		doc    Doc
		consts []Const
		imps   []Import
		ins    Inst
		typ    Type
	)

tokenLoop:
	for {
		switch t := tokens.peek(0).(type) {
		case token.Eof:
			break tokenLoop
		case token.Newline:
			tokens.next()
		case token.Comment:
			doc = buildDoc(&tokens)
		case token.Const:
			consts, err = buildConst(&tokens)
			if len(consts) > 0 {
				if !doc.isEmpty() {
					consts[0].Doc = doc
				}
				f.Consts = append(f.Consts, consts...)
			}
		case token.Ident:
			ins, err = buildInst(&tokens)
			if !doc.isEmpty() {
				ins.Doc = doc
			}
			f.Insts = append(f.Insts, ins)
		case token.Import:
			imps, err = buildImport(&tokens)
			if len(imps) > 0 {
				f.Imports = append(f.Imports, imps...)
			}
		case token.Type:
			typ, err = buildType(&tokens)
			f.Types = append(f.Types, typ)
		default:
			return f, unexpected(t, "const, type, identifier, import or comment")
		}

		if err != nil {
			return f, err
		}
	}

	return f, nil
}
