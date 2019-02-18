package main

import "fmt"

func main() {
	scores := make([][]allocation, 10)
	for i := range scores {
		scores[i] = ascend(1000, 0.00000001, 0.00001)
		fmt.Println("ascent complete", i)
	}
	writeManyToCSV(scores, "manyGradients.csv")
}
