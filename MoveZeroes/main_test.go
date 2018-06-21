package main

import (
	"testing"
	"reflect"
)

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{[]int{0,1,0,3,12}, []int{1,3,12,0,0}},
		{[]int{0,0,1}, []int{1,0,0}},
		{[]int{0}, []int{0}},
	}
	for _, test := range tests {
		moveZeroes(test.nums)
		if ! reflect.DeepEqual(test.nums, test.want) {
			t.Errorf("moveZeroes(%v) = %v", test.nums, test.want)
		}
	}
}
