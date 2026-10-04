package hnsw

import (
	"math"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.M != 16 {
		t.Errorf("M = %d, want 16", cfg.M)
	}
	if cfg.MMax0 != 32 {
		t.Errorf("MMax0 = %d, want 32", cfg.MMax0)
	}
	if cfg.EfConstruction != 200 {
		t.Errorf("EfConstruction = %d, want 200", cfg.EfConstruction)
	}
	if cfg.EfSearch != 50 {
		t.Errorf("EfSearch = %d, want 50", cfg.EfSearch)
	}

	wantLevelMult := 1 / math.Log(16)
	if math.Abs(cfg.LevelMult-wantLevelMult) > 1e-9 {
		t.Errorf("LevelMult = %v, want %v", cfg.LevelMult, wantLevelMult)
	}
}

func TestEuclideanDistance(t *testing.T) {
	tests := []struct {
		name string
		a, b []float64
		want float64
	}{
		{"identical points", []float64{1, 2, 3}, []float64{1, 2, 3}, 0},
		{"3-4-5 triangle", []float64{0, 0}, []float64{3, 4}, 5},
		{"single dimension", []float64{2}, []float64{7}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EuclideanDistance(tt.a, tt.b)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("EuclideanDistance(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
