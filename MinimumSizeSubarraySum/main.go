/*
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum = s. If there isn't one, return 0 instead.

Example: 

Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2,3,1,2,4,3}))
	fmt.Println(minSubArrayLen(0, []int{}))
	fmt.Println(minSubArrayLen(3, []int{1,1}))
}

func minSubArrayLen(s int, nums []int) int {
	min := len(nums)
	hit := false
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= s && min > j - i {
				min = j - i + 1
				hit = true
				break
			}
		}
	}
	if hit {
		return min
	}
	return 0
}
