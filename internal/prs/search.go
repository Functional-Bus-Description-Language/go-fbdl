package prs

type Searchable interface {
	GetSymbol(s string) (Symbol, error)
}
