package main

func hillClimb(iters int, delta float64) []allocation {
	scen := getBigDevelopingScenario()
	alloc := getRandomWeight(scen)
	scores := make([]allocation, iters)
	var diffs [][]float64
	for i := range scores {
		alloc.score, _ = simulate(scen, alloc.fireAllocation, 1)
		scores[i] = copyAllocation(alloc)
		diffs = diff(alloc.fireAllocation, 1, scen, delta)
		alloc.fireAllocation = matAdd(alloc.fireAllocation, diffs, true)
		normalise(alloc.fireAllocation)
	}
	return scores
}
