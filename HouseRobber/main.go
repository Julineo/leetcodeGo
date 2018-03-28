package main 

import (
)

func rob(nums []int) int {
	a := 0
	b := 0

	for i := 0; i < len(nums); i++ {
		if i % 2 == 0 {
			a = max(a + nums[i], b)
		} else {
			b = max(a, b + nums[i])
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
