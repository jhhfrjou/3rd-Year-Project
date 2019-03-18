package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	scen := readScenFromFile("scenJsons/1stTest.json")
	singleScen(scen, "scen1")

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
	iters := 10
	samples := 3
	wgJ := sync.WaitGroup{}
	for j := 1; j < samples; j++ {
		wgJ.Add(3)
		go func(indexJ int) {
			allocsAnneal := make([][]allocation, samples)
			for i := 1; i < samples; i++ {
				allocsAnneal[i] = anneal(iters, samples, float64(indexJ*10), scen)
			}
			writeManyToCSV(allocsAnneal, "results/anneal"+scenName+"Run"+fmt.Sprint(indexJ)+"test.csv")
			wgJ.Done()
		}(j)
		go func(indexJ int) {
			allocsGrad := make([][]allocation, samples)
			for i := 1; i < samples; i++ {
				allocsGrad[i] = ascend(iters, 0.0001, float64(indexJ)*0.00001, scen)
			}
			writeManyToCSV(allocsGrad, "results/gradient"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
			wgJ.Done()
		}(j)
		go func(indexJ int) {
			allocsGen := make([][]allocation, samples)
			for i := 1; i < samples; i++ {
				allocsGen[i] = geneticAlgo(iters, 100, scen, 0.1*float64(indexJ))
			}
			writeManyToCSV(allocsGen, "results/gen"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
			wgJ.Done()
		}(j)
	}
	wgJ.Add(1)
	go func() {
		allocsHillClimb := make([][]allocation, samples)
		wgHill := sync.WaitGroup{}
		wgHill.Add(samples)
		for i := 1; i < samples; i++ {
			go func(indexI int) {
				allocsHillClimb[indexI], _, _ = hillClimb(iters, scen, timeOut)
				wgHill.Done()
			}(i)
		}
		wgHill.Wait()
		writeManyToCSV(allocsHillClimb, "results/hillClimb"+scenName+"Run"+fmt.Sprint(1)+".csv")
		wgJ.Done()
	}()
	wgJ.Wait()
	/*

		for j := 1; j < samples; j++ {
			allocsAnneal := make([][]allocation, samples)
			allocsHillClimb := make([][]allocation, samples)
			allocsGrad := make([][]allocation, samples)
			allocsGen := make([][]allocation, samples)
			wg := sync.WaitGroup{}
			for i := 1; i < samples; i++ {
				wg.Add(4)
				go func(indexI, indexJ int) {
					allocsAnneal[indexI] = anneal(iters, samples, float64(indexJ*10), scen)
					wg.Done()
				}(i, j)
				go func(indexI, indexJ int) {
					allocsHillClimb[indexI], _, _ = hillClimb(iters, scen, timeOut)
					wg.Done()
				}(i, j)
				go func(indexI, indexJ int) {
					allocsGrad[indexI] = ascend(iters, 0.0001, float64(indexJ)*0.00001, scen)
					wg.Done()
				}(i, j)
				go func(indexI, indexJ int) {
					wg.Done()
				}(i, j)

			}
			wg.Wait()
			writeManyToCSV(allocsAnneal, "results/anneal"+scenName+"Run"+fmt.Sprint(j)+".csv")
		}
	*/
	//writeScorestoCSV(alls,"gradient3", false)
}
