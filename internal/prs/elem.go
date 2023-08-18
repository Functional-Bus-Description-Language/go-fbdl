package prs

// Element is common interface for Inst and Type structs.
// Type is actually an element, but not instantiated.
type Element interface {
	Searchable
	Symbol
	IsArray() bool
	Count() Expr
	Type() string
	Args() []Arg
	Params() []Param
	SetResolvedArgs(args map[string]Expr)
	ResolvedArgs() map[string]Expr
	Props() PropContainer
	Symbols() SymbolContainer
}
