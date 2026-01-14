package main

import (
	"slices"
)

func maxCoins(nums []int) int {
	nums = slices.Insert(nums, 0, 1)
	nums = append(nums, 1)
	n := len(nums)
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	ans := burst(0, n-1, nums, matrix)

	return ans
}

func burst(i, j int, v []int, m [][]int) int {

	if i+1 == j {
		return 0
	}

	if m[i][j] != 0 {
		return m[i][j]
	}

	if i+2 == j {

		return v[i] * v[i+1] * v[j]
	}

	t := 0

	for k := i + 1; k < j; k++ {
		a := burst(i, k, v, m)
		b := burst(k, j, v, m)
		c := a + b + v[i]*v[k]*v[j]

		if t < c {
			t = c
		}

	}

	m[i][j] = t

	return t
}
