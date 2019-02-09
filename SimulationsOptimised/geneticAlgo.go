package main

import (
	"math/rand"
	"sort"
	"sync"
)

var randomBool = new()

func geneticAlgo(iters, samples int) []float64 {
	gen := getRandomWeights(samples)
	getScores(gen, getTestingScenario())
	bestWeight := gen[0]
	for i := 0; i < iters; i++ {
		sort.Slice(gen, func(i, j int) bool {
			return gen[i][15] > gen[j][15]
		})
		if bestWeight[15] < gen[0][15] {
			bestWeight = gen[0]
		}
		gen = getNextGen(gen, 0.3)
	}
	return bestWeight
}

func getScores(samples [][]float64, scenario scenario) float64 {
	wg := sync.WaitGroup{}
	wg.Add(len(samples))
	for i := 0; i < len(samples); i++ {
		go func(index int) {
			samples[index][15], _ = simulate(scenario, samples[index], 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	sum := 0.0
	for _, v := range samples {
		sum += v[15]
	}
	return sum

}

func crossOver(p1, p2 []float64) []float64 {
	child := make([]float64, len(p1))
	for i := range p1 {
		if randomBool.Bool() {
			child[i] = p1[i]
		} else {
			child[i] = p2[i]
		}
	}
	return child
}

func mutate(w []float64, mutateFactor float64) {
	for i := range w {
		if rand.Float64() < mutateFactor {
			w[i] = rand.Float64()
		}
	}
}

func getNextGen(currentGen [][]float64, mutateFactor float64) [][]float64 {
	numSample := len(currentGen)
	nextGen := make([][]float64, numSample)
	wg := sync.WaitGroup{}
	wg.Add(numSample)
	for i := 0; i < numSample; i++ {
		go func(index int) {
			if index < 10 {
				nextGen[index] = currentGen[index]
			} else {
				i1 := rand.Intn(numSample / 2)
				i2 := rand.Intn(numSample / 2)
				child := crossOver(currentGen[i1], currentGen[i2])
				mutate(child, mutateFactor)
				nextGen[index] = child
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return nextGen
}
