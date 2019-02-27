package main

import (
	"fmt"
	"time"
)

func hillClimb(iters int, scen scenario, timeOut time.Duration) ([]allocation, bool) {
	allocs := []allocation{}
	var better bool
	randomBool := new()
	alloc := getRandomWeight(scen)
	alloc.score, _ = simulate(scen, alloc.fireAllocation, 1)
	start := time.Now()
	var timer time.Duration
	i := 0
	j := 0
	avgStep := 0.0
	for i < iters && time.Since(start) < timeOut {
		loopTimer := time.Now()
		if i%100 == 0 {
			fmt.Println(i, avgStep/100, j, alloc.score, timer)
		}
		j = 0
		better = false
		for !better && time.Since(start) < timeOut {
			add := matScale(logistic(j, 10000, 1), getRandomWeight(scen).fireAllocation)
			newAllocs := matAdd(alloc.fireAllocation, add, randomBool.Bool())
			normalise(newAllocs)
			score, _ := simulate(scen, newAllocs, 1)
			if score > alloc.score {
				alloc.fireAllocation = newAllocs
				alloc.score = score
				better = true
			}

			avgStep += float64(j)
			j++
		}
		allocs = append(allocs, copyAllocation(alloc))
		i++
		timer = time.Since(loopTimer)

	}
	return allocs, time.Since(start) > timeOut
}
