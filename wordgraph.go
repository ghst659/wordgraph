// Package wordgraph provides word-ladder lookup capability.
package wordgraph

// wordSet is a shorthand for a set of strings.
type wordSet map[string]struct{}

// Graph is a struct representing a word-ladder graph.
type Graph struct {
	adjacents map[string]wordSet
}

// NewGraph returns a pointer to a new word graph.
func NewGraph() *Graph {
	return &Graph{
		adjacents: make(map[string]wordSet),
	}
}

// dirEdge creates a directional edge from X to Y.
func (g *Graph) dirEdge(x, y string) {
	if _, ok := g.adjacents[x]; !ok {
		g.adjacents[x] = make(wordSet)
	}
	g.adjacents[x][y] = struct{}{}
}

// AddEdge creates a non-directional edge between two words.
func (g *Graph) AddEdge(a, b string) {
	if a != b {
		g.dirEdge(a, b)
		g.dirEdge(b, a)
	}
}

// Neighbours returns a slice of words adjacent to a given word.
func (g *Graph) Neighbours(x string) (n []string) {
	if s, ok := g.adjacents[x]; ok {
		for k := range s {
			n = append(n, k)
		}
	}
	return
}
