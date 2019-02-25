package main

import (
	"math"
	"math/rand"
	"sort"
	"sync"
)

var randomBool = new()

func geneticAlgo(iters, samples int, scenario scenario) []allocation {
	bestWeights := make([]allocation, iters)
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
		bestWeights[i] = copyAllocation(bestWeight)
		gen = getNextGen(gen, 0.8)
	}
	return bestWeights
}

func getNextGen(currentGen []allocation, mutateFactor float64) []allocation {
	numSample := len(currentGen)
	nextGen := make([]allocation, numSample)
	wg := sync.WaitGroup{}
	wg.Add(numSample)
	for i := 0; i < numSample; i++ {
		go func(index int) {
			if index < len(currentGen) /10 {
				nextGen[index] = currentGen[index]
			} else {
				i1 := rand.Intn(numSample / 2)
				i2 := rand.Intn(numSample / 2)
				child := crossOver(currentGen[i1].fireAllocation, currentGen[i2].fireAllocation)
				child = mutate(child, mutateFactor)
				nextGen[index] = allocation{child, -math.MaxFloat64}
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
	child := make([][]float64, len(p1[0]))
	sumed := 0.0
	for j := range p1[0] {
		child[j] = make([]float64, len(p1))
		for i := range p1 {
			if randomBool.Bool() {
				child[j][i] = p1[i][j]
			} else {
				child[j][i] = p2[i][j]
			}
		}
		sumed = sum(child[j])
		child[j] = scale(1.0/sumed, child[j])
	}
	return transpose(child)
}

func mutate(w [][]float64, mutateFactor float64) [][]float64 {
	transposed := transpose(w)
	sumed := 0.0
	for i, vec := range transposed {
		for j := range vec {
			if rand.Float64() < mutateFactor {
				if rand.Float64() < 0.3 {
					transposed[i][j] = rand.Float64()
				} else {
					transposed[i][j] = 0
				}

			}
		}
		sumed = sum(transposed[i])
		transposed[i] = scale(1.0/sumed, transposed[i])
	}
	return transpose(transposed)
}
