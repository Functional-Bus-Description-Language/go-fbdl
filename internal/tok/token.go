package tok

import "fmt"

// Various token types
type (
	Token interface {
		Start() int
		End() int
		Line() int
		Column() int
		Name() string
	}

	Functionality interface {
		Token
		functionality()
	}

	Number interface {
		Token
		number()
	}

	Operator interface {
		Token
		Precedence() int
	}

	Property interface {
		Token
		property()
	}
)

// Loc returns location of the token within the file in "line:column" format.
func Loc(t Token) string {
	return fmt.Sprintf("%d:%d", t.Line(), t.Column())
}

// Text returns token text from the source.
func Text(t Token, src []byte) string {
	return string(src[t.Start() : t.End()+1])
}

type None struct {
	position
}

func (n None) Name() string { return "" }

type Comment struct {
	position
}

func (c Comment) Name() string { return "comment" }

// Indent increment
type Indent struct {
	position
}

func (i Indent) Name() string { return "indent increment" }

// Indent decrement
type Dedent struct {
	position
}

func (d Dedent) Name() string { return "indent decrement" }

type Newline struct {
	position
}

func (n Newline) Name() string { return "newline" }

// End of file
type Eof struct {
	position
}

func (e Eof) Name() string { return "end of file" }

// Identifier
type Ident struct {
	position
}

func (i Ident) Name() string { return "identifier" }

// Qualified Identifier
type QualIdent struct {
	position
}

func (qi QualIdent) Name() string { return "qualified identifier" }

type Bool struct {
	position
}

func (b Bool) Name() string { return "bool" }

type Int struct {
	position
}

func (i Int) Name() string { return "integer" }

func (i Int) number() {}

type Real struct {
	position
}

func (r Real) Name() string { return "real" }

func (r Real) number() {}

type String struct {
	position
}

func (s String) Name() string { return "string" }

type BitString struct {
	position
}

func (bs BitString) Name() string { return "bit string" }

type Time struct {
	position
}

func (t Time) Name() string { return "time" }

// !
type Neg struct {
	position
}

func (n Neg) Name() string { return "!" }

// =
type Ass struct {
	position
}

func (a Ass) Name() string { return "'='" }

// +
type Add struct {
	position
}

func (a Add) Name() string    { return "'+'" }
func (a Add) Precedence() int { return 4 }

// -
type Sub struct {
	position
}

func (s Sub) Name() string    { return "'-'" }
func (s Sub) Precedence() int { return 4 }

// *
type Mul struct {
	position
}

func (m Mul) Name() string    { return "'*'" }
func (m Mul) Precedence() int { return 5 }

// /
type Div struct {
	position
}

func (d Div) Name() string    { return "'/'" }
func (d Div) Precedence() int { return 5 }

// %
type Rem struct {
	position
}

func (r Rem) Name() string    { return "'%'" }
func (r Rem) Precedence() int { return 5 }

// **
type Exp struct {
	position
}

func (e Exp) Name() string    { return "'**'" }
func (e Exp) Precedence() int { return 6 }

// ==
type Eq struct {
	position
}

func (e Eq) Name() string    { return "'=='" }
func (e Eq) Precedence() int { return 3 }

// !=
type Neq struct {
	position
}

func (n Neq) Name() string    { return "'!='" }
func (n Neq) Precedence() int { return 3 }

// <
type Less struct {
	position
}

func (l Less) Name() string    { return "'<'" }
func (l Less) Precedence() int { return 3 }

// <=
type LessEq struct {
	position
}

func (le LessEq) Name() string    { return "'<='" }
func (le LessEq) Precedence() int { return 3 }

// >
type Greater struct {
	position
}

func (g Greater) Name() string    { return "'>'" }
func (g Greater) Precedence() int { return 3 }

// >=
type GreaterEq struct {
	position
}

func (ge GreaterEq) Name() string    { return "'>='" }
func (ge GreaterEq) Precedence() int { return 3 }

// &&
type And struct {
	position
}

func (a And) Name() string    { return "'&&'" }
func (a And) Precedence() int { return 2 }

// ||
type Or struct {
	position
}

func (o Or) Name() string    { return "'||'" }
func (o Or) Precedence() int { return 1 }

// <<
type LeftShift struct {
	position
}

func (ls LeftShift) Name() string    { return "'<<'" }
func (ls LeftShift) Precedence() int { return 5 }

// >>
type RightShift struct {
	position
}

func (ls RightShift) Name() string    { return "'>>'" }
func (ls RightShift) Precedence() int { return 5 }

// &
type BitAnd struct {
	position
}

func (ba BitAnd) Name() string    { return "'&'" }
func (ba BitAnd) Precedence() int { return 5 }

// |
type BitOr struct {
	position
}

func (bo BitOr) Name() string    { return "'|'" }
func (bo BitOr) Precedence() int { return 4 }

// ^
type Xor struct {
	position
}

func (x Xor) Name() string    { return "'^'" }
func (x Xor) Precedence() int { return 4 }

// (
type LeftParen struct {
	position
}

func (lp LeftParen) Name() string { return "'('" }

// )
type RightParen struct {
	position
}

func (lp RightParen) Name() string { return "')'" }

