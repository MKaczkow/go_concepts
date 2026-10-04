package hnsw

import (
	"math"
	"math/rand"
	"slices"
)

// assignLevel draws a random level using the exponential distribution from
// the HNSW paper, normalized by levelMult (mL).
//
// Uses 1-rng.Float64() rather than rng.Float64() as the input to Log: the
// range of 1-rng.Float64() is (0, 1], excluding 0, so -math.Log(...) can
// never produce +Inf (which rng.Float64() returning exactly 0 would cause,
// astronomically rarely).
func assignLevel(rng *rand.Rand, levelMult float64) int {
	return int(math.Floor(-math.Log(1-rng.Float64()) * levelMult))
}

// insert adds vector to the graph, wiring it into every layer from its
// assigned level down to 0, and returns its new node ID. Callers must hold
// idx.mu for writing.
func (idx *Index[T]) insert(vector []T) uint64 {
	id := idx.nextID
	idx.nextID++

	level := assignLevel(idx.rng, idx.cfg.LevelMult)
	n := newNode(id, slices.Clone(vector), level)
	idx.nodes[id] = n

	if len(idx.nodes) == 1 {
		idx.entryPoint = id
		idx.topLevel = level
		return id
	}

	entry := greedySearch(idx.nodes, idx.dist, vector, idx.entryPoint, idx.topLevel, level)

	top := level
	if idx.topLevel < top {
		top = idx.topLevel
	}
	for lc := top; lc >= 0; lc-- {
		candidates := searchLayer(idx.nodes, idx.dist, vector, []uint64{entry}, lc, idx.cfg.EfConstruction)

		maxNeighbors := idx.cfg.M
		if lc == 0 {
			maxNeighbors = idx.cfg.MMax0
		}
		neighborIDs := selectNeighbors(candidates, maxNeighbors)

		n.neighbors[lc] = neighborIDs
		for _, nid := range neighborIDs {
			idx.addLink(nid, id, lc)
		}

		if len(candidates) > 0 {
			entry = candidates[0].id
		}
	}

	if level > idx.topLevel {
		idx.topLevel = level
		idx.entryPoint = id
	}

	return id
}

// addLink adds a bidirectional link from id to neighborID at level,
// pruning neighborID's link list back down to its max if needed.
func (idx *Index[T]) addLink(neighborID, id uint64, level int) {
	neighbor := idx.nodes[neighborID]
	if level > neighbor.topLevel() {
		return
	}
	neighbor.neighbors[level] = append(neighbor.neighbors[level], id)

	maxNeighbors := idx.cfg.M
	if level == 0 {
		maxNeighbors = idx.cfg.MMax0
	}
	if len(neighbor.neighbors[level]) > maxNeighbors {
		candidates := make([]candidate, len(neighbor.neighbors[level]))
		for i, nid := range neighbor.neighbors[level] {
			candidates[i] = candidate{id: nid, distance: idx.dist(idx.nodes[nid].vector, neighbor.vector)}
		}
		neighbor.neighbors[level] = selectNeighbors(candidates, maxNeighbors)
	}
}
