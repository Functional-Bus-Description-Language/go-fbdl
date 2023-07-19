package token

import "fmt"

type Token interface {
	Start() int
	End() int
	Line() int
	Column() int
	Kind() string
}

// Loc returns location of the token within the file in "line:column" format.
func Loc(t Token) string {
	return fmt.Sprintf("%d:%d", t.Line(), t.Column())
}

type Number interface {
	Token
	numberToken()
}

type None struct {
	start  int
	end    int
	line   int
	column int
}

func (n None) Start() int   { return n.start }
func (n None) End() int     { return n.end }
func (n None) Line() int    { return n.line }
func (n None) Column() int  { return n.column }
func (n None) Kind() string { return "" }

type Comment struct {
	start  int
	end    int
	line   int
	column int
}

func (c Comment) Start() int   { return c.start }
func (c Comment) End() int     { return c.end }
func (c Comment) Line() int    { return c.line }
func (c Comment) Column() int  { return c.column }
func (c Comment) Kind() string { return "comment" }

// Indent increment
type IndentInc struct {
	start  int
	end    int
	line   int
	column int
}

func (ii IndentInc) Start() int   { return ii.start }
func (ii IndentInc) End() int     { return ii.end }
func (ii IndentInc) Line() int    { return ii.line }
func (ii IndentInc) Column() int  { return ii.column }
func (ii IndentInc) Kind() string { return "indent increment" }

// Indent decrement
type IndentDec struct {
	start  int
	end    int
	line   int
	column int
}

func (id IndentDec) Start() int   { return id.start }
func (id IndentDec) End() int     { return id.end }
func (id IndentDec) Line() int    { return id.line }
func (id IndentDec) Column() int  { return id.column }
func (id IndentDec) Kind() string { return "indent decrement" }

type Newline struct {
	start  int
	end    int
	line   int
	column int
}

func (n Newline) Start() int   { return n.start }
func (n Newline) End() int     { return n.end }
func (n Newline) Line() int    { return n.line }
func (n Newline) Column() int  { return n.column }
func (n Newline) Kind() string { return "newline" }

// End of file
type Eof struct {
	start  int
	end    int
	line   int
	column int
}

func (e Eof) Start() int   { return e.start }
func (e Eof) End() int     { return e.end }
func (e Eof) Line() int    { return e.line }
func (e Eof) Column() int  { return e.column }
func (e Eof) Kind() string { return "end of file" }

// Identifier
type Ident struct {
	start  int
	end    int
	line   int
	column int
}

func (i Ident) Start() int   { return i.start }
func (i Ident) End() int     { return i.end }
func (i Ident) Line() int    { return i.line }
func (i Ident) Column() int  { return i.column }
func (i Ident) Kind() string { return "identifier" }

type Bool struct {
	start  int
	end    int
	line   int
	column int
}

func (b Bool) Start() int   { return b.start }
func (b Bool) End() int     { return b.end }
func (b Bool) Line() int    { return b.line }
func (b Bool) Column() int  { return b.column }
func (b Bool) Kind() string { return "bool" }

type Int struct {
	start  int
	end    int
	line   int
	column int
}

func (i Int) Start() int   { return i.start }
func (i Int) End() int     { return i.end }
func (i Int) Line() int    { return i.line }
func (i Int) Column() int  { return i.column }
func (i Int) Kind() string { return "integer" }

func (i Int) numberToken() {}

type Real struct {
	start  int
	end    int
	line   int
	column int
}

func (r Real) Start() int   { return r.start }
func (r Real) End() int     { return r.end }
func (r Real) Line() int    { return r.line }
func (r Real) Column() int  { return r.column }
func (r Real) Kind() string { return "real" }

func (r Real) numberToken() {}

type String struct {
	start  int
	end    int
	line   int
	column int
}

func (s String) Start() int   { return s.start }
func (s String) End() int     { return s.end }
func (s String) Line() int    { return s.line }
func (s String) Column() int  { return s.column }
func (s String) Kind() string { return "string" }

type BitString struct {
	start  int
	end    int
	line   int
	column int
}

func (bs BitString) Start() int   { return bs.start }
func (bs BitString) End() int     { return bs.end }
func (bs BitString) Line() int    { return bs.line }
func (bs BitString) Column() int  { return bs.column }
func (bs BitString) Kind() string { return "bit string" }

type Time struct {
	start  int
	end    int
	line   int
	column int
}

