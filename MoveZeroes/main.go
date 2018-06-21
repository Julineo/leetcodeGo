/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
package main
*/

import "fmt"

func moveZeroes(nums []int)  {
	i := 0
	count := 0
	for {
		if i == len(nums) {
			break
		}
		if nums[i] == 0	{
			nums = append(nums[:i],nums[i+1:]...)
			count++
			continue
		}
		i++
	}
	for j := 0; j < count; j++ {
		nums = append(nums, 0)
	}
}
