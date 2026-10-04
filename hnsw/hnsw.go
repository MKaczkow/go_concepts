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
		cfg:   cfg,
		nodes: make(map[uint64]*node[T]),
		rng:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
