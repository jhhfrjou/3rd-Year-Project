package main

import (
	"math"
	"math/rand"
)

func anneal(iters int, samples int) ([]float64, float64) {
	weight := getRandomWeight()
	scenario := getTestingScenario()
	oldScore, _ := simulate(scenario, weight, 1)
	r := new()
	for i := 0; i < iters; i++ {
		for j := 0; j < samples; j++ {
			deltaWeight := scale(0.1, getRandomWeight())
			newWeight := vecAdd(weight, deltaWeight, r.Bool())
			newScore, _ := simulate(scenario, newWeight, 1)
			if newScore > oldScore {
				oldScore = newScore
				weight = newWeight
			} else {
				chance := math.Exp(-sum(deltaWeight) / float64(iters-i))
				p := rand.Float64()
				if chance < p {
					oldScore = newScore
					weight = newWeight
				}
			}
		}
	}
	return weight, oldScore
}
