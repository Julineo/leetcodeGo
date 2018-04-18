package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", countAndSay(7))
}

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	res := "1"
	for n := n - 1; n > 0; n-- {
		cur := ""
		for i := 0; i < len(res); i++ {
			count := 1
			for ; i + 1 < len(res) && res[i] == res[i+1]; {
				count++
				i++
			}
			cur = cur + strconv.Itoa(count) + string(res[i])
		}
	res = cur
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
