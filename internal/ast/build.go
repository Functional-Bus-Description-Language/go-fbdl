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

func buildDoc(toks []token.Token, c *ctx) Doc {
	doc := Doc{}
	doc.Lines = append(doc.Lines, toks[c.i].(token.Comment))

	prevNewline := false
	for {
		c.i++
		switch t := toks[c.i].(type) {
		case token.Newline:
			if prevNewline {
				break
			} else {
				prevNewline = true
			}
		case token.Comment:
			doc.Lines = append(doc.Lines, t)
			prevNewline = false
		default:
			return doc
		}
	}
}

func buildConst(toks []token.Token, c *ctx) (Const, error) {
	switch t := toks[c.i+1].(type) {
	case token.Ident:
		return buildSingleConst(toks, c)
	case token.Newline:
		return buildMultiConst(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(toks []token.Token, c *ctx) (SingleConst, error) {
	sc := SingleConst{Name: toks[c.i+1].(token.Ident)}

	c.i += 2
	if t, ok := toks[c.i].(token.Ass); !ok {
		return sc, unexpected(t, "=")
	}

	c.i++
	expr, err := buildExpr(toks, c, nil)
	if err != nil {
		return sc, err
	}
	sc.Expr = expr

	return sc, nil
}

func buildMultiConst(toks []token.Token, c *ctx) (MultiConst, error) {
	mc := MultiConst{}
	sc := SingleConst{}

	const (
		Indent int = iota
		FirstId
		Ass
		Exp
		Id
	)
	state := Indent

	c.i += 1
tokenLoop:
	for {
		c.i++
		switch state {
		case Indent:
			switch t := toks[c.i].(type) {
			case token.Newline:
				continue
			case token.Indent:
				state = FirstId
			default:
				return mc, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := toks[c.i].(type) {
			case token.Ident:
				sc.Name = t
				state = Ass
			default:
				return mc, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case token.Ass:
				state = Exp
			default:
				return mc, unexpected(t, "=")
			}
		case Exp:
			var (
				err  error
				expr Expr
			)
			expr, err = buildExpr(toks, c, nil)
			if err != nil {
				return mc, err
			}
			sc.Expr = expr
			mc.Consts = append(mc.Consts, sc)
			sc = SingleConst{}
			c.i--
			state = Id
		case Id:
			switch t := toks[c.i].(type) {
			case token.Ident:
				sc.Name = t
				state = Ass
			case token.Comment:
				doc := buildDoc(toks, c)
				sc.Doc = doc
				c.i--
			case token.Newline:
				continue
			case token.Dedent, token.Eof:
				break tokenLoop
			default:
				return mc, unexpected(t, "identifier or dedent")
			}
		}
	}

	return mc, nil
}

func buildImport(toks []token.Token, c *ctx) (Import, error) {
	switch t := toks[c.i+1].(type) {
	case token.Ident, token.String:
		return buildSingleImport(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(toks []token.Token, c *ctx) (SingleImport, error) {
	si := SingleImport{}

	c.i++
	switch t := toks[c.i].(type) {
	case token.Ident:
		si.Name = t
		c.i++
		switch t := toks[c.i].(type) {
		case token.String:
			si.Path = t
			c.i++
		default:
			return si, unexpected(t, "string")
		}
	case token.String:
		si.Path = t
		c.i++
	}

	return si, nil
}
