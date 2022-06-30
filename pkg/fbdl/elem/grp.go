package elem

type Group interface {
	Name() string
	Statuses() []*Status
}
