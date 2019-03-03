package main

func ascend(iters int, delta, rate float64, scen scenario) []allocation {
	alloc := getRandomWeight(scen)
	scores := make([]allocation, iters)
	var diffs [][]float64
	for i := range scores {
		alloc.Score, _ = simulate(scen, alloc.FireAllocation, 1)
		scores[i] = copyAllocation(alloc)
		diffs = diff(alloc.FireAllocation, 1, scen, delta)
		alloc.FireAllocation = matAdd(alloc.FireAllocation, matScale(rate, diffs), true)
		normalise(alloc.FireAllocation)
	}
	return scores
}
