package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestIncreasingPath(t *testing.T) {

	tests := []struct {
		input  [][]int
		expect int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{[][]int{{1, 2, 3}, {5, 6, 8}}, 4},

		//{[][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
	}
	for _, v := range tests {
		c := longestIncreasingPath(v.input)
		assert.Equal(t, v.expect, c, "they should  e equal")
	}

}
