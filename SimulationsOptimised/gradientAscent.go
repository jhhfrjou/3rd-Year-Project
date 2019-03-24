package main

import "fmt"

func ascend(iters int, delta, rate float64, scen scenario) []allocation {
	alloc := getRandomWeight(scen)
	scores := make([]allocation, iters)
	var diffs [][]float64
	for i := range scores {
		diffs = diff(alloc.FireAllocation, 1, scen, delta)
		alloc.FireAllocation = matAdd(alloc.FireAllocation, matScale(rate, diffs), true)
		normalise(alloc.FireAllocation)
		alloc.Score, _ = simulate(scen, alloc.FireAllocation, 1)
		scores[i] = copyAllocation(alloc)
		if i%100 == 0 {
			fmt.Println(i, alloc.Score)
		}
		if i > 0 && scores[i].Score == scores[i-1].Score {
			break
		}
	}
	return scores
}
