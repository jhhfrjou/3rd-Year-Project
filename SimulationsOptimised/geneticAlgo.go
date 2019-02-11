package main

import (
	"math"
	"math/rand"
	"sort"
	"sync"
)

var randomBool = new()

func geneticAlgo(iters, samples int) allocation {
	scenario := getBigDevelopingScenario()
	gen := getRandomWeights(scenario, samples)
	getScores(gen, scenario)
	bestWeight := gen[0]
	for i := 0; i < iters; i++ {
		sort.Slice(gen, func(i, j int) bool {
			return gen[i].score > gen[j].score
		})
		if bestWeight.score < gen[0].score {
			bestWeight = gen[0]
		}
		gen = getNextGen(gen, 0.3)
	}
	return bestWeight
}

func getNextGen(currentGen []allocation, mutateFactor float64) []allocation {
	numSample := len(currentGen)
	nextGen := make([]allocation, numSample)
	wg := sync.WaitGroup{}
	wg.Add(numSample)
	for i := 0; i < numSample; i++ {
		go func(index int) {
			if index < 10 {
				nextGen[index] = currentGen[index]
			} else {
				i1 := rand.Intn(numSample / 2)
				i2 := rand.Intn(numSample / 2)
				child := allocation{crossOver(currentGen[i1].fireAllocation, currentGen[i2].fireAllocation), -math.MaxFloat64}
				mutate(child.fireAllocation, mutateFactor)
				nextGen[index] = child
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return nextGen
}

func getScores(samples []allocation, scenario scenario) float64 {
	wg := sync.WaitGroup{}
	wg.Add(len(samples))
	for i := 0; i < len(samples); i++ {
		go func(index int) {
			samples[index].score, _ = simulate(scenario, samples[index].fireAllocation, 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	sum := 0.0
	for _, v := range samples {
		sum += v.score
	}
	return sum

}

func crossOver(p1, p2 [][]float64) [][]float64 {
	child := make([][]float64, len(p1))
	for i := range p1 {
		child[i] = make([]float64,len(p1[i]))
		for j := range p1[i] {
			if randomBool.Bool() {
				child[i][j] = p1[i][j]
			} else {
				child[i][j] = p2[i][j]
			}
		}
	}
	return child
}

func mutate(w [][]float64, mutateFactor float64) {
	for i := range w {
		for j := range w[i] {
			if rand.Float64() < mutateFactor {
				w[i][j] = rand.Float64()
			}
		}
	}
}
