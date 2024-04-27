package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Building context
type context struct {
	i int // Current token index
}

// Build builds ast from provided source.
func Build(src []byte, path string) (File, error) {
	var (
		err    error
		f      File
		ctx    context
		doc    Doc
		consts []Const
		imps   []Import
		ins    Instance
		typ    Type
	)

	toks, err := tok.Parse([]byte(src), path)
	if err != nil {
		return File{}, err
	}

	for {
		if _, ok := toks[ctx.i].(tok.Eof); ok {
			break
		}

		switch t := toks[ctx.i].(type) {
		case tok.Newline:
			ctx.i++
		case tok.Comment:
			doc = buildDoc(toks, &ctx)
		case tok.Const:
			consts, err = buildConst(toks, &ctx)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()-1 {
					consts[0].Doc = doc
				}
				f.Consts = append(f.Consts, consts...)
			}
		case tok.Ident:
			ins, err = buildInst(toks, &ctx)
			if doc.endLine() == ins.Name.Line()-1 {
				ins.Doc = doc
			}
			f.Insts = append(f.Insts, ins)
		case tok.Import:
			imps, err = buildImport(toks, &ctx)
			if len(imps) > 0 {
				f.Imports = append(f.Imports, imps...)
			}
		case tok.Type:
			typ, err = buildType(toks, &ctx)
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
