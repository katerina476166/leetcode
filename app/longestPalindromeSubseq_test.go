package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSubseq(t *testing.T) {

	tests := []struct {
		text1  string
		expect int
	}{
		{"aaa", 3},
		{"a", 1},
		{"ab", 1},
		{"bbbab", 4},
		{"cbbd", 2},
	}
	for _, v := range tests {
		c := longestPalindromeSubseq(v.text1)
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
