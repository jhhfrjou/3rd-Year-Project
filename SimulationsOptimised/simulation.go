package main

import matrix "github.com/skelterjohn/go.matrix"

/*
	Policy
	1. Static
	2. Semi-Dynamic
	3. Dynamic
*/
func simulate(orig scenario, weights []float64, policyCode int) (float64, int) {
	scenario := copyScen(orig)
	rAllGone := false
	bAllGone := false
	change := false
	count := 0
	r := matrix.MakeDenseMatrix(scenario.r, len(scenario.r), 1)
	b := matrix.MakeDenseMatrix(scenario.b, len(scenario.b), 1)
	var alteredConsts *matrix.DenseMatrix
	var alteredEnemy *matrix.DenseMatrix
	if policyCode == 1 {
		alteredConsts = alterFactors(scenario, weights)
		alteredEnemy = enemyEqualSplit(scenario.kB, scenario.r)
	}
	if policyCode == 3 {
		change = true
	}
	for !(rAllGone || bAllGone) {
		if policyCode != 1 && change {
			alteredConsts = alterFactors(scenario, weights)
			alteredEnemy = enemyEqualSplit(scenario.kB, scenario.r)
		}
		dB, _ := alteredConsts.Times(r)
		dR, _ := alteredEnemy.Times(b)
		dB.Scale(0.001)
		dR.Scale(0.001)
		r.Subtract(dR)
		b.Subtract(dB)
		rAllGone, scenario.r = winCondition(scenario.r)
		bAllGone, scenario.b = winCondition(scenario.b)
		count++
	}

	return sum(scenario.r), count
}

func enemyEqualSplit(factors [][]float64, own []float64) *matrix.DenseMatrix {
	count := 0.0
	for i := range own {
		if own[i] >= 0 {
			count++
		}
	}
	factor := matrix.MakeDenseMatrixStacked(factors)
	factor.Scale(1.0 / count)
	return factor
}

func alterFactors(scen scenario, weights []float64) *matrix.DenseMatrix {
	altered := [][]float64{}
	for j := range scen.kR[0] {
		scores := []float64{}
		for i := range scen.kR {
			scored := score([]float64{scen.kR[i][j], scen.kB[j][i], scen.r[j], scen.b[i]}, weights)
			scores = append(scores, scored)
		}
		sum := sum(scores)
		for i := range scores {
			scores[i] = scores[i] * scen.kR[i][j] / sum
		}
		altered = append(altered, scores)
	}
	return matrix.MakeDenseMatrixStacked(altered).Transpose()
}

func score(factors []float64, weights []float64) float64 {
	combos := combs(factors)
	score := 0.0
	for i, v := range combos {
		score += v * weights[i]
	}
	return score
}

func winCondition(side []float64) (bool, []float64) {
	allGone := true
	for i, v := range side {
		if v <= 0 {
			side[i] = 0
		} else {
			allGone = false
		}
	}

	return allGone, side
}
