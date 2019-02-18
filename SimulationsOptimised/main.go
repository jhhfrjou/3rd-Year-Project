package main

import "fmt"

func main() {
	scores := make([][]allocation, 10)
	for i := range scores {
		scores[i] = anneal(1000, 100, 0.1)
		fmt.Println("ascent complete", i)
	}
	writeManyToCSV(scores, "man.csv")
}
