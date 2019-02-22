package main

import (
	"fmt"
	"sync"
)

func main() {
	scores := make([][]allocation, 40)
	wg := sync.WaitGroup{}
	wg.Add(len(scores))
	for i := range scores {
		go func(index int) {
			scores[index] = anneal(10000, 1000, 0.1*float64(index+1))
			fmt.Println("ascent complete", index)
			wg.Done()
		}(i)
	}
	wg.Wait()
	//writeManyToCSV(scores, "annneals.csv")

}
