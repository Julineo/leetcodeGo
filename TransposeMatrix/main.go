/*
https://leetcode.com/problems/transpose-matrix/description/
Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

 

Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]
Example 2:

Input: [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]
*/

package main

import "fmt"

func transpose(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	B := make([][]int, n)
	for i := 0; i < n; i++ {
		B[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			B[j][i] = A[i][j]
		}
	}
	return B
}