// [
type LeftBracket struct {
	position
}

func (lb LeftBracket) Name() string { return "'['" }

// ]
type RightBracket struct {
	position
}

func (rb RightBracket) Name() string { return "']'" }

// ,
type Comma struct {
	position
}

func (c Comma) Name() string { return "','" }

// ;
type Semicolon struct {
	position
}

func (s Semicolon) Name() string { return "';'" }

type Const struct {
	position
}

func (c Const) Name() string { return "'const'" }

type Import struct {
	position
}

func (i Import) Name() string { return "'import'" }

type Type struct {
	position
}

func (t Type) Name() string { return "'type'" }

type Block struct {
	position
}

func (b Block) Name() string { return "'block'" }

func (b Block) functionality() {}

type Bus struct {
	position
}

func (b Bus) Name() string { return "'bus'" }

func (b Bus) functionality() {}

type Config struct {
	position
}

func (c Config) Name() string   { return "'config'" }
func (c Config) functionality() {}

type Irq struct {
	position
}

func (i Irq) Name() string   { return "'irq'" }
func (i Irq) functionality() {}

type Mask struct {
	position
}

func (m Mask) Name() string   { return "'mask'" }
func (m Mask) functionality() {}

type Memory struct {
	position
}

func (m Memory) Name() string   { return "'memory'" }
func (m Memory) functionality() {}

type Param struct {
	position
}

func (p Param) Name() string   { return "'param'" }
func (p Param) functionality() {}

type Proc struct {
	position
}

func (p Proc) Name() string   { return "'proc'" }
func (p Proc) functionality() {}

type Return struct {
	position
}

func (r Return) Name() string   { return "'return'" }
func (r Return) functionality() {}

type Static struct {
	position
}

func (s Static) Name() string   { return "'static'" }
func (s Static) functionality() {}

type Status struct {
	position
}

func (s Status) Name() string   { return "'status'" }
func (s Status) functionality() {}

type Stream struct {
	position
}

func (s Stream) Name() string   { return "'stream'" }
func (s Stream) functionality() {}

type Access struct {
	position
}

func (a Access) Name() string { return "'access'" }
func (a Access) property()    {}

type AddEnable struct {
	position
}

func (ae AddEnable) Name() string { return "'add-enable'" }
func (ae AddEnable) property()    {}

type Atomic struct {
	position
}

func (a Atomic) Name() string { return "'atomic'" }
func (a Atomic) property()    {}

type ByteWriteEnable struct {
	position
}

func (bwe ByteWriteEnable) Name() string { return "'byte-write-enable'" }
func (bwe ByteWriteEnable) property()    {}

type Clear struct {
	position
}

func (c Clear) Name() string { return "'clear'" }
func (c Clear) property()    {}

type Delay struct {
	position
}

func (d Delay) Name() string { return "'delay'" }
func (d Delay) property()    {}

type EnableInitValue struct {
	position
}

func (eiv EnableInitValue) Name() string { return "'enable-init-value'" }
func (eiv EnableInitValue) property()    {}

type EnableResetValue struct {
	position
}

func (erv EnableResetValue) Name() string { return "'enable-reset-value'" }
func (erv EnableResetValue) property()    {}

type Groups struct {
	position
}

func (g Groups) Name() string { return "'groups'" }
func (g Groups) property()    {}

type InitValue struct {
	position
}

func (iv InitValue) Name() string { return "'init-value'" }
func (iv InitValue) property()    {}

type InTrigger struct {
	position
}

func (it InTrigger) Name() string { return "'in-trigger'" }
func (it InTrigger) property()    {}

type Masters struct {
	position
}

func (m Masters) Name() string { return "'masters'" }
func (m Masters) property()    {}

type OutTrigger struct {
	position
}

func (ot OutTrigger) Name() string { return "'out-trigger'" }
func (ot OutTrigger) property()    {}

type Range struct {
	position
}

func (r Range) Name() string { return "'range'" }
func (r Range) property()    {}

type ReadLatency struct {
	position
}

func (rl ReadLatency) Name() string { return "'read-latency'" }
func (r ReadLatency) property()     {}

type ReadValue struct {
	position
}

func (rv ReadValue) Name() string { return "'read-value'" }
func (rv ReadValue) property()    {}

type Reset struct {
	position
}

func (r Reset) Name() string { return "'reset'" }
func (r Reset) property()    {}

type ResetValue struct {
	position
}

func (rv ResetValue) Name() string { return "'reset-value'" }
func (rv ResetValue) property()    {}

type Size struct {
	position
}

func (s Size) Name() string { return "'size'" }
func (s Size) property()    {}

type Width struct {
	position
}

func (w Width) Name() string { return "'width'" }
func (w Width) property()    {}

// . - currently unused
type Period struct {
	position
}

func (p Period) Name() string { return "'.'" }

// : - currently unused
type Colon struct {
	position
}

func (c Colon) Name() string { return "':'" }

// { - currently unused
type LeftBrace struct {
	position
}

func (lb LeftBrace) Name() string { return "'{'" }

// } - currently unused
type RightBrace struct {
	position
}

func (rb RightBrace) Name() string { return "'}'" }
