package prs

// Arg represents argument in the argument list.
type Arg struct {
	HasName bool
	Name    string
	Value   Expr
}

type Property struct {
	LineNumber uint32
	Value      Expr
}

type Element interface {
	Searchable
	Symbol
	Type() string
	Args() []Arg
	Params() []Param
	SetResolvedArgs(args map[string]Expr)
	ResolvedArgs() map[string]Expr
	Properties() map[string]Property
	Symbols() SymbolContainer
}
