package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Build builds an AST from a non-empty token stream ending in token.Eof.
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
			continue
		case token.Const:
			consts, err = buildConst(&tokens)
			if len(consts) > 0 {
				if consts[0].Doc.isEmpty() && !doc.isEmpty() {
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
			typ.Doc = doc
			f.Types = append(f.Types, typ)
		default:
			return f, unexpected(t, "'const', 'import', 'type', identifier, or comment")
		}

		doc = emptyDoc()

		if err != nil {
			return f, err
		}
	}

	return f, nil
}
