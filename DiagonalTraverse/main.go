/*
Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

Example:
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output:  [1,2,4,7,5,3,6,8,9]
*/

package main

import "fmt"

func main() {
	m := [][]int{
		 { 1, 2, 3 },
		 { 4, 5, 6 },
		 { 7, 8, 9 },
		}
	fmt.Println(findDiagonalOrder(m))
}

func findDiagonalOrder(matrix [][]int) []int {
	res := []int{}
	fmt.Println(matrix)
	if matrix == nil || len(matrix) == 0 {
		return res
	}
	m := len(matrix) //rows
	n := len(matrix[0]) //columns
	i, j := 0, 0 //indexes
	di, dj := -1, 1 //directions
	for k := 0; k < m * n; k++ {
		res = append(res, matrix[i][j])
		i = i + di
		j = j + dj

		//bottom border switch direction
		if i >= m {
			i = m - 1
			j += 2
			di = -1
			dj = 1
		}
		//right border switch direction
		if j >= n {
			j = n - 1
			i += 2
			di = 1
			dj = -1
		}

		//top border switch direction
		if i < 0 {
			i = 0
			di = 1
			dj = -1
		}
		//left border switch direction
		if j < 0 {
			j = 0
			di = -1
			dj = 1
		}
	}
	return res
}
