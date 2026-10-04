# HNSW (Hierarchical Navigable Small World) — Design

## Purpose

A from-scratch Go implementation of the HNSW approximate nearest-neighbor algorithm, built as a standalone learning module, prioritizing understanding the algorithm over building a production-grade vector search library.

## Package layout

New top-level directory `hnsw/`, own `go.mod`, following the repo's existing
per-concept convention:

```
hnsw/
  go.mod          // module hnsw, go 1.25.0
  hnsw.go         // public API: Index, Config, New, Insert, Search, Save, Load
  node.go         // internal node/layer representation
  build.go        // insertion + level-assignment logic
  search.go       // greedy search + layer search (searchLayer, selectNeighbors)
  persist.go      // gob-based Save/Load
  hnsw_test.go
  README.md       // brief notes + link back to the source article
```

## Public API

```go
package hnsw

type Numeric interface {
    ~float32 | ~float64
}

type DistanceFunc[T Numeric] func(a, b []T) float64

type Config struct {
    M              int     // max neighbors per node (layers > 0)
    MMax0          int     // max neighbors at layer 0 (typically 2*M)
    EfConstruction int     // candidate list size during insertion
    EfSearch       int     // candidate list size during search
    LevelMult      float64 // mL, level-assignment normalizer (default 1/ln(M))
}

func DefaultConfig() Config // M=16, EfConstruction=200, EfSearch=50, LevelMult=1/ln(16)

type Index[T Numeric] struct { /* unexported */ }

func New[T Numeric](dist DistanceFunc[T], cfg Config) *Index[T]

func (idx *Index[T]) Insert(vector []T) (id uint64)
func (idx *Index[T]) Search(query []T, k int) []SearchResult

type SearchResult struct {
    ID       uint64
    Distance float64
}

func (idx *Index[T]) Save(w io.Writer) error
func Load[T Numeric](r io.Reader, dist DistanceFunc[T]) (*Index[T], error)
```

Notes:
- `dist` is passed to both `New` and `Load` rather than serialized — functions
  aren't gob-encodable, and it keeps "what does distance mean" explicit at
  every call site.
- IDs are auto-assigned, monotonically increasing `uint64`s, stable across
  Save/Load.
- No `Delete` — keeps focus on the paper's core construction/search
  algorithm; node removal isn't fully specified by the original paper anyway.

## Concurrency

A single `sync.RWMutex` on `Index`:
- `Search` and `Save` take `RLock` (read-only / consistent snapshot).
- `Insert` takes `Lock` (exclusive).
- `Load` needs no lock (it builds a fresh index).

## Data flow

**Insert:**
1. Assign a random level per `LevelMult`.
2. If the graph is empty, the new node becomes the entry point; done.
3. Greedy-descend from the entry point's top layer down to `level+1`
   (`ef=1` per layer) to find the best entry point into the insertion layer.
4. From `level` down to `0`: run `searchLayer` with `ef=EfConstruction`,
   select up to `M` (or `MMax0` at layer 0) neighbors, add bidirectional
   links, and prune any neighbor list that now exceeds its max.
5. If `level` exceeds the current top layer, the new node becomes the new
   entry point.

**Search:**
1. Greedy-descend from the entry point's top layer to layer 1 (`ef=1`).
2. At layer 0, run `searchLayer` with `ef = max(EfSearch, k)`.
3. Return the top `k` results by distance.

## Error handling

- `Search`/`Insert` on an empty index return an empty/nil result — "no
  neighbors yet" is a normal state, not a failure.
- `Load` returns an error on format/version mismatch or gob decode failure.
- No panics in the public API.

## Testing

- Table-driven tests building small known point sets (e.g. points on a line
  or grid) where exact nearest neighbors are computable by brute force;
  assert `Search` recall against brute-force ground truth at varying `k`.
- A construction test checking the level distribution roughly matches
  `LevelMult` statistically.
- A Save/Load round-trip test (build index, save, load, confirm identical
  search results).

## CI and README integration

Following the existing per-directory pattern (see `regex-engine-ci.yml`,
`web-crawler-ci.yml`):

- New workflow `.github/workflows/hnsw-ci.yml`: triggers on push/PR to `main`
  touching `hnsw/**`, sets up Go 1.25, runs `go test ./...` inside `hnsw/`.
- New badge added to `README.md` alongside the existing CI badges, linking to
  the new workflow.