func (t Time) Start() int   { return t.start }
func (t Time) End() int     { return t.end }
func (t Time) Line() int    { return t.line }
func (t Time) Column() int  { return t.column }
func (t Time) Kind() string { return "time" }

// !
type Neg struct {
	start  int
	end    int
	line   int
	column int
}

func (n Neg) Start() int   { return n.start }
func (n Neg) End() int     { return n.end }
func (n Neg) Line() int    { return n.line }
func (n Neg) Column() int  { return n.column }
func (n Neg) Kind() string { return "!" }

// =
type Ass struct {
	start  int
	end    int
	line   int
	column int
}

func (a Ass) Start() int   { return a.start }
func (a Ass) End() int     { return a.end }
func (a Ass) Line() int    { return a.line }
func (a Ass) Column() int  { return a.column }
func (a Ass) Kind() string { return "=" }

// +
type Add struct {
	start  int
	end    int
	line   int
	column int
}

func (a Add) Start() int   { return a.start }
func (a Add) End() int     { return a.end }
func (a Add) Line() int    { return a.line }
func (a Add) Column() int  { return a.column }
func (a Add) Kind() string { return "+" }

// -
type Sub struct {
	start  int
	end    int
	line   int
	column int
}

func (s Sub) Start() int   { return s.start }
func (s Sub) End() int     { return s.end }
func (s Sub) Line() int    { return s.line }
func (s Sub) Column() int  { return s.column }
func (s Sub) Kind() string { return "-" }

// *
type Mul struct {
	start  int
	end    int
	line   int
	column int
}

func (m Mul) Start() int   { return m.start }
func (m Mul) End() int     { return m.end }
func (m Mul) Line() int    { return m.line }
func (m Mul) Column() int  { return m.column }
func (m Mul) Kind() string { return "*" }

// /
type Div struct {
	start  int
	end    int
	line   int
	column int
}

func (d Div) Start() int   { return d.start }
func (d Div) End() int     { return d.end }
func (d Div) Line() int    { return d.line }
func (d Div) Column() int  { return d.column }
func (d Div) Kind() string { return "/" }

// %
type Rem struct {
	start  int
	end    int
	line   int
	column int
}

func (r Rem) Start() int   { return r.start }
func (r Rem) End() int     { return r.end }
func (r Rem) Line() int    { return r.line }
func (r Rem) Column() int  { return r.column }
func (r Rem) Kind() string { return "%" }

// **
type Exp struct {
	start  int
	end    int
	line   int
	column int
}

func (e Exp) Start() int   { return e.start }
func (e Exp) End() int     { return e.end }
func (e Exp) Line() int    { return e.line }
func (e Exp) Column() int  { return e.column }
func (e Exp) Kind() string { return "**" }

// ==
type Eq struct {
	start  int
	end    int
	line   int
	column int
}

func (e Eq) Start() int   { return e.start }
func (e Eq) End() int     { return e.end }
func (e Eq) Line() int    { return e.line }
func (e Eq) Column() int  { return e.column }
func (e Eq) Kind() string { return "==" }

// !=
type Neq struct {
	start  int
	end    int
	line   int
	column int
}

func (n Neq) Start() int   { return n.start }
func (n Neq) End() int     { return n.end }
func (n Neq) Line() int    { return n.line }
func (n Neq) Column() int  { return n.column }
func (n Neq) Kind() string { return "!=" }

// <
type Less struct {
	start  int
	end    int
	line   int
	column int
}

func (l Less) Start() int   { return l.start }
func (l Less) End() int     { return l.end }
func (l Less) Line() int    { return l.line }
func (l Less) Column() int  { return l.column }
func (l Less) Kind() string { return "<" }

// <=
type LessEq struct {
	start  int
	end    int
	line   int
	column int
}

func (le LessEq) Start() int   { return le.start }
func (le LessEq) End() int     { return le.end }
func (le LessEq) Line() int    { return le.line }
func (le LessEq) Column() int  { return le.column }
func (le LessEq) Kind() string { return "<=" }

// >
type Greater struct {
	start  int
	end    int
	line   int
	column int
}

func (g Greater) Start() int   { return g.start }
func (g Greater) End() int     { return g.end }
func (g Greater) Line() int    { return g.line }
func (g Greater) Column() int  { return g.column }
func (g Greater) Kind() string { return ">" }

