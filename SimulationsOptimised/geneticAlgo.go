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

func geneticAlgo(iters, samples int, scenario scenario, mutateFactor float64) []allocation {
	rand.Seed(time.Now().Unix())
	gen := getRandomWeights(scenario, samples)
	return geneticAlgoAllocs(iters, samples, scenario, gen, mutateFactor)
}

func geneticAlgoS(iters, samples int, scenario scenario, alloc allocation, mutateFactor float64) []allocation {
	gen := make([]allocation, samples)
	gen[0] = copyAllocation(alloc)
	for i := 1; i < samples; i++ {
		gen[i] = copyAllocation(alloc)
		mutate(gen[i].FireAllocation, mutateFactor, logistic(rand.Intn(2000), 1000, 50), scenario)
		gen[i].Score = -math.MaxFloat64
	}
	return geneticAlgoAllocs(iters, samples, scenario, gen, mutateFactor)

}

func geneticAlgoAllocs(iters, samples int, scenario scenario, gen []allocation, mutateFactor float64) []allocation {
	bestWeights := make([]allocation, iters)
	getScores(gen, scenario)
	fmt.Println("initial Scores")
	bestWeight := gen[0]
	for i := 0; i < iters; i++ {
		sort.SliceStable(gen, func(i, j int) bool {
			return gen[i].Score > gen[j].Score
		})
		if i%10 == 0 {
			j := rand.Intn(samples)
			fmt.Println("Iter: ", i, " Best: ", gen[0].Score, "Rand: ", j, gen[j].Score, "Worst", gen[samples-1].Score)
		}
		if bestWeight.Score < gen[0].Score {
			bestWeight = gen[0]
		}
		bestWeights[i] = copyAllocation(bestWeight)
		gen = getNextGen(gen, mutateFactor, scenario)
	}
	return bestWeights
}

func getNextGen(currentGen []allocation, mutateFactor float64, scenario scenario) []allocation {
	numSample := len(currentGen)
	nextGen := make([]allocation, numSample)
	wg := sync.WaitGroup{}
	wg.Add(numSample - 1)
	nextGen[0] = copyAllocation(currentGen[0])
	for i := 1; i < numSample; i++ {
		go func(index int) {
			i1, i2 := rand.Intn(int(float64(numSample)*0.7)), rand.Intn(int(float64(numSample)*0.7))
			child := crossOver(currentGen[i1].FireAllocation, currentGen[i2].FireAllocation)
			child = mutate(child, mutateFactor, logistic(rand.Intn(2000), 1000, 50), scenario)
			nextGen[index] = allocation{child, -math.MaxFloat64}
			nextGen[index].Score, _ = simulate(scenario, child, 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return nextGen
}

func getScores(samples []allocation, scenario scenario) {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(len(samples))

	overCount := 0
	for i := 0; i < len(samples); i++ {
		go func(index int) {
			var count int
			samples[index].Score, count = simulate(scenario, samples[index].FireAllocation, 1)
			if count > 50000 {
				overCount++
				fmt.Println(samples[index].Score)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("OverCounts :", overCount)
	fmt.Println(time.Since(start))
}

func crossOver(p1, p2 [][]float64) [][]float64 {
	p1T := transpose(p1)
	p2T := transpose(p2)
	child := make([][]float64, len(p1T))
	for i := range p1T {
		child[i] = make([]float64, len(p1T[i]))
		if randomBool.Bool() {
			copy(child[i], p1T[i])
		} else {
			copy(child[i], p2T[i])
		}
	}
	return transpose(child)
}

func mutate(w [][]float64, mutateFactor, scale float64, scen scenario) [][]float64 {
	if rand.Float64() < mutateFactor {
		add := matScale(scale, getRandomWeight(scen).FireAllocation)
		w = matAdd(w, add, randomBool.Bool())
	}
	normalise(w)
	return w
}
