package main

import (
	"fmt"
	"time"
)

func main() {
	scen := readScenFromFile("scenJsons/4thTest.json")
	iters := 10000
	timeOut, _ := time.ParseDuration("20m")
	allocs, timeOuted := hillClimb(iters, scen, timeOut)
	fmt.Println(timeOuted)
	prettyPrintAllocation(allocs[len(allocs)-1])
}
