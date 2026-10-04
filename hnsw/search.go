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
