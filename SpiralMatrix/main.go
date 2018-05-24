/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

package main

import (
	"fmt"
)

func main() {
	mx := [][]int{{1,2,3,11},{4,5,6,12},{7,8,9,13},}
	fmt.Println(spiralOrder(mx))
}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if matrix == nil || len(matrix) == 0 {
		return res
	}
	m, n := len(matrix), len(matrix[0])
	i, j := 0, -1
	d := 0
	k := 0
	for {
			//right
			for c := 0; c < n-d; c++ {
				k++
				j++
				res = append(res, matrix[i][j])
			}
			if k >= n * m {break}
			//down
			d++
			for c := 0; c < m-d; c++ {
				k++
				i++
				res = append(res, matrix[i][j])

			}
			if k >= n * m {break}
			//left
			for c := 0; c < n-d; c++ {
				k++
				j--
				res = append(res, matrix[i][j])

			}
			if k >= n * m {break}
			//up
			d++
			for c := 0; c < m-d; c++ {
				k++
				i--
				res = append(res, matrix[i][j])

			}
			if k >= n * m {break}
	}
	return res
}
