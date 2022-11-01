package elem

type Func struct {
	Elem

	Params  []*Param
	Returns []*Return

	StbAddr int64
	AckAddr int64
}

func (f *Func) ParamsBufSize() int64 {
	params := f.Params
	l := len(params)

	if l == 0 {
		return 0
	}

	return params[l-1].Access.EndAddr() - params[0].Access.StartAddr() + 1
}

func (f *Func) ParamsStartAddr() int64 {
	if len(f.Params) == 0 {
		return f.StbAddr
	}

	return f.Params[0].Access.StartAddr()
}
