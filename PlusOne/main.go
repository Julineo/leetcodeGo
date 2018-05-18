/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

package main

import (
	. "fmt"
)

func main () {
	Println(plusOne([]int{1,2,3}))
	Println(plusOne([]int{4,3,2,1}))
	Println(plusOne([]int{9,1,0}))
	Println(plusOne([]int{1,9,3}))
	Println(plusOne([]int{1,3,4,9}))
	Println(plusOne([]int{9}))
	Println(plusOne([]int{9,9,9}))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	//add new number in case of 9s
	digits = append([]int{1}, digits...)
	return digits
}

