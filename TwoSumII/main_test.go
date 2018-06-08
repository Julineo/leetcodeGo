package main

import (
	"testing"
	"reflect"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		numbers[]int
		target int
		want []int
	}{
		{[]int{2,7,11,15}, 9, []int{1,2}},
		{[]int{1,2}, 3, []int{1,2}},
}
	for _, test := range tests {
		if got := twoSum(test.numbers, test.target); ! reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v, %v) = %v", test.numbers, test.target, got)
		}
	}
}
