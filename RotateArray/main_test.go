package main

import (
	"testing"
	. "reflect"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		nums[]int
		k int
		want []int
	}{
		{[]int{1,2,3,4,5,6,7}, 3, []int{5,6,7,1,2,3,4}},
		{[]int{}, 3, []int{}},
		{[]int{1,2,3,4,5,6,7}, 8, []int{7,1,2,3,4,5,6}},
		{[]int{1,2,3,4,5,6,7}, 0, []int{1,2,3,4,5,6,7}},

	}
	for _, test := range tests {
		rotate(test.nums, test.k)
		if ! DeepEqual(test.nums, test.want) {
			t.Errorf("%v != %v", test.nums, test.want)
		}
	}
}
