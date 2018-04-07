package main

import "testing"

func TestCanJump(t *testing.T) {
	var tests = []struct {
		a[]int
		want bool
	}{
		{[]int{2,3,1,1,4}, true},
		{[]int{3,2,1,0,4}, false}}
	for _, test := range tests {
		if got := canJump(test.a); got != test.want {
			t.Errorf("canJump(%v) = %v", test.a, got)
		}
	}
}
