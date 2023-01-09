package elem

type Proc struct {
	Elem

	Params  []*Param
	Returns []*Return

	CallAddr int64
	ExitAddr int64
}

func (p *Proc) ParamsBufSize() int64 {
	params := p.Params
	l := len(params)

	if l == 0 {
		return 0
	}

	return params[l-1].Access.EndAddr() - params[0].Access.StartAddr() + 1
}

func (p *Proc) ParamsStartAddr() int64 {
	if len(p.Params) == 0 {
		return p.CallAddr
	}

	return p.Params[0].Access.StartAddr()
}
