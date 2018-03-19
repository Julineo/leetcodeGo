package main

import (
	"fmt"
)

func main() {
	generateParenthesis(1)
	fmt.Println("ans:", ans)
}

var ans []string

func generateParenthesis(n int) []string {
	ans = []string{}
	backtrack("", 0, 0, n)
	return ans
}

func backtrack(s string, left int, right int, n int) {
	if len(s) == 2 * n {
		fmt.Println(s)
		ans = append(ans, s)
		return
	}
	if left < n {
		backtrack(s+"(", left + 1, right, n)
	}
	if right < left {
		backtrack(s+")", left, right + 1, n)
	}
}
