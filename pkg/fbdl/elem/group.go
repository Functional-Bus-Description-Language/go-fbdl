package elem

type Groupable interface {
	Element
	GroupNames() []string
}

type GroupHolder interface {
	GroupedElems() []Groupable
}
