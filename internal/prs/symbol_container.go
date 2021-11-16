package prs

type SymbolContainer []Symbol

// Add adds symbol to SymbolContainer.
// If symbol with given name already exists it returns false.
// If the operation is successful it returns true.
func (sc *SymbolContainer) Add(sym Symbol) bool {
	for _, s := range *sc {
		if s.Name() == sym.Name() {
			return false
		}
	}

	*sc = append(*sc, sym)

	return true
}

func (sc *SymbolContainer) Get(name string) (Symbol, bool) {
	for _, s := range *sc {
		if s.Name() == name {
			return s, true
		}
	}

	return nil, false
}
