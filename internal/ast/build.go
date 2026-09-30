package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Build builds ast from provided source.
func Build(src []byte, path string) (File, error) {
	var (
		err    error
		f      File
		ctx    context
		doc    Doc
		consts []Const
		imps   []Import
		ins    Inst
		typ    Type
	)

	toks, err := token.Parse(src, path)
	if err != nil {
		return File{}, err
	}
	ctx.toks = toks

	for {
		if _, ok := ctx.token().(token.Eof); ok {
			break
		}

		switch t := ctx.token().(type) {
		case token.Newline:
			ctx.idx++
		case token.Comment:
			doc = buildDoc(&ctx)
		case token.Const:
			consts, err = buildConst(&ctx)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()-1 {
					consts[0].Doc = doc
				}
				f.Consts = append(f.Consts, consts...)
			}
		case token.Ident:
			ins, err = buildInst(&ctx)
			if doc.endLine() == ins.Name.Line()-1 {
				ins.Doc = doc
			}
			f.Insts = append(f.Insts, ins)
		case token.Import:
			imps, err = buildImport(&ctx)
			if len(imps) > 0 {
				f.Imports = append(f.Imports, imps...)
			}
		case token.Type:
			typ, err = buildType(&ctx)
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
