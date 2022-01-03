package prs

type Searchable interface {
	GetSymbol(name string, kind SymbolKind) (Symbol, error)
}
