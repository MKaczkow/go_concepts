# HNSW Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a from-scratch, generic Go implementation of the HNSW approximate nearest-neighbor algorithm as a standalone module (`hnsw/`), plus CI and README integration matching this repo's existing per-concept conventions.

**Architecture:** A single `hnsw` package, split by responsibility (not by layer): shared types, an internal node/graph representation, the search algorithms (greedy descent + layer beam search), the insertion/construction algorithm, the public `Index` API wrapping those with a `sync.RWMutex`, and gob-based persistence. Every piece is built bottom-up so each task compiles and is independently testable: types → node → search → Index skeleton → build (insert) → public API (Insert/Search) → persistence → docs/CI.

**Tech Stack:** Go 1.25, standard library only (`testing`, `sort`, `math`, `math/rand`, `sync`, `encoding/gob`). No third-party test libraries — this repo's algorithm modules (e.g. `regex_engine/`) use stdlib `testing` directly, not `testify`.

**Spec:** `docs/superpowers/specs/2026-10-04-hnsw-design.md`

## Global Constraints

- Module: `module hnsw`, `go 1.25.0` (matches the newest `go.mod` versions already in this repo, e.g. `web_app_bis/go.mod`).
- No `Delete` operation — out of scope per spec.
- IDs are auto-assigned, monotonically increasing `uint64`, starting at 0.
- Distance function is always passed in by the caller (to `New` and `Load`), never stored/serialized as code.
- Concurrency: single `sync.RWMutex` on `Index` — `Insert` takes `Lock`, `Search`/`Save` take `RLock`, `Load` takes no lock.
- No panics in the public API; `Search`/`Insert` on an empty index return an empty/nil result, not an error.

---

### Task 1: Core types and config

**Files:**
- Create: `hnsw/go.mod`
- Create: `hnsw/types.go`
- Test: `hnsw/types_test.go`

**Interfaces:**
- Consumes: nothing (foundational task).
- Produces:
  - `type Numeric interface { ~float32 | ~float64 }`
  - `type DistanceFunc[T Numeric] func(a, b []T) float64`
  - `type Config struct { M, MMax0, EfConstruction, EfSearch int; LevelMult float64 }`
  - `func DefaultConfig() Config`
  - `func EuclideanDistance[T Numeric](a, b []T) float64`

- [ ] **Step 1: Create the module**

Create `hnsw/go.mod`:

```
module hnsw

go 1.25.0
```

- [ ] **Step 2: Write the failing test**

Create `hnsw/types_test.go`:

```go
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
```

- [ ] **Step 3: Run test to verify it fails**

Run: `cd hnsw && go test ./... -run 'TestDefaultConfig|TestEuclideanDistance' -v`
Expected: FAIL (build fails — `DefaultConfig`, `EuclideanDistance` undefined)

- [ ] **Step 4: Write minimal implementation**

Create `hnsw/types.go`:

```go
package hnsw

import "math"

// Numeric constrains the element type of vectors stored in the index.
type Numeric interface {
	~float32 | ~float64
}

// DistanceFunc computes the distance between two vectors of the same length.
// Smaller values mean "closer".
type DistanceFunc[T Numeric] func(a, b []T) float64

// Config holds the tunable HNSW parameters.
type Config struct {
	M              int     // max neighbors per node (layers > 0)
	MMax0          int     // max neighbors at layer 0
	EfConstruction int     // candidate list size during insertion
	EfSearch       int     // candidate list size during search
	LevelMult      float64 // mL, level-assignment normalizer
}

// DefaultConfig returns the paper-recommended defaults for M=16.
func DefaultConfig() Config {
	m := 16
	return Config{
		M:              m,
		MMax0:          2 * m,
		EfConstruction: 200,
		EfSearch:       50,
		LevelMult:      1 / math.Log(float64(m)),
	}
}

// EuclideanDistance is a ready-to-use DistanceFunc for L2 distance.
func EuclideanDistance[T Numeric](a, b []T) float64 {
	var sum float64
	for i := range a {
		d := float64(a[i]) - float64(b[i])
		sum += d * d
	}
	return math.Sqrt(sum)
}
```

- [ ] **Step 5: Run test to verify it passes**

Run: `cd hnsw && go test ./... -run 'TestDefaultConfig|TestEuclideanDistance' -v`
Expected: PASS

- [ ] **Step 6: Commit**

```bash
git add hnsw/go.mod hnsw/types.go hnsw/types_test.go
git commit -m "feat(hnsw): add core types and default config"
```

