package main

import (
	"fmt"
	"math"
	"math/rand"
)

func anneal(iters int, samples int, rate float64, scenario scenario) []allocation {
	var bestWeight allocation
	weight := getRandomWeight(scenario)
	weight.Score, _ = simulate(scenario, weight.FireAllocation, 1)
	weights := make([]allocation, iters)
	r := new()
	//var timer time.Duration
	for i := range weights {
		//start := time.Now()
		if i%100 == 0 {
			fmt.Println(i, weight.Score)
		}
		for j := 0; j < samples; j++ {
			deltaWeight := matScale(logistic(rand.Intn(32000), 16000, 50), getRandomWeight(scenario).FireAllocation)
			newWeight := matAdd(weight.FireAllocation, deltaWeight, r.Bool())
			normalise(newWeight)
			newScore, _ := simulate(scenario, newWeight, 1)
			if newScore > weight.Score {
				weight.Score = newScore
				weight.FireAllocation = newWeight
				if newScore > bestWeight.Score {
					bestWeight = copyAllocation(weight)
				}
			} else {
				deltaScore := newScore - weight.Score
				chance := math.Exp(deltaScore * float64(i) / rate)
				p := rand.Float64()
				if chance > p {
					weight.Score = newScore
					weight.FireAllocation = newWeight
				}
			}
		}
		//timer = time.Since(start)
		//fmt.Println(bestWeight.Score, "Time", timer)
		weights[i] = copyAllocation(bestWeight)
	}
	return weights
}
