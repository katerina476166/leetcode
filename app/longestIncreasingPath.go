package main

func longestIncreasingPath(matrix [][]int) int {

	m := len(matrix)
	n := len(matrix[0])

	var mem [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = 0
		}
	}

	max := 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := longestIncreasingPathCalc(matrix, mem, i, j)
			if v > max {
				max = v
			}
		}
	}

	return max

}

func longestIncreasingPathCalc(matrix [][]int, memory [][]int, i, j int) int {
	m := len(matrix)
	n := len(matrix[0])

	if memory[i][j] != 0 {
		return memory[i][j]
	}

	max := 0
	if i > 0 && matrix[i][j] > matrix[i-1][j] {
		v := longestIncreasingPathCalc(matrix, memory, i-1, j)
		if v > max {
			max = v
		}
	}

	if i < m-1 && matrix[i][j] > matrix[i+1][j] {
		v := longestIncreasingPathCalc(matrix, memory, i+1, j)
		if v > max {
			max = v
		}
	}

	if j > 0 && matrix[i][j] > matrix[i][j-1] {
		v := longestIncreasingPathCalc(matrix, memory, i, j-1)
		if v > max {
			max = v
		}
	}

	if j < n-1 && matrix[i][j] > matrix[i][j+1] {
		v := longestIncreasingPathCalc(matrix, memory, i, j+1)
		if v > max {
			max = v
		}
	}

	memory[i][j] = max + 1
	return memory[i][j]

}
