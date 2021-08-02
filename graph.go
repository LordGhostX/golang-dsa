package main

type Graph struct {
	edges map[string]map[string]interface{}
}

func (g *Graph) Add(from, to string) {
	if len(g.edges) == 0 {
		g.edges = make(map[string]map[string]interface{})
	}

	_, ok := g.edges[from]
	if !ok {
		g.edges[from] = make(map[string]interface{})
	}
	g.edges[from][to] = nil
}

func (g *Graph) Remove(v string) {
	delete(g.edges, v)
	for _, k := range g.edges {
		if _, ok := k[v]; ok {
			delete(k, v)
		}
	}
}

func (g *Graph) Lookup(from, to string) bool {
	_, ok := g.edges[from][to]
	return ok
}
