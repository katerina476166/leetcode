package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCoins(t *testing.T) {

	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{3, 1, 5, 8}, 167},
		{[]int{1, 5}, 10},
		{[]int{3, 5, 8}, 152},
		{[]int{5}, 5},
		{[]int{9, 76, 64, 21, 97, 60, 5}, 1088290},
	}
	for _, v := range tests {
		c := maxCoins(v.input[:])
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
