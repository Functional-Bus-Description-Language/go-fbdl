// Package tok implements Functional Bus Description Language tokens.
package token

// Various token types
type (
	Token interface {
		Start() int
		End() int
		Name() string
		Text(src []byte) string
	}

	Functionality interface {
		Token
		// Transforms functionality token into identifier token
		Ident() Ident
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

	// Properties being a single word
	WordProperty interface {
		Property
		// Transforms functionality token into identifier token
		Ident() Ident
	}

	// Properties containing at least one hyphen
	HyphenProperty interface {
		Property
		// Splits the hyphen property input multiple identifiers and subtractions
		Split() []Token
	}
)

// New constructs a token with the source span [start, end).
// It is useful mostly for tests in ast package.
func New[T interface {
	Token
	~struct{ span }
}](start, end int) T {
	return T{span{start, end}}
}

type (
	None      struct{ span }
	Comment   struct{ span }
	Indent    struct{ span } // Indent increment
	Dedent    struct{ span } // Indent decrement
	Newline   struct{ span }
	Eof       struct{ span } // End of file
	Ident     struct{ span } // Identifier
	QualIdent struct{ span } // Qualified Identifier
	Bool      struct{ span }
	Int       struct{ span }
	Float     struct{ span }
	String    struct{ span }
	BitString struct{ span }
	Time      struct{ span }
	Neg       struct{ span } // !
	Ass       struct{ span } // =
	Add       struct{ span } // +
	Sub       struct{ span } // -
	Mul       struct{ span } // *
	Div       struct{ span } // /
	Rem       struct{ span } // %
	Exp       struct{ span } // **
	Eq        struct{ span } // ==
	Neq       struct{ span } // !=
	Less      struct{ span } // <
	LessEq    struct{ span } // <=
	Greater   struct{ span } // >
	GreaterEq struct{ span } // >=
	And       struct{ span } // &&
	Or        struct{ span } // ||
	LShift    struct{ span } // <<
	RShift    struct{ span } // >>
	BitAnd    struct{ span } // &
	BitOr     struct{ span } // |
	Xor       struct{ span } // ^
	LParen    struct{ span } // (
	RParen    struct{ span } // )
	LBracket  struct{ span } // [
	RBracket  struct{ span } // ]
	Comma     struct{ span } // ,
	Semicolon struct{ span } // ;
	Colon     struct{ span } // :
	// Keyword tokens
	Const  struct{ span }
	Import struct{ span }
	Type   struct{ span }
	// Functionality tokens
	Blackbox struct{ span }
	Block    struct{ span }
	Bus      struct{ span }
	Config   struct{ span }
	Group    struct{ span }
	Irq      struct{ span }
	Mask     struct{ span }
	Param    struct{ span }
	Proc     struct{ span }
	Return   struct{ span }
	Static   struct{ span }
	Status   struct{ span }
	Stream   struct{ span }
	// Property tokens
	Access           struct{ span }
	AddEnable        struct{ span }
	Align            struct{ span }
	Atomic           struct{ span }
	ByteWriteEnable  struct{ span }
	Clear            struct{ span }
	Delay            struct{ span }
	EnableInitValue  struct{ span }
	EnableResetValue struct{ span }
	InitValue        struct{ span }
	InTrigger        struct{ span }
	Masters          struct{ span }
	OutTrigger       struct{ span }
	Range            struct{ span }
	ReadLatency      struct{ span }
	ReadValue        struct{ span }
	Reset            struct{ span }
	ResetValue       struct{ span }
	Size             struct{ span }
	Virtual          struct{ span }
	Width            struct{ span }
	// Currently unused tokens
	Period struct{ span } // .
	LBrace struct{ span } // {
	RBrace struct{ span } // }
)

func (n None) Name() string { return "" }

func (c Comment) Name() string { return "comment" }

func (i Indent) Name() string { return "indent increment" }

func (d Dedent) Name() string { return "indent decrement" }

func (n Newline) Name() string { return "newline" }

func (e Eof) Name() string { return "end of file" }

func (e Eof) Text(src []byte) string {
	panic("calling Text() on Eof indicates logic error")
}

func (i Ident) Name() string { return "identifier" }

func (qi QualIdent) Name() string { return "qualified identifier" }

func (b Bool) Name() string { return "bool" }

func (i Int) Name() string { return "integer" }
func (i Int) number()      {}

func (f Float) Name() string { return "float" }

func (f Float) number() {}

func (s String) Name() string { return "string" }

func (bs BitString) Name() string { return "bit string" }

func (t Time) Name() string { return "time" }

func (n Neg) Name() string    { return "'!'" }
func (n Neg) Precedence() int { return 7 }

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

func (bb Blackbox) Name() string   { return "'blackbox'" }
func (bb Blackbox) functionality() {}
func (bb Blackbox) Ident() Ident   { return Ident(bb) }

func (b Block) Name() string   { return "'block'" }
func (b Block) functionality() {}
func (b Block) Ident() Ident   { return Ident(b) }

func (b Bus) Name() string   { return "'bus'" }
func (b Bus) functionality() {}
func (b Bus) Ident() Ident   { return Ident(b) }

func (c Config) Name() string   { return "'config'" }
func (c Config) functionality() {}
func (c Config) Ident() Ident   { return Ident(c) }

func (g Group) Name() string   { return "'group'" }
func (g Group) functionality() {}
func (g Group) Ident() Ident   { return Ident(g) }

func (i Irq) Name() string   { return "'irq'" }
func (i Irq) functionality() {}
func (i Irq) Ident() Ident   { return Ident(i) }

func (m Mask) Name() string   { return "'mask'" }
func (m Mask) functionality() {}
func (m Mask) Ident() Ident   { return Ident(m) }

func (p Param) Name() string   { return "'param'" }
func (p Param) functionality() {}
func (p Param) Ident() Ident   { return Ident(p) }

func (p Proc) Name() string   { return "'proc'" }
func (p Proc) functionality() {}
func (p Proc) Ident() Ident   { return Ident(p) }

func (r Return) Name() string   { return "'return'" }
func (r Return) functionality() {}
func (r Return) Ident() Ident   { return Ident(r) }

func (s Static) Name() string   { return "'static'" }
func (s Static) functionality() {}
func (s Static) Ident() Ident   { return Ident(s) }

func (s Status) Name() string   { return "'status'" }
func (s Status) functionality() {}
func (s Status) Ident() Ident   { return Ident(s) }

func (s Stream) Name() string   { return "'stream'" }
func (s Stream) functionality() {}
func (s Stream) Ident() Ident   { return Ident(s) }

func (a Access) Name() string { return "'access'" }
func (a Access) property()    {}
func (a Access) Ident() Ident { return Ident(a) }

func (ae AddEnable) Name() string { return "'add-enable'" }
func (ae AddEnable) property()    {}

func (ae AddEnable) Split() []Token {
	return []Token{
		Ident{span{ae.start, ae.start + 3}},
		Sub{span{ae.start + 3, ae.start + 4}},
		Ident{span{ae.start + 4, ae.end}},
	}
}

func (a Align) Name() string { return "'align'" }
func (a Align) property()    {}
func (a Align) Ident() Ident { return Ident(a) }

func (a Atomic) Name() string { return "'atomic'" }
func (a Atomic) property()    {}
func (a Atomic) Ident() Ident { return Ident(a) }

func (bwe ByteWriteEnable) Name() string { return "'byte-write-enable'" }
func (bwe ByteWriteEnable) property()    {}

func (bwe ByteWriteEnable) Split() []Token {
	return []Token{
		Ident{span{bwe.start, bwe.start + 4}},
		Sub{span{bwe.start + 4, bwe.start + 5}},
		Ident{span{bwe.start + 5, bwe.start + 10}},
		Sub{span{bwe.start + 10, bwe.start + 11}},
		Ident{span{bwe.start + 11, bwe.end}},
	}
}

func (c Clear) Name() string { return "'clear'" }
func (c Clear) property()    {}
func (c Clear) Ident() Ident { return Ident(c) }

func (d Delay) Name() string { return "'delay'" }
func (d Delay) property()    {}
func (d Delay) Ident() Ident { return Ident(d) }

func (eiv EnableInitValue) Name() string { return "'enable-init-value'" }
func (eiv EnableInitValue) property()    {}

func (eiv EnableInitValue) Split() []Token {
	return []Token{
		Ident{span{eiv.start, eiv.start + 6}},
		Sub{span{eiv.start + 6, eiv.start + 7}},
		Ident{span{eiv.start + 7, eiv.start + 11}},
		Sub{span{eiv.start + 11, eiv.start + 12}},
		Ident{span{eiv.start + 12, eiv.end}},
	}
}

func (erv EnableResetValue) Name() string { return "'enable-reset-value'" }
func (erv EnableResetValue) property()    {}

func (erv EnableResetValue) Split() []Token {
	return []Token{
		Ident{span{erv.start, erv.start + 6}},
		Sub{span{erv.start + 6, erv.start + 7}},
		Ident{span{erv.start + 7, erv.start + 12}},
		Sub{span{erv.start + 12, erv.start + 13}},
		Ident{span{erv.start + 13, erv.end}},
	}
}

func (iv InitValue) Name() string { return "'init-value'" }
func (iv InitValue) property()    {}

func (iv InitValue) Split() []Token {
	return []Token{
		Ident{span{iv.start, iv.start + 4}},
		Sub{span{iv.start + 4, iv.start + 5}},
		Ident{span{iv.start + 5, iv.end}},
	}
}

func (it InTrigger) Name() string { return "'in-trigger'" }
func (it InTrigger) property()    {}

func (it InTrigger) Split() []Token {
	return []Token{
		Ident{span{it.start, it.start + 2}},
		Sub{span{it.start + 2, it.start + 3}},
		Ident{span{it.start + 3, it.end}},
	}
}

func (m Masters) Name() string { return "'masters'" }
func (m Masters) property()    {}
func (m Masters) Ident() Ident { return Ident(m) }

func (ot OutTrigger) Name() string { return "'out-trigger'" }
func (ot OutTrigger) property()    {}

func (ot OutTrigger) Split() []Token {
	return []Token{
		Ident{span{ot.start, ot.start + 3}},
		Sub{span{ot.start + 3, ot.start + 4}},
		Ident{span{ot.start + 4, ot.end}},
	}
}

func (r Range) Name() string { return "'range'" }
func (r Range) property()    {}
func (r Range) Ident() Ident { return Ident(r) }

func (rl ReadLatency) Name() string { return "'read-latency'" }
func (r ReadLatency) property()     {}

func (r ReadLatency) Split() []Token {
	return []Token{
		Ident{span{r.start, r.start + 4}},
		Sub{span{r.start + 4, r.start + 5}},
		Ident{span{r.start + 5, r.end}},
	}
}

func (rv ReadValue) Name() string { return "'read-value'" }
func (rv ReadValue) property()    {}

func (rv ReadValue) Split() []Token {
	return []Token{
		Ident{span{rv.start, rv.start + 4}},
		Sub{span{rv.start + 4, rv.start + 5}},
		Ident{span{rv.start + 5, rv.end}},
	}
}

func (r Reset) Name() string { return "'reset'" }
func (r Reset) property()    {}
func (r Reset) Ident() Ident { return Ident(r) }

func (rv ResetValue) Name() string { return "'reset-value'" }
func (rv ResetValue) property()    {}

func (rv ResetValue) Split() []Token {
	return []Token{
		Ident{span{rv.start, rv.start + 5}},
		Sub{span{rv.start + 5, rv.start + 6}},
		Ident{span{rv.start + 6, rv.end}},
	}
}

func (s Size) Name() string { return "'size'" }
func (s Size) property()    {}
func (s Size) Ident() Ident { return Ident(s) }

func (v Virtual) Name() string { return "'virtual'" }
func (v Virtual) property()    {}
func (v Virtual) Ident() Ident { return Ident(v) }

func (w Width) Name() string { return "'width'" }
func (w Width) property()    {}
func (w Width) Ident() Ident { return Ident(w) }

func (p Period) Name() string { return "'.'" }

func (lb LBrace) Name() string { return "'{'" }

func (rb RBrace) Name() string { return "'}'" }
