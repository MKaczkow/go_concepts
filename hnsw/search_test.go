package hnsw

import (
	"math"
	"testing"
)

// buildLine creates 5 nodes at x=0..4, each linked to its immediate
// neighbor(s) at level 0 only (a line graph).
func buildLine(t *testing.T) map[uint64]*node[float64] {
	t.Helper()
	nodes := make(map[uint64]*node[float64])
	for i := uint64(0); i < 5; i++ {
		nodes[i] = newNode[float64](i, []float64{float64(i)}, 0)
	}
	links := map[uint64][]uint64{
		0: {1},
		1: {0, 2},
		2: {1, 3},
		3: {2, 4},
		4: {3},
	}
	for id, neighbors := range links {
		nodes[id].neighbors[0] = neighbors
	}
	return nodes
}

func TestSearchLayer(t *testing.T) {
	nodes := buildLine(t)
	query := []float64{3.6}

	results := searchLayer(nodes, EuclideanDistance[float64], query, []uint64{0}, 0, 2)

	if len(results) != 2 {
		t.Fatalf("len(results) = %d, want 2", len(results))
	}
	if results[0].id != 4 || math.Abs(results[0].distance-0.4) > 1e-9 {
		t.Errorf("results[0] = %+v, want {id:4 distance:0.4}", results[0])
	}
	if results[1].id != 3 || math.Abs(results[1].distance-0.6) > 1e-9 {
		t.Errorf("results[1] = %+v, want {id:3 distance:0.6}", results[1])
	}
}

func TestGreedySearch(t *testing.T) {
	// Two nodes present at levels 0-2; node 1 (x=10) is closer than
	// node 0 (x=0) to the query (x=9), so greedy descent should move
	// from the entry point (0) to 1 while descending from level 2 to
	// just above level 0.
	n0 := newNode[float64](0, []float64{0}, 2)
	n1 := newNode[float64](1, []float64{10}, 2)
	n0.neighbors[2] = []uint64{1}
	n0.neighbors[1] = []uint64{1}
	n1.neighbors[2] = []uint64{0}
	n1.neighbors[1] = []uint64{0}
	nodes := map[uint64]*node[float64]{0: n0, 1: n1}

	got := greedySearch(nodes, EuclideanDistance[float64], []float64{9}, 0, 2, 0)

	if got != 1 {
		t.Errorf("greedySearch() = %d, want 1", got)
	}
}

func TestSelectNeighbors(t *testing.T) {
	candidates := []candidate{
		{id: 3, distance: 5},
		{id: 1, distance: 1},
		{id: 2, distance: 3},
	}

	got := selectNeighbors(candidates, 2)

	want := []uint64{1, 2}
	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}
