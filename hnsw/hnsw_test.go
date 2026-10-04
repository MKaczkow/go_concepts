package hnsw

import "testing"

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
