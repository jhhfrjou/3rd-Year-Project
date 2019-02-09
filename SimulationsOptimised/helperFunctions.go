package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type scenario struct {
	r  []float64
	b  []float64
	kR [][]float64
	kB [][]float64
}

func scale(scalor float64, vector []float64) []float64 {
	wg := sync.WaitGroup{}
	wg.Add(len(vector))
	for i := range vector {
		go func(index int) {
			vector[index] = scalor * vector[index]
			wg.Done()
		}(i)
	}
	wg.Wait()
	return vector
}

func vecAdd(v1, v2 []float64, add bool) []float64 {
	v3 := make([]float64, len(v1))
	wg := sync.WaitGroup{}
	wg.Add(len(v1))
	if add {
		for i := range v1 {
			go func(index int) {
				v3[index] = v1[index] + v2[index]
				wg.Done()
			}(i)
		}
	} else {
		for i := range v1 {
			go func(index int) {
				v3[index] = v1[index] - v2[index]
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
	return v3
}

func copyScen(original scenario) scenario {
	r := make([]float64, len(original.r))
	b := make([]float64, len(original.b))
	copy(r, original.r)
	copy(b, original.b)
	kR := make([][]float64, len(original.kR))
	for i := range original.kR {
		kR[i] = make([]float64, len(original.kR[i]))
		copy(kR[i], original.kR[i])
	}
	kB := make([][]float64, len(original.kB))
	for i := range original.kB {
		kB[i] = make([]float64, len(original.kB[i]))
		copy(kB[i], original.kB[i])
	}
	return scenario{r, b, kR, kB}
}

func sum(array []float64) (sum float64) {
	for _, v := range array {
		sum += v
	}
	return sum
}

func getRandomScenario() scenario {
	rSize := 2 //rand.Intn(4)+2
	bSize := 2 //rand.Intn(4) + 2
	kR := make([][]float64, bSize)
	kB := make([][]float64, rSize)
	r := make([]float64, rSize)
	b := make([]float64, bSize)
	for i := 0; i < rSize; i++ {
		r[i] = float64(rand.Intn(40) + 10)
		kR[i] = make([]float64, rSize)
		for j := 0; j < bSize; j++ {
			if i == 0 {
				b[j] = float64(rand.Intn(20) + 10)
				kB[j] = make([]float64, bSize)
			}
			kR[i][j] = rand.Float64()
			kB[j][i] = rand.Float64()
		}
	}
	return scenario{r, b, kR, kB}
}

func getTestingScenario() scenario {
	kR := [][]float64{
		{0.023, 0.08},
		{0.023, 0.05},
		{0.0024, 0.03},
	}
	kB := [][]float64{
		{0.023, 0.08, 0.01},
		{0.023, 0.05, 0.05},
	}
	R := []float64{300, 5}
	B := []float64{250, 10, 20}
	return scenario{R, B, kR, kB}
}

func getRandomWeight() []float64 {
	weights := make([]float64, 16)
	weights[15] = math.SmallestNonzeroFloat64
	for i := 0; i < 15; i++ {
		weights[i] = rand.Float64()
	}
	return weights
}

func getRandomWeights(num int) [][]float64 {
	samples := [][]float64{}
	for i := 0; i < num; i++ {
		samples = append(samples, getRandomWeight())
	}
	return samples
}

//Adapted from https://github.com/mxschmitt/golang-combinations
func combs(inputs []float64) (combs []float64) {
	length := uint(len(inputs))
	for combsBits := 1; combsBits < (1 << length); combsBits++ {
		comb := 1.0

		for factor := uint(0); factor < length; factor++ {
			if (combsBits>>factor)&1 == 1 {
				comb *= inputs[factor]
			}
		}
		combs = append(combs, comb)
	}
	return combs
}

type boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *boolgen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func new() *boolgen {
	return &boolgen{src: rand.NewSource(time.Now().UnixNano())}
}

func staticWeight(number float64) []float64 {
	weight := make([]float64, 16)
	for i := 0; i < 16; i++ {
		weight[i] = number
	}
	return weight
}

func diff(weights []float64, policyCode int, scen scenario) []float64 {
	score, _ := simulate(scen, weights, policyCode)
	diffs := make([]float64, len(weights))
	wg := sync.WaitGroup{}
	wg.Add(len(weights))
	for i := range weights {
		go func(index int) {
			diffs[index] = diffWeightScore(index, policyCode, 0.000000001, score, weights, scen)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return diffs
}

func diffWeightScore(index, policyCode int, delta, origScore float64, original []float64, scen scenario) float64 {
	weights := make([]float64, len(original))
	copy(weights, original)

	weights[index] = delta + weights[index]
	score, _ := simulate(scen, weights, policyCode)
	diff := score - origScore
	return diff / delta
}
