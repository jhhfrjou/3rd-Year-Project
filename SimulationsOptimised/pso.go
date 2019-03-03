package main

import (
	"math"
	"math/rand"
	"sync"
)

func pso(iters, samples int, scenario scenario) []allocation {
	currentWeights := getRandomWeights(scenario, samples)
	bestWeight := getRandomWeight(scenario)
	bestWeights := getRandomWeights(scenario, samples)
	prevVelocity := getRandomWeights(scenario, samples)
	progression := make([]allocation, iters)
	wg := sync.WaitGroup{}

	ownBestBias := 0.01
	allBestBias := 0.04
	for i := 0; i < iters; i++ {
		wg.Add(samples)
		for j := 0; j < samples; j++ {
			go func(thatLoop int) {
				prevVelocity[thatLoop], currentWeights[thatLoop] = indPso(currentWeights[thatLoop], bestWeights[thatLoop], bestWeight, prevVelocity[thatLoop], ownBestBias, allBestBias, i, iters, scenario)
				if currentWeights[thatLoop].Score > bestWeights[thatLoop].Score {
					bestWeights[thatLoop] = copyAllocation(currentWeights[thatLoop])
					if currentWeights[thatLoop].Score > bestWeight.Score {
						bestWeight = bestWeights[thatLoop]
					}
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
		progression[i] = copyAllocation(bestWeight)
	}
	return progression
}

func indPso(currentWeight, pbestWeight, bestWeight, prevVelocity allocation, ownBias, allBias float64, iter, totalIters int, scenario scenario) (allocation, allocation) {
	ownBestRand := rand.Float64()
	allBestRand := rand.Float64()
	inertia := float64(totalIters-iter) * 0.00001
	inertial := matScale(inertia, prevVelocity.FireAllocation)
	ownBestVec := matScale(ownBias*ownBestRand, matAdd(pbestWeight.FireAllocation, currentWeight.FireAllocation, false))
	allBestVec := matScale(allBias*allBestRand, matAdd(bestWeight.FireAllocation, currentWeight.FireAllocation, false))
	newV := matAdd(inertial, matAdd(ownBestVec, allBestVec, true), true)
	newFireAlloc := matAdd(currentWeight.FireAllocation, newV, true)
	normalise(newFireAlloc)
	newFireScore, _ := simulate(scenario, newFireAlloc, 1)
	return allocation{newV, -math.MaxFloat64}, allocation{newFireAlloc, newFireScore}
}