---

### Task 2: Node representation

**Files:**
- Create: `hnsw/node.go`
- Test: `hnsw/node_test.go`

**Interfaces:**
- Consumes: `Numeric` (Task 1)
- Produces:
  - `type node[T Numeric] struct { id uint64; vector []T; neighbors [][]uint64 }`
  - `func newNode[T Numeric](id uint64, vector []T, level int) *node[T]`
  - `func (n *node[T]) topLevel() int`

- [ ] **Step 1: Write the failing test**

Create `hnsw/node_test.go`:

```go
package hnsw

import "testing"

func TestNewNode(t *testing.T) {
	n := newNode[float64](7, []float64{1, 2, 3}, 2)

	if n.id != 7 {
		t.Errorf("id = %d, want 7", n.id)
	}
	if len(n.vector) != 3 {
		t.Errorf("len(vector) = %d, want 3", len(n.vector))
	}
	if len(n.neighbors) != 3 {
		t.Fatalf("len(neighbors) = %d, want 3 (levels 0,1,2)", len(n.neighbors))
	}
	for level, neighbors := range n.neighbors {
		if len(neighbors) != 0 {
			t.Errorf("neighbors[%d] = %v, want empty", level, neighbors)
		}
	}
}

func TestNodeTopLevel(t *testing.T) {
	n := newNode[float64](1, []float64{0}, 4)

	if got := n.topLevel(); got != 4 {
		t.Errorf("topLevel() = %d, want 4", got)
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd hnsw && go test ./... -run 'TestNewNode|TestNodeTopLevel' -v`
Expected: FAIL (build fails — `newNode` undefined)

- [ ] **Step 3: Write minimal implementation**

Create `hnsw/node.go`:

```go
package hnsw

// node is a single point stored in the graph. neighbors[level] holds the
// IDs of this node's neighbors at that layer; a node exists in every layer
// from 0 up to its topLevel().
type node[T Numeric] struct {
	id        uint64
	vector    []T
	neighbors [][]uint64
}

func newNode[T Numeric](id uint64, vector []T, level int) *node[T] {
	neighbors := make([][]uint64, level+1)
	for i := range neighbors {
		neighbors[i] = []uint64{}
	}
	return &node[T]{id: id, vector: vector, neighbors: neighbors}
}

func (n *node[T]) topLevel() int {
	return len(n.neighbors) - 1
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd hnsw && go test ./... -run 'TestNewNode|TestNodeTopLevel' -v`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add hnsw/node.go hnsw/node_test.go
git commit -m "feat(hnsw): add internal node representation"
```

---

### Task 3: Search algorithms

**Files:**
- Create: `hnsw/search.go`
- Test: `hnsw/search_test.go`

**Interfaces:**
- Consumes: `node[T]`, `newNode` (Task 2); `Numeric`, `DistanceFunc[T]`, `EuclideanDistance` (Task 1)
- Produces:
  - `type candidate struct { id uint64; distance float64 }`
  - `func searchLayer[T Numeric](nodes map[uint64]*node[T], dist DistanceFunc[T], query []T, entryPoints []uint64, level int, ef int) []candidate`
  - `func greedySearch[T Numeric](nodes map[uint64]*node[T], dist DistanceFunc[T], query []T, entryID uint64, startLevel int, targetLevel int) uint64`
  - `func selectNeighbors(candidates []candidate, max int) []uint64`

- [ ] **Step 1: Write the failing test**

Create `hnsw/search_test.go`:

```go
package hnsw

import (
	"math"
	"testing"
)

// buildLine creates 5 nodes at x=0..4, each linked to its immediate
// neighbor(s) at level 0 only (a line graph).
func buildLine(t *testing.T) map[uint64]*node[float64] {
	t.Helper()
	nodes := make(map[uint64]*node[float64])
	for i := uint64(0); i < 5; i++ {
		nodes[i] = newNode[float64](i, []float64{float64(i)}, 0)
	}
	links := map[uint64][]uint64{
		0: {1},
		1: {0, 2},
		2: {1, 3},
		3: {2, 4},
		4: {3},
	}
	for id, neighbors := range links {
		nodes[id].neighbors[0] = neighbors
	}
	return nodes
}

