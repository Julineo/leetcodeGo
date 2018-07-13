package main

import (
	"testing",
	
)

func TestTranspose(t *testing.T) {
	var tests = []struct {
		[][]A int
		[][]want int
	}{
		{[][]int{{1,2,3},{4,5,6},{7,8,9}}, [][]int{{1,4,7},{2,5,8},{3,6,9}}},
		{[][]int{{1,2,3},{4,5,6}}, [][]int{{1,4},{2,5},{3,6}}}
	for _, test := range tests {
		if got := transpose(test.A); got != test.want {
			t.Errorf("transpose(%v) = %v", test.A, got)
		}
	}
}