// >=
type GreaterEq struct {
	start  int
	end    int
	line   int
	column int
}

func (ge GreaterEq) Start() int   { return ge.start }
func (ge GreaterEq) End() int     { return ge.end }
func (ge GreaterEq) Line() int    { return ge.line }
func (ge GreaterEq) Column() int  { return ge.column }
func (ge GreaterEq) Kind() string { return ">=" }

// &&
type And struct {
	start  int
	end    int
	line   int
	column int
}

func (a And) Start() int   { return a.start }
func (a And) End() int     { return a.end }
func (a And) Line() int    { return a.line }
func (a And) Column() int  { return a.column }
func (a And) Kind() string { return "&&" }

// ||
type Or struct {
	start  int
	end    int
	line   int
	column int
}

func (o Or) Start() int   { return o.start }
func (o Or) End() int     { return o.end }
func (o Or) Line() int    { return o.line }
func (o Or) Column() int  { return o.column }
func (o Or) Kind() string { return "||" }

// <<
type LeftShift struct {
	start  int
	end    int
	line   int
	column int
}

func (ls LeftShift) Start() int   { return ls.start }
func (ls LeftShift) End() int     { return ls.end }
func (ls LeftShift) Line() int    { return ls.line }
func (ls LeftShift) Column() int  { return ls.column }
func (ls LeftShift) Kind() string { return "<<" }

// >>
type RightShift struct {
	start  int
	end    int
	line   int
	column int
}

func (rs RightShift) Start() int   { return rs.start }
func (rs RightShift) End() int     { return rs.end }
func (rs RightShift) Line() int    { return rs.line }
func (rs RightShift) Column() int  { return rs.column }
func (ls RightShift) Kind() string { return ">>" }

// &
type BitAnd struct {
	start  int
	end    int
	line   int
	column int
}

func (ba BitAnd) Start() int   { return ba.start }
func (ba BitAnd) End() int     { return ba.end }
func (ba BitAnd) Line() int    { return ba.line }
func (ba BitAnd) Column() int  { return ba.column }
func (ba BitAnd) Kind() string { return "&" }

// |
type BitOr struct {
	start  int
	end    int
	line   int
	column int
}

func (bo BitOr) Start() int   { return bo.start }
func (bo BitOr) End() int     { return bo.end }
func (bo BitOr) Line() int    { return bo.line }
func (bo BitOr) Column() int  { return bo.column }
func (bo BitOr) Kind() string { return "|" }

// ^
type Xor struct {
	start  int
	end    int
	line   int
	column int
}

func (x Xor) Start() int   { return x.start }
func (x Xor) End() int     { return x.end }
func (x Xor) Line() int    { return x.line }
func (x Xor) Column() int  { return x.column }
func (x Xor) Kind() string { return "^" }

// (
type LeftParen struct {
	start  int
	end    int
	line   int
	column int
}

func (lp LeftParen) Start() int   { return lp.start }
func (lp LeftParen) End() int     { return lp.end }
func (lp LeftParen) Line() int    { return lp.line }
func (lp LeftParen) Column() int  { return lp.column }
func (lp LeftParen) Kind() string { return "(" }

// )
type RightParen struct {
	start  int
	end    int
	line   int
	column int
}

func (rp RightParen) Start() int   { return rp.start }
func (rp RightParen) End() int     { return rp.end }
func (rp RightParen) Line() int    { return rp.line }
func (rp RightParen) Column() int  { return rp.column }
func (lp RightParen) Kind() string { return ")" }

// [
type LeftBracket struct {
	start  int
	end    int
	line   int
	column int
}

func (lb LeftBracket) Start() int   { return lb.start }
func (lb LeftBracket) End() int     { return lb.end }
func (lb LeftBracket) Line() int    { return lb.line }
func (lb LeftBracket) Column() int  { return lb.column }
func (lb LeftBracket) Kind() string { return "[" }

// ]
type RightBracket struct {
	start  int
	end    int
	line   int
	column int
}

func (rb RightBracket) Start() int   { return rb.start }
func (rb RightBracket) End() int     { return rb.end }
func (rb RightBracket) Line() int    { return rb.line }
func (rb RightBracket) Column() int  { return rb.column }
func (rb RightBracket) Kind() string { return "]" }

// ,
type Comma struct {
	start  int
	end    int
	line   int
	column int
}

