package hnsw

import (
	"math/rand"
	"testing"
)

func TestAssignLevelDistribution(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	levelMult := DefaultConfig().LevelMult

	const trials = 10000
	level0 := 0
	for i := 0; i < trials; i++ {
		if assignLevel(rng, levelMult) == 0 {
			level0++
		}
	}

	// Theoretical P(level == 0) = 1 - exp(-1/levelMult) ~= 0.937 for M=16.
	frac := float64(level0) / float64(trials)
	if frac < 0.85 || frac > 0.98 {
		t.Errorf("fraction at level 0 = %v, want roughly 0.937 (between 0.85 and 0.98)", frac)
	}
}

func TestInsertFirstNode(t *testing.T) {
	idx := New[float64](EuclideanDistance[float64], DefaultConfig())

	id := idx.insert([]float64{1, 2})

	if id != 0 {
		t.Errorf("id = %d, want 0", id)
	}
	if idx.entryPoint != 0 {
		t.Errorf("entryPoint = %d, want 0", idx.entryPoint)
	}
	if len(idx.nodes) != 1 {
		t.Errorf("len(nodes) = %d, want 1", len(idx.nodes))
	}
}

func TestInsertClonesVector(t *testing.T) {
	idx := New[float64](EuclideanDistance[float64], DefaultConfig())

	buf := []float64{1, 2, 3}
	id := idx.insert(buf)

	buf[0] = 999
	buf[1] = 999
	buf[2] = 999

	stored := idx.nodes[id].vector
	want := []float64{1, 2, 3}
	for i := range want {
		if stored[i] != want[i] {
			t.Errorf("stored vector[%d] = %v, want %v (mutating caller's buffer after Insert affected stored vector)", i, stored[i], want[i])
		}
	}
}

func TestInsertWiresNeighbors(t *testing.T) {
	cfg := DefaultConfig()
	cfg.M = 2
	cfg.MMax0 = 4
	idx := New[float64](EuclideanDistance[float64], cfg)

	for i := 0; i < 10; i++ {
		idx.insert([]float64{float64(i), float64(i * i % 5)})
	}

	if len(idx.nodes) != 10 {
		t.Fatalf("len(nodes) = %d, want 10", len(idx.nodes))
	}

	for id, n := range idx.nodes {
		for level, neighbors := range n.neighbors {
			max := cfg.M
			if level == 0 {
				max = cfg.MMax0
			}
			if len(neighbors) > max {
				t.Errorf("node %d level %d has %d neighbors, want <= %d", id, level, len(neighbors), max)
			}
			for _, nb := range neighbors {
				if _, ok := idx.nodes[nb]; !ok {
					t.Errorf("node %d level %d references missing neighbor %d", id, level, nb)
				}
			}
		}
	}

	// Every non-entry node should have at least one neighbor at level 0
	// once more than one node exists.
	for id, n := range idx.nodes {
		if len(n.neighbors[0]) == 0 {
			t.Errorf("node %d has no level-0 neighbors", id)
		}
	}
}
