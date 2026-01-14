package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {

	tests := []struct {
		input  []int
		expect bool
	}{
		{[]int{1, 2, 5}, false},
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 56}, false},
	}
	for _, v := range tests {
		c := canPartition(v.input[:])
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
