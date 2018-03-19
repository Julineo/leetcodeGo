package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

//Heap's algorithm
func permute(nums []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(nums []int, n int) {
		if n == 1 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(nums, n-1)
				if n%2 == 1 {
					nums[i], nums[n-1] = nums[n-1], nums[i]
				} else {
					nums[0], nums[n-1] = nums[n-1], nums[0]
				}
			}
		}
	}

	helper(nums, len(nums))
	return res
}
