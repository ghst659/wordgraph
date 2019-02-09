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

// Clear clears all edge information in the Graph.
func (g *Graph) Clear() {
	for k := range g.adjacents {
		delete(g.adjacents, k)
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

// BuildFromList populates the graph from the list of words.   The first
// boolean argument, if true, indicates that single insertions/deletions
// are to be considered adjacency edits, in addition to substitutions.
func (g *Graph) BuildFromList(insdel bool, words []string) (err error) {
	buckets := make(map[string][]string)
	for _, w := range words {
		for p := range w {
			pre := w[:p]
			suf := w[p+1:]
			kSub := pre + " " + suf
			buckets[kSub] = append(buckets[kSub], w)
			if insdel {
				aft := w[p:]
				kInsDel := pre + " " + aft
				buckets[kInsDel] = append(buckets[kInsDel], w)
			}
		}
		if insdel {
			kInsDel := w + " "
			buckets[kInsDel] = append(buckets[kInsDel], w)
		}
	}
	for _, wList := range buckets {
		if len(wList) > 1 {
			for _, a := range wList {
				for _, b := range wList {
					g.AddEdge(a, b)
				}
			}
		}
	}
	return
}
