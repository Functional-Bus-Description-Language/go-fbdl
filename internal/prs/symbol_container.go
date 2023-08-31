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

func (sc *SymbolContainer) GetConst(name string) (*Const, bool) {
	for _, s := range *sc {
		if s.Name() == name && s.Kind() == ConstDef {
			return s.(*Const), true
		}
	}

	return nil, false
}

func (sc *SymbolContainer) GetInst(name string) (*Inst, bool) {
	for _, s := range *sc {
		if s.Name() == name && s.Kind() == FuncInst {
			return s.(*Inst), true
		}
	}

	return nil, false
}

func (sc *SymbolContainer) GetType(name string) (*Type, bool) {
	for _, s := range *sc {
		if s.Name() == name && s.Kind() == TypeDef {
			return s.(*Type), true
		}
	}

	return nil, false
}

func (sc *SymbolContainer) GetByName(name string) (Symbol, bool) {
	for _, s := range *sc {
		if s.Name() == name {
			return s, true
		}
	}

	return nil, false
}
