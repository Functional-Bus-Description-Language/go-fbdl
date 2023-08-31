package prs

type Scope interface {
	GetConst(name string) (*Const, error)
	GetInst(name string) (*Inst, error)
	GetType(name string) (*Type, error)
}
