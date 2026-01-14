package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestCommonSupersequence(t *testing.T) {

	tests := []struct {
		input  string
		input2 string
		expect string
	}{
		{"cab", "abac", "cabac"},
		{"abac", "cab", "cabac"},
		{"aaaaaaaa", "aaaaaaaa", "aaaaaaaa"},
	}
	for _, v := range tests {
		c := shortestCommonSupersequence(v.input, v.input2)
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
