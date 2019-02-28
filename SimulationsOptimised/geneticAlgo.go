package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var randomBool = new()

func geneticAlgo(iters, samples int, scenario scenario) []allocation {
	gen := getRandomWeights(scenario, samples)
	return geneticAlgoAllocs(iters, samples, scenario, gen)
}

func geneticAlgoS(iters, samples int, scenario scenario, alloc allocation) []allocation {
	gen := make([]allocation, samples)
	gen[0] = alloc
	for i := 1; i < samples; i++ {
		gen[i] = allocation{mutate(alloc.fireAllocation, 0.1), -math.MaxFloat64}
	}
	return geneticAlgoAllocs(iters, samples, scenario, gen)

}

func geneticAlgoAllocs(iters, samples int, scenario scenario, gen []allocation) []allocation {
	bestWeights := make([]allocation, iters)
	start := time.Now()
	getScores(gen, scenario)
	fmt.Println(time.Since(start))
	fmt.Println("initial Scores")
	bestWeight := gen[0]
	for i := 0; i < iters; i++ {
		if i%100 == 0 {
			fmt.Println(i, bestWeight.score)
		}
		sort.Slice(gen, func(i, j int) bool {
			return gen[i].score > gen[j].score
		})
		if bestWeight.score < gen[0].score {
			bestWeight = gen[0]
		}
		bestWeights[i] = copyAllocation(bestWeight)
		gen = getNextGen(gen, 0.8, scenario)
	}
	return bestWeights
}

func getNextGen(currentGen []allocation, mutateFactor float64, scenario scenario) []allocation {
	numSample := len(currentGen)
	nextGen := make([]allocation, numSample)
	wg := sync.WaitGroup{}
	wg.Add(numSample)
	for i := 0; i < numSample; i++ {
		go func(index int) {
			if index < len(currentGen)/10 {
				nextGen[index] = currentGen[index]
			} else if index > len(currentGen)/2 {
				nextGen[index] = getRandomWeight(scenario)
			} else {
				i1 := rand.Intn(index)
				i2 := rand.Intn(index)
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
	child := make([][]float64, len(p1))
	for i := range p1 {
		child[i] = make([]float64, len(p1[0]))
		for j := range p1 {
			if randomBool.Bool() {
				child[i][j] = p1[i][j]
			} else {
				child[i][j] = p2[i][j]
			}
		}
	}
	normalise(child)
	return child
}

func mutate(w [][]float64, mutateFactor float64) [][]float64 {
	for i, vec := range w {
		for j := range vec {
			if rand.Float64() < mutateFactor {
				if rand.Float64() < 0.5 {
					w[i][j] = rand.Float64()
				} else {
					if randomBool.Bool() {
						w[i][j] = 0
					} else {
						w[i][j] = 1
					}
				}

			}
		}
	}
	normalise(w)
	return w
}
