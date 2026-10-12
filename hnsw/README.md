# HNSW

From-scratch Go implementation of Hierarchical Navigable Small World (HNSW)
approximate nearest-neighbor search, built as a learning exercise.

## Usage

```go
idx := hnsw.New[float64](hnsw.EuclideanDistance[float64], hnsw.DefaultConfig())
id := idx.Insert([]float64{1, 2, 3})
results := idx.Search([]float64{1, 2, 3}, 5)
```

## Save and load

An index can be persisted to any `io.Writer` and restored from any
`io.Reader` (a file, a `bytes.Buffer`, etc.). `Load` needs a `DistanceFunc`
because functions aren't serializable.

```go
var buf bytes.Buffer
if err := idx.Save(&buf); err != nil {
    log.Fatal(err)
}

loaded, err := hnsw.Load[float64](&buf, hnsw.EuclideanDistance[float64])
if err != nil {
    log.Fatal(err)
}
```

## Config

`DefaultConfig()` returns paper-recommended defaults for `M=16`. Any field
left at its zero value (e.g. in a partial struct literal) is filled in from
the default when the `Index` is created or loaded.

- `M` — max neighbors per node at layers above 0.
- `MMax0` — max neighbors per node at layer 0 (usually `2*M`).
- `EfConstruction` — candidate list size while inserting; higher gives a
  better-connected graph at the cost of slower inserts.
- `EfSearch` — candidate list size while searching; higher gives better
  recall at the cost of slower searches.
- `LevelMult` — `mL`, the normalizer for the random level-assignment
  distribution.

## Design

See [`docs/superpowers/specs/2026-10-04-hnsw-design.md`](../docs/superpowers/specs/2026-10-04-hnsw-design.md)
for the interface design and algorithm walkthrough.

## Reference

Malkov, Y. A., & Yashunin, D. A. (2018). *Efficient and robust approximate
nearest neighbor search using Hierarchical Navigable Small World graphs.*
https://arxiv.org/abs/1603.09320
