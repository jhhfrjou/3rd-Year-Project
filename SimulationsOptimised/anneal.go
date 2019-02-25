package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func anneal(iters int, samples int, rate float64, scenario scenario) []allocation {
	var bestWeight allocation
	weight := getRandomWeight(scenario)
	weights := make([]allocation, iters)
	oldScore, _ := simulate(scenario, weight.fireAllocation, 1)
	r := new()
	var rec time.Duration
	for i := range weights {
		start := time.Now()
		for j := 0; j < samples; j++ {
			deltaWeight := matScale(0.1, getRandomWeight(scenario).fireAllocation)
			newWeight := matAdd(weight.fireAllocation, deltaWeight, r.Bool())
			normalise(newWeight)
			newScore, _ := simulate(scenario, newWeight, 1)
			if newScore > oldScore {
				oldScore = newScore
				weight = allocation{newWeight, newScore}
				if newScore > bestWeight.score {
					bestWeight = copyAllocation(weight)
				}
			} else {
				deltaScore := newScore - oldScore
				chance := math.Exp(deltaScore * rate / float64(iters-i))
				p := rand.Float64()
				if chance > p {
					oldScore = newScore
					weight = allocation{newWeight, newScore}
				}
			}
		}
		rec = time.Since(start)
		fmt.Println(rec)
		fmt.Println(bestWeight.score)
		weights[i] = copyAllocation(bestWeight)
	}
	return weights
}
