package hnsw

import (
	"encoding/gob"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type snapshotNode[T Numeric] struct {
	ID        uint64
	Vector    []T
	Neighbors [][]uint64
}

type snapshot[T Numeric] struct {
	Cfg        Config
	Nodes      []snapshotNode[T]
	EntryPoint uint64
	TopLevel   int
	NextID     uint64
}

// Save writes a snapshot of the index to w, readable by Load.
func (idx *Index[T]) Save(w io.Writer) error {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	snap := snapshot[T]{
		Cfg:        idx.cfg,
		Nodes:      make([]snapshotNode[T], 0, len(idx.nodes)),
		EntryPoint: idx.entryPoint,
		TopLevel:   idx.topLevel,
		NextID:     idx.nextID,
	}
	for _, n := range idx.nodes {
		snap.Nodes = append(snap.Nodes, snapshotNode[T]{
			ID:        n.id,
			Vector:    n.vector,
			Neighbors: n.neighbors,
		})
	}

	if err := gob.NewEncoder(w).Encode(snap); err != nil {
		return fmt.Errorf("hnsw: encode index: %w", err)
	}
	return nil
}

// Load reads a snapshot written by Save and reconstructs an Index, using
// dist to compare vectors (distance functions are not serializable).
func Load[T Numeric](r io.Reader, dist DistanceFunc[T]) (*Index[T], error) {
	var snap snapshot[T]
	if err := gob.NewDecoder(r).Decode(&snap); err != nil {
		return nil, fmt.Errorf("hnsw: decode index: %w", err)
	}

	idx := &Index[T]{
		dist:       dist,
		cfg:        normalizeConfig(snap.Cfg),
		nodes:      make(map[uint64]*node[T], len(snap.Nodes)),
		entryPoint: snap.EntryPoint,
		topLevel:   snap.TopLevel,
		nextID:     snap.NextID,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for _, sn := range snap.Nodes {
		idx.nodes[sn.ID] = &node[T]{id: sn.ID, vector: sn.Vector, neighbors: sn.Neighbors}
	}

	if len(snap.Nodes) > 0 {
		if _, ok := idx.nodes[snap.EntryPoint]; !ok {
			return nil, fmt.Errorf("hnsw: invalid snapshot: entry point does not refer to a node in the snapshot")
		}
	}
	for _, n := range idx.nodes {
		for _, neighbors := range n.neighbors {
			for _, nid := range neighbors {
				if _, ok := idx.nodes[nid]; !ok {
					return nil, fmt.Errorf("hnsw: invalid snapshot: neighbor list references a node not in the snapshot")
				}
			}
		}
	}

	return idx, nil
}
