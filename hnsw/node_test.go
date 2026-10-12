package hnsw

import "testing"

func TestNewNode(t *testing.T) {
	n := newNode[float64](7, []float64{1, 2, 3}, 2)

	if n.id != 7 {
		t.Errorf("id = %d, want 7", n.id)
	}
	if len(n.vector) != 3 {
		t.Errorf("len(vector) = %d, want 3", len(n.vector))
	}
	if len(n.neighbors) != 3 {
		t.Fatalf("len(neighbors) = %d, want 3 (levels 0,1,2)", len(n.neighbors))
	}
	for level, neighbors := range n.neighbors {
		if len(neighbors) != 0 {
			t.Errorf("neighbors[%d] = %v, want empty", level, neighbors)
		}
	}
}

func TestNodeTopLevel(t *testing.T) {
	n := newNode[float64](1, []float64{0}, 4)

	if got := n.topLevel(); got != 4 {
		t.Errorf("topLevel() = %d, want 4", got)
	}
}
