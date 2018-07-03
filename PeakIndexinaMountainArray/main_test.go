package main

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	var tests = []struct {
		A []int
		want int
	}{
		{[]int{0,1,0},1},
		{[]int{0,2,1,0},1},
	}
	for _, test := range tests {
		if got := peakIndexInMountainArray(test.A); got != test.want  {
			t.Errorf("peakIndexInMountainArray(%v) = %v", test.A, got)
		}
	}
}