func (c Comma) Start() int   { return c.start }
func (c Comma) End() int     { return c.end }
func (c Comma) Line() int    { return c.line }
func (c Comma) Column() int  { return c.column }
func (c Comma) Kind() string { return "," }

// ;
type Semicolon struct {
	start  int
	end    int
	line   int
	column int
}

func (s Semicolon) Start() int   { return s.start }
func (s Semicolon) End() int     { return s.end }
func (s Semicolon) Line() int    { return s.line }
func (s Semicolon) Column() int  { return s.column }
func (s Semicolon) Kind() string { return ";" }

type Const struct {
	start  int
	end    int
	line   int
	column int
}

func (c Const) Start() int   { return c.start }
func (c Const) End() int     { return c.end }
func (c Const) Line() int    { return c.line }
func (c Const) Column() int  { return c.column }
func (c Const) Kind() string { return "const" }

type Import struct {
	start  int
	end    int
	line   int
	column int
}

func (i Import) Start() int   { return i.start }
func (i Import) End() int     { return i.end }
func (i Import) Line() int    { return i.line }
func (i Import) Column() int  { return i.column }
func (i Import) Kind() string { return "import" }

type Type struct {
	start  int
	end    int
	line   int
	column int
}

func (t Type) Start() int   { return t.start }
func (t Type) End() int     { return t.end }
func (t Type) Line() int    { return t.line }
func (t Type) Column() int  { return t.column }
func (t Type) Kind() string { return "type" }

type Block struct {
	start  int
	end    int
	line   int
	column int
}

func (b Block) Start() int   { return b.start }
func (b Block) End() int     { return b.end }
func (b Block) Line() int    { return b.line }
func (b Block) Column() int  { return b.column }
func (b Block) Kind() string { return "block" }

type Bus struct {
	start  int
	end    int
	line   int
	column int
}

func (b Bus) Start() int   { return b.start }
func (b Bus) End() int     { return b.end }
func (b Bus) Line() int    { return b.line }
func (b Bus) Column() int  { return b.column }
func (b Bus) Kind() string { return "bus" }

type Config struct {
	start  int
	end    int
	line   int
	column int
}

func (c Config) Start() int   { return c.start }
func (c Config) End() int     { return c.end }
func (c Config) Line() int    { return c.line }
func (c Config) Column() int  { return c.column }
func (c Config) Kind() string { return "config" }

type Irq struct {
	start  int
	end    int
	line   int
	column int
}

func (i Irq) Start() int   { return i.start }
func (i Irq) End() int     { return i.end }
func (i Irq) Line() int    { return i.line }
func (i Irq) Column() int  { return i.column }
func (i Irq) Kind() string { return "irq" }

type Mask struct {
	start  int
	end    int
	line   int
	column int
}

func (m Mask) Start() int   { return m.start }
func (m Mask) End() int     { return m.end }
func (m Mask) Line() int    { return m.line }
func (m Mask) Column() int  { return m.column }
func (m Mask) Kind() string { return "mask" }

type Memory struct {
	start  int
	end    int
	line   int
	column int
}

func (m Memory) Start() int   { return m.start }
func (m Memory) End() int     { return m.end }
func (m Memory) Line() int    { return m.line }
func (m Memory) Column() int  { return m.column }
func (m Memory) Kind() string { return "memory" }

type Param struct {
	start  int
	end    int
	line   int
	column int
}

func (p Param) Start() int   { return p.start }
func (p Param) End() int     { return p.end }
func (p Param) Line() int    { return p.line }
func (p Param) Column() int  { return p.column }
func (p Param) Kind() string { return "param" }

type Proc struct {
	start  int
	end    int
	line   int
	column int
}

func (p Proc) Start() int   { return p.start }
func (p Proc) End() int     { return p.end }
func (p Proc) Line() int    { return p.line }
func (p Proc) Column() int  { return p.column }
func (p Proc) Kind() string { return "proc" }

type Return struct {
	start  int
	end    int
	line   int
	column int
}

func (r Return) Start() int   { return r.start }
func (r Return) End() int     { return r.end }
func (r Return) Line() int    { return r.line }
func (r Return) Column() int  { return r.column }
func (r Return) Kind() string { return "return" }

type Static struct {
	start  int
	end    int
	line   int
	column int
}

func (s Static) Start() int   { return s.start }
func (s Static) End() int     { return s.end }
func (s Static) Line() int    { return s.line }
func (s Static) Column() int  { return s.column }
func (s Static) Kind() string { return "static" }

