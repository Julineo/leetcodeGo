package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", uniquePaths(3,7))
}

func uniquePaths(m int, n int) int {
	path := make([][]int, m)

	for k := 0; k < m; k++ {
		path[k] = make([]int, n)
		path[k][0] = 1
	}

	for k := 0; k < n; k++ {
		path[0][k] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i - 1][j] + path[i][j - 1]
		}
	}

	return path[m - 1][n - 1]
}
