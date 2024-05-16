// Package tok implements Functional Bus Description Language tokens.
package tok

import "fmt"

// Various token types
type (
	Token interface {
		Start() int
		End() int
		Line() int
		Column() int
		Src() []byte
		Path() string
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

type (
	None      struct{ position }
	Comment   struct{ position }
	Indent    struct{ position } // Indent increment
	Dedent    struct{ position } // Indent decrement
	Newline   struct{ position }
	Eof       struct{ position } // End of file
	Ident     struct{ position } // Identifier
	QualIdent struct{ position } // Qualified Identifier
	Bool      struct{ position }
	Int       struct{ position }
	Real      struct{ position }
	String    struct{ position }
	BitString struct{ position }
	Time      struct{ position }
	Neg       struct{ position } // !
	Ass       struct{ position } // =
	Add       struct{ position } // +
	Sub       struct{ position } // -
	Mul       struct{ position } // *
	Div       struct{ position } // /
	Rem       struct{ position } // %
	Exp       struct{ position } // **
	Eq        struct{ position } // ==
	Neq       struct{ position } // !=
	Less      struct{ position } // <
	LessEq    struct{ position } // <=
	Greater   struct{ position } // >
	GreaterEq struct{ position } // >=
	And       struct{ position } // &&
	Or        struct{ position } // ||
	LShift    struct{ position } // <<
	RShift    struct{ position } // >>
	BitAnd    struct{ position } // &
	BitOr     struct{ position } // |
	Xor       struct{ position } // ^
	LParen    struct{ position } // (
	RParen    struct{ position } // )
	LBracket  struct{ position } // [
	RBracket  struct{ position } // ]
	Comma     struct{ position } // ,
	Semicolon struct{ position } // ;
	Colon     struct{ position } // :
	// Keyword tokens
	Const  struct{ position }
	Import struct{ position }
	Type   struct{ position }
	// Functionality tokens
	Block  struct{ position }
	Bus    struct{ position }
	Config struct{ position }
	Irq    struct{ position }
	Mask   struct{ position }
	Memory struct{ position }
	Param  struct{ position }
	Proc   struct{ position }
	Return struct{ position }
	Static struct{ position }
	Status struct{ position }
	Stream struct{ position }
	// Property tokens
	Access           struct{ position }
	AddEnable        struct{ position }
	Atomic           struct{ position }
	ByteWriteEnable  struct{ position }
	Clear            struct{ position }
	Delay            struct{ position }
	EnableInitValue  struct{ position }
	EnableResetValue struct{ position }
	Groups           struct{ position }
	InitValue        struct{ position }
	InTrigger        struct{ position }
	Masters          struct{ position }
	OutTrigger       struct{ position }
	Range            struct{ position }
	ReadLatency      struct{ position }
	ReadValue        struct{ position }
	Reset            struct{ position }
	ResetValue       struct{ position }
	Size             struct{ position }
	Width            struct{ position }
	// Currently unused tokens
	Period struct{ position } // .
	LBrace struct{ position } // {
	RBrace struct{ position } // }
)

func (n None) Name() string { return "" }

func (c Comment) Name() string { return "comment" }

func (i Indent) Name() string { return "indent increment" }

func (d Dedent) Name() string { return "indent decrement" }

func (n Newline) Name() string { return "newline" }

func (e Eof) Name() string { return "end of file" }

func (i Ident) Name() string { return "identifier" }

func (qi QualIdent) Name() string { return "qualified identifier" }

func (b Bool) Name() string { return "bool" }

func (i Int) Name() string { return "integer" }
func (i Int) number()      {}

func (r Real) Name() string { return "real" }

func (r Real) number() {}

func (s String) Name() string { return "string" }

func (bs BitString) Name() string { return "bit string" }

func (t Time) Name() string { return "time" }

func (n Neg) Name() string { return "!" }

func (a Ass) Name() string { return "'='" }

func (a Add) Name() string    { return "'+'" }
func (a Add) Precedence() int { return 4 }

func (s Sub) Name() string    { return "'-'" }
func (s Sub) Precedence() int { return 4 }

func (m Mul) Name() string    { return "'*'" }
func (m Mul) Precedence() int { return 5 }

func (d Div) Name() string    { return "'/'" }
func (d Div) Precedence() int { return 5 }

func (r Rem) Name() string    { return "'%'" }
func (r Rem) Precedence() int { return 5 }

func (e Exp) Name() string    { return "'**'" }
func (e Exp) Precedence() int { return 6 }

func (e Eq) Name() string    { return "'=='" }
func (e Eq) Precedence() int { return 3 }

func (n Neq) Name() string    { return "'!='" }
func (n Neq) Precedence() int { return 3 }

func (l Less) Name() string    { return "'<'" }
func (l Less) Precedence() int { return 3 }

func (le LessEq) Name() string    { return "'<='" }
func (le LessEq) Precedence() int { return 3 }

func (g Greater) Name() string    { return "'>'" }
func (g Greater) Precedence() int { return 3 }

func (ge GreaterEq) Name() string    { return "'>='" }
func (ge GreaterEq) Precedence() int { return 3 }

func (a And) Name() string    { return "'&&'" }
func (a And) Precedence() int { return 2 }

func (o Or) Name() string    { return "'||'" }
func (o Or) Precedence() int { return 1 }

func (ls LShift) Name() string    { return "'<<'" }
func (ls LShift) Precedence() int { return 5 }

func (ls RShift) Name() string    { return "'>>'" }
func (ls RShift) Precedence() int { return 5 }

func (ba BitAnd) Name() string    { return "'&'" }
func (ba BitAnd) Precedence() int { return 5 }

func (bo BitOr) Name() string    { return "'|'" }
func (bo BitOr) Precedence() int { return 4 }

func (x Xor) Name() string    { return "'^'" }
func (x Xor) Precedence() int { return 4 }

func (lp LParen) Name() string { return "'('" }

func (lp RParen) Name() string { return "')'" }

func (lb LBracket) Name() string { return "'['" }

func (rb RBracket) Name() string { return "']'" }

func (c Comma) Name() string { return "','" }

func (s Semicolon) Name() string { return "';'" }

func (c Colon) Name() string    { return "':'" }
func (c Colon) Precedence() int { return 0 }

func (c Const) Name() string { return "'const'" }

func (i Import) Name() string { return "'import'" }

func (t Type) Name() string { return "'type'" }

func (b Block) Name() string { return "'block'" }

func (b Block) functionality() {}

func (b Bus) Name() string { return "'bus'" }

func (b Bus) functionality() {}

func (c Config) Name() string   { return "'config'" }
func (c Config) functionality() {}

func (i Irq) Name() string   { return "'irq'" }
func (i Irq) functionality() {}

func (m Mask) Name() string   { return "'mask'" }
func (m Mask) functionality() {}

func (m Memory) Name() string   { return "'memory'" }
func (m Memory) functionality() {}

func (p Param) Name() string   { return "'param'" }
func (p Param) functionality() {}

func (p Proc) Name() string   { return "'proc'" }
func (p Proc) functionality() {}

func (r Return) Name() string   { return "'return'" }
func (r Return) functionality() {}

func (s Static) Name() string   { return "'static'" }
func (s Static) functionality() {}

func (s Status) Name() string   { return "'status'" }
func (s Status) functionality() {}

func (s Stream) Name() string   { return "'stream'" }
func (s Stream) functionality() {}

func (a Access) Name() string { return "'access'" }
func (a Access) property()    {}

func (ae AddEnable) Name() string { return "'add-enable'" }
func (ae AddEnable) property()    {}

func (a Atomic) Name() string { return "'atomic'" }
func (a Atomic) property()    {}

func (bwe ByteWriteEnable) Name() string { return "'byte-write-enable'" }
func (bwe ByteWriteEnable) property()    {}

func (c Clear) Name() string { return "'clear'" }
func (c Clear) property()    {}

func (d Delay) Name() string { return "'delay'" }
func (d Delay) property()    {}

func (eiv EnableInitValue) Name() string { return "'enable-init-value'" }
func (eiv EnableInitValue) property()    {}

func (erv EnableResetValue) Name() string { return "'enable-reset-value'" }
func (erv EnableResetValue) property()    {}

func (g Groups) Name() string { return "'groups'" }
func (g Groups) property()    {}

func (iv InitValue) Name() string { return "'init-value'" }
func (iv InitValue) property()    {}

func (it InTrigger) Name() string { return "'in-trigger'" }
func (it InTrigger) property()    {}

func (m Masters) Name() string { return "'masters'" }
func (m Masters) property()    {}

func (ot OutTrigger) Name() string { return "'out-trigger'" }
func (ot OutTrigger) property()    {}

func (r Range) Name() string { return "'range'" }
func (r Range) property()    {}

func (rl ReadLatency) Name() string { return "'read-latency'" }
func (r ReadLatency) property()     {}

func (rv ReadValue) Name() string { return "'read-value'" }
func (rv ReadValue) property()    {}

func (r Reset) Name() string { return "'reset'" }
func (r Reset) property()    {}

func (rv ResetValue) Name() string { return "'reset-value'" }
func (rv ResetValue) property()    {}

func (s Size) Name() string { return "'size'" }
func (s Size) property()    {}

func (w Width) Name() string { return "'width'" }
func (w Width) property()    {}

func (p Period) Name() string { return "'.'" }

func (lb LBrace) Name() string { return "'{'" }

func (rb RBrace) Name() string { return "'}'" }
