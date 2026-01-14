package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {

	tests := []struct {
		input  []int
		target int
		expect int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
	}
	for _, v := range tests {
		c := findTargetSumWays(v.input[:], v.target)
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
