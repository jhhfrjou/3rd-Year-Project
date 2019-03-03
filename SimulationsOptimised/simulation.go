package main

import (
	matrix "github.com/skelterjohn/go.matrix"
)

/*
	Policy
	1. Static
	2. Semi-Dynamic
	3. Dynamic
*/
func simulate(orig scenario, allocation [][]float64, policyCode int) (float64, int) {
	scenario := copyScen(orig)
	rAllGone := false
	bAllGone := false
	change := false
	count := 0
	r := matrix.MakeDenseMatrix(scenario.R, len(scenario.R), 1)
	b := matrix.MakeDenseMatrix(scenario.B, len(scenario.B), 1)
	var dB matrix.Matrix
	var dR matrix.Matrix
	var alteredConsts *matrix.DenseMatrix
	var alteredEnemy *matrix.DenseMatrix
	if policyCode == 1 {
		alteredConsts = alterFactors(scenario.KR, allocation)
		alteredEnemy = enemyEqualSplit(scenario.KB, scenario.R)
	}
	if policyCode == 3 {
		change = true
	}
	for !(rAllGone || bAllGone) {
		if policyCode != 1 && change {
			alteredConsts = alterFactors(scenario.KR, allocation)
			alteredEnemy = enemyEqualSplit(scenario.KB, scenario.R)
		}
		dB, _ = alteredConsts.Times(r)
		dR, _ = alteredEnemy.Times(b)
		dB.Scale(0.001)
		dR.Scale(0.001)
		r.Subtract(dR)
		b.Subtract(dB)
		rAllGone, scenario.R = winCondition(scenario.R)
		bAllGone, scenario.B = winCondition(scenario.B)
		count++
	}

	return sum(scenario.R) - sum(scenario.B), count
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

func alterFactors(friendlyKill, weights [][]float64) *matrix.DenseMatrix {
	altered := [][]float64{}
	for i := range weights {
		row := make([]float64, len(weights[i]))
		for j := range weights[i] {
			row[j] = friendlyKill[i][j] * weights[i][j]
		}
		altered = append(altered, row)
	}
	return matrix.MakeDenseMatrixStacked(altered)
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
