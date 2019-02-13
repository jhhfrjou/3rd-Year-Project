package main

import (
	"math"
	"math/rand"
)

func anneal(iters int, samples int) allocation {
	scenario := getDevelopingScenario()
	weight := getRandomWeight(scenario)
	oldScore, _ := simulate(scenario, weight.fireAllocation, 1)
	r := new()
	for i := 0; i < iters; i++ {
		for j := 0; j < samples; j++ {
			deltaWeight := matScale(0.1, getRandomWeight(scenario).fireAllocation)
			newWeight := matAdd(weight.fireAllocation, deltaWeight, r.Bool())
			newScore, _ := simulate(scenario, newWeight, 1)
			if newScore > oldScore {
				oldScore = newScore
				weight = allocation{newWeight, newScore}
			} else {
				deltaScore := oldScore - newScore
				chance := math.Exp(-deltaScore / float64(iters-i))
				p := rand.Float64()
				if chance < p {
					oldScore = newScore
					weight = allocation{newWeight, newScore}
				}
			}
		}
	}
	return weight
}
