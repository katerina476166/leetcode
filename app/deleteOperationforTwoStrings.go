package main

import (
	"unicode/utf8"
)

func minDistance(word1 string, word2 string) int {
	i := 1
	j := 1
	n1 := utf8.RuneCountInString(word1)
	n2 := utf8.RuneCountInString(word2)
	matrix := make([][]int, n1+1)

	for i := 0; i < n1+1; i++ {
		matrix[i] = make([]int, n2+1)
		matrix[i][0] = i
	}

	for i := 0; i < n2+1; i++ {
		matrix[0][i] = i
	}

	for i := 0; i < n1+1; i++ {
	}
	for _, rw1 := range word1 {
		for _, rw2 := range word2 {
			if rw1 == rw2 {
				matrix[i][j] = matrix[i-1][j-1] //то записывае
			} else {
				matrix[i][j] = min(matrix[i-1][j], matrix[i][j-1]) + 1
			}
			j++
		}
		j = 1
		i++
	}

	return matrix[n1][n2]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
