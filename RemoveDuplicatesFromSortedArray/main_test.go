package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	var tests = []struct {
		nums[]int
		want int
	}{
		{[]int{1,1,2}, 2},
		{[]int{2}, 1}}
	for _, test := range tests {
		if got := removeDuplicates(test.nums); got != test.want {
			t.Errorf("removeDuplicates(%v) = %v", test.nums, got)
		}
	}
}
