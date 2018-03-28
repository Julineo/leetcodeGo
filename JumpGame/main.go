package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
}

func canJump(nums []int) bool {
	dist := 0
	for i := 0; i <= dist; i++  {
		fmt.Printf("dist: %v\n", dist)
		dist = max(dist, i + nums[i])
		if dist >= len(nums) - 1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
