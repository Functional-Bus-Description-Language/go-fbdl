package prs

type Scope interface {
	GetSymbol(name string, kind SymbolKind) (Symbol, error)
}