func TestSearchLayer(t *testing.T) {
	nodes := buildLine(t)
	query := []float64{3.6}

	results := searchLayer(nodes, EuclideanDistance[float64], query, []uint64{0}, 0, 2)

	if len(results) != 2 {
		t.Fatalf("len(results) = %d, want 2", len(results))
	}
	if results[0].id != 4 || math.Abs(results[0].distance-0.4) > 1e-9 {
		t.Errorf("results[0] = %+v, want {id:4 distance:0.4}", results[0])
	}
	if results[1].id != 3 || math.Abs(results[1].distance-0.6) > 1e-9 {
		t.Errorf("results[1] = %+v, want {id:3 distance:0.6}", results[1])
	}
}

func TestGreedySearch(t *testing.T) {
	// Two nodes present at levels 0-2; node 1 (x=10) is closer than
	// node 0 (x=0) to the query (x=9), so greedy descent should move
	// from the entry point (0) to 1 while descending from level 2 to
	// just above level 0.
	n0 := newNode[float64](0, []float64{0}, 2)
	n1 := newNode[float64](1, []float64{10}, 2)
	n0.neighbors[2] = []uint64{1}
	n0.neighbors[1] = []uint64{1}
	n1.neighbors[2] = []uint64{0}
	n1.neighbors[1] = []uint64{0}
	nodes := map[uint64]*node[float64]{0: n0, 1: n1}

	got := greedySearch(nodes, EuclideanDistance[float64], []float64{9}, 0, 2, 0)

	if got != 1 {
		t.Errorf("greedySearch() = %d, want 1", got)
	}
}

func TestSelectNeighbors(t *testing.T) {
	candidates := []candidate{
		{id: 3, distance: 5},
		{id: 1, distance: 1},
		{id: 2, distance: 3},
	}

	got := selectNeighbors(candidates, 2)

	want := []uint64{1, 2}
	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd hnsw && go test ./... -run 'TestSearchLayer|TestGreedySearch|TestSelectNeighbors' -v`
Expected: FAIL (build fails — `searchLayer`, `greedySearch`, `selectNeighbors`, `candidate` undefined)

- [ ] **Step 3: Write minimal implementation**

Create `hnsw/search.go`:

```go
package hnsw

import "sort"

// candidate pairs a node ID with its distance to some query vector.
type candidate struct {
	id       uint64
	distance float64
}

// searchLayer performs a greedy beam search within a single layer, starting
// from entryPoints, and returns up to ef candidates nearest to query,
// sorted by ascending distance.
func searchLayer[T Numeric](
	nodes map[uint64]*node[T],
	dist DistanceFunc[T],
	query []T,
	entryPoints []uint64,
	level int,
	ef int,
) []candidate {
	visited := make(map[uint64]bool, len(entryPoints))
	candidates := make([]candidate, 0, len(entryPoints))
	results := make([]candidate, 0, len(entryPoints))

	for _, id := range entryPoints {
		d := dist(nodes[id].vector, query)
		c := candidate{id: id, distance: d}
		candidates = append(candidates, c)
		results = append(results, c)
		visited[id] = true
	}

	for len(candidates) > 0 {
		sort.Slice(candidates, func(i, j int) bool { return candidates[i].distance < candidates[j].distance })
		curr := candidates[0]
		candidates = candidates[1:]

		sort.Slice(results, func(i, j int) bool { return results[i].distance < results[j].distance })
		if len(results) >= ef && curr.distance > results[len(results)-1].distance {
			break
		}

		currNode := nodes[curr.id]
		if level > currNode.topLevel() {
			continue
		}

		for _, neighborID := range currNode.neighbors[level] {
			if visited[neighborID] {
				continue
			}
			visited[neighborID] = true

			d := dist(nodes[neighborID].vector, query)
			sort.Slice(results, func(i, j int) bool { return results[i].distance < results[j].distance })

			if len(results) < ef || d < results[len(results)-1].distance {
				nc := candidate{id: neighborID, distance: d}
				candidates = append(candidates, nc)
				results = append(results, nc)
				if len(results) > ef {
					sort.Slice(results, func(i, j int) bool { return results[i].distance < results[j].distance })
					results = results[:ef]
				}
			}
		}
	}

	sort.Slice(results, func(i, j int) bool { return results[i].distance < results[j].distance })
	if len(results) > ef {
		results = results[:ef]
	}
	return results
}

// greedySearch descends from entryID's layer startLevel down to
// targetLevel+1, at each layer greedily moving to the closest neighbor
// until no closer neighbor exists. It returns the ID of the closest node
// found, to be used as the entry point for the next phase of search.
func greedySearch[T Numeric](
	nodes map[uint64]*node[T],
	dist DistanceFunc[T],
	query []T,
	entryID uint64,
	startLevel int,
	targetLevel int,
) uint64 {
	curr := entryID
	currDist := dist(nodes[curr].vector, query)

	for level := startLevel; level > targetLevel; level-- {
		changed := true
		for changed {
			changed = false
			currNode := nodes[curr]
			if level > currNode.topLevel() {
				continue
			}
			for _, neighborID := range currNode.neighbors[level] {
				d := dist(nodes[neighborID].vector, query)
				if d < currDist {
					curr = neighborID
					currDist = d
					changed = true
				}
			}
		}
	}
	return curr
}

// selectNeighbors picks up to max candidates with the smallest distance,
// returning their IDs sorted by ascending distance.
func selectNeighbors(candidates []candidate, max int) []uint64 {
	sorted := make([]candidate, len(candidates))
	copy(sorted, candidates)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].distance < sorted[j].distance })
	if len(sorted) > max {
		sorted = sorted[:max]
	}
	ids := make([]uint64, len(sorted))
	for i, c := range sorted {
		ids[i] = c.id
	}
	return ids
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd hnsw && go test ./... -run 'TestSearchLayer|TestGreedySearch|TestSelectNeighbors' -v`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add hnsw/search.go hnsw/search_test.go
git commit -m "feat(hnsw): add greedy and layer search algorithms"
```

