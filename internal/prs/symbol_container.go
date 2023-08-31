package prs

type SymbolContainer struct {
	Consts []*Const
	Insts  []*Inst
	Types  []*Type
}

func (sc *SymbolContainer) addConst(cnst *Const) bool {
	for _, c := range sc.Consts {
		if c.name == cnst.name {
			return false
		}
	}
	sc.Consts = append(sc.Consts, cnst)
	return true
}

func (sc *SymbolContainer) addInst(ins *Inst) bool {
	for _, i := range sc.Insts {
		if i.name == ins.name {
			return false
		}
	}
	sc.Insts = append(sc.Insts, ins)
	return true
}

func (sc *SymbolContainer) addType(typ *Type) bool {
	for _, t := range sc.Types {
		if t.name == typ.name {
			return false
		}
	}
	sc.Types = append(sc.Types, typ)
	return true
}

func (sc SymbolContainer) GetConst(name string) (*Const, bool) {
	for _, c := range sc.Consts {
		if c.name == name {
			return c, true
		}
	}

	return nil, false
}

func (sc SymbolContainer) GetInst(name string) (*Inst, bool) {
	for _, i := range sc.Insts {
		if i.Name() == name {
			return i, true
		}
	}

	return nil, false
}

func (sc SymbolContainer) GetType(name string) (*Type, bool) {
	for _, t := range sc.Types {
		if t.name == name {
			return t, true
		}
	}

	return nil, false
}

func (sc SymbolContainer) Symbols() []Symbol {
	syms := []Symbol{}

	for _, s := range sc.Consts {
		syms = append(syms, s)
	}
	for _, s := range sc.Insts {
		syms = append(syms, s)
	}
	for _, s := range sc.Types {
		syms = append(syms, s)
	}

	return syms
}
