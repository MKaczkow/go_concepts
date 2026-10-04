package hnsw

import (
	"math"
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := DefaultConfig()
	idx := New[float64](EuclideanDistance[float64], cfg)

	if idx == nil {
		t.Fatal("New() returned nil")
	}
	if len(idx.nodes) != 0 {
		t.Errorf("len(nodes) = %d, want 0", len(idx.nodes))
	}
	if idx.nextID != 0 {
		t.Errorf("nextID = %d, want 0", idx.nextID)
	}
	if idx.cfg != cfg {
		t.Errorf("cfg = %+v, want %+v", idx.cfg, cfg)
	}
	if idx.rng == nil {
		t.Error("rng is nil, want initialized")
	}
}

func bruteForceKNN(points map[uint64][]float64, query []float64, k int) []SearchResult {
	results := make([]SearchResult, 0, len(points))
	for id, v := range points {
		results = append(results, SearchResult{ID: id, Distance: EuclideanDistance(v, query)})
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Distance < results[j].Distance })
	if len(results) > k {
		results = results[:k]
	}
	return results
}

func TestSearchOnEmptyIndex(t *testing.T) {
	idx := New[float64](EuclideanDistance[float64], DefaultConfig())

	got := idx.Search([]float64{1, 2}, 3)

	if got != nil {
		t.Errorf("Search() on empty index = %v, want nil", got)
	}
}

func TestInsertAndSearchRecall(t *testing.T) {
	idx := New[float64](EuclideanDistance[float64], DefaultConfig())
	points := make(map[uint64][]float64)

	// 10x5 grid of points, inserted in index order 0..49.
	for i := 0; i < 50; i++ {
		v := []float64{float64(i % 10), float64(i / 10)}
		id := idx.Insert(v)
		points[id] = v
	}

	// (4.3, 1.9) is chosen so the top-3 distances are strictly distinct
	// (0.10, 0.50, 0.90 to grid points (4,2), (5,2), (4,1)); a tied
	// distance would make the brute-force/HNSW comparison order-dependent.
	query := []float64{4.3, 1.9}
	want := bruteForceKNN(points, query, 3)
	got := idx.Search(query, 3)

	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i].ID != want[i].ID {
			t.Errorf("got[%d].ID = %d, want %d", i, got[i].ID, want[i].ID)
		}
		if math.Abs(got[i].Distance-want[i].Distance) > 1e-9 {
			t.Errorf("got[%d].Distance = %v, want %v", i, got[i].Distance, want[i].Distance)
		}
	}
}
