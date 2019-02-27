package main

import (
	"fmt"
	"math"
	"math/rand"
)

func anneal(iters int, samples int, rate float64, scenario scenario) []allocation {
	var bestWeight allocation
	weight := getRandomWeight(scenario)
	weight.score, _ = simulate(scenario, weight.fireAllocation, 1)
	weights := make([]allocation, iters)
	r := new()
	//var timer time.Duration
	for i := range weights {
		//start := time.Now()
		if i%100 == 0 {
			fmt.Println(i, weight.score)
		}
		for j := 0; j < samples; j++ {
			deltaWeight := matScale(0.01, getRandomWeight(scenario).fireAllocation)
			newWeight := matAdd(weight.fireAllocation, deltaWeight, r.Bool())
			normalise(newWeight)
			newScore, _ := simulate(scenario, newWeight, 1)
			if newScore > weight.score {
				weight.score = newScore
				weight.fireAllocation = newWeight
				if newScore > bestWeight.score {
					bestWeight = copyAllocation(weight)
				}
			} else {
				deltaScore := newScore - weight.score
				chance := math.Exp(deltaScore * rate / float64(iters-i))
				p := rand.Float64()
				if chance > p {
					weight.score = newScore
					weight.fireAllocation = newWeight
				}
			}
		}
		//timer = time.Since(start)
		//fmt.Println(bestWeight.score, "Time", timer)
		weights[i] = copyAllocation(bestWeight)
	}
	return weights
}
