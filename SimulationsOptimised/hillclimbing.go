package main

func hillClimb(iters int) []float64{
	scen := getBigDevelopingScenario()
	alloc := getRandomWeight(scen)
	scores := make([]float64,iters)
	for i := range scores {
		scores[i],_ = simulate(scen,alloc.fireAllocation,1)
		diffs := diff(alloc.fireAllocation,1,scen,0.0001)
		alloc.fireAllocation = matAdd(alloc.fireAllocation,diffs,true)
	}
	return scores
}