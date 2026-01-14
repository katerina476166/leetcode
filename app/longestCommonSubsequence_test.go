package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {

	tests := []struct {
		text1  string
		text2  string
		expect int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, v := range tests {
		c := longestCommonSubsequence(v.text1, v.text2)
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
