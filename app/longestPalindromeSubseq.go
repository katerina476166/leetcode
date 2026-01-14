package main

func longestPalindromeSubseq(s string) int {
	n := len(s)
	var matrix [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				matrix[i][j] = matrix[i+1][j-1] + 2

			} else {
				matrix[i][j] = max(matrix[i][j-1], matrix[i+1][j])
			}

		}

	}
	return matrix[0][n-1]
}
