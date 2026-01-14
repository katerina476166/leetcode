package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {

	tests := []struct {
		text1  string
		text2  string
		expect int
	}{
		{"sea", "eat", 2},
		{"leetcode", "etco", 4},
	}
	for _, v := range tests {
		c := minDistance(v.text1, v.text2)
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
