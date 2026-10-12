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
