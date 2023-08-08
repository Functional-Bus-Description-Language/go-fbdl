package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Arg struct represents instantiation argument.
type Arg struct {
	Name  token.Token // token.Ident or nil
	Value Expr
}

// The Prop struct represents functionality property.
type Prop struct {
	Name  token.Property
	Value Expr
}

// The Body struct represents functionality body.
type Body struct {
	Consts []Const
	Insts  []Inst
	Props  []Prop
}

func (b Body) eq(b2 Body) bool {
	if len(b.Consts) != len(b2.Consts) ||
		len(b.Insts) != len(b2.Insts) ||
		len(b.Props) != len(b2.Props) {
		return false
	}

	for i := range b.Consts {
		if !b.Consts[i].eq(b2.Consts[i]) {
			return false
		}
	}
	for i := range b.Insts {
		if !b.Insts[i].eq(b2.Insts[i]) {
			return false
		}
	}
	for i := range b.Props {
		if b.Props[i] != b2.Props[i] {
			return false
		}
	}

	return true
}

// The Inst struct represents functionality instantiation.
type Inst struct {
	Doc   Doc
	Name  token.Ident
	Count Expr        // If not nil, then it is a list
	Type  token.Token // Basic type, identifier or qualified identifier
	Args  []Arg
	Body  Body
}

func (i Inst) eq(i2 Inst) bool {
	if !i.Doc.eq(i2.Doc) ||
		i.Name != i2.Name ||
		i.Count != i2.Count ||
		i.Type != i2.Type ||
		len(i.Args) != len(i2.Args) ||
		!i.Body.eq(i2.Body) {
		return false
	}

	for n := range i.Args {
		if i.Args[n] != i2.Args[n] {
			return false
		}
	}

	return true
}

func buildInst(toks []token.Token, c *ctx) (Inst, error) {
	inst := Inst{Name: toks[c.i].(token.Ident)}
	c.i++

	if _, ok := toks[c.i].(token.LeftBracket); ok {
		c.i++
		expr, err := buildExpr(toks, c, nil)
		if err != nil {
			return inst, err
		}
		inst.Count = expr

		if _, ok := toks[c.i].(token.RightBracket); !ok {
			return inst, unexpected(toks[c.i], "]")
		}
		c.i++
	}

	switch t := toks[c.i].(type) {
	case token.Functionality, token.Ident, token.QualIdent:
		inst.Type = t
		c.i++
	default:
		return inst, unexpected(t, "functionality type")
	}

	args, err := buildArgList(toks, c)
	if err != nil {
		return inst, err
	}
	inst.Args = args

	switch t := toks[c.i].(type) {
	case token.Semicolon:
		c.i++
		props, err := buildPropAssignments(toks, c)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	case token.Newline:
		if _, ok := toks[c.i+1].(token.Indent); ok {
			c.i += 2
			body, err := buildBody(toks, c)
			if err != nil {
				return inst, err
			}
			inst.Body = body
		}
	default:
		return inst, unexpected(t, "; or newline")
	}

	return inst, nil
}

func buildArgList(toks []token.Token, c *ctx) ([]Arg, error) {
	if _, ok := toks[c.i].(token.LeftParen); !ok {
		return nil, nil
	}
	if _, ok := toks[c.i+1].(token.RightParen); ok {
		return nil, fmt.Errorf(
			"%s: empty argument list", token.Loc(toks[c.i]),
		)
	}

	args := []Arg{}
	a := Arg{}

	const (
		Name = iota
		Ass
		Comma
		Exp
	)
	state := Name

tokenLoop:
	for {
		c.i++
		switch state {
		case Name:
			switch t := toks[c.i].(type) {
			case token.Ident:
				a.Name = t
				state = Ass
			default:
				a.Name = nil
				expr, err := buildExpr(toks, c, nil)
				if err != nil {
					return nil, err
				}
				c.i--
				a.Value = expr
				args = append(args, a)
				state = Comma
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "=")
			}
		case Comma:
			switch t := toks[c.i].(type) {
			case token.Comma:
				state = Name
			case token.RightParen:
				c.i++
				break tokenLoop
			default:
				return nil, unexpected(t, ", or )")
			}
		case Exp:
			expr, err := buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			c.i--
			a.Value = expr
			args = append(args, a)
			state = Comma
		}
	}

	return args, nil
}

func buildPropAssignments(toks []token.Token, c *ctx) ([]Prop, error) {
	props := []Prop{}
	p := Prop{}

	const (
		Prop = iota
		Ass
		Exp
		Semicolon
	)
	state := Prop

	// Decrement contex index as it is incremented at the beginnig of the for loop.
	c.i--
tokenLoop:
	for {
		c.i++
		switch state {
		case Prop:
			switch t := toks[c.i].(type) {
			case token.Property:
				p.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "property name")
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "=")
			}
		case Exp:
			expr, err := buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			c.i--
			p.Value = expr
			props = append(props, p)
			state = Semicolon
		case Semicolon:
			switch t := toks[c.i].(type) {
			case token.Newline, token.Eof:
				break tokenLoop
			case token.Semicolon:
				state = Prop
			default:
				return nil, unexpected(t, "; or newline")
			}
		}
	}

	return props, nil
}

func buildBody(toks []token.Token, c *ctx) (Body, error) {
	var (
		err    error
		body   Body
		consts []Const
		doc    Doc
		ins    Inst
		props  []Prop
	)

	for {
		if _, ok := toks[c.i].(token.Eof); ok {
			break
		}

		switch t := toks[c.i].(type) {
		case token.Newline:
			c.i++
		case token.Comment:
			doc = buildDoc(toks, c)
		case token.Const:
			consts, err = buildConst(toks, c)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()+1 {
					consts[0].Doc = doc
				}
				body.Consts = append(body.Consts, consts...)
			}
		case token.Ident:
			ins, err = buildInst(toks, c)
			body.Insts = append(body.Insts, ins)
		case token.Property:
			props, err = buildPropAssignments(toks, c)
			if err != nil {
				return body, err
			}
			if props != nil {
				body.Props = append(body.Props, props...)
			}
		default:
			panic(fmt.Sprintf("%s: unhandled token %s", token.Loc(t), t.Kind()))
		}

		if err != nil {
			return body, err
		}
	}

	return body, nil
}