---

### Task 4: Index struct and constructor

**Files:**
- Create: `hnsw/hnsw.go`
- Test: `hnsw/hnsw_test.go`

**Interfaces:**
- Consumes: `node[T]` (Task 2); `Numeric`, `DistanceFunc[T]`, `Config`, `EuclideanDistance` (Task 1)
- Produces:
  - `type Index[T Numeric] struct { mu sync.RWMutex; dist DistanceFunc[T]; cfg Config; nodes map[uint64]*node[T]; entryPoint uint64; topLevel int; nextID uint64; rng *rand.Rand }`
  - `func New[T Numeric](dist DistanceFunc[T], cfg Config) *Index[T]`

- [ ] **Step 1: Write the failing test**

Create `hnsw/hnsw_test.go`:

```go
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
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd hnsw && go test ./... -run TestNew -v`
Expected: FAIL (build fails — `Index`, `New` undefined)

- [ ] **Step 3: Write minimal implementation**

Create `hnsw/hnsw.go`:

```go
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
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd hnsw && go test ./... -run TestNew -v`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add hnsw/hnsw.go hnsw/hnsw_test.go
git commit -m "feat(hnsw): add Index struct and constructor"
```

---

### Task 5: Insertion / graph construction

**Files:**
- Modify: `hnsw/build.go` (create)
- Test: `hnsw/build_test.go`

**Interfaces:**
- Consumes:
  - `Index[T]` struct fields (Task 4)
  - `node[T]`, `newNode` (Task 2)
  - `searchLayer`, `greedySearch`, `selectNeighbors`, `candidate` (Task 3)
- Produces:
  - `func assignLevel(rng *rand.Rand, levelMult float64) int`
  - `func (idx *Index[T]) insert(vector []T) uint64`
  - `func (idx *Index[T]) addLink(neighborID, id uint64, level int)`

- [ ] **Step 1: Write the failing test**

Create `hnsw/build_test.go`:

```go
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
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd hnsw && go test ./... -run 'TestAssignLevelDistribution|TestInsertFirstNode|TestInsertWiresNeighbors' -v`
Expected: FAIL (build fails — `assignLevel`, `idx.insert` undefined)

- [ ] **Step 3: Write minimal implementation**

Create `hnsw/build.go`:

```go
package hnsw

import (
	"math"
	"math/rand"
)

// assignLevel draws a random level using the exponential distribution from
// the HNSW paper, normalized by levelMult (mL).
func assignLevel(rng *rand.Rand, levelMult float64) int {
	return int(math.Floor(-math.Log(rng.Float64()) * levelMult))
}

