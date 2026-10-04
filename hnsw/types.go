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
