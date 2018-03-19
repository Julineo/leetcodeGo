package main

import (
	"fmt"
)

func main() {
	generateParenthesis(3)
	fmt.Println("ans:", ans)
}

var ans []string

func generateParenthesis(n int) []string {
	backtrack("))))))", 0, 0, n)
	return []string{}
}

func backtrack(s string, left int, right int, n int) {
	if len(s) == 2 * n {
		fmt.Println(s)
		ans = append(ans, s)
		return
	}
}
