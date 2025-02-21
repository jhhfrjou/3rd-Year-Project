package main

import (
	"fmt"
	"time"
)

func hillClimb(iters int, scen scenario, timeOut time.Duration) ([]allocation, bool, int) {
	allocs := make([]allocation, iters)
	var better bool
	randomBool := new()
	alloc := getRandomWeight(scen)
	alloc.Score, _ = simulate(scen, alloc.FireAllocation, 1)
	start := time.Now()
	var timer time.Duration
	i := 0
	j := 0
	k := 0
	avgStep := 0.0
	for i < iters && time.Since(start) < timeOut {
		loopTimer := time.Now()
		if i%100 == 0 {
			fmt.Println(i, avgStep/100, j, alloc.Score, timer)
		}
		k += j
		j = 0
		better = false
		for !better && time.Since(start) < timeOut {
			add := matScale(logistic(j, 10000, 1), getRandomWeight(scen).FireAllocation)
			newAllocs := matAdd(alloc.FireAllocation, add, randomBool.Bool())
			normalise(newAllocs)
			score, _ := simulate(scen, newAllocs, 1)
			if score > alloc.Score {
				alloc.FireAllocation = newAllocs
				alloc.Score = score
				better = true
			}

			avgStep += float64(j)
			j++
		}
		allocs[i] = copyAllocation(alloc)
		i++
		timer = time.Since(loopTimer)

	}
	return allocs, time.Since(start) > timeOut, k
}
