package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func pso(iters, samples int, scenario scenario, ownBestBias, allBestBias float64) []allocation {
	currentWeights := getRandomWeights(scenario, samples)
	bestWeight := getRandomWeight(scenario)
	bestWeights := getRandomWeights(scenario, samples)
	prevVelocity := getRandomVelocs(scenario, samples)
	progression := make([]allocation, iters)
	wg := sync.WaitGroup{}

	for i := 0; i < iters; i++ {
		wg.Add(samples)
		for j := 0; j < samples; j++ {
			go func(sample int) {
				prevVelocity[sample], currentWeights[sample] = indPso(currentWeights[sample], bestWeights[sample], bestWeight, prevVelocity[sample], ownBestBias, allBestBias, i, iters, scenario)
				if currentWeights[sample].Score > bestWeights[sample].Score {
					bestWeights[sample] = copyAllocation(currentWeights[sample])
					if currentWeights[sample].Score > bestWeight.Score {
						bestWeight = bestWeights[sample]
					}
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
		progression[i] = copyAllocation(bestWeight)
		if i%100 == 0 {
			fmt.Println(i, progression[i].Score)
		}
	}
	return progression
}

func indPso(currentWeight, pBestWeight, bestWeight allocation, prevVelocity [][]float64, ownBias, allBias float64, iter, totalIters int, scenario scenario) ([][]float64, allocation) {
	ownBestRand := rand.Float64()
	allBestRand := rand.Float64()
	inertia := math.Exp(float64(iter)/float64(totalIters)) * 0.1
	inertial := matScale(inertia, prevVelocity)
	ownBestVec := matScale(ownBias*ownBestRand, matAdd(pBestWeight.FireAllocation, currentWeight.FireAllocation, false))
	allBestVec := matScale(allBias*allBestRand, matAdd(bestWeight.FireAllocation, currentWeight.FireAllocation, false))
	newV := matAdd(inertial, matAdd(ownBestVec, allBestVec, true), true)
	newFireAlloc := matAdd(currentWeight.FireAllocation, newV, true)
	normalise(newFireAlloc)
	newFireScore, _ := simulate(scenario, newFireAlloc, 1)
	return newV, allocation{newFireAlloc, newFireScore}
}
