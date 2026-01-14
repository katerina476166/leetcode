package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {

	tests := []struct {
		input  int
		expect int
	}{
		{1, 1},
		{3, 5},
	}
	for _, v := range tests {
		c := generateParenthesis(v.input)
		assert.Equal(t, v.expect, len(c), "they should be equal")

	}
}
