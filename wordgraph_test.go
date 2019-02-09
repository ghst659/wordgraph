package wordgraph

import (
	"testing"
)

func gotIsWant(got, want []string) bool {
	wantFound := make(map[string]bool)
	for _, w := range want {
		wantFound[w] = false
	}
	for _, g := range got {
		seen, ok := wantFound[g]
		if !ok {
			return false
		}
		if seen {
			return false
		}
		wantFound[g] = true
	}
	for _, seen := range wantFound {
		if !seen {
			return false
		}
	}
	return true
}

func TestGraphEdgesNeighbours(t *testing.T) {
	g := NewGraph()
	g.AddEdge("a", "b")
	g.AddEdge("a", "c")
	g.AddEdge("b", "d")
	g.AddEdge("c", "e")
	na := g.Neighbours("a")
	nb := g.Neighbours("b")
	ne := g.Neighbours("e")
	wantA := []string{"b", "c"}
	wantB := []string{"a", "d"}
	wantE := []string{"c"}
	if !gotIsWant(na, wantA) {
		t.Errorf("a neighbour mismatch: %v vs %v", na, wantA)
	}
	if !gotIsWant(nb, wantB) {
		t.Errorf("b neighbour mismatch: %v vs %v", nb, wantB)
	}
	if !gotIsWant(ne, wantE) {
		t.Errorf("e neighbour mismatch: %v vs %v", ne, wantE)
	}
}

func TestBuildFromListNoInsertions(t *testing.T) {
	g := NewGraph()
	g.BuildFromList(false, []string{
		"cat", "hat", "car", "cur", "cut", "cart",
	})
	got := g.Neighbours("cat")
	want := []string{
		"hat", "car", "cut",
	}
	if !gotIsWant(got, want) {
		t.Errorf("%q neighbour mismatch: %v vs %v", "cat", got, want)
	}
}

func TestBuildFromListWithInsertions(t *testing.T) {
	g := NewGraph()
	g.BuildFromList(true, []string{
		"cat", "hat", "car", "cur", "cut", "cart", "chat", "chats", "cats",
	})
	got := g.Neighbours("cat")
	want := []string{
		"hat", "car", "cut", "cart", "chat", "cats",
	}
	if !gotIsWant(got, want) {
		t.Errorf("%q neighbour mismatch: %v vs %v", "cat", got, want)
	}
}
