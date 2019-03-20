package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	scen := readScenFromFile("scenJsons/1stTest.json")
	singleScen(scen, "scen1")
	/*samples := 10
	allocsPso := make([][]allocation, samples*samples)
	wg := sync.WaitGroup{}
	wg.Add(samples * samples)
	for i := 0; i < samples; i++ {
		for j := 0; j < samples; j++ {
			go func(i, j int) {
				allocsPso[i*samples+j] = pso(3000, 100, scen, float64(i+1), float64(j+1))
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	writeManyToCSV(allocsPso, "results/testPsoscen3Test2.csv")*/

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
	timeOut, _ := time.ParseDuration("10h")
	iters := 10000
	samples := 5
	wgJ := sync.WaitGroup{}
	for j := 1; j < samples; j++ {
		wgJ.Add(4)
		go func(indexJ int) {
			allocsAnneal := make([][]allocation, samples)
			for i := 1; i < samples; i++ {
				allocsAnneal[i] = anneal(iters, samples, float64(indexJ*10), scen)
			}
			writeManyToCSV(allocsAnneal, "results/testanneal"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
			wgJ.Done()
		}(j)
		go func(indexJ int) {
			allocsGrad := make([][]allocation, samples)
			for i := 1; i < samples; i++ {
				allocsGrad[i] = ascend(iters, 0.0001, float64(indexJ)*0.00001, scen)
			}
			writeManyToCSV(allocsGrad, "results/testgradient"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
			wgJ.Done()
		}(j)
		go func(indexJ int) {
			allocsGen := make([][]allocation, samples)
			for i := 1; i < samples; i++ {
				allocsGen[i] = geneticAlgo(iters, 1000, scen, 0.2*float64(indexJ))
			}
			writeManyToCSV(allocsGen, "results/testgen"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
			wgJ.Done()
		}(j)
		go func(indexJ int) {
			allocsPso := make([][]allocation, samples)
			for i := 1; i < samples; i++ {
				allocsPso[i] = pso(iters, 1000, scen, 5, 2*float64(indexJ))
			}
			writeManyToCSV(allocsPso, "results/testpso"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
			wgJ.Done()
		}(j)
	}
	wgJ.Add(1)
	go func() {
		allocsHillClimb := make([][]allocation, samples)
		wgHill := sync.WaitGroup{}
		wgHill.Add(samples)
		for i := 0; i < samples; i++ {
			go func(indexI int) {
				allocsHillClimb[indexI], _, _ = hillClimb(iters, scen, timeOut)
				wgHill.Done()
			}(i)
		}
		wgHill.Wait()
		writeManyToCSV(allocsHillClimb, "results/testhillClimb"+scenName+"Run"+fmt.Sprint(1)+".csv")
		wgJ.Done()
	}()
	wgJ.Wait()
}
