package main

import (
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
	iters := 10000
	/*
		fmt.Println("Anneal")
		alloc := readAllocFromFile("allocations/1stTestAscend.json")
		fmt.Println("Hill Climb")
		prettyPrintAllocation(allocst86
		[len(allocs)-1])
		writeScorestoCSV(allocs, "genetic.csv", true)
		alloc := getRandomWeight(scen)
		fmt.Println("Genetic")
		allocs[index] = anneal(iters, 100, float64(10*index), scen)
		allocs = geneticAlgo(iters, 1000, scen, 1)
		prettyPrintAllocation(allocs[len(allocs)-1])
	*/
	//allocs := ascend(iters, 0.0001, 0.0001, scen)
	samples := 10
	for j := 0; j < samples; j++ {
		wg := sync.WaitGroup{}
		allocs := make([][]allocation, samples)
		wg.Add(samples)
		for i := 0; i < samples; i++ {
			go func(index int) {
				allocs[index] = anneal(iters, samples, float64(index*10), scen)
				prettyPrintAllocation(allocs[index][iters-1])
				wg.Done()
			}(j)
		}
		wg.Wait()
		writeManyToCSV(allocs, "hillClimb3Run"+string(j)+".csv")
	}
	//alls := ascend(10000, 0.0001, 0.00001, scen)
	//writeScorestoCSV(alls,"gradient3", false)
}
