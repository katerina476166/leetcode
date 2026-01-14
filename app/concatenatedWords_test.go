package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllConcatenatedWordsInADict(t *testing.T) {

	tests := []struct {
		text1  []string
		expect []string
	}{
		{[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
			[]string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}},

		//{[]string{"cat", "dog", "catdog"}, []string{"catdog"}},
	}

	for _, v := range tests {
		c := findAllConcatenatedWordsInADict(v.text1)

		assert.True(t, areEqualIgnoreCase(c, v.expect), "they should be equal")

	}
}

func areEqualIgnoreCase(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if !strings.EqualFold(arr1[i], arr2[i]) { // Case-insensitive comparison
			return false
		}
	}
	return true
}
