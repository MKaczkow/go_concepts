package hnsw

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestSaveLoadRoundTrip(t *testing.T) {
	idx := New[float64](EuclideanDistance[float64], DefaultConfig())
	for i := 0; i < 20; i++ {
		idx.Insert([]float64{float64(i), float64(i % 4)})
	}

	var buf bytes.Buffer
	if err := idx.Save(&buf); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	loaded, err := Load[float64](&buf, EuclideanDistance[float64])
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	query := []float64{5.3, 1.7}
	want := idx.Search(query, 5)
	got := loaded.Search(query, 5)

	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}

	// NextID must survive the round trip: the next Insert should get ID 20,
	// not collide with any of the 20 original IDs (0..19).
	newID := loaded.Insert([]float64{100, 100})
	if newID != 20 {
		t.Errorf("Insert() after Load() = %d, want 20", newID)
	}
}

func TestLoadInvalidData(t *testing.T) {
	_, err := Load[float64](bytes.NewReader([]byte("not a valid gob stream")), EuclideanDistance[float64])
	if err == nil {
		t.Error("Load() error = nil, want non-nil for invalid data")
	}
}

func TestLoadRejectsDanglingNeighbor(t *testing.T) {
	snap := snapshot[float64]{
		Cfg:        DefaultConfig(),
		EntryPoint: 0,
		TopLevel:   0,
		NextID:     2,
		Nodes: []snapshotNode[float64]{
			{ID: 0, Vector: []float64{0}, Neighbors: [][]uint64{{1, 42}}},
			{ID: 1, Vector: []float64{1}, Neighbors: [][]uint64{{0}}},
		},
	}

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(snap); err != nil {
		t.Fatalf("encode snapshot: %v", err)
	}

	idx, err := Load[float64](&buf, EuclideanDistance[float64])
	if err == nil {
		t.Error("Load() error = nil, want non-nil for dangling neighbor reference")
	}
	if idx != nil {
		t.Error("Load() index = non-nil, want nil on error")
	}
}

func TestLoadRejectsBadEntryPoint(t *testing.T) {
	snap := snapshot[float64]{
		Cfg:        DefaultConfig(),
		EntryPoint: 99,
		TopLevel:   0,
		NextID:     1,
		Nodes: []snapshotNode[float64]{
			{ID: 0, Vector: []float64{0}, Neighbors: [][]uint64{{}}},
		},
	}

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(snap); err != nil {
		t.Fatalf("encode snapshot: %v", err)
	}

	idx, err := Load[float64](&buf, EuclideanDistance[float64])
	if err == nil {
		t.Error("Load() error = nil, want non-nil for bad entry point")
	}
	if idx != nil {
		t.Error("Load() index = non-nil, want nil on error")
	}
}
