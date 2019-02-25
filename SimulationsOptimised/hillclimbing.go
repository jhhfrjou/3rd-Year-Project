package main

import "time"

func hillClimb(iters int, scen scenario, timeOut time.Duration) ([]allocation, bool) {
	allocs := []allocation{}
	var better bool
	randomBool := new()
	alloc := getRandomWeight(scen)
	alloc.score, _ = simulate(scen, alloc.fireAllocation, 1)
	start := time.Now()
	i := 0
	for i < iters && time.Since(start) < timeOut {
		better = false
		for !better && time.Since(start) < timeOut {
			add := matScale(0.01, getRandomWeight(scen).fireAllocation)
			newAllocs := matAdd(alloc.fireAllocation, add, randomBool.Bool())
			normalise(newAllocs)
			score, _ := simulate(scen, newAllocs, 1)
			if score > alloc.score {
				alloc.fireAllocation = newAllocs
				alloc.score = score
				better = true
			}
		}
		allocs = append(allocs, copyAllocation(alloc))
		i++
	}
	return allocs, !(time.Since(start) < timeOut)
}
