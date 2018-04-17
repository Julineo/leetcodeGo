package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", countAndSay(5))
}

func countAndSay(n int) string {
	ns := []string{"1"}
	letter, letterP, res := "", "", ""
	for i := 0; i < n; i++ {

		count := 1
		for _, a := range ns[i] {
			letter = string(a)
			if letterP == letter {
				count++
			} else {
				res = res + strconv.Itoa(count) + letter
				letterP = letter
				count = 1
			}
		}
	ns = append(ns, res)
	}

	return res
}

func countAndSayReversed(n int) string {
	ns := strconv.Itoa(n)
	ns = "13112221"
	cFlag := true
	count, letter, res := "", "", ""
	for _, a := range ns {
		if cFlag {
			count = string(a)
			cFlag = false
			fmt.Println(count)
		} else {
			letter = string(a)
			cFlag = true
			fmt.Println(letter)
			n, _ := strconv.Atoi(count)
			for i := 0; i < n; i++ {
				res = res + letter
			}
		}
	}
	return res
}
