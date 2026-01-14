package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEnvelopes(t *testing.T) {

	tests := []struct {
		input  [][]int
		expect int
	}{
		{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{[][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
	}
	for _, v := range tests {
		c := maxEnvelopes(v.input)
		assert.Equal(t, v.expect, c, "they should  e equal")
	}

}
