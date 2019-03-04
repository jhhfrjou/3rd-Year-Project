package main

import (
	"sync"
	"time"
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
	iters := 10000
	/*
		fmt.Println("Anneal")
		alloc := readAllocFromFile("allocations/1stTestAscend.json")
		fmt.Println("Hill Climb")
		prettyPrintAllocation(allocs[len(allocs)-1])
		writeScorestoCSV(allocs, "genetic.csv", true)
		alloc := getRandomWeight(scen)
		fmt.Println("Genetic")
		allocs[index] = anneal(iters, 100, float64(10*index), scen)
		allocs = geneticAlgo(iters, 1000, scen, 1)
		prettyPrintAllocation(allocs[len(allocs)-1])
	*/
	allocs := make([][]allocation, 1)
	//allocs := ascend(iters, 0.0001, 0.0001, scen)

	wg := sync.WaitGroup{}
	timeO, _ := time.ParseDuration("1h")
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(index int) {
			allocs, _, _ := hillClimb(iters, scen, timeO)
			prettyPrintAllocation(allocs[len(allocs)-1])
			wg.Done()
		}(i)
	}
	wg.Wait()
	writeManyToCSV(allocs, "hillClimb3.csv")

}
