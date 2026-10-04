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
