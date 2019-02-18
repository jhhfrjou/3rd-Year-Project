package main

import "fmt"

func main() {
	scores := make([][]allocation,3)
	for i := range scores {
		scores[i] = ascend(100, 0.00000001, 0.00001)
		fmt.Println("ascent complete", i)
	}
	writeManyToCSV(scores, "manyGradients.csv")
}
