package main

import (
	"fmt"
	"sync"
)

func main() {
	scen := readScenFromFile("scenJsons/1stTest.json")
	singleScen(scen)

}

func allScens() {
	iters := 100
	scens := getAllScenariosFromFile("scenJsons")
	wg := sync.WaitGroup{}
	wg.Add(len(scens))
	for _, scen := range scens {
		go func(scen scenario) {
			allocs := geneticAlgo(iters, 1000, scen)
			prettyPrintAllocation(allocs[len(allocs)-1])
			wg.Done()
		}(scen)
	}
	wg.Wait()
}

func singleScen(scen scenario) {
	iters := 1000
	/*allocs := anneal(iters, 100, 50, scen)
	fmt.Println("Anneal")
	prettyPrintAllocation(allocs[len(allocs)-1])*/
	allocs := ascend(iters, 0.0001, 0.0001, scen)
	fmt.Println("Ascend")
	prettyPrintAllocation(allocs[len(allocs)-1])
	allocs = geneticAlgoS(iters, 10000, scen, allocs[len(allocs)-1])
	fmt.Println("Genetic")
	prettyPrintAllocation(allocs[len(allocs)-1])
	/*timeO, _ := time.ParseDuration("5m")
	allocs, _ = hillClimb(iters, scen, timeO)
	fmt.Println("Hill Climb")
	prettyPrintAllocation(allocs[len(allocs)-1])*/

}
