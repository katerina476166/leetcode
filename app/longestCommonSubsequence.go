package main

import (
	"unicode/utf8"
)

func longestCommonSubsequence(text1 string, text2 string) int {

	i := 1
	j := 1
	n1 := utf8.RuneCountInString(text1)
	n2 := utf8.RuneCountInString(text2)
	matrix := make([][]int, n1+1)

	// Initialize each inner slice (row) with 'cols' number of elements
	for i := 0; i < n1+1; i++ {
		matrix[i] = make([]int, n2+1)
	}

	for _, rw1 := range text1 {
		for _, rw2 := range text2 {
			if rw1 == rw2 {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = max(matrix[i-1][j], matrix[i][j-1])
			}
			j++
		}
		j = 1
		i++
	}

	return matrix[n1][n2]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
