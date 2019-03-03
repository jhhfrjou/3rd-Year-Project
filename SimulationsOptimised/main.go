package main

import (
	"fmt"
	"sync"
)

func main() {
	scen := readScenFromFile("scenJsons/3rdTest.json")
	singleScen(scen)

}

func allScens() {
	iters := 1000
	scens := getAllScenariosFromFile("scenJsons")
	wg := sync.WaitGroup{}
	wg.Add(len(scens))
	for _, scen := range scens {
		go func(scen scenario) {
			allocs := geneticAlgo(iters, 1000, scen, 0.1)
			prettyPrintAllocation(allocs[len(allocs)-1])
			wg.Done()
		}(scen)
	}
	wg.Wait()
}

func singleScen(scen scenario) {
	iters := 100
	/*
		allocs, _ := hillClimb(iters, scen, timeO)
		timeO, _ := time.ParseDuration("5m")
		fmt.Println("Anneal")
		prettyPrintAllocation(allocs[len(allocs)-1])
		alloc := readAllocFromFile("allocations/1stTestAscend.json")
		fmt.Println("Hill Climb")
		prettyPrintAllocation(allocs[len(allocs)-1])
		writeScorestoCSV(allocs, "genetic.csv", true)
		alloc := getRandomWeight(scen)
		fmt.Println("Genetic")
		allocs[index] = anneal(iters, 100, float64(10*index), scen)
	*/
	allocs := make([][]allocation, 10)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(index int) {
			allocs[index] = geneticAlgo(iters, 1000, scen, 0.1*float64(index))
			fmt.Println(index, allocs[index][iters-1].Score)
			wg.Done()
		}(i)
	}
	wg.Wait()
	writeManyToCSV(allocs, "MutationFactor3.csv")

}
