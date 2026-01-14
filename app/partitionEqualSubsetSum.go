package main

/*

f(i, w) = f(i-1, w-n[i]) or f(i-1, w)

f(0, 0) = true
f(0, w) = false
f(i, 0) = true
w-n[i]<0 = false
*/

func canPartition(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	s := sum(nums)
	if s%2 != 0 {
		return false
	}

	return check(nums, s/2)
}

func sum(nums []int) int {
	s := 0
	for index := 0; index < len(nums); index++ {
		s += nums[index]
	}
	return s
}

func check(nums []int, s int) bool {
	z := make(map[f]bool)
	return check2(nums, len(nums)-1, s, z)

}
func check2(nums []int, i int, s int, m map[f]bool) bool {
	if s < 0 {
		return false
	}
	if s == 0 {
		return true
	}

	if i == 0 {
		return false
	}

	if info, ok := m[f{i, s}]; ok {
		return info
	}
	m[f{i, s}] = check2(nums, i-1, s-nums[i-1], m) || check2(nums, i-1, s, m)

	return m[f{i, s}]
}

type f struct {
	i int
	w int
}
