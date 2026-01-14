package main

/*

f(i, w) = f(i-1, w-n[i])+ f(i-1, w+n[i])

f(0,0) = 1
f(0, w) = -1
f(i, 0) = true
w-n[i]<0 = false
*/

func findTargetSumWays(nums []int, target int) int {
	z := make(map[f]int)
	return findTargetSumWays2(nums, len(nums), target, z)
}

func findTargetSumWays2(nums []int, i int, target int, m map[f]int) int {

	if i == 0 {
		if target == 0 {
			return 1
		}

		return 0
	}

	if info, ok := m[f{i, target}]; ok {
		return info
	}

	m[f{i, target}] = findTargetSumWays2(nums, i-1, target-nums[i-1], m) + findTargetSumWays2(nums, i-1, target+nums[i-1], m)

	return m[f{i, target}]

}

type info struct {
	i int
	w int
}
