package elem

type Group interface {
	Name() string
	Statuses() []*Status
}

type Groupable interface {
	GroupNames() []string
}
