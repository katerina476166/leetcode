package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {

	tests := []struct {
		text1  string
		dic    []string
		expect bool
	}{
		{"abcd", []string{"a", "abc", "b", "cd"}, true},

		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			[]string{"aa"}, false},

		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, false},

		{"leetcode", []string{"leet", "code"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, v := range tests {
		c := wordBreak(v.text1, v.dic)
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
