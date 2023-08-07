package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Argument struct {
	Name  token.Token // token.Ident or nil
	Value Expr
}

type Property struct {
	Name  token.Property
	Value Expr
}

type Body struct {
	Consts []Const
	Insts  []Instantiation
	Props  []Property
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

type Instantiation struct {
	Doc   Doc
	Name  token.Ident
	Count Expr        // If not nil, then it is a list
	Type  token.Token // Basic type, identifier or qualified identifier
	Args  []Argument
	Body  Body
}

func (i Instantiation) eq(i2 Instantiation) bool {
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

func buildInstantiation(toks []token.Token, c *ctx) (Instantiation, error) {
	inst := Instantiation{Name: toks[c.i].(token.Ident)}
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

	if _, ok := toks[c.i].(token.Semicolon); ok {
		props, err := buildPropAssignments(toks, c)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	}

	return inst, nil
}

func buildArgList(toks []token.Token, c *ctx) ([]Argument, error) {
	if _, ok := toks[c.i].(token.LeftParen); !ok {
		return nil, nil
	}
	if _, ok := toks[c.i+1].(token.RightParen); ok {
		return nil, fmt.Errorf(
			"%s: empty argument list", token.Loc(toks[c.i]),
		)
	}

	args := []Argument{}
	a := Argument{}

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
				var (
					err  error
					expr Expr
				)
				expr, err = buildExpr(toks, c, nil)
				if err != nil {
					return nil, err
				}
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
			var (
				err  error
				expr Expr
			)
			expr, err = buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			a.Value = expr
			args = append(args, a)
			state = Comma
		}
	}

	return args, nil
}

func buildPropAssignments(toks []token.Token, c *ctx) ([]Property, error) {
	props := []Property{}
	p := Property{}

	const (
		Prop = iota
		Ass
		Exp
		Semicolon
	)
	state := Prop

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
			var (
				err  error
				expr Expr
			)
			expr, err = buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			p.Value = expr
			props = append(props, p)
			c.i--
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
