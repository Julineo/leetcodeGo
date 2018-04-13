package main

import "testing"

func TestHammingWeight(t *testing.T) {
	var tests = []struct {
			n uint32
			want int
		}{
			{11,3},
			{1,1},
		}

	for _, test := range tests {
		if got := hammingWeight(test.n); got != test.want {
			t.Errorf("hammingWeight(%032b) = %v" , test.n, got)
		}
	}
}
