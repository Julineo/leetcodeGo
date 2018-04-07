package main

import "testing"

func TestUniquePaths (t *testing.T) {
	var tests = []struct {
		m int
		n int
		want int
	}{
		{3,7,28}}
	for _, test := range tests {
		if got := uniquePaths(test.m, test.n); got != test.want {
			t.Errorf("uniquePaths(%v,%v) = %v", test.m, test.n, test.want)
		}
	}
}
