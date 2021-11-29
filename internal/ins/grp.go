package ins

type Group struct {
	Name     string
	Elements []*Element
}

// IsStatus returns true if group contains only status elements.
func (g *Group) IsStatus() bool {
	for _, e := range g.Elements {
		if e.BaseType != "status" {
			return false
		}
	}

	return true
}

// IsConfig returns true if group contains only config elements.
func (g *Group) IsConfig() bool {
	for _, e := range g.Elements {
		if e.BaseType != "config" {
			return false
		}
	}

	return true
}

type grpNode map[string]bool

type grpGraph struct {
	startNodes []string
	nodes      map[string]grpNode
}

func newGrpGraph() *grpGraph {
	return &grpGraph{
		startNodes: []string{},
		nodes:      make(map[string]grpNode),
	}
}

func (g *grpGraph) addEdge(from, to string) {
	if from == "" {
		exist := false
		for _, sn := range g.startNodes {
			if to == sn {
				exist = true
				break
			}
		}
		if !exist {
			g.startNodes = append(g.startNodes, to)
		}
	} else {
		if _, ok := g.nodes[from]; !ok {
			g.nodes[from] = make(grpNode)
		}
	}

	if _, ok := g.nodes[to]; !ok {
		g.nodes[to] = make(grpNode)
	}

	if from != "" {
		g.nodes[from][to] = true
	}
}

func (g *grpGraph) topSort(node string, sorted []string, visited map[string]bool) []string {
	visited[node] = true

	for child, _ := range g.nodes[node] {
		if visited[child] == true {
			continue
		}

		sorted = g.topSort(child, sorted, visited)
	}

	// After the children are visited push into the sorted stack.
	sorted = append(sorted, node)

	return sorted
}

func (g *grpGraph) sort() []string {
	visited := map[string]bool{}

	sorted := []string{}

	for i := len(g.startNodes) - 1; i >= 0; i-- {
		node := g.startNodes[i]

		if visited[node] {
			continue
		}

		sorted = g.topSort(node, sorted, visited)
	}

	// Reverse
	for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
		sorted[i], sorted[j] = sorted[j], sorted[i]
	}

	return sorted
}
