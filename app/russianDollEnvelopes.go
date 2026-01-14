package main

import (
	"sort"
)

func lengthOfLIS2(items [][]int) int {
	tails := []int{items[0][1]}
	for i := 1; i < len(items); i++ {
		v := items[i][1]
		j := findSuccessor(tails, 0, len(tails)-1, v)
		if j >= 0 {
			tails[j] = v
		} else {
			tails = append(tails, v)
		}
	}
	return len(tails)
}

func findSuccessor(nums []int, left, right, v int) int {
	if nums[right] < v {
		return -1
	}

	if nums[left] >= v {
		return left
	}

	if left == right {
		return -1
	}

	mid := (left + right) / 2
	if nums[mid] >= v {
		return findSuccessor(nums, left, mid, v)
	}

	return findSuccessor(nums, mid+1, right, v)
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		w1, h1 := envelopes[i][0], envelopes[i][1]
		w2, h2 := envelopes[j][0], envelopes[j][1]
		if w1 < w2 {
			return true
		}
		if w1 > w2 {
			return false
		}
		return h1 > h2
	})
	return lengthOfLIS2(envelopes)
}
