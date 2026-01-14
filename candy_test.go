package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandy(t *testing.T) {

	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 5, 4, 3, 2, 1, 4, 3, 2, 1, 3, 2, 1, 1, 2, 3, 4}, 47}, //1234 321

		{[]int{4, 4, 4, 4}, 4}, ///1111
		{[]int{1, 2, 4, 4}, 7}, ///1231

		{[]int{3, 2, 1, 1, 2, 3, 4}, 16},

		{[]int{1, 6, 10, 8, 7, 3, 2}, 18},   //125 4321
		{[]int{1, 6, 10, 11, 7, 3, 2}, 16},  //1234 321
		{[]int{1, 2, 87, 87, 87, 2, 1}, 13}, //1231321

		{[]int{29, 51, 87, 87, 72, 12}, 12}, //123321
		{[]int{1}, 1},
		{[]int{1, 2, 3}, 6},
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4}, //1,2,1
		{[]int{4, 0, 4}, 5},
		{[]int{2, 1}, 3},
		{[]int{1, 2, 5}, 6},
		{[]int{2, 6, 8, 1, 4, 7}, 12}, //1, 2, 3, 1, 2, 3
		{[]int{2, 6, 8, 1, 34}, 9},    // 1, 3, 2,1, 2 - 9
		{[]int{4, 4}, 2},
		{[]int{1, 2, 2, 3}, 6},
	}
	for _, v := range tests {
		c := candy(v.input[:])
		if v.expect != c {
			c = candy(v.input[:])
			assert.Equal(t, v.expect, c, "they should be equal")

		}

	}
}

func TestCalcSum(t *testing.T) {

	tests := []struct {
		input  []int
		pmax   int
		expect int
	}{
		{[]int{6, 10, 8, 7, 3, 2}, 1, 22}, //26 5432
		{[]int{6, 10, 11, 7, 3}, 2, 14},   //234 32
		{[]int{51, 87, 87, 72}, 1, 10},    //2332
		{[]int{1}, 0, 2},                  //2
	}
	for _, v := range tests {
		c := calcSum(v.pmax, v.input[:])
		if v.expect != c {
			assert.Equal(t, v.expect, c, "they should be equal")

		}

	}
}
