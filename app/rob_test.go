package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {

	tests := []struct {
		input  *TreeNode
		expect int
	}{
		{&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 5}},
			}}, 12},
		{&TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val: 3,
			}}, 7},
		{&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,

				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 1,
				},
			}}, 7},
		{&TreeNode{
			Val: 3}, 3},
	}
	for _, v := range tests {
		c := rob(v.input)
		assert.Equal(t, v.expect, c, "they should be equal")

	}

}
