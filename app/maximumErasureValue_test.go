package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumUniqueSubarray(t *testing.T) {

	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{10000, 1, 10000, 1, 1, 1, 1, 1, 1}, 10001},
		{[]int{4, 2, 4, 5, 6}, 17},
		{[]int{4, 2, 1}, 7},

		{[]int{187, 470, 25, 436, 538, 809, 441, 167, 477, 110, 275, 133, 666, 345,
			411, 459, 490, 266, 987, 965, 429, 166, 809, 340, 467, 318, 125, 165,
			809, 610, 31, 585, 970, 306, 42, 189, 169, 743, 78, 810, 70, 382, 367,
			490, 787, 670, 476, 278, 775, 673, 299, 19, 893, 817, 971, 458, 409, 886, 434}, 16911},
	}
	for _, v := range tests {
		c := maximumUniqueSubarray(v.input[:])
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
