package main

import (
	"math/rand"
	"sync"
)

func pso(iters, samples int) []float64 {
	currentWeights := getRandomWeights(samples)
	bestWeight := getRandomWeight()
	bestWeights := getRandomWeights(samples)
	prevVelocity := getRandomWeights(samples)
	scenario := getTestingScenario()
	wg := sync.WaitGroup{}

	ownBestBias := 0.01
	allBestBias := 0.04
	for i := 0; i < iters; i++ {
		wg.Add(samples)
		for j := 0; j < samples; j++ {
			go func(thatLoop int) {
				prevVelocity[thatLoop], currentWeights[thatLoop] = indPso(currentWeights[thatLoop], bestWeights[thatLoop], bestWeight, prevVelocity[thatLoop], ownBestBias, allBestBias, i, iters, scenario)
				if currentWeights[thatLoop][15] > bestWeights[thatLoop][15] {
					copy(bestWeights[thatLoop], currentWeights[thatLoop])
					if currentWeights[thatLoop][15] > bestWeight[15] {
						bestWeight = bestWeights[thatLoop]
					}
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
	}
	return bestWeight
}

func indPso(currentWeight, pbestWeight, bestWeight, prevVelocity []float64, ownBias, allBias float64, iter, totalIters int, scenario scenario) (newV, newWeight []float64) {
	ownBestRand := rand.Float64()
	allBestRand := rand.Float64()
	inertia := float64(totalIters-iter) * 0.00001
	inertial := scale(inertia, prevVelocity)
	ownBestVec := scale(ownBias*ownBestRand, vecAdd(pbestWeight, currentWeight, false))
	allBestVec := scale(allBias*allBestRand, vecAdd(bestWeight, currentWeight, false))
	newV = vecAdd(inertial, vecAdd(ownBestVec, allBestVec, true), true)
	currentWeight = vecAdd(currentWeight, newV, true)
	currentWeight[15], _ = simulate(scenario, currentWeight, 1)
	return newV, currentWeight
}
