/*
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.
*/

package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := []int{1, 1}
	for k := 2; k < rowIndex + 1; k++ {
		resPrev := res
		tp := make([]int, k+1)
		tp[0], tp[len(tp)-1] = 1, 1
		for i, j := 1, len(tp)-2; i <= j; i, j = i+1, j-1 {
			tp[i], tp[j] = resPrev[i-1] + resPrev[i], resPrev[j] + resPrev[j-1]
		}
		res = tp
	}
	return res
}
