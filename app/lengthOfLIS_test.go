package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {

	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{17, 7, 7, 7, 7, 7, 7}, 1},
	}
	for _, v := range tests {
		c := lengthOfLIS(v.input[:])
		assert.Equal(t, v.expect, c, "they should be equal")

	}
}
