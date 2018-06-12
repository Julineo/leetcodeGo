/*
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
*/
package main

import (
	"fmt"
)

func main () {
	fmt.Println(findMaxConsecutiveOnes([]int{1,1,0,1,1,1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1}))
	fmt.Println(findMaxConsecutiveOnes([]int{0}))
	fmt.Println(findMaxConsecutiveOnes([]int{}))
}

func findMaxConsecutiveOnes(nums []int) (max int) {
	start := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			start = i
		}
		if nums[i] == 1 {
			if i - start > max {
				max = i - start
			}
		}
	}
	return
}
