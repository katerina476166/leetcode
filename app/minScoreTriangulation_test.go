package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinScoreTriangulation(t *testing.T) {

	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{3, 7, 4, 5}, 144},
		{[]int{1, 3, 1, 4, 1, 5}, 13},
	}
	for _, v := range tests {
		c := minScoreTriangulation(v.input[:])
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
