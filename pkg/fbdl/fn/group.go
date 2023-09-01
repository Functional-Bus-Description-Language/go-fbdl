package fn

type Groupable interface {
	Functionality
	GroupNames() []string
}

type GroupHolder interface {
	GroupedInsts() []Groupable
}
