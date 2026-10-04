package hnsw

import (
	"bytes"
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
}

func TestLoadInvalidData(t *testing.T) {
	_, err := Load[float64](bytes.NewReader([]byte("not a valid gob stream")), EuclideanDistance[float64])
	if err == nil {
		t.Error("Load() error = nil, want non-nil for invalid data")
	}
}