type Stream struct {
	start  int
	end    int
	line   int
	column int
}

func (s Stream) Start() int   { return s.start }
func (s Stream) End() int     { return s.end }
func (s Stream) Line() int    { return s.line }
func (s Stream) Column() int  { return s.column }
func (s Stream) Kind() string { return "stream" }

type Access struct {
	start  int
	end    int
	line   int
	column int
}

func (a Access) Start() int   { return a.start }
func (a Access) End() int     { return a.end }
func (a Access) Line() int    { return a.line }
func (a Access) Column() int  { return a.column }
func (a Access) Kind() string { return "access" }

type AddEnable struct {
	start  int
	end    int
	line   int
	column int
}

func (ae AddEnable) Start() int   { return ae.start }
func (ae AddEnable) End() int     { return ae.end }
func (ae AddEnable) Line() int    { return ae.line }
func (ae AddEnable) Column() int  { return ae.column }
func (ae AddEnable) Kind() string { return "add-enable" }

type Atomic struct {
	start  int
	end    int
	line   int
	column int
}

func (a Atomic) Start() int   { return a.start }
func (a Atomic) End() int     { return a.end }
func (a Atomic) Line() int    { return a.line }
func (a Atomic) Column() int  { return a.column }
func (a Atomic) Kind() string { return "atomic" }

type ByteWriteEnable struct {
	start  int
	end    int
	line   int
	column int
}

func (bwe ByteWriteEnable) Start() int   { return bwe.start }
func (bwe ByteWriteEnable) End() int     { return bwe.end }
func (bwe ByteWriteEnable) Line() int    { return bwe.line }
func (bwe ByteWriteEnable) Column() int  { return bwe.column }
func (bwe ByteWriteEnable) Kind() string { return "byte-write-enable" }

type Clear struct {
	start  int
	end    int
	line   int
	column int
}

func (c Clear) Start() int   { return c.start }
func (c Clear) End() int     { return c.end }
func (c Clear) Line() int    { return c.line }
func (c Clear) Column() int  { return c.column }
func (c Clear) Kind() string { return "clear" }

type Delay struct {
	start  int
	end    int
	line   int
	column int
}

func (d Delay) Start() int   { return d.start }
func (d Delay) End() int     { return d.end }
func (d Delay) Line() int    { return d.line }
func (d Delay) Column() int  { return d.column }
func (d Delay) Kind() string { return "delay" }

type EnableInitValue struct {
	start  int
	end    int
	line   int
	column int
}

func (eiv EnableInitValue) Start() int   { return eiv.start }
func (eiv EnableInitValue) End() int     { return eiv.end }
func (eiv EnableInitValue) Line() int    { return eiv.line }
func (eiv EnableInitValue) Column() int  { return eiv.column }
func (eiv EnableInitValue) Kind() string { return "enable-init-value" }

type EnableResetValue struct {
	start  int
	end    int
	line   int
	column int
}

func (erv EnableResetValue) Start() int   { return erv.start }
func (erv EnableResetValue) End() int     { return erv.end }
func (erv EnableResetValue) Line() int    { return erv.line }
func (erv EnableResetValue) Column() int  { return erv.column }
func (erv EnableResetValue) Kind() string { return "enable-reset-value" }

type Groups struct {
	start  int
	end    int
	line   int
	column int
}

func (g Groups) Start() int   { return g.start }
func (g Groups) End() int     { return g.end }
func (g Groups) Line() int    { return g.line }
func (g Groups) Column() int  { return g.column }
func (g Groups) Kind() string { return "groups" }

type InitValue struct {
	start  int
	end    int
	line   int
	column int
}

func (iv InitValue) Start() int   { return iv.start }
func (iv InitValue) End() int     { return iv.end }
func (iv InitValue) Line() int    { return iv.line }
func (iv InitValue) Column() int  { return iv.column }
func (iv InitValue) Kind() string { return "init-value" }

type InTrigger struct {
	start  int
	end    int
	line   int
	column int
}

func (it InTrigger) Start() int   { return it.start }
func (it InTrigger) End() int     { return it.end }
func (it InTrigger) Line() int    { return it.line }
func (it InTrigger) Column() int  { return it.column }
func (it InTrigger) Kind() string { return "in-trigger" }

type Masters struct {
	start  int
	end    int
	line   int
	column int
}

