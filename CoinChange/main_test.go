package main

import "testing"

func TestCoinChange(t *testing.T) {
	var tests = []struct {
		coins[]int
		amount int
		want int
	}{
		{[]int{1,2,5}, 11, 3},
		{[]int{2}, 3, -1}}
	for _, test := range tests {
		if got := coinChange(test.coins, test.amount); got != test.want {
			t.Errorf("coinChange(%v) = %v", test.coins, got)
		}
	}
}
