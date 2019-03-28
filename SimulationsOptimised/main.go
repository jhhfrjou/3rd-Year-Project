package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	scen := readScenFromFile("scens/5thTest.json")
	singleScen(scen, "scen5")
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
	samples := 6
	wgJ := sync.WaitGroup{}
	wgJ.Add(4)
	go func() {
		for indexJ := 1; indexJ < samples; indexJ++ {
			allocsAnneal := make([][]allocation, 5)
			wgI := sync.WaitGroup{}
			wgI.Add(5)
			for i := 0; i < 5; i++ {
				go func(indexI int) {
					allocsAnneal[indexI] = anneal(iters, 100, float64(indexJ*10), scen)
					fmt.Println("finished anneal", indexJ, indexI)
					wgI.Done()
				}(i)
			}
			wgI.Wait()
			writeManyToCSV(allocsAnneal, "results/anneal"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
		}
		wgJ.Done()
	}()
	go func() {
		for indexJ := 1; indexJ < samples; indexJ++ {
			allocsGrad := make([][]allocation, 5)
			wgI := sync.WaitGroup{}
			wgI.Add(5)
			for i := 0; i < 5; i++ {
				go func(indexI int) {
					allocsGrad[indexI] = ascend(iters, 0.0001, float64(indexJ)*0.00001, scen)
					fmt.Println("finished ascent", indexJ, indexI)
					wgI.Done()
				}(i)
			}
			wgI.Wait()
			writeManyToCSV(allocsGrad, "results/gradient"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
		}
		wgJ.Done()
	}()
	go func() {
		for indexJ := 1; indexJ < samples; indexJ++ {
			allocsGen := make([][]allocation, 5)
			wgI := sync.WaitGroup{}
			wgI.Add(5)
			for i := 0; i < 5; i++ {
				go func(indexI int) {
					allocsGen[indexI] = geneticAlgo(iters, 1000, scen, 0.2*float64(indexJ))
					fmt.Println("finished ge", indexJ, indexI)
					wgI.Done()
				}(i)
			}
			wgI.Wait()
			writeManyToCSV(allocsGen, "results/gen"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
		}
		wgJ.Done()
	}()
	go func() {
		for indexJ := 1; indexJ < samples; indexJ++ {
			allocsPso := make([][]allocation, 5)
			wgI := sync.WaitGroup{}
			wgI.Add(5)
			for i := 0; i < 5; i++ {
				go func(indexI int) {
					allocsPso[indexI] = pso(iters, 1000, scen, 5, 2*float64(indexJ))
					fmt.Println("finished pso", indexJ, indexI)
					wgI.Done()
				}(i)
			}
			wgI.Wait()
			writeManyToCSV(allocsPso, "results/pso"+scenName+"Run"+fmt.Sprint(indexJ)+".csv")
		}
		wgJ.Done()
	}()
	wgJ.Add(1)
	go func() {
		allocsHillClimb := make([][]allocation, 5)
		wgHill := sync.WaitGroup{}
		wgHill.Add(5)
		for i := 0; i < 5; i++ {
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
}
