package main

import "testing"

func TestOne(t *testing.T) {
	num := []int{1,2,3,4,5,6}

	sum := 12

	if rob(num) != sum {
		t.Error(`testOne fails`, "1,2,3,4,5,6")
	}

	if rob(num) == sum {
		t.Error(`testOne pass`, "1,2,3,4,5,6")
	}
}

