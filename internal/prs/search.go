package prs

type Searchable interface {
	GetSymbol(name string) (Symbol, error)
}
