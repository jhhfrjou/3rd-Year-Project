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
	*/
	iters := 10000
	samples := 100
	timeOut, _ := time.ParseDuration("10h")
	for j := 0; j < samples; j++ {
		wg := sync.WaitGroup{}
		allocsAnneal := make([][]allocation, samples)
		allocsHillClimb := make([][]allocation, samples)
		allocsGrad := make([][]allocation, samples)
		wg.Add(samples)
		for i := 0; i < samples; i++ {
			go func(indexI, indexJ int) {
				allocsAnneal[indexI] = anneal(iters, samples, float64(indexJ*10), scen)
				allocsHillClimb[indexI], _, _ = hillClimb(iters, scen, timeOut)
				allocsGrad[indexI] = ascend(iters, 0.0001, float64(indexJ)*0.00001, scen)
				fmt.Println(allocsAnneal[indexI][iters-1].Score)
				wg.Done()
			}(i, j)
		}
		wg.Wait()
		writeManyToCSV(allocsAnneal, "results/anneal"+scenName+"Run"+fmt.Sprint(j)+".csv")
		writeManyToCSV(allocsHillClimb, "results/hillClimb"+scenName+"Run"+fmt.Sprint(j)+".csv")
		writeManyToCSV(allocsGrad, "results/gradient"+scenName+"Run"+fmt.Sprint(j)+".csv")
	}
	//writeScorestoCSV(alls,"gradient3", false)
}
