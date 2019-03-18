package main

import "fmt"

func cames(iters, samples int, scenario scenario) []allocation {
	allocs := make([]allocation, iters)
	sample := getRandomWeights(scenario, samples)
	fmt.Println(sample)
	return allocs
}
