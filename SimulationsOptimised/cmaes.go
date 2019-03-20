package main

import "sort"

func cmaes(iters, samples int, scenario scenario) []allocation {
	allocs := make([]allocation, iters)
	samplePop := getRandomWeights(scenario, samples)
	mean := getRandomVeloc(scenario)
	sig := getRandomVeloc(scenario)
	c := getRandomVeloc(scenario)
	ps, pc := 0.0, 0.0
	for i := range allocs {
		for j := range samplePop {
			samplePop[j] = allocFromNormal(mean, sig, c, scenario)
		}
		sort.SliceStable(samplePop, func(i, j int) bool {
			return samplePop[i].Score > samplePop[j].Score
		})
		oldMean := mean
		mean = updateMean(samplePop)
		ps = updatePs(ps, sig, c, mean, oldMean)
		pc = updatePc(pc, ps, sig, mean, oldMean)
		c = updateC(c, pc, samplePop, oldMean, sig)
		sig = updateSig(sig, ps)
		allocs[i] = copyAllocation(samplePop[0])
	}
	return allocs
}

func allocFromNormal(m, sig, c [][]float64, scenario scenario) allocation {
	alloc := getRandomWeight(scenario)
	alloc.Score, _ = simulate(scenario, alloc.FireAllocation, 1)
	return getRandomWeight(scenario)
}

func updateMean(pop []allocation) (mean [][]float64) {
	mean = make([][]float64, len(pop[0].FireAllocation))
	for i := range mean {
		mean[i] = make([]float64, len(pop[0].FireAllocation[i]))
	}
	for _, v := range pop {
		for i, vec := range v.FireAllocation {
			for j, val := range vec {
				mean[i][j] += val
			}
		}
	}
	normalise(mean)
	return
}

func updatePs(ps float64, sig, c, mean, oldMean [][]float64) float64 {
	return 0.0
}

func updatePc(pc, ps float64, sig, mean, oldMean [][]float64) float64 {
	return 0.0
}

func updateC(c [][]float64, pc float64, samplePop []allocation, oldMean [][]float64, sig [][]float64) [][]float64 {
	return nil
}

func updateSig(sig [][]float64, ps float64) [][]float64 {
	return nil
}
