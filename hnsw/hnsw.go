package hnsw

import (
	"math/rand"
	"sync"
	"time"
)

// Index is an HNSW approximate nearest-neighbor index over vectors of
// element type T, compared with a caller-supplied DistanceFunc.
type Index[T Numeric] struct {
	mu         sync.RWMutex
	dist       DistanceFunc[T]
	cfg        Config
	nodes      map[uint64]*node[T]
	entryPoint uint64
	topLevel   int
	nextID     uint64
	rng        *rand.Rand
}

// New creates an empty index using dist to compare vectors and cfg to
// control graph construction and search.
func New[T Numeric](dist DistanceFunc[T], cfg Config) *Index[T] {
	return &Index[T]{
		dist:  dist,
		cfg:   normalizeConfig(cfg),
		nodes: make(map[uint64]*node[T]),
		rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// SearchResult is one match returned by Index.Search.
type SearchResult struct {
	ID       uint64
	Distance float64
}

// Insert adds vector to the index and returns its assigned ID.
//
// vector must have the same length as every other vector passed to this
// Index (and to its DistanceFunc generally); a mismatched length causes a
// panic rather than an error.
func (idx *Index[T]) Insert(vector []T) uint64 {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	return idx.insert(vector)
}

// Search returns up to k nearest neighbors of query, ordered by ascending
// distance. It returns nil if the index is empty or if k <= 0.
//
// query must have the same length as every vector previously inserted into
// this Index; a mismatched length causes a panic rather than an error.
func (idx *Index[T]) Search(query []T, k int) []SearchResult {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if len(idx.nodes) == 0 || k <= 0 {
		return nil
	}

	ef := idx.cfg.EfSearch
	if k > ef {
		ef = k
	}

	entry := greedySearch(idx.nodes, idx.dist, query, idx.entryPoint, idx.topLevel, 0)
	candidates := searchLayer(idx.nodes, idx.dist, query, []uint64{entry}, 0, ef)
	if len(candidates) > k {
		candidates = candidates[:k]
	}

	results := make([]SearchResult, len(candidates))
	for i, c := range candidates {
		results[i] = SearchResult{ID: c.id, Distance: c.distance}
	}
	return results
}
