package main

import (
	"fmt"
)

func main() {
	bestWeight := geneticAlgo(100, 100)
	for _, v := range bestWeight {
		fmt.Println(v)
	}
}