package main

func lengthOfLIS(nums []int) int {
	n := len(nums) + 1
	matrix := make([]int, n)
	max := 1
	for i := 1; i < n; i++ {
		matrix[i] = 1
		for j := 1; j < n; j++ {
			if nums[i-1] > nums[j-1] {
				matrix[i] = maxl(matrix[j]+1, matrix[i])
			}
		}
		if max < matrix[i] {
			max = matrix[i]
		}
	}

	return max

}
func maxl(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