func (m Masters) Start() int   { return m.start }
func (m Masters) End() int     { return m.end }
func (m Masters) Line() int    { return m.line }
func (m Masters) Column() int  { return m.column }
func (m Masters) Kind() string { return "masters" }

type OutTrigger struct {
	start  int
	end    int
	line   int
	column int
}

func (ot OutTrigger) Start() int   { return ot.start }
func (ot OutTrigger) End() int     { return ot.end }
func (ot OutTrigger) Line() int    { return ot.line }
func (ot OutTrigger) Column() int  { return ot.column }
func (ot OutTrigger) Kind() string { return "out-trigger" }

type Range struct {
	start  int
	end    int
	line   int
	column int
}

func (r Range) Start() int   { return r.start }
func (r Range) End() int     { return r.end }
func (r Range) Line() int    { return r.line }
func (r Range) Column() int  { return r.column }
func (r Range) Kind() string { return "range" }

type ReadLatency struct {
	start  int
	end    int
	line   int
	column int
}

func (rl ReadLatency) Start() int   { return rl.start }
func (rl ReadLatency) End() int     { return rl.end }
func (rl ReadLatency) Line() int    { return rl.line }
func (rl ReadLatency) Column() int  { return rl.column }
func (rl ReadLatency) Kind() string { return "read-latency" }

type ReadValue struct {
	start  int
	end    int
	line   int
	column int
}

func (rv ReadValue) Start() int   { return rv.start }
func (rv ReadValue) End() int     { return rv.end }
func (rv ReadValue) Line() int    { return rv.line }
func (rv ReadValue) Column() int  { return rv.column }
func (rv ReadValue) Kind() string { return "read-value" }

type Reset struct {
	start  int
	end    int
	line   int
	column int
}

func (r Reset) Start() int   { return r.start }
func (r Reset) End() int     { return r.end }
func (r Reset) Line() int    { return r.line }
func (r Reset) Column() int  { return r.column }
func (r Reset) Kind() string { return "reset" }

type ResetValue struct {
	start  int
	end    int
	line   int
	column int
}

func (rv ResetValue) Start() int   { return rv.start }
func (rv ResetValue) End() int     { return rv.end }
func (rv ResetValue) Line() int    { return rv.line }
func (rv ResetValue) Column() int  { return rv.column }
func (rv ResetValue) Kind() string { return "reset-value" }

type Size struct {
	start  int
	end    int
	line   int
	column int
}

func (s Size) Start() int   { return s.start }
func (s Size) End() int     { return s.end }
func (s Size) Line() int    { return s.line }
func (s Size) Column() int  { return s.column }
func (s Size) Kind() string { return "size" }

type Width struct {
	start  int
	end    int
	line   int
	column int
}

func (w Width) Start() int   { return w.start }
func (w Width) End() int     { return w.end }
func (w Width) Line() int    { return w.line }
func (w Width) Column() int  { return w.column }
func (w Width) Kind() string { return "width" }

// . - currently unused
type Period struct {
	start  int
	end    int
	line   int
	column int
}

func (p Period) Start() int   { return p.start }
func (p Period) End() int     { return p.end }
func (p Period) Line() int    { return p.line }
func (p Period) Column() int  { return p.column }
func (p Period) Kind() string { return "." }

// : - currently unused
type Colon struct {
	start  int
	end    int
	line   int
	column int
}

func (c Colon) Start() int   { return c.start }
func (c Colon) End() int     { return c.end }
func (c Colon) Line() int    { return c.line }
func (c Colon) Column() int  { return c.column }
func (c Colon) Kind() string { return ":" }

// { - currently unused
type LeftBrace struct {
	start  int
	end    int
	line   int
	column int
}

func (lb LeftBrace) Start() int   { return lb.start }
func (lb LeftBrace) End() int     { return lb.end }
func (lb LeftBrace) Line() int    { return lb.line }
func (lb LeftBrace) Column() int  { return lb.column }
func (lb LeftBrace) Kind() string { return "{" }

// } - currently unused
type RightBrace struct {
	start  int
	end    int
	line   int
	column int
}

func (rb RightBrace) Start() int   { return rb.start }
func (rb RightBrace) End() int     { return rb.end }
func (rb RightBrace) Line() int    { return rb.line }
func (rb RightBrace) Column() int  { return rb.column }
func (rb RightBrace) Kind() string { return "}" }