// insert adds vector to the graph, wiring it into every layer from its
// assigned level down to 0, and returns its new node ID. Callers must hold
// idx.mu for writing.
func (idx *Index[T]) insert(vector []T) uint64 {
	id := idx.nextID
	idx.nextID++

	level := assignLevel(idx.rng, idx.cfg.LevelMult)
	n := newNode(id, vector, level)
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
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd hnsw && go test ./... -run 'TestAssignLevelDistribution|TestInsertFirstNode|TestInsertWiresNeighbors' -v`
Expected: PASS

- [ ] **Step 5: Commit**

```bash
git add hnsw/build.go hnsw/build_test.go
git commit -m "feat(hnsw): add insertion and graph construction"
```

---

### Task 6: Public Insert/Search API

**Files:**
- Modify: `hnsw/hnsw.go`
- Modify: `hnsw/hnsw_test.go`

**Interfaces:**
- Consumes: `idx.insert` (Task 5); `searchLayer`, `greedySearch` (Task 3); `Index[T]` (Task 4)
- Produces:
  - `type SearchResult struct { ID uint64; Distance float64 }`
  - `func (idx *Index[T]) Insert(vector []T) uint64`
  - `func (idx *Index[T]) Search(query []T, k int) []SearchResult`

- [ ] **Step 1: Write the failing test**

Append to `hnsw/hnsw_test.go`:

```go
func bruteForceKNN(points map[uint64][]float64, query []float64, k int) []SearchResult {
	results := make([]SearchResult, 0, len(points))
	for id, v := range points {
		results = append(results, SearchResult{ID: id, Distance: EuclideanDistance(v, query)})
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Distance < results[j].Distance })
	if len(results) > k {
		results = results[:k]
	}
	return results
}

func TestSearchOnEmptyIndex(t *testing.T) {
	idx := New[float64](EuclideanDistance[float64], DefaultConfig())

	got := idx.Search([]float64{1, 2}, 3)

	if got != nil {
		t.Errorf("Search() on empty index = %v, want nil", got)
	}
}

func TestInsertAndSearchRecall(t *testing.T) {
	idx := New[float64](EuclideanDistance[float64], DefaultConfig())
	points := make(map[uint64][]float64)

	// 10x5 grid of points, inserted in index order 0..49.
	for i := 0; i < 50; i++ {
		v := []float64{float64(i % 10), float64(i / 10)}
		id := idx.Insert(v)
		points[id] = v
	}

	// (4.3, 1.9) is chosen so the top-3 distances are strictly distinct
	// (0.10, 0.50, 0.90 to grid points (4,2), (5,2), (4,1)); a tied
	// distance would make the brute-force/HNSW comparison order-dependent.
	query := []float64{4.3, 1.9}
	want := bruteForceKNN(points, query, 3)
	got := idx.Search(query, 3)

	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i].ID != want[i].ID {
			t.Errorf("got[%d].ID = %d, want %d", i, got[i].ID, want[i].ID)
		}
		if math.Abs(got[i].Distance-want[i].Distance) > 1e-9 {
			t.Errorf("got[%d].Distance = %v, want %v", i, got[i].Distance, want[i].Distance)
		}
	}
}
```

Add `"math"` and `"sort"` to the existing `import` block in `hnsw/hnsw_test.go` (alongside `"testing"`).

- [ ] **Step 2: Run test to verify it fails**

Run: `cd hnsw && go test ./... -run 'TestSearchOnEmptyIndex|TestInsertAndSearchRecall' -v`
Expected: FAIL (build fails — `idx.Insert`, `idx.Search`, `SearchResult` undefined)

- [ ] **Step 3: Write minimal implementation**

Append to `hnsw/hnsw.go`:

```go
// SearchResult is one match returned by Index.Search.
type SearchResult struct {
	ID       uint64
	Distance float64
}

// Insert adds vector to the index and returns its assigned ID.
func (idx *Index[T]) Insert(vector []T) uint64 {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	return idx.insert(vector)
}

