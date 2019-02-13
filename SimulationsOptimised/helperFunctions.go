package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type scenario struct {
	R  []float64
	B  []float64
	KR [][]float64
	KB [][]float64
}

type allocation struct {
	fireAllocation [][]float64
	score          float64
}

var bigScen scenario

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

func matScale(scalor float64, m [][]float64) [][]float64 {
	for i := range m {
		m[i] = scale(scalor, m[i])
	}
	return m
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

func matAdd(m1, m2 [][]float64, add bool) [][]float64 {
	m3 := make([][]float64, len(m1))
	for i := range m3 {
		m3[i] = vecAdd(m1[i], m2[i], add)
	}
	return m3

}

func sum(array []float64) (sum float64) {
	for _, v := range array {
		sum += v
	}
	return sum
}

func getRandomDefaultScenario() scenario {
	return getRandomScenario(4, 3)
}

func getRandomScenario(rSize, bSize int) scenario {
	kR := make([][]float64, bSize)
	kB := make([][]float64, rSize)
	r := make([]float64, rSize)
	b := make([]float64, bSize)
	for i := 0; i < rSize; i++ {
		r[i] = float64(rand.Intn(40) + 10)
		kB[i] = make([]float64, bSize)
		for j := 0; j < bSize; j++ {
			if i == 0 {
				kR[j] = make([]float64, rSize)
				b[j] = float64(rand.Intn(20) + 10)

			}
			kR[j][i] = rand.Float64()
			kB[i][j] = rand.Float64()
		}
	}
	return scenario{r, b, kR, kB}
}

func getDevelopingScenario() scenario {
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

func getBigDevelopingScenario() scenario {
	if len(bigScen.B) == 0 {
		bigScen = readScenFromFile("largeScen.json")
	}
	return bigScen
}

func writeScentoFile(scen scenario, file string) {
	bytes, _ := json.MarshalIndent(scen, "", " ")
	ioutil.WriteFile(file, bytes, 0644)
}

func readScenFromFile(file string) scenario {
	bytes, _ := ioutil.ReadFile(file)
	var scen scenario
	json.Unmarshal(bytes, &scen)
	return scen
}

func writeScorestoCSV(scores []allocation, fileName string) {
	output := make([][]string, len(scores))
	for i, v := range scores {

		output[i] = append([]string{fmt.Sprint(i), fmt.Sprint(v.score)}, allocsToString(v.fireAllocation)...)
	}
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	csvWrite := csv.NewWriter(file)
	csvWrite.WriteAll(output)
	csvWrite.Flush()
	file.Close()
}

func allocsToString(matrix [][]float64) []string {
	rows := len(matrix)
	cols := len(matrix[0])
	elements := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			elements[i*cols+j] = matrix[i][j]
		}
	}
	return []string{fmt.Sprint(rows), fmt.Sprint(cols), arrayToString(elements, ";")}
}

func arrayToString(a []float64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func copyScen(original scenario) scenario {
	r := make([]float64, len(original.R))
	b := make([]float64, len(original.B))
	copy(r, original.R)
	copy(b, original.B)
	kR := copyMatrix(original.KR)
	kB := copyMatrix(original.KB)
	return scenario{r, b, kR, kB}
}

func prettyPrintScen(scenario scenario) {
	fmt.Println("R", scenario.R)
	fmt.Println("B", scenario.B)
	fmt.Println("kR")
	for _, v := range scenario.KR {
		fmt.Println(v)
	}
	fmt.Println("kB")
	for _, v := range scenario.KB {
		fmt.Println(v)
	}
}

func copyAllocation(original allocation) allocation {
	fireDis := copyMatrix(original.fireAllocation)
	return allocation{fireDis, original.score}
}

func prettyPrintAllocation(allocation allocation) {
	fmt.Println("Score: ", allocation.score)
	fmt.Println("Allocation")
	for _, v := range allocation.fireAllocation {
		fmt.Println(v)
	}
}

func getRandomWeight(scenario scenario) allocation {
	weights := make([][]float64, len(scenario.KR[0]))
	for i := range weights {
		weights[i] = getRandomVector(len(scenario.KR))
	}
	return allocation{transpose(weights), -math.MaxFloat64}
}

func getRandomWeights(scenario scenario, sampleSize int) []allocation {
	weights := make([]allocation, sampleSize)
	for i := range weights {
		weights[i] = getRandomWeight(scenario)
	}
	return weights
}

func getRandomVector(num int) []float64 {
	samples := []float64{}
	for i := 0; i < num; i++ {
		samples = append(samples, rand.Float64())
	}
	sumed := sum(samples)
	return scale(1.0/sumed, samples)
}

func transpose(original [][]float64) [][]float64 {
	transposed := make([][]float64, len(original[0]))
	for j, row := range original {
		for i, val := range row {
			if j == 0 {
				transposed[i] = make([]float64, len(original))
			}
			transposed[i][j] = val
		}
	}
	return transposed
}

func normalise(original [][]float64) {
	colSums := make([]float64, len(original[0]))
	wg := sync.WaitGroup{}
	wg.Add(len(colSums))
	for i := range colSums {
		go func(index int) {
			colSums[index] = colSum(original, index)
			if colSums[index] != 1 {
				colDiv(original, index, colSums[index])
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func colDiv(mat [][]float64, col int, divFactor float64) {
	wg := sync.WaitGroup{}
	wg.Add(len(mat))
	for i := range mat {
		go func(index int) {
			mat[index][col] = mat[index][col] / divFactor
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func colSum(mat [][]float64, col int) float64 {
	sum := 0.0
	for i := range mat {
		sum += mat[i][col]
	}
	return sum
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

func copyMatrix(original [][]float64) [][]float64 {
	copied := make([][]float64, len(original))
	for i := range original {
		copied[i] = make([]float64, len(original[i]))
		copy(copied[i], original[i])
	}
	return copied
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

func diff(allocation [][]float64, policyCode int, scen scenario, delta float64) [][]float64 {
	score, _ := simulate(scen, allocation, policyCode)
	diffs := make([][]float64, len(allocation))
	wg := sync.WaitGroup{}
	wg.Add(len(allocation) * len(allocation[0]))
	for i := range allocation {
		diffs[i] = make([]float64, len(allocation[i]))
		for j := range allocation[i] {
			go func(indexI, indexJ int) {
				diffs[indexI][indexJ] = diffWeightScore(indexI, indexJ, policyCode, delta, score, allocation, scen)
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()

	return diffs
}

func diffWeightScore(indexX, indexY, policyCode int, delta, origScore float64, original [][]float64, scen scenario) float64 {
	differential := copyMatrix(original)
	differential[indexX][indexY] = delta + differential[indexX][indexY]
	score, _ := simulate(scen, differential, policyCode)
	diff := score - origScore
	return diff / delta
}
