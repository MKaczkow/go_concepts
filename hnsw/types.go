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

// normalizeConfig returns cfg with any zero-valued field replaced by the
// corresponding field from DefaultConfig(). This protects against a
// caller-constructed Config (e.g. a struct literal that omits a field, or
// a zero-value Config{}) silently producing a broken or panicking index.
func normalizeConfig(cfg Config) Config {
	def := DefaultConfig()
	if cfg.M == 0 {
		cfg.M = def.M
	}
	if cfg.MMax0 == 0 {
		cfg.MMax0 = def.MMax0
	}
	if cfg.EfConstruction == 0 {
		cfg.EfConstruction = def.EfConstruction
	}
	if cfg.EfSearch == 0 {
		cfg.EfSearch = def.EfSearch
	}
	if cfg.LevelMult == 0 {
		cfg.LevelMult = def.LevelMult
	}
	return cfg
}

// EuclideanDistance is a ready-to-use DistanceFunc for L2 distance.
//
// As with any DistanceFunc, a and b must have the same length; a and all
// other vectors passed to the same Index must also share that length.
// Mismatched lengths cause a panic (index out of range), not an error.
func EuclideanDistance[T Numeric](a, b []T) float64 {
	var sum float64
	for i := range a {
		d := float64(a[i]) - float64(b[i])
		sum += d * d
	}
	return math.Sqrt(sum)
}
