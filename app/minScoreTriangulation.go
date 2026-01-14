package main

import (
	"math"
)

func minScoreTriangulation(values []int) int {
	n := len(values)
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	return minScoreTriang(0, n-1, values, matrix)
}

func minScoreTriang(i, j int, v []int, m [][]int) int {

	if i+1 == j {
		return 0
	}

	if m[i][j] != 0 {
		return m[i][j]
	}
	if i+2 == j {
		return v[i] * v[i+1] * v[j]
	}

	t := math.MaxInt32

	for k := i + 1; k < j; k++ {
		a := minScoreTriang(i, k, v, m)
		b := minScoreTriang(k, j, v, m)
		c := a + b + v[i]*v[k]*v[j]

		if t > c {
			t = c
		}

	}
	m[i][j] = t
	return t
}
