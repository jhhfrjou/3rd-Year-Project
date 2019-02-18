package main

import (
	"fmt"
	"math"
	"math/rand"
)

func anneal(iters int, samples int, rate float64) []allocation {
	scenario := getBigDevelopingScenario()
	var bestWeight allocation
	weight := getRandomWeight(scenario)
	weights := make([]allocation, iters)
	oldScore, _ := simulate(scenario, weight.fireAllocation, 1)
	r := new()
	for i := range weights {
		for j := 0; j < samples; j++ {
			deltaWeight := matScale(0.1, getRandomWeight(scenario).fireAllocation)
			newWeight := matAdd(weight.fireAllocation, deltaWeight, r.Bool())
			normalise(newWeight)
			newScore, _ := simulate(scenario, newWeight, 1)
			if newScore > oldScore {
				oldScore = newScore
				weight = allocation{newWeight, newScore}
				bestWeight = weight
			} else {
				deltaScore := oldScore - newScore
				chance := math.Exp(deltaScore * rate / float64(iters-i))
				p := rand.Float64()
				if chance > p {
					oldScore = newScore
					weight = allocation{newWeight, newScore}
				}
			}
		}
		if i%100 == 0 {
			fmt.Println(i) 
		}
		weights[i] = copyAllocation(bestWeight)
	}
	return weights
}