// Search returns up to k nearest neighbors of query, ordered by ascending
// distance. It returns nil if the index is empty.
func (idx *Index[T]) Search(query []T, k int) []SearchResult {
	idx.mu.RLock()
	defer idx.mu.RUnlock()

	if len(idx.nodes) == 0 {
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
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd hnsw && go test ./... -run 'TestSearchOnEmptyIndex|TestInsertAndSearchRecall' -v`
Expected: PASS

- [ ] **Step 5: Run the full package test suite**

Run: `cd hnsw && go test ./... -v`
Expected: PASS (all tests from Tasks 1-6)

- [ ] **Step 6: Commit**

```bash
git add hnsw/hnsw.go hnsw/hnsw_test.go
git commit -m "feat(hnsw): add public Insert/Search API"
```

---

### Task 7: Persistence (Save/Load)

**Files:**
- Create: `hnsw/persist.go`
- Test: `hnsw/persist_test.go`

**Interfaces:**
- Consumes: `Index[T]` (Task 4), `node[T]` (Task 2), `Config` (Task 1), `Insert`/`Search` (Task 6)
- Produces:
  - `func (idx *Index[T]) Save(w io.Writer) error`
  - `func Load[T Numeric](r io.Reader, dist DistanceFunc[T]) (*Index[T], error)`

- [ ] **Step 1: Write the failing test**

Create `hnsw/persist_test.go`:

```go
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
```

- [ ] **Step 2: Run test to verify it fails**

Run: `cd hnsw && go test ./... -run 'TestSaveLoadRoundTrip|TestLoadInvalidData' -v`
Expected: FAIL (build fails — `idx.Save`, `Load` undefined)

- [ ] **Step 3: Write minimal implementation**

Create `hnsw/persist.go`:

```go
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
		cfg:        snap.Cfg,
		nodes:      make(map[uint64]*node[T], len(snap.Nodes)),
		entryPoint: snap.EntryPoint,
		topLevel:   snap.TopLevel,
		nextID:     snap.NextID,
		rng:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for _, sn := range snap.Nodes {
		idx.nodes[sn.ID] = &node[T]{id: sn.ID, vector: sn.Vector, neighbors: sn.Neighbors}
	}

	return idx, nil
}
```

- [ ] **Step 4: Run test to verify it passes**

Run: `cd hnsw && go test ./... -run 'TestSaveLoadRoundTrip|TestLoadInvalidData' -v`
Expected: PASS

- [ ] **Step 5: Run the full package test suite**

Run: `cd hnsw && go test ./... -v`
Expected: PASS (all tests from Tasks 1-7)

- [ ] **Step 6: Commit**

```bash
git add hnsw/persist.go hnsw/persist_test.go
git commit -m "feat(hnsw): add gob-based Save/Load persistence"
```

---

### Task 8: README and CI integration

**Files:**
- Create: `hnsw/README.md`
- Create: `.github/workflows/hnsw-ci.yml`
- Modify: `README.md`

**Interfaces:**
- Consumes: nothing new (documentation/CI only).
- Produces: nothing consumed by later tasks (final task).

- [ ] **Step 1: Write `hnsw/README.md`**

```markdown
# HNSW

From-scratch Go implementation of Hierarchical Navigable Small World (HNSW)
approximate nearest-neighbor search, built as a learning exercise.

## Usage

```go
idx := hnsw.New[float64](hnsw.EuclideanDistance[float64], hnsw.DefaultConfig())
id := idx.Insert([]float64{1, 2, 3})
results := idx.Search([]float64{1, 2, 3}, 5)
```

## Design

See [`docs/superpowers/specs/2026-10-04-hnsw-design.md`](../docs/superpowers/specs/2026-10-04-hnsw-design.md)
for the interface design and algorithm walkthrough.

## Reference

Malkov, Y. A., & Yashunin, D. A. (2018). *Efficient and robust approximate
nearest neighbor search using Hierarchical Navigable Small World graphs.*
https://arxiv.org/abs/1603.09320
```

- [ ] **Step 2: Create the CI workflow**

Create `.github/workflows/hnsw-ci.yml`, following the existing pattern (compare `.github/workflows/web-crawler-ci.yml`):

```yaml
name: HNSW CI

on:
  push:
    branches: [ "main" ]
    paths:
      - 'hnsw/**'
  pull_request:
    branches: [ "main" ]
    paths:
      - 'hnsw/**'

jobs:
  test:
    name: Test
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.25'

    - name: Run tests in 'hnsw' directory
      run: |
        cd hnsw
        go test ./...
```

- [ ] **Step 3: Add the badge to the root README**

In `README.md`, add a new badge line after the existing `Regex Engine CI` badge (around line 8), before the `Gilded Rose CI` badge:

```markdown
[![HNSW CI](https://github.com/MKaczkow/go_concepts/actions/workflows/hnsw-ci.yml/badge.svg)](https://github.com/MKaczkow/go_concepts/actions/workflows/hnsw-ci.yml)  
```

- [ ] **Step 4: Add a todo entry, following the existing checklist style**

In `README.md`, under the `### todo` section, add:

```markdown
- [x] HNSW (Hierarchical Navigable Small World) vector index from scratch
```

- [ ] **Step 5: Verify the whole module still builds and tests pass**

Run: `cd hnsw && go vet ./... && go test ./...`
Expected: PASS, no vet warnings

- [ ] **Step 6: Commit**

```bash
git add hnsw/README.md .github/workflows/hnsw-ci.yml README.md
git commit -m "docs(hnsw): add README, CI workflow, and badge"
```