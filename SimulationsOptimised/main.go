package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	scen := readScenFromFile("scenJsons/3rdTest.json")
	singleScen(scen, "scen3")

}

func allScens() {
	scens := getAllScenariosFromFile("scenJsons")
	wg := sync.WaitGroup{}
	wg.Add(len(scens))
	for i, scen := range scens {
		go func(index int, scen scenario) {
			singleScen(scen, fmt.Sprint(index))
		}(i, scen)
	}
	wg.Wait()
}

func singleScen(scen scenario, scenName string) {
	//alls := ascend(10000, 0.0001, 0.00001, scen)
	//fmt.Println(alls[9999].Score)

	//geneticAlgo(iters, 1000, scen, 1)
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
				allocs := ascend(iters, 0.0001, 0.0001, scen)
				allocs := ascend(iters, 0.0001, 0.0001, scen)
	*/
	timeOut, _ := time.ParseDuration("10h")
	iters := 10000
	samples := 5 + 1
	for j := 1; j < samples; j++ {
		allocsAnneal := make([][]allocation, samples)
		allocsHillClimb := make([][]allocation, samples)
		allocsGrad := make([][]allocation, samples)
		allocsGen := make([][]allocation, samples)
		for i := 1; i < samples; i++ {
			allocsAnneal[i] = anneal(iters, samples, float64(j*10), scen)
			allocsHillClimb[i], _, _ = hillClimb(iters, scen, timeOut)
			allocsGrad[i] = ascend(iters, 0.0001, float64(j)*0.00001, scen)
			allocsGen[i] = geneticAlgo(iters, 100, scen, 0.1*float64(j))

		}
		writeManyToCSV(allocsAnneal, "results/anneal"+scenName+"Run"+fmt.Sprint(j)+".csv")
		writeManyToCSV(allocsHillClimb, "results/hillClimb"+scenName+"Run"+fmt.Sprint(j)+".csv")
		writeManyToCSV(allocsGrad, "results/gradient"+scenName+"Run"+fmt.Sprint(j)+".csv")
		writeManyToCSV(allocsGen, "results/gen"+scenName+"Run"+fmt.Sprint(j)+".csv")
	}
	//writeScorestoCSV(alls,"gradient3", false)
}
