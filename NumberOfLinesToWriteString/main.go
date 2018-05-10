package main

import "fmt"

func main() {
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func numberOfLines(widths []int, S string) []int {
	ss := "abcdefghijklmnopqrstuvwxyz"
	m := map[rune]int{}
	for i, s := range ss {
		m[s] = widths[i]
	}
	rows := 1
	rest := 0
	for i := 0; i < len(S); i++ {
		if rest+m[rune(S[i])] > 100 {
			rows++
			rest = 0
			i--
			continue
		}
		rest = rest + m[rune(S[i])]
	}
	return []int{rows, rest}
}
