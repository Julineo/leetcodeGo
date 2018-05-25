/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	if numRows == 0 {return [][]int{}}
	if numRows == 1 {return [][]int{{1}}}
    res := [][]int{{1},{1,1}}
    for k := 2; k < numRows; k++ {
        tp := make([]int, k+1)
        tp[0], tp[len(tp)-1] = 1, 1
        for i, j := 1, len(tp)-2; i <= j; i, j = i+1, j-1 {
            tp[i], tp[j] = res[k-1][i-1] + res[k-1][i], res[k-1][j]+res[k-1][j-1]
        }
        res = append(res, tp)
    }
    return res
}